package helpers

import (
	"os"
	"strings"
)

func Getlines(path string) []string {
	data, _ := os.ReadFile(path)
	return strings.Split(string(data), "\r\n")
}
