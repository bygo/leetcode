package main

// https://leetcode.cn/problems/simple-bank-system

type Bank struct {
	balance []int64
	maxId   int
}

func Constructor(balance []int64) Bank {
	return Bank{balance: balance, maxId: len(balance) - 1}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	id1 := account1 - 1
	id2 := account2 - 1
	if !this.valid(id2) || !this.money(id1, money) {
		return false
	}
	this.balance[id1] -= money
	this.balance[id2] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	id := account - 1
	if !this.valid(id) {
		return false
	}
	this.balance[id] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	id := account - 1
	if !this.money(id, money) {
		return false
	}
	this.balance[id] -= money
	return true
}

func (this *Bank) valid(id int) bool {
	return 0 <= id && id <= this.maxId
}

func (this *Bank) money(id int, money int64) bool {
	return this.valid(id) && money <= this.balance[id]
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
