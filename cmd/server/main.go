// main.go
package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

type StaticFileServer struct {
	rootDir string
}

func NewStaticFileServer(rootDir string) *StaticFileServer {
	absPath, err := filepath.Abs(rootDir)
	if err != nil {
		log.Fatalf("Failed to get absolute path: %v", err)
	}
	return &StaticFileServer{rootDir: absPath}
}

func (s *StaticFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Prevent directory traversal attacks
	if containsDotDot(r.URL.Path) {
		http.Error(w, "Invalid file path", http.StatusBadRequest)
		return
	}

	// Clean the path and serve the file
	path := filepath.Join(s.rootDir, filepath.Clean(r.URL.Path))
	http.ServeFile(w, r, path)
}

func containsDotDot(path string) bool {
	for _, element := range filepath.SplitList(path) {
		if element == ".." {
			return true
		}
	}
	return false
}

func main() {
	port := flag.String("port", "8080", "Port to serve on")
	directory := flag.String("dir", ".", "Directory to serve files from")
	flag.Parse()

	server := NewStaticFileServer(*directory)

	log.Printf("Starting server on port %s serving directory %s", *port, *directory)
	if err := http.ListenAndServe(":"+*port, server); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
