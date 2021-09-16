package arbitrator

import (
	"github.com/djedjethai/multithreading/deadlocks_train/common"
	// "sort"
	"sync"
	"time"
)

var (
	controller = sync.Mutex{}
	cond       = sync.NewCond(&controller)
)

func allFree(intersectionsToLock []*common.Intersection) bool {
	for _, it := range intersectionsToLock {
		if it.LockedBy >= 0 {
			return false
		}
	}

	return true
}

// demonstrate locking 2 crossing at the time does not work
func lockIntersectionsInDistance(id, reserveStart, reserveEnd int, crossings []*common.Crossing) {
	var intersectionsToLock []*common.Intersection
	for _, crossing := range crossings {
		if reserveEnd >= crossing.Position &&
			reserveStart <= crossing.Position &&
			crossing.Intersection.LockedBy != id {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}

	// lock the mutex to make sure only one thread proceed
	controller.Lock()
	// infinite loop while allFree() return false (out of the loop asa return true)
	for !allFree(intersectionsToLock) {
		// lock the current thread
		cond.Wait()
	}

	for _, it := range intersectionsToLock {
		it.LockedBy = id
		time.Sleep(10 * time.Millisecond)

	}
	controller.Unlock()
}

func MoveTrain(train *common.Train, distance int, crossings []*common.Crossing) {
	for train.Front < distance {
		train.Front += 1
		for _, crossing := range crossings {
			// if train at crossing position try to acquire the lock
			if train.Front == crossing.Position {
				lockIntersectionsInDistance(train.Id, crossing.Position, crossing.Position+train.TrainLength, crossings)
				// crossing.Intersection.Mutex.Lock()
				// crossing.Intersection.LockedBy = train.Id
			}
			// unlock the mutex when the train passed the crossing
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				// as one cross-road is release
				// lock the controller for the condition to be re-evaluated
				controller.Lock()
				// release the locked intersection
				crossing.Intersection.LockedBy = -1
				// broadcast to inform waiting threads to wake up
				// as one intersection became available
				cond.Broadcast()
				// release the controller locked mutex
				// for waiting threads to be able to access the condition
				// in case of new Broadcast
				controller.Unlock()
			}
		}
		// sleep 30millisecond, otherwise the simulation will run too fast
		time.Sleep(30 * time.Millisecond)
	}
}
