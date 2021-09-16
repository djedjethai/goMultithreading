package hierarchy

import (
	"github.com/djedjethai/multithreading/deadlocks_train/common"
	"sort"
	"time"
)

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

	sort.Slice(intersectionsToLock, func(i, j int) bool {
		return intersectionsToLock[i].Id < intersectionsToLock[j].Id
	})

	// create a deadlock situation
	for _, it := range intersectionsToLock {
		it.Mutex.Lock()
		it.LockedBy = id
		// in this situation the deadlock situation will be difficult to simulate
		// let increase the risk by locking for longer time
		time.Sleep(10 * time.Millisecond)
	}
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
				crossing.Intersection.Mutex.Unlock()
				crossing.Intersection.LockedBy = -1
			}
		}
		// sleep 30millisecond, otherwise the simulation will run too fast
		time.Sleep(30 * time.Millisecond)
	}
}
