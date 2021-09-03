package main

import (
	"github.com/djedjethai/multithreading/deadlocks_train/common"
	// "github.com/djedjethai/multithreading/deadlocks_train/deadlock"
	"github.com/djedjethai/multithreading/deadlocks_train/hierarchy"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"sync"
)

var (
	trains        [4]*common.Train
	intersections [4]*common.Intersection
)

const trainLength = 70

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawTracks(screen)
	DrawIntersections(screen)
	DrawTrains(screen)
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return 320, 320
}

func main() {

	// at first set each train
	for i := 0; i < 4; i++ {
		trains[i] = &common.Train{Id: i, TrainLength: trainLength, Front: 0}
	}

	// create the 4 intersections
	for i := 0; i < 4; i++ {
		intersections[i] = &common.Intersection{Id: i, Mutex: sync.Mutex{}, LockedBy: -1}
	}

	// start the threads(start moving the trains)
	go hierarchy.MoveTrain(trains[0], 300, []*common.Crossing{{Position: 125, Intersection: intersections[0]}, {Position: 175, Intersection: intersections[1]}})

	go hierarchy.MoveTrain(trains[1], 300, []*common.Crossing{{Position: 125, Intersection: intersections[1]}, {Position: 175, Intersection: intersections[2]}})

	go hierarchy.MoveTrain(trains[2], 300, []*common.Crossing{{Position: 125, Intersection: intersections[2]}, {Position: 175, Intersection: intersections[3]}})

	go hierarchy.MoveTrain(trains[3], 300, []*common.Crossing{{Position: 125, Intersection: intersections[3]}, {Position: 175, Intersection: intersections[0]}})

	ebiten.SetWindowSize(320*3, 320*3)
	ebiten.SetWindowTitle("Trains in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

// deadlock situation
// package main
//
// import (
// 	"github.com/djedjethai/multithreading/deadlocks_train/common"
// 	"github.com/djedjethai/multithreading/deadlocks_train/deadlock"
// 	"github.com/hajimehoshi/ebiten/v2"
// 	"log"
// 	"sync"
// )
//
// var (
// 	trains        [4]*common.Train
// 	intersections [4]*common.Intersection
// )
//
// const trainLength = 70
//
// type Game struct{}
//
// func (g *Game) Update() error {
// 	return nil
// }
//
// func (g *Game) Draw(screen *ebiten.Image) {
// 	DrawTracks(screen)
// 	DrawIntersections(screen)
// 	DrawTrains(screen)
// }
//
// func (g *Game) Layout(_, _ int) (w, h int) {
// 	return 320, 320
// }
//
// func main() {
//
// 	// at first set each train
// 	for i := 0; i < 4; i++ {
// 		trains[i] = &common.Train{Id: i, TrainLength: trainLength, Front: 0}
// 	}
//
// 	// create the 4 intersections
// 	for i := 0; i < 4; i++ {
// 		intersections[i] = &common.Intersection{Id: i, Mutex: sync.Mutex{}, LockedBy: -1}
// 	}
//
// 	// start the threads(start moving the trains)
// 	go deadlock.MoveTrain(trains[0], 300, []*common.Crossing{{Position: 125, Intersection: intersections[0]}, {Position: 175, Intersection: intersections[1]}})
//
// 	go deadlock.MoveTrain(trains[1], 300, []*common.Crossing{{Position: 125, Intersection: intersections[1]}, {Position: 175, Intersection: intersections[2]}})
//
// 	go deadlock.MoveTrain(trains[2], 300, []*common.Crossing{{Position: 125, Intersection: intersections[2]}, {Position: 175, Intersection: intersections[3]}})
//
// 	go deadlock.MoveTrain(trains[3], 300, []*common.Crossing{{Position: 125, Intersection: intersections[3]}, {Position: 175, Intersection: intersections[0]}})
//
// 	ebiten.SetWindowSize(320*3, 320*3)
// 	ebiten.SetWindowTitle("Trains in a box")
// 	if err := ebiten.RunGame(&Game{}); err != nil {
// 		log.Fatal(err)
// 	}
// }
