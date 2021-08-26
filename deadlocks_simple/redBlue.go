package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock1 = sync.Mutex{}
	lock2 = sync.Mutex{}
)

func blueRobot() {
	for {
		fmt.Println("Blue: acquiring lock1")
		lock1.Lock()
		fmt.Println("Blue: acquiring lock2")
		lock2.Lock()
		fmt.Println("Blue: both lock acquired")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Blue: locks has been released")
	}
}

func redRobot() {
	for {
		fmt.Println("Red: acquiring lock1")
		lock2.Lock()
		fmt.Println("Red: acquiring lock2")
		lock1.Lock()
		fmt.Println("Red: both lock acquired")
		lock2.Unlock()
		lock1.Unlock()
		fmt.Println("Red: locks has been released")
	}
}

func main() {
	go blueRobot()
	go redRobot()
	time.Sleep(20 * time.Second)
	fmt.Println("done")
}
