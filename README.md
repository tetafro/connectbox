# ConnectBox Client

[![License](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/tetafro/connectbox/master/LICENSE)
[![Github CI](https://img.shields.io/github/actions/workflow/status/tetafro/connectbox/push.yml)](https://github.com/tetafro/connectbox/actions)
[![Go Report](https://goreportcard.com/badge/github.com/tetafro/connectbox)](https://goreportcard.com/report/github.com/tetafro/connectbox)
[![Codecov](https://codecov.io/gh/tetafro/connectbox/branch/master/graph/badge.svg)](https://codecov.io/gh/tetafro/connectbox)

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
