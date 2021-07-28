package main

import (
	"fmt"
	"sync"
	"time"
)

type Backend struct {
	db Database
}


func (b Backend) onAddClick(userId int, month time.Month) {
	valueOfSingleClick := 0.001
	b.db.addUserTransaction(userId, month, valueOfSingleClick)

	//recalculate balance
	balance := b.db.getUserBalance(userId)
	balance += valueOfSingleClick
	b.db.setUserBalance(userId, balance)
}

func (b Backend) onAddClickConcurrent(userId int, month time.Month, wg *sync.WaitGroup) {
	valueOfSingleClick := 0.001
	b.db.addUserTransaction(userId, month, valueOfSingleClick)

	//recalculate balance
	balance := b.db.getUserBalance(userId)
	balance += valueOfSingleClick
	b.db.setUserBalance(userId, balance)
	wg.Done()
}

func sum(transactions []float64) (total float64) {
	for _, transaction := range transactions {
		total += transaction
	}
	return
}

func (b Backend) cashOut(userId int, month time.Month) {
	balance := b.db.getUserBalance(userId)

	if balance > 0 {
		transactions := b.db.getUserTransactions(userId, month)
		amountToBePaid := sum(transactions)

		b.makePayment(month, userId, amountToBePaid)

		//recalculate balance
		balance := b.db.getUserBalance(userId)
		balance -= amountToBePaid
		b.db.setUserBalance(userId, balance)
	}
}

func (b Backend) makePayment(month time.Month, userId int, amount float64)  {
	//Some logic to do actual payment
	balance := b.db.getUserBalance(userId)
	fmt.Printf(
		"Payment made (%.3s): userId=%d, amount=%.3f, balance_before_payment=%.3f, balance_after_payment=%.3f\n",
		month, userId, amount, balance, balance - amount);
}
