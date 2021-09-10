package main

import "math/rand"

// https://leetcode-cn.com/problems/insert-delete-getrandom-o1

type RandomizedSet struct {
	m map[int]int
	a []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{m: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.m[val]
	if ok {
		return false
	}

	this.a = append(this.a, val)
	this.m[val] = len(this.a) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.m[val]
	if !ok {
		return false
	}
	n := len(this.a) - 1
	// 先改 可能删掉自己
	this.m[this.a[n]] = index
	this.a[index], this.a[n] = this.a[n], this.a[index]

	// 后删
	this.a = this.a[:n]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	n := rand.Intn(len(this.a))
	return this.a[n]
}
