// USING MUTEX
package main

import (
	"math/rand"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

const (
	totalAccounts  = 50000
	maxAmountMoved = 10
	initialMoney   = 100
	threads        = 4
)

func perform_movements(ledger *[totalAccounts]int32, locks *[totalAccounts]sync.Locker, totalTrans *int64) {
	for {
		accountA := rand.Intn(totalAccounts)
		accountB := rand.Intn(totalAccounts)
		for accountA == accountB {
			accountB = rand.Intn(totalAccounts)
		}
		amountToMove := rand.Int31n(maxAmountMoved)
		toLock := []int{accountA, accountB}
		sort.Ints(toLock)

		locks[toLock[0]].Lock()
		locks[toLock[1]].Lock()

		atomic.AddInt32(&ledger[accountA], -amountToMove)
		atomic.AddInt32(&ledger[accountB], amountToMove)
		atomic.AddInt64(totalTrans, 1)

		locks[toLock[1]].Unlock()
		locks[toLock[0]].Unlock()
	}
}

func main() {
	println("Total accounts:", totalAccounts, " total threads:", threads, "using Mutex")
	var ledger [totalAccounts]int32
	var locks [totalAccounts]sync.Locker
	var totalTrans int64
	for i := 0; i < totalAccounts; i++ {
		ledger[i] = initialMoney
		locks[i] = &sync.Mutex{} // NewSpinLock() //&sync.Mutex{}
	}
	for i := 0; i < threads; i++ {
		go perform_movements(&ledger, &locks, &totalTrans)
	}
	for {
		time.Sleep(2000 * time.Millisecond)
		var sum int32
		for i := 0; i < totalAccounts; i++ {
			locks[i].Lock()
		}
		for i := 0; i < totalAccounts; i++ {
			sum += ledger[i]
		}
		for i := 0; i < totalAccounts; i++ {
			locks[i].Unlock()
		}
		println(totalTrans, sum)
	}
}

// USING SPIN-LOCKS
// package main
//
// import (
// 	"math/rand"
// 	"sort"
// 	"sync"
// 	"sync/atomic"
// 	"time"
// )
//
// const (
// 	totalAccounts  = 50000
// 	maxAmountMoved = 10
// 	initialMoney   = 100
// 	threads        = 4
// )
//
// func perform_movements(ledger *[totalAccounts]int32, locks *[totalAccounts]sync.Locker, totalTrans *int64) {
// 	for {
// 		accountA := rand.Intn(totalAccounts)
// 		accountB := rand.Intn(totalAccounts)
// 		for accountA == accountB {
// 			accountB = rand.Intn(totalAccounts)
// 		}
// 		amountToMove := rand.Int31n(maxAmountMoved)
// 		toLock := []int{accountA, accountB}
// 		sort.Ints(toLock)
//
// 		locks[toLock[0]].Lock()
// 		locks[toLock[1]].Lock()
//
// 		atomic.AddInt32(&ledger[accountA], -amountToMove)
// 		atomic.AddInt32(&ledger[accountB], amountToMove)
// 		atomic.AddInt64(totalTrans, 1)
//
// 		locks[toLock[1]].Unlock()
// 		locks[toLock[0]].Unlock()
// 	}
// }
//
// func main() {
// 	println("Total accounts:", totalAccounts, " total threads:", threads, "using SpinLocks")
// 	var ledger [totalAccounts]int32
// 	var locks [totalAccounts]sync.Locker
// 	var totalTrans int64
// 	for i := 0; i < totalAccounts; i++ {
// 		ledger[i] = initialMoney
// 		locks[i] = NewSpinLock() //&sync.Mutex{}
// 	}
// 	for i := 0; i < threads; i++ {
// 		go perform_movements(&ledger, &locks, &totalTrans)
// 	}
// 	for {
// 		time.Sleep(2000 * time.Millisecond)
// 		var sum int32
// 		for i := 0; i < totalAccounts; i++ {
// 			locks[i].Lock()
// 		}
// 		for i := 0; i < totalAccounts; i++ {
// 			sum += ledger[i]
// 		}
// 		for i := 0; i < totalAccounts; i++ {
// 			locks[i].Unlock()
// 		}
// 		println(totalTrans, sum)
// 	}
// }

// TO DEBUG
// package main
//
// import (
// 	"math/rand"
// 	"sort"
// 	"sync"
// 	"sync/atomic"
// 	"time"
// )
//
// const (
// 	totalAccounts  = 50000
// 	maxAmountMoved = 10
// 	initialMoney   = 100
// 	threads        = 4
// )
//
// func perform_movements(ledger *[totalAccounts]int32, locks *[totalAccounts]sync.Locker, totalTrans *int64) {
// 	for {
// 		accountA := rand.Intn(totalAccounts)
// 		accountB := rand.Intn(totalAccounts)
// 		for accountA == accountB {
// 			accountB = rand.Intn(totalAccounts)
// 		}
//
// 		amountToMove := rand.Int31n(maxAmountMoved)
//
// 		// create an array of the account to lock for transaction
// 		// sort the array to avoid deadlocks
// 		toLock := []int{accountA, accountB}
// 		sort.Ints(toLock)
// 		// lock the account by hierarchy
// 		locks[toLock[0]].Lock()
// 		locks[toLock[1]].Lock()
//
// 		// do the transfert using the atomic package
// 		atomic.AddInt32(&ledger[accountA], -amountToMove)
// 		atomic.AddInt32(&ledger[accountB], amountToMove)
//
// 		// the order to unlock do not matter
// 		locks[toLock[1]].Lock()
// 		locks[toLock[0]].Lock()
// 	}
// }
//
// func main() {
// 	println("total accounts:", totalAccounts, " totalThreads:", threads, " using SpinLocks")
// 	var ledger [totalAccounts]int32
// 	// interface we have implemented in the other file
// 	var locks [totalAccounts]sync.Locker
// 	var totalTrans int64
//
// 	// set all the account with the initial money
// 	for i := 0; i < totalAccounts; i++ {
// 		ledger[i] = initialMoney
// 		// and initialise the lock(per account)
// 		locks[i] = NewSpinLock()
// 	}
//
// 	println("allo")
//
// 	// start our threads
// 	for i := 0; i < threads; i++ {
// 		go perform_movements(&ledger, &locks, &totalTrans)
// 	}
//
// 	println("allo1")
// 	// check if our ledger is in an consistent state or not
// 	for {
// 		println("al")
// 		// check the total value of our ledger every 2s
// 		time.Sleep(2000 * time.Millisecond)
// 		var sum int32
// 		// lock all the account the time to check the sum
// 		for i := 0; i < totalAccounts; i++ {
// 			locks[i].Lock()
// 		}
// 		println("allo2")
// 		// sum all the account
// 		for i := 0; i < totalAccounts; i++ {
// 			sum += ledger[i]
// 		}
// 		println("allo3")
// 		// unlock all the account
// 		for i := 0; i < totalAccounts; i++ {
// 			locks[i].Unlock()
// 		}
// 		println("allo4")
// 		println(totalTrans, sum)
// 	}
// }
