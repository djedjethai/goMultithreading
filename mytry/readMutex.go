package main

import (
	"fmt"
	"sync"
	"time"
)

var rwLock = sync.RWMutex{}

func oneTwoThreeB() {
	rwLock.RLock()
	for i := 1; i <= 300; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}
	rwLock.RLock()
}

func StartThreadsB() {
	for i := 1; i <= 2; i++ {
		go oneTwoThreeB()
	}
	time.Sleep(1 * time.Second)
}

func main() {
	StartThreadsB()
}
