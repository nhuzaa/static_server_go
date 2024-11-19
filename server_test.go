// server_test.go
package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestStaticFileServer(t *testing.T) {
	// Create a temporary directory for test files
	tmpDir, err := os.MkdirTemp("", "static-server-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a test file
	testContent := "Hello, World!"
	testFilePath := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(testFilePath, []byte(testContent), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create subdirectory with a file
	subDir := filepath.Join(tmpDir, "sub")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}
	subFileContent := "Sub file content"
	subFilePath := filepath.Join(subDir, "sub.txt")
	if err := os.WriteFile(subFilePath, []byte(subFileContent), 0644); err != nil {
		t.Fatalf("Failed to create sub file: %v", err)
	}

	server := NewStaticFileServer(tmpDir)

	tests := []struct {
		name         string
		path         string
		expectedCode int
		expectedBody string
		expectedType string
	}{
		{
			name:         "Existing file",
			path:         "/test.txt",
			expectedCode: http.StatusOK,
			expectedBody: testContent,
			expectedType: "text/plain; charset=utf-8",
		},
		{
			name:         "Subdirectory file",
			path:         "/sub/sub.txt",
			expectedCode: http.StatusOK,
			expectedBody: subFileContent,
			expectedType: "text/plain; charset=utf-8",
		},
		{
			name:         "Non-existent file",
			path:         "/notfound.txt",
			expectedCode: http.StatusNotFound,
			expectedBody: "",
			expectedType: "text/plain; charset=utf-8",
		},
		{
			name:         "Directory traversal attempt",
			path:         "/../secret.txt",
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid file path\n",
			expectedType: "text/plain; charset=utf-8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			w := httptest.NewRecorder()

			server.ServeHTTP(w, req)

			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedCode {
				t.Errorf("Expected status code %d, got %d", tt.expectedCode, resp.StatusCode)
			}

			if tt.expectedBody != "" && string(body) != tt.expectedBody {
				t.Errorf("Expected body %q, got %q", tt.expectedBody, string(body))
			}

			if tt.expectedType != "" && resp.Header.Get("Content-Type") != tt.expectedType {
				t.Errorf("Expected Content-Type %q, got %q", tt.expectedType, resp.Header.Get("Content-Type"))
			}
		})
	}
}
