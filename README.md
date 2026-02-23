### go mod 

- go mod is a package and dependency manager 
- every project must have the go mod file (once it was not there but from version 1.15 > it is must)

```bash
go mod init demo
```

- can create two kinds of projects, 1. binary 2. library
- binary project contains main function in main package

### to run go application

```bash
go run main.go
go run .
```
- It compiles , create .o(object file) file by assemler and linked using linker -> binary in temp location , and then it runs the binary file and deletes that binary 
- Where all binary is created 
x
```bash
go run --work .
```

### go build 

- build the application means building the binary

```bash
go build .
go build main.go
# named binary
go build -o hello main.go

# binary for the release mode 

go build -ldflags="-s -w" -o hello_small main.go
```
