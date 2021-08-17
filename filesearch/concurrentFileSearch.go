package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var (
	matches []string
)

func fileSearch(root string, filename string) {
	files, _ := ioutil.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			matches = append(matches, filepath.Join(root, filename))
		}
		if file.IsDir() {
			fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}
}

func main() {
	p := filepath.Join("/", "home", "jerome", "Documents", "cours")

	fileSearch(p, "reportHeadAch.odt")
	for _, file := range matches {
		fmt.Println("Matched: ", file)
	}
}
