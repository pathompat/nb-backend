# Tickbook backend

This project uses the **Gin** web framework for Go to build fast and lightweight web applications. For live reloading during development, we use **Air**.

---

## Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.23 or later)
- [Air](https://github.com/cosmtrek/air) (for live reloading)

---

## Run project

1. Option 1: simple run

```
  go run main.go
```
2. Option 2: hot reload (need install [Air](https://github.com/cosmtrek/air) first)
- install air (macOS)
```
  curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
- run go with air
```
  air -c .air.toml
```
## Migrate database
1. install [go-migrate](https://github.com/golang-migrate/migrate)
2. run script migrate database
```
  migrate -database 'postgres://{{DB_USER}}:{{DB_PASSWORD}}@{{DB_HOST}}:{{DB_PORT}}/{{DB_NAME}}' -path db/migrations up
```