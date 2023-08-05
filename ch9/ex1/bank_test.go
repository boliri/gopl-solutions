// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	"bank"
)

// empties the bank account
func setup() {
	bank.Withdraw(bank.Balance())
}

func TestBankDeposit(t *testing.T) {
	setup()

	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestBankWithdrawOK(t *testing.T) {
	setup()

	success := make(chan bool)

	// Pre-load some funds
	bank.Deposit(1000)

	// Alice
	go func() {
		success <- bank.Withdraw(300)
	}()

	// Bob
	go func() {
		success <- bank.Withdraw(300)
	}()

	// Wait for both transactions.
	ok := true
	ok = ok && <-success
	ok = ok && <-success

	if !ok {
		t.Errorf("One or more withdrawals failed")
	}

	if got, want := bank.Balance(), 400; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestBankWithdrawInsufficientFunds(t *testing.T) {
	setup()

	success := make(chan bool)

	// Pre-load some funds
	bank.Deposit(400)

	// Alice
	go func() {
		success <- bank.Withdraw(300)
	}()

	// Bob
	go func() {
		success <- bank.Withdraw(300)
	}()

	// Wait for both transactions.
	ok := true
	ok = ok && <-success
	ok = ok && <-success

	if ok {
		t.Error("All withdrawals succeeded, current balance is below zero")
	}

	if got, want := bank.Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
