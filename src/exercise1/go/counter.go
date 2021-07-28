package main

import (
	"fmt"
	"sync"
)

var counter = 0

func run(wg *sync.WaitGroup, rounds int) {
	defer wg.Done()

	for i := 0; i < rounds; i++ {
		counter++
	}
}

func main() {
	var wg sync.WaitGroup

    wg.Add(2)
    go run(&wg, 100000)
    go run(&wg, 100000)

	wg.Wait()
    fmt.Println(counter)
}
