package main

import (
	"time"
)

func main() {
	backend := Backend{NewDatabase()}
	userId := 1

	backend.onAddClick(userId, time.July)
	backend.onAddClick(userId, time.July)
	backend.onAddClick(userId, time.July)
	backend.onAddClick(userId, time.July)
	backend.onAddClick(userId, time.July)
	backend.onAddClick(userId, time.August)
	backend.onAddClick(userId, time.August)
	backend.onAddClick(userId, time.August)
	backend.onAddClick(userId, time.August)

	backend.cashOut(userId, time.July)
	backend.cashOut(userId, time.August)
}
