package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	matrixSize = 250
)

var (
	matrixA = [matrixSize][matrixSize]int{}
	matrixB = [matrixSize][matrixSize]int{}
	result  = [matrixSize][matrixSize]int{}
	// have matrixSize thread + 1 matrix writer thread
	workStart    = NewBarrier(matrixSize + 1)
	workComplete = NewBarrier(matrixSize + 1)
)

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			result[row][col] += rand.Intn(10) - 5
		}
	}
}

func workOutRow(row int) {
	for {
		// before to work on the row, make sure the input matrices are populated
		workStart.Wait()
		for col := 0; col < matrixSize; col++ {
			for i := 0; i < matrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
		// signal the job is done, and will wait all the other thread to finish
		workComplete.Wait()
	}
}

func main() {
	fmt.Println("---- start ----")

	// it will create a thread per row
	for row := 0; row < matrixSize; row++ {
		go workOutRow(row)
	}

	// generate 100 times the random matrix
	start := time.Now()
	for i := 0; i < 100; i++ {

		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		// signal when everyone is waiting on that first barrier
		// and release all thread to compute their own role
		workStart.Wait()
		// immediatly block the writer(of matrix) thread
		// when this is unblock, it means all threads have finish their job
		workComplete.Wait()
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
