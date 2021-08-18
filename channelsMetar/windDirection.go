package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"
)

func main() {
	absPath, _ := filepath.Abs("./multithreading/metarfiles")
	files, _ := ioutil.ReadDir(absPath)
	start := time.Now()

	for _, file := range files {
		dat, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			panic(err)
		}
		text := string(dat)

		// change to array, each metar report is a separate item  in the array
		metarsReports := parseToArray(text)

		// extract wind direction
		windsDirections := extractWindDirection(metarsReports)

		// assign to N, NE, E, ES, etc..
		mineWindDistribution(windsDirections)
	}

	elapsed := time.Since(start)
	fmt.Printf("%v\n", windDist)
	fmt.Printf("processing took: %s\n", elapsed)
}
