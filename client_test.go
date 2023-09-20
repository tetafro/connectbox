package connectbox

import (
	"context"
	"net/http"
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

	t.Run("failed xml request", func(t *testing.T) {
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
			BodyString("token=token1&fun=15&Username=bob&Password=" +
				"65e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5").
			Reply(http.StatusInternalServerError)

		err = client.Login(context.Background())
		require.ErrorContains(t, err, "xml request")
	})

	t.Run("error response", func(t *testing.T) {
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
			BodyString("fail")

		err = client.Login(context.Background())
		require.ErrorContains(t, err, "invalid response")
	})

	t.Run("missing sid in response", func(t *testing.T) {
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
			BodyString("success;key1=value1;key2=value2")

		err = client.Login(context.Background())
		require.ErrorContains(t, err, "missing SID")
	})

	t.Run("wrong address", func(t *testing.T) {
		client, err := NewClient("http://127.0.0.100", "bob", "qwerty")
		require.NoError(t, err)

		err = client.Login(context.Background())
		require.ErrorContains(t, err, "connection refused")
	})
}

func TestClient_Logout(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		defer gock.Off()

		client, err := NewClient("http://127.0.0.1", "bob", "qwerty")
		require.NoError(t, err)
		client.token = "token1"

		gock.InterceptClient(client.http)

		gock.New("http://127.0.0.1").
			Post(xmlSetter).
			BodyString("token=token1&fun=16").
			Reply(http.StatusOK)

		err = client.Logout(context.Background())
		require.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		defer gock.Off()

		client, err := NewClient("http://127.0.0.1", "bob", "qwerty")
		require.NoError(t, err)
		client.token = "token1"

		gock.InterceptClient(client.http)

		gock.New("http://127.0.0.1").
			Post(xmlSetter).
			BodyString("token=token1&fun=16").
			Reply(http.StatusInternalServerError)

		err = client.Logout(context.Background())
		require.ErrorContains(t, err, "invalid response status")
	})
}

func TestClient_Get(t *testing.T) {
	t.Run("valid response", func(t *testing.T) {
		defer gock.Off()

		client, err := NewClient("http://127.0.0.1", "bob", "qwerty")
		require.NoError(t, err)
		client.token = "token1"

		gock.InterceptClient(client.http)

		gock.New("http://127.0.0.1").
			Post(xmlGetter).
			BodyString("token=token1&fun=999").
			Reply(http.StatusOK).
			BodyString(`<?xml version="1.0"?><root><field>50</field></root>`)

		var data struct {
			Field string `xml:"field"`
		}
		err = client.Get(context.Background(), "999", &data)
		require.NoError(t, err)
		require.Equal(t, "50", data.Field)
	})

	t.Run("invalid xml in response", func(t *testing.T) {
		defer gock.Off()

		client, err := NewClient("http://127.0.0.1", "bob", "qwerty")
		require.NoError(t, err)
		client.token = "token1"

		gock.InterceptClient(client.http)

		gock.New("http://127.0.0.1").
			Post(xmlGetter).
			BodyString("token=token1&fun=999").
			Reply(http.StatusOK).
			BodyString("<?xml")

		var data struct {
			Field string `xml:"field"`
		}
		err = client.Get(context.Background(), "999", &data)
		require.ErrorContains(t, err, "unmarshal response")
	})

	t.Run("error response code", func(t *testing.T) {
		defer gock.Off()

		client, err := NewClient("http://127.0.0.1", "bob", "qwerty")
		require.NoError(t, err)
		client.token = "token1"

		gock.InterceptClient(client.http)

		gock.New("http://127.0.0.1").
			Post(xmlGetter).
			BodyString("token=token1&fun=999").
			Reply(http.StatusInternalServerError)

		var data struct {
			Field string `xml:"field"`
		}
		err = client.Get(context.Background(), "999", &data)
		require.ErrorContains(t, err, "invalid response status")
	})

	t.Run("wrong address", func(t *testing.T) {
		defer gock.Off()

		client, err := NewClient("http://127.0.0.100", "bob", "qwerty")
		require.NoError(t, err)
		client.token = "token1"

		var data struct {
			Field string `xml:"field"`
		}
		err = client.Get(context.Background(), "999", &data)
		require.ErrorContains(t, err, "connection refused")
	})
}
