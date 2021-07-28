package main

import "time"

type Database struct {
	transactions map[int]map[time.Month][]float64
	balances map[int]float64
}

func NewDatabase() Database {
	var db Database
	db.transactions = make(map[int]map[time.Month][]float64)
	db.transactions[1] = make(map[time.Month][]float64) //user 1
	db.transactions[2] = make(map[time.Month][]float64) //user 2
	db.balances = make(map[int]float64)
	return db
}

func (d Database) getUserBalance(userId int) float64 {
	balance :=  d.balances[userId]
	time.Sleep(100 * time.Millisecond)
	return balance
}

func (d Database) setUserBalance(userId int, balance float64) {
	d.balances[userId] = balance
}

func (d Database) addUserTransaction(userId int, month time.Month, transaction float64) {
	d.transactions[userId][month] = append(d.transactions[userId][month], transaction)
}

func (d Database) getUserTransactions(userId int, month time.Month) []float64 {
	return d.transactions[userId][month]
}

