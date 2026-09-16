### Recommend by instructor

```text

Project-root/
├── cmd/ (This folder contains the main entry point of your application. Typically one per binary.)
│   └── api/
│       ├── server.go (Entry point)
│       └── .env
├── internal/ (This folder contains the private application code, including your API handlers, models, repos, etc.)
│   ├── api/
│   │   ├── handlers/ (Use plural)
│   │   ├── middlewares/ (Use plural)
│   │   └── routers.go (Consider keeping it singular if it represents a single module)
│   ├── models/ (Use plural)
│   └── repositories/ (Use plural)
│       ├── mongodb/
│       │   └── mongoconnect.go
│       └── sqlconnect/
│           └── sqlconfig.go
├── pkg/
│   └── utils/ (Use plural)
│       ├── error_handling.go
│       └── jwt_processing.go
├── proto/
├── go.mod
└── go.sum
```
