
# Static File Server

A lightweight static file server written in Go.

## Install

```bash
git clone https://github.com/nhuzaa/static-file-server
cd static-file-server
go build -o static-server cmd/server/main.go
```

## Usage

Start the server:
```bash
./static-server -dir /static-files port 8080
```

Options:
- `-dir`: Directory to serve (default: current directory)
- `-port`: Port number (default: 8080)

## Test

```bash
go test -v ./pkg/server
```

## Project Structure

```
├── cmd/server/    # Main application
├── pkg/server/    # Server implementation
└── test-files/    # Example files
```

## License

MIT

## Project Structure
```
static-file-server/
├── cmd/
│   └── server/
│       └── main.go
├── pkg/
│   └── server/
│       ├── server.go
│       └── server_test.go
├── test-files/
│   ├── test.txt
│   └── subdir/
│       └── sub.txt
├── go.mod
└── static-server (binary)
```
