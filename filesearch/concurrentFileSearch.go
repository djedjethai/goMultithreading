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
	fmt.Println("searching in: ", root)
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

// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"path/filepath"
// 	"strings"
// 	"sync"
// )
//
// var (
// 	matches   []string
// 	waitgroup = sync.WaitGroup{}
// 	lock      = sync.Mutex{}
// )
//
// func fileSearch(root string, filename string) {
// 	fmt.Println("searching in: ", root)
// 	files, _ := ioutil.ReadDir(root)
// 	for _, file := range files {
// 		if strings.Contains(file.Name(), filename) {
// 			lock.Lock()
// 			matches = append(matches, filepath.Join(root, filename))
// 			lock.Unlock()
// 		}
// 		if file.IsDir() {
// 			waitgroup.Add(1)
// 			go fileSearch(filepath.Join(root, file.Name()), filename)
// 		}
// 	}
// 	// means all threads are finish, the waitgroup will unlock
// 	waitgroup.Done()
// }
//
// func main() {
// 	p := filepath.Join("/", "home", "jerome", "Documents", "cours")
//
// 	// initialize the waitgroup with just one
// 	waitgroup.Add(1)
// 	go fileSearch(p, "reportHeadAch.odt")
// 	// wait for everything to complete
// 	waitgroup.Wait()
//
// 	for _, file := range matches {
// 		fmt.Println("Matched: ", file)
// 	}
// }
