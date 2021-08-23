package main

import (
	"fmt"
	// "sync"
	// "time"
)

func runConsumer(channel chan string) {
	msg := <-channel
	fmt.Println("Consumer, received", msg)
	channel <- "Bye"
}

func RunProducer() {
	channel := make(chan string)
	go runConsumer(channel)
	fmt.Println("Producer Sending Hello")
	channel <- "Hello"
	fmt.Println("Producer, received", <-channel)
}

func main() {
	RunProducer()
	// affiche
	/* Producer Sending Hello
	Consumer, received Hello
	Producer, received Bye */
}

// var rwLock = sync.RWMutex{}
//
// func oneTwoThreeB() {
// 	rwLock.RLock()
// 	for i := 1; i <= 300; i++ {
// 		fmt.Println(i)
// 		time.Sleep(1 * time.Millisecond)
// 	}
// 	rwLock.RLock()
// }
//
// func StartThreadsB() {
// 	for i := 1; i <= 2; i++ {
// 		go oneTwoThreeB()
// 	}
// 	time.Sleep(1 * time.Second)
// }
//
// func main() {
// 	StartThreadsB()
// }
