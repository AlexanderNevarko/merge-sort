package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"./mergesort"
)

func main() {
	numProcPtr := flag.Int("np", runtime.NumCPU(), "number of workers")
	arrSizePtr := flag.Int("s", 100000, "size of an array to sort")
	flag.Parse()
	numProc := *numProcPtr
	arrSize := *arrSizePtr

	arr := rand.Perm(arrSize)

	start := time.Now()
	array := mergesort.Mergesort(arr, numProc)
	duration := time.Since(start)

	fmt.Printf("time spent for parallel mergesort: %s\n", duration)
	for i, value := range array {
		if i == len(array)-1 {
			break
		}
		if value > array[i+1] {
			fmt.Printf("Error in mergesort occured!\n")
		}
	}
}
