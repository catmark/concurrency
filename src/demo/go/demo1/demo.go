package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var JOBS = 100

func serial() time.Duration {
	start := time.Now()
	for i := 0; i < JOBS; i++ {
		run(i)
	}
	return time.Since(start)
}

func parallel() time.Duration {
	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < JOBS; i++ {
		wg.Add(1)
		go runParallel(i, &wg)
	}
	wg.Wait()
	return time.Since(start)
}

func runParallel(jobNo int, wg *sync.WaitGroup) {
	run(jobNo)
	wg.Done()
}

func run(jobNo int) {
	total := 0
	for i := 1; i < 100000000; i++ {
		total += i
	}
}

func main() {
	serialTime := serial()
	parallelTime := parallel()
	fmt.Println("Serial  ", serialTime)
	fmt.Println("Parallel", parallelTime)
	fmt.Printf("Using %d jobs with %v CPUs is %.2f time faster\n", JOBS, runtime.NumCPU(), float64(serialTime.Nanoseconds())/float64(parallelTime.Nanoseconds()))
}
