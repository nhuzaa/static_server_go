package server

import (
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
	if containsDotDot(r.URL.Path) {
		http.Error(w, "Invalid file path", http.StatusBadRequest)
		return
	}

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
