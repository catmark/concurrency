package main

import (
	"sync"
	"time"
)

func main() {
	backend := Backend{NewDatabase()}
	userId := 1
	var wg sync.WaitGroup
	wg.Add(9)

	go backend.onAddClickConcurrent(userId, time.July, &wg)
	time.Sleep(30 * time.Millisecond)
	go backend.onAddClickConcurrent(userId, time.July, &wg)
	time.Sleep(30 * time.Millisecond)
	go backend.onAddClickConcurrent(userId, time.July, &wg)
	time.Sleep(30 * time.Millisecond)
	go backend.onAddClickConcurrent(userId, time.July, &wg)
	time.Sleep(30 * time.Millisecond)
	go backend.onAddClickConcurrent(userId, time.August, &wg)
	time.Sleep(30 * time.Millisecond)
	go backend.onAddClickConcurrent(userId, time.August, &wg)
	time.Sleep(30 * time.Millisecond)
	go backend.onAddClickConcurrent(userId, time.August, &wg)
	time.Sleep(30 * time.Millisecond)
	go backend.onAddClickConcurrent(userId, time.August, &wg)
	time.Sleep(30 * time.Millisecond)
	go backend.onAddClickConcurrent(userId, time.August, &wg)

	wg.Wait()

	backend.cashOut(userId, time.July)
	backend.cashOut(userId, time.August)
}
