package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// has i have 16 cores on this machine
const numberOfthreads int = 16

var (
	r         = regexp.MustCompile(`\((\d*),(\d*)\)`)
	waitGroup = sync.WaitGroup{}
)

type Point2D struct {
	x int
	y int
}

func findArea(inputChannel chan string) {
	// we exit this loop asa the channel is close
	for pointsStr := range inputChannel {
		var points []Point2D
		for _, p := range r.FindAllStringSubmatch(pointsStr, -1) {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])
			points = append(points, Point2D{x, y})
		}

		area := 0.0
		for i := 0; i < len(points); i++ {
			// le modulo de % len(points) return 0 at i + 1 == len(points)
			// so return to the points[i = 0] (first item of the array)
			a, b := points[i], points[(i+1)%len(points)]

			// calcul the sum of one way and substract to the other way
			// this is the algo to calcule the full surface
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}
		// final result
		fmt.Println(math.Abs(area / 2))
	}
	waitGroup.Done()
}

func main() {
	// absPath, _ := filepath.Abs("./polygones.txt")
	absPath, _ := filepath.Abs("./polygones.txt")
	dat, _ := ioutil.ReadFile(absPath)
	text := string(dat)

	// bufferChannel, this will queue all req into a buffer of 1000 entry
	// means that 1000 channels can queue up,
	// waiting for the worker thread to pick them up
	inputChannel := make(chan string, 1000)
	// pre-start our worker threads
	for i := 0; i < numberOfthreads; i++ {
		go findArea(inputChannel)
	}
	waitGroup.Add(numberOfthreads)

	start := time.Now()
	for _, line := range strings.Split(text, "\n") {
		// inputChannel is the equivalent of the master thread
		inputChannel <- line
	}
	close(inputChannel)
	// wait all worker thread to finish
	waitGroup.Wait()

	elapsed := time.Since(start)
	fmt.Printf("processing took: %v\n", elapsed)
}

// func findArea(pointStr string) {
// 	var points []Point2D
// 	for _, p := range r.FindAllStringSubmatch(pointStr, -1) {
// 		x, _ := strconv.Atoi(p[1])
// 		y, _ := strconv.Atoi(p[2])
// 		points = append(points, Point2D{x, y})
// 	}
//
// 	area := 0.0
// 	for i := 0; i < len(points); i++ {
// 		// le modulo de % len(points) return 0 at i + 1 == len(points)
// 		// so return to the points[i = 0] (first item of the array)
// 		a, b := points[i], points[(i+1)%len(points)]
//
// 		// calcul the sum of one way and substract to the other way
// 		// this is the algo to calcule the full surface
// 		area += float64(a.x*b.y) - float64(a.y*b.x)
// 	}
// 	// final result
// 	fmt.Println(math.Abs(area / 2))
//
// }

// func main() {
// 	// absPath, _ := filepath.Abs("./polygones.txt")
// 	absPath, _ := filepath.Abs("./polygones.txt")
// 	dat, _ := ioutil.ReadFile(absPath)
// 	text := string(dat)
//
// 	start := time.Now()
// 	for _, line := range strings.Split(text, "\n") {
// 		// line := "(4,10),(12,8),(10,3),(2,2),(7,5)"
// 		findArea(line)
// 	}
// 	elapsed := time.Since(start)
// 	fmt.Printf("processing took: %v\n", elapsed)
// }
