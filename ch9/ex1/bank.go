// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package bank

import "log"

type withdrawOp struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)           // send amount to deposit
var withdrawals = make(chan withdrawOp) // send amount to withdraw
var balances = make(chan int)           // receive balance

func Deposit(amount int) { deposits <- amount }

func Withdraw(amount int) bool {
	success := make(chan bool)
	op := withdrawOp{amount, success}
	withdrawals <- op

	return <-success
}

func Balance() int { return <-balances }

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case op := <-withdrawals:
			if op.amount > balance {
				log.Printf("bank: cannot withdraw %d from account (%d)", op.amount, balance)
				op.success <- false
				break
			}

			balance -= op.amount
			op.success <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
