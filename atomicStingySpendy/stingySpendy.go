package main

import (
	"fmt"
	"time"
)

var (
	money int32 = 100
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		fmt.Println("Stingy sees balance of: ", money)
		time.Sleep(1 * time.Millisecond)
	}
	println("stingy done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		fmt.Println("Spendy sees balance of: ", money)
		time.Sleep(1 * time.Millisecond)
	}
	println("spendy done")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(money)
}

// package main
//
// import (
// 	"fmt"
// 	"sync/atomic"
// 	"time"
// )
//
// var (
// 	money int32 = 100
// )
//
// func stingy() {
// 	for i := 1; i <= 1000; i++ {
// 		atomic.AddInt32(&money, 10)
// 		fmt.Println("Stingy sees balance of: ", money)
// 		time.Sleep(1 * time.Millisecond)
// 	}
// 	println("stingy done")
// }
//
// func spendy() {
// 	for i := 1; i <= 1000; i++ {
// 		atomic.AddInt32(&money, -10)
// 		fmt.Println("Spendy sees balance of: ", money)
// 		time.Sleep(1 * time.Millisecond)
// 	}
// 	println("spendy done")
// }
//
// func main() {
// 	go stingy()
// 	go spendy()
// 	time.Sleep(3000 * time.Millisecond)
// 	print(money)
// }
