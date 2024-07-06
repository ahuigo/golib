package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if err := filepath.Walk(".", watchDir); err != nil {
		fmt.Println("ERROR", err)
	}

}

func watchDir(path string, fi os.FileInfo, err error) error {
	fmt.Println(path, "is dir?", fi.Mode().IsDir())
	return nil
}
