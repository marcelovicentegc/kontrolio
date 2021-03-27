package utils

import (
	"fmt"
	"os"
)

func CreateFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return file
}

func WriteFile(file *os.File, content string) {
	fmt.Fprintln(file, content)
}

func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
		return
	}
}
