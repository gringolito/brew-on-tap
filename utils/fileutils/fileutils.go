package fileutils

import (
	"fmt"
	"os"
)

func WriteFile(path string, format string, args ...interface{}) (n int, err error) {
	n = 0
	file, err := os.OpenFile(path, os.O_WRONLY, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	n, err = fmt.Fprintf(file, format, args...)
	return
}
