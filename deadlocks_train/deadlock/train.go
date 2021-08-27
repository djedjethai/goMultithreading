package deadlock

import (
	"github.com/djedjethai/multithreading/deadlocks_train/common"
	"time"
)

func MoveTrain(train *common.Train, distance int, crossings []*common.Crossing) {
	for train.Front < distance {
		train.Front += 1
		for _, crossing := range crossings {
			// if train at crossing position try to acquire the lock
			if train.Front == crossing.Position {
				crossing.Intersection.Mutex.Lock()
				crossing.Intersection.LockedBy = train.Id
			}
			// unlock the mutex when the train passed the crossing
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				crossing.Intersection.LockedBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
		}
		// sleep 30millisecond, otherwise the simulation will run too fast
		time.Sleep(30 * time.Millisecond)
	}
}
