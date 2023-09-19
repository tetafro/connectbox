# ConnectBox Client

HTTP client for ConnectBox routers used by Ziggo internet provider in the
Netherlands.

## Example

```go
client, err := NewConnectBox("http://192.168.178.1", "NULL", "password")
if err != nil {
    log.Fatalf("Failed to init ConnectBox client: %v", err)
}

if err := client.Login(ctx); err != nil {
    log.Fatalf("Failed to login: %v", err)
}

var data CMSystemInfo
err := client.GetMetrics(ctx, FnCMSystemInfo, &data)
if err != nil {
    log.Fatalf("Failed to get CMSystemInfo: %v", err)
}
log.Printf("System uptime: %s", data.SystemUptime)

if err := client.Logout(ctx); err != nil {
    log.Fatalf("Failed to logout: %v", err)
}
```
