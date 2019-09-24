package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 은행계좌
type Account struct {
	balance int
	mutext  *sync.Mutex
}

// 출금
func (a *Account) Widthdraw(val int) {
	a.mutext.Lock()
	a.balance -= val
	a.mutext.Unlock()
}

// 입금
func (a *Account) Deposit(val int) {
	a.mutext.Lock()
	a.balance += val
	a.mutext.Unlock()
}

// 잔액조회
func (a *Account) Balance() int {
	a.mutext.Lock()
	balance := a.balance
	a.mutext.Unlock()
	return balance
}

var accounts []*Account
var globalLock *sync.Mutex

func Transfer(sender, receiver int, money int) {
	globalLock.Lock()
	accounts[sender].Widthdraw(money)
	accounts[receiver].Deposit(money)
	globalLock.Unlock()
}

// 총 계좌 잔액 조회
func GetTotalBalance() int {
	globalLock.Lock()
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	globalLock.Unlock()
	return total
}

// 랜덤하게 송금
func RandomTransfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts))
		if sender != receiver {
			break
		}
	}
	money := rand.Intn(balance)
	Transfer(sender, receiver, money)
}

func GoTransfer() {
	for {
		RandomTransfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total %d \n", GetTotalBalance())
}

func main() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000, mutext: &sync.Mutex{}})
	}
	globalLock = &sync.Mutex{}

	PrintTotalBalance()

	for i := 0; i < 10; i++ {
		go GoTransfer()
	}

	for {
		PrintTotalBalance()
		time.Sleep(100 * time.Millisecond)
	}
}
