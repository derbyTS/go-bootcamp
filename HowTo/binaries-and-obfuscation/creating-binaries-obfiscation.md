```sh
GOOS=darwin GOARCH=arm64 go build -o binaries/rest_api_macOs_with_middleware_http2_tls server.go
GOOS=darwin GOARCH=arm64 go build -o binaries/rest_api_macOs_TLS_No_Middleware server.go
GOOS=darwin GOARCH=arm64 go build -o binaries/rest_api_macOs_no_TLS_with_middleware server.go
GOOS=darwin GOARCH=arm64 go build -o binaries/rest_api_macOs_no_TLS_No_Middleware server.go
```

```sh
garble build -o obfuscated/rest_api_macOs_with_middleware_http2_tls server.go
garble build -o obfuscated/rest_api_macOs_TLS_No_Middleware server.go
garble build -o obfuscated/rest_api_macOs_no_TLS_with_middleware server.go
garble build -o obfuscated/rest_api_macOs_no_TLS_No_Middleware server.go
```
