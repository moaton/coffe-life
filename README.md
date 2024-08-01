# Coffe-Life
> Серверная часть приложения кофейни

```bash
├── coffe-life/
│   ├── cmd/
│   │   └── main.go
│   ├── config/
│   │   ├── config.go
│   │   └── config.yaml
│   ├── internal/
│   │   ├── application/
│   │   │   └── application.go
│   │   ├── controller/
│   │   │   └── http/
│   │   │       ├── v1/
│   │   │       │   ├── admin.go
│   │   │       │   ├── coffe.go
│   │   │       │   ├── errors.go
│   │   │       │   └── handler.go
│   │   │       └── router.go
│   │   ├── dto
│   │   ├── entity
│   │   ├── interfaces
│   │   ├── repository
│   │   ├── usecase/
│   │   │   └── admin
│   │   └── utils
│   ├── mocks
│   ├── migrations/
│   │   ├── *.sql
│   │   └── migrations.go
│   └── pkg
├── go.mod
├── go.sum
└── Makefile
```
