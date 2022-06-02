### Getting Started

```
## run service

DSN=file::memory:?cache=shared DIALECT=sqlite go run cmd/main.go
```

```
## sample upload
curl -i -X POST -F "file=@sample/ports.json" http://localhost:5050/upload
```

```
## test

go test ./...
```

```
## with docker-compose
 docker-compose up
```