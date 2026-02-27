- To Get a package

```bash
go get -u github.com/gin-gonic/gin
go mod tidy
```

docker network create  shell_demo_network

docker run -d --name pgui --network shell_demo_network -p 18080:8080 adminer

 docker run -d --name pg --network shell_demo_network -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=demodb postgres


- test a func

```bash
    go test -test.fullpath=true -timeout 30s -run ^TestUserValidateSuccess$ http-demo/models
```

- test all test files in the project

```bash
go test ./...
```
-- test a pacakge with coverage

```bash
    go test -test.fullpath=true -timeout 30s -coverprofile=cover.out http-demo/models
```

-- to open test coverage in html format

```go tool cover -html=coverage.out
```