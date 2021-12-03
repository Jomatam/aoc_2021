package helpers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetFile(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	lines := string(data)
	return strings.Split(string(lines), "\n")
}

func Getlines(path string, fallback string) []string {
	if path == "" {
		path = fallback
	}

	_, b, _, _ := runtime.Caller(0)
	fullpath := filepath.Join(filepath.Dir(b), "..\\", path)
	data, _ := os.ReadFile(fullpath)
	return strings.Split(string(data), "\r\n")
}
