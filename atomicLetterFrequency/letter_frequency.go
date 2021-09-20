package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency *[26]int32, waitGroup *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("err Get req")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// for the purpuse of the example, to increase the processing time
	// parse 20 times each text
	for i := 0; i <= 20; i++ {
		// parse each byte of the dowloaded text
		for b := range body {
			c := strings.ToLower(string(b))
			// get the index of the position of each letter
			index := strings.Index(allLetters, c)
			if index >= 0 {
				atomic.AddInt32(&frequency[index], 1)
			}
		}
	}
	waitGroup.Done()
}

func main() {
	var frequency [26]int32
	waitGroup := sync.WaitGroup{}

	start := time.Now()
	for i := 1000; i < 1200; i++ {
		waitGroup.Add(1)
		go countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency, &waitGroup)
	}
	waitGroup.Wait()

	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Printf("Processing took: %s\n", elapsed)
	for i, f := range frequency {
		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
	}
	//  loop 20 times on each doc(using ATOMIC) takes: Processing took: 3.651901934s
	//  loop 20 times on each doc(using mutex) takes: Processing took: 37.007677996s

	// single thread: Processing took: 1m6.786431343s
	// multithread with mutex: Processing took: 3.757501678s
}

// multithreading with mutex
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// 	"sync"
// 	"time"
// )
//
// var lock = sync.Mutex{}
//
// const allLetters = "abcdefghijklmnopqrstuvwxyz"
//
// func countLetters(url string, frequency *[26]int32, waitGroup *sync.WaitGroup) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("err Get req")
// 	}
// 	defer resp.Body.Close()
//
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	// for the purpuse of the example, to increase the processing time
// 	// parse 20 times each text
// 	for i := 0; i <= 20; i++ {
// 		// parse each byte of the dowloaded text
// 		for b := range body {
// 			c := strings.ToLower(string(b))
// 			// get the index of the position of each letter
// 			lock.Lock()
// 			index := strings.Index(allLetters, c)
// 			if index >= 0 {
// 				frequency[index] += 1
// 			}
// 			lock.Unlock()
// 		}
// 	}
// 	waitGroup.Done()
// }
//
// func main() {
// 	var frequency [26]int32
// 	waitGroup := sync.WaitGroup{}
//
// 	start := time.Now()
// 	for i := 1000; i < 1200; i++ {
// 		waitGroup.Add(1)
// 		go countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency, &waitGroup)
// 	}
// 	waitGroup.Wait()
//
// 	elapsed := time.Since(start)
// 	fmt.Println("Done")
// 	fmt.Printf("Processing took: %s\n", elapsed)
// 	for i, f := range frequency {
// 		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
// 	}
// 	//  loop 20 times on each doc(using mutex) takes: Processing took: 37.007677996s
//
// 	// single thread: Processing took: 1m6.786431343s
// 	// multithread with mutex: Processing took: 3.757501678s
// }

// ==========================================
// single thread
// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// 	"time"
// )
//
// const allLetters = "abcdefghijklmnopqrstuvwxyz"
//
// func countLetters(url string, frequency *[26]int32) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("err Get req")
// 	}
// 	defer resp.Body.Close()
//
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	// parse each byte of the dowloaded text
// 	for b := range body {
// 		c := strings.ToLower(string(b))
// 		// get the index of the position of each letter
// 		index := strings.Index(allLetters, c)
// 		if index >= 0 {
// 			frequency[index] += 1
// 		}
// 	}
// }
//
// func main() {
// 	var frequency [26]int32
//
// 	start := time.Now()
// 	for i := 1000; i < 1200; i++ {
// 		countLetters(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &frequency)
// 	}
//
// 	elapsed := time.Since(start)
// 	fmt.Println("Done")
// 	fmt.Printf("Processing took: %s\n", elapsed)
// 	for i, f := range frequency {
// 		fmt.Printf("%s -> %d\n", string(allLetters[i]), f)
// 	}
// 	// single thread: Processing took: 1m6.786431343s
// }
