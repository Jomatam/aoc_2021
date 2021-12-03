package helpers

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func Getlines(path string, fallback string) []string {
	if path == "" {
		path = fallback
	}

	_, b, _, _ := runtime.Caller(0)
	fullpath := filepath.Join(filepath.Dir(b), "..\\", path)
	data, _ := os.ReadFile(fullpath)
	return strings.Split(string(data), "\r\n")
}
