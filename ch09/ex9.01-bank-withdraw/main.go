// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

type withdrawal struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)           // send amount to deposit
var balances = make(chan int)           // receive balance
var withdrawals = make(chan withdrawal) // send withdrawal

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	r := make(chan bool)
	w := withdrawal{amount, r}
	withdrawals <- w
	return <-r
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case w := <-withdrawals:
			if w.amount <= balance {
				balance -= w.amount
				w.success <- true
			} else {
				w.success <- false
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
