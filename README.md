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

# binary for complete debug , go compiler flags
go build -gcflags "all=-N -l" -o demo .
```

### go env

```bash
# list of env
go env

GOROOT -> Factory installation
GOPATH -> Third party or user defined libraries , cache of packages
GOBIN  -> The binaries are installed in this directory, if GOBIN is not setup then $GOPATH/bin is used 
GOARCH -> amd64,arm64 ..etc..
GOOS   -> windows, darwin, linux, android , ios, mips...

# To check the GOOS/GOARCH for the cross compilation
go tool dist list

# To do the cross compilation

GOOS=linux GOARCH=arm64 go build -o demo_linux_arm64 main.go

GOOS=windows GOARCH=amd64 go build -o demo.exe main.go
```
### go install 

- It compiles , builds, generate binary and installs/copies in the GOBIN directory

```bash
GOBIN=/Users/jiten/workspace/training/shell-golang-training/7-if-else/bin go install github.com/josharian/impl@v1.4.0

# compile,build and install into the GOPATH/bin directory if the go bin is not setup
go install . 
```
