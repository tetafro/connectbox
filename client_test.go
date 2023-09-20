package connectbox

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	t.Run("valid config", func(t *testing.T) {
		client, err := NewClient("127.0.0.1:8080", "bob", "qwerty")
		require.NoError(t, err)
		require.Equal(t, "http://127.0.0.1:8080", client.addr)
		require.Equal(t, "bob", client.username)
		require.Equal(t,
			"65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5",
			client.password)
	})

	t.Run("invalid address", func(t *testing.T) {
		_, err := NewClient("hello, world!", "bob", "qwerty")
		require.ErrorContains(t, err, "invalid address")
	})
}

func TestClient_Login(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer gock.Off()

		client, err := NewClient("http://127.0.0.1", "bob", "qwerty")
		require.NoError(t, err)

		gock.InterceptClient(client.http)

		gock.New("http://127.0.0.1").
			Get(loginPage).
			Reply(http.StatusOK).
			AddHeader("Set-Cookie", "sessionToken=token1; Path=/")
		gock.New("http://127.0.0.1").
			Post(xmlSetter).
			MatchHeader("Cookie", "sessionToken=token1").
			BodyString("token=token1&fun=15&Username=bob&Password="+
				"65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5").
			Reply(http.StatusOK).
			AddHeader("Set-Cookie", "sessionToken=token2; Path=/").
			BodyString("success;SID=sid1")

		err = client.Login(context.Background())
		require.NoError(t, err)
		require.Equal(t, "token2", client.getCookie(sessionTokenName))
		require.Equal(t, "sid1", client.getCookie(sessionIDName))
	})

	t.Run("failed to get initial token", func(t *testing.T) {
		defer gock.Off()

		client, err := NewClient("http://127.0.0.1", "bob", "qwerty")
		require.NoError(t, err)

		gock.InterceptClient(client.http)

		gock.New("http://127.0.0.1").
			Get(loginPage).
			Reply(http.StatusInternalServerError)

		err = client.Login(context.Background())
		require.ErrorContains(t, err, "get initial token")
	})
}

func TestClient_Logout(t *testing.T) {
	ctx := context.Background()

	connectbox := testConnectBox{
		status: http.StatusOK,
	}
	server := httptest.NewServer(&connectbox)
	defer server.Close()

	client, err := NewClient(server.URL, "bob", "qwerty")
	require.NoError(t, err)
	client.token = "abc"

	err = client.Logout(ctx)
	require.NoError(t, err)

	want := "token=abc&fun=16"
	require.Equal(t, want, connectbox.req)
}

func TestClient_Get(t *testing.T) {
	t.Run("valid response", func(t *testing.T) {
		ctx := context.Background()

		connectbox := testConnectBox{
			status: http.StatusOK,
			resp:   `<?xml version="1.0"?><root><field>50</field></root>`,
		}
		server := httptest.NewServer(&connectbox)
		defer server.Close()

		client, err := NewClient(server.URL, "bob", "qwerty")
		require.NoError(t, err)
		client.token = "abc"

		var data struct {
			Field string `xml:"field"`
		}
		err = client.Get(ctx, "100", &data)
		require.NoError(t, err)
		require.Equal(t, "token=abc&fun=100", connectbox.req)
		require.Equal(t, "50", data.Field)
	})

	t.Run("invalid response", func(t *testing.T) {
		ctx := context.Background()

		connectbox := testConnectBox{
			status: http.StatusOK,
			resp:   `<?xml`,
		}
		server := httptest.NewServer(&connectbox)
		defer server.Close()

		client, err := NewClient(server.URL, "bob", "qwerty")
		require.NoError(t, err)
		client.token = "abc"

		var data struct {
			Field string `xml:"field"`
		}
		err = client.Get(ctx, "100", &data)
		require.ErrorContains(t, err, "unmarshal response")
	})
}

func TestClient_xmlRequest(t *testing.T) {
	t.Run("valid response", func(t *testing.T) {
		ctx := context.Background()

		connectbox := testConnectBox{
			cookies: map[string]string{"token": "def"},
			status:  http.StatusOK,
			resp:    "hello, world",
		}
		server := httptest.NewServer(&connectbox)
		defer server.Close()

		client, err := NewClient(server.URL, "bob", "qwerty")
		require.NoError(t, err)
		client.token = "abc"

		args := xmlArgs{{"key", "value"}}
		resp, err := client.xmlRequest(ctx, "/test", "100", args)
		require.NoError(t, err)

		want := "token=abc&fun=100&key=value"
		require.Equal(t, want, connectbox.req)
		want = "hello, world"
		require.Equal(t, want, resp)
		require.Equal(t, "def", client.getCookie("token"))
	})

	t.Run("invalid status code", func(t *testing.T) {
		ctx := context.Background()

		connectbox := testConnectBox{
			status: http.StatusInternalServerError,
		}
		server := httptest.NewServer(&connectbox)
		defer server.Close()

		client, err := NewClient(server.URL, "bob", "qwerty")
		require.NoError(t, err)
		client.token = "abc"

		args := xmlArgs{{"key", "value"}}
		_, err = client.xmlRequest(ctx, "/test", "100", args)
		require.ErrorContains(t, err, "invalid response status")
	})
}

func TestClient_get(t *testing.T) {
	ctx := context.Background()

	connectbox := testConnectBox{
		cookies: map[string]string{"token": "def"},
		status:  http.StatusOK,
		resp:    "hello, world",
	}
	server := httptest.NewServer(&connectbox)
	defer server.Close()

	client, err := NewClient(server.URL, "bob", "qwerty")
	require.NoError(t, err)
	client.token = "abc"

	resp, err := client.get(ctx, "/test")
	require.NoError(t, err)
	require.Equal(t, "hello, world", resp)
	require.Equal(t, "def", client.getCookie("token"))
}

type testConnectBox struct {
	// Save input data
	path string
	req  string
	// Respond with provided data
	cookies map[string]string
	status  int
	resp    string
}

func (t *testConnectBox) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Save input data
	t.path = r.URL.Path
	body, err := io.ReadAll(r.Body)
	if err == nil {
		t.req = string(body)
	}
	// Respond with provided data
	for name, val := range t.cookies {
		http.SetCookie(w, &http.Cookie{Name: name, Value: val})
	}
	w.WriteHeader(t.status)
	w.Write([]byte(t.resp))
}
