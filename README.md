# README.md

# Static File Server
A lightweight and secure static file server written in Go, designed for serving files with robust security features and comprehensive testing.

## Features
- 🚀 Fast static file serving
- 🛡️ Built-in security against directory traversal attacks
- 📁 Support for serving files from any directory
- ⚙️ Configurable port and root directory
- 📝 Comprehensive logging
- ✅ Thoroughly tested

## Quick Start
1. Clone the repository:
```bash
git clone https://github.com/nhuzaa/static-file-server.git
cd static-file-server
```

2. Build the server:
```bash
go build -o main main.go 
```

3. Run the server:
```bash
./static-server -dir /path/to/files -port 8080
```

## Configuration
The server accepts the following command-line flags:
- `-dir`: Directory to serve files from (default: current directory)
- `-port`: Port to listen on (default: 8080)

## Usage Examples
1. Serve files from the current directory:
```bash
./static-server
```

2. Serve files from a specific directory on port 3000:
```bash
./static-server -dir /var/www/html -port 3000
```

## License
MIT License

---

# README.dev.md

# Developer Documentation
Technical documentation for developers working on the Static File Server project.

## Project Structure
```
├── main.go           # Main server implementation
├── server_test.go    # Server tests
├── README.md         # Main documentation
├── README.dev.md     # Developer documentation
└── README.test.md    # Testing documentation
```

## Architecture
The server is built around the `StaticFileServer` struct, which implements the `http.Handler` interface:

```go
type StaticFileServer struct {
    rootDir string
}
```

### Key Components
1. **StaticFileServer**: Core server struct
2. **ServeHTTP**: Main request handler
3. **containsDotDot**: Security utility function

## Development Setup
1. Install Go 1.21 or later
2. Install development tools:
```bash
go install golang.org/x/tools/cmd/goimports@latest
go install golang.org/x/lint/golint@latest
```

## Coding Standards
- Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `goimports` for consistent formatting
- Write tests for all new features
- Maintain 80% or higher code coverage

## Contributing
1. Create a feature branch
2. Write tests for new features
3. Ensure all tests pass
4. Submit a pull request

---

# README.test.md

# Testing Documentation
Guide for testing the Static File Server project.

## Test Structure
The test suite covers:
1. Basic file serving
2. Security features
3. Error handling
4. Content type verification

## Running Tests
### Basic Test Run
```bash
go test
```

### Verbose Output
```bash
go test -v
```

### With Coverage
```bash
go test -cover
```

### Generate Coverage Report
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Test Cases
### Core Functionality Tests
- Serving existing files
- Handling non-existent files
- Serving files from subdirectories

### Security Tests
- Directory traversal prevention
- Invalid path handling
- Permission checks

### Content Type Tests
- Text files
- Binary files
- Various MIME types

## Adding New Tests
1. Create test files in temporary directories
2. Define expected outcomes
3. Verify status codes, content types, and bodies
4. Clean up temporary files

## Example Test Pattern
```go
func TestNewFeature(t *testing.T) {
    // Setup
    tmpDir, err := os.MkdirTemp("", "test-")
    if err != nil {
        t.Fatal(err)
    }
    defer os.RemoveAll(tmpDir)

    // Create test files
    // Run tests
    // Verify results
}
```

## CI/CD Integration
The test suite is integrated with:
- GitHub Actions
- Local pre-commit hooks
- Continuous integration pipelines