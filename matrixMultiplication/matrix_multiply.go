package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	matrixSize = 250
)

var (
	matrixA = [matrixSize][matrixSize]int{}
	matrixB = [matrixSize][matrixSize]int{}
	result  = [matrixSize][matrixSize]int{}
	rwLock  = sync.RWMutex{}
	// need to give the reader portion on this condition variable
	// doing this when worker thread(reader) are waiting for work
	// they unlock the reader portion of this mutex
	cond      = sync.NewCond(rwLock.RLocker())
	waitGroup = sync.WaitGroup{}
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			result[row][col] += rand.Intn(10) - 5
		}
	}
}

func workOutRow(row int) {
	rwLock.RLock()
	for {
		// to make sure the broadcast from master thread
		// wait the thread to be Done(otherwise, it won t be received)
		waitGroup.Done()
		cond.Wait()
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
	}
}

func main() {
	fmt.Println("---- start ----")

	waitGroup.Add(matrixSize)
	// it will create a thread per row
	for row := 0; row < matrixSize; row++ {
		go workOutRow(row)
	}

	// generate 100 times the random matrix
	start := time.Now()
	for i := 0; i < 100; i++ {
		// make sure our worker thread are waiting on this conditional wait
		waitGroup.Wait()
		// try to aquire this readerWriter's lock from the writer point of view
		// and if can it means the reader's thread(workers) have release their lock
		// and are waiting
		rwLock.Lock()
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		// need to reset the waitGroup as the previous one has been released
		waitGroup.Add(matrixSize)
		// release the writer lock
		rwLock.Unlock()
		// broadcast for the computation to start
		cond.Broadcast()
	}

	fmt.Println("---- done ----")
	elapsed := time.Since(start)

	fmt.Println("elapsed time: ", elapsed)
	// no concurrency: elapsed time:  2.888642001s
}

// package main
//
// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )
//
// const (
// 	matrixSize = 250
// )
//
// var (
// 	matrixA = [matrixSize][matrixSize]int{}
// 	matrixB = [matrixSize][matrixSize]int{}
// 	result  = [matrixSize][matrixSize]int{}
// )
//
// func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
// 	for row := 0; row < matrixSize; row++ {
// 		for col := 0; col < matrixSize; col++ {
// 			result[row][col] += rand.Intn(10) - 5
// 		}
// 	}
// }
//
// func workOutRow(row int) {
// 	for col := 0; col < matrixSize; col++ {
// 		for i := 0; i < matrixSize; i++ {
// 			result[row][col] += matrixA[row][i] * matrixB[i][col]
// 		}
// 	}
// }
//
// func main() {
// 	fmt.Println("---- start ----")
//
// 	// generate 100 times the random matrix
// 	start := time.Now()
// 	for i := 0; i < 100; i++ {
// 		generateRandomMatrix(&matrixA)
// 		generateRandomMatrix(&matrixB)
// 		for row := 0; row < matrixSize; row++ {
// 			workOutRow(row)
// 		}
// 	}
//
// 	fmt.Println("---- done ----")
// 	elapsed := time.Since(start)
//
// 	fmt.Println("elapsed time: ", elapsed)
// 	// no concurrency: elapsed time:  2.888642001s
// }
