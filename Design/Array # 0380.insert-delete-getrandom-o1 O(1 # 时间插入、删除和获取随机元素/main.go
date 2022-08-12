package main

import "math/rand"

// https://leetcode.cn/problems/insert-delete-getrandom-o1

type RandomizedSet struct {
	m map[int]int
	a []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{m: map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	_, ok := rs.m[val]
	if ok {
		return false
	}

	rs.a = append(rs.a, val)
	rs.m[val] = len(rs.a) - 1
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	index, ok := rs.m[val]
	if !ok {
		return false
	}
	n := len(rs.a) - 1
	// 先改 可能删掉自己
	rs.m[rs.a[n]] = index
	rs.a[index], rs.a[n] = rs.a[n], rs.a[index]

	// 后删
	rs.a = rs.a[:n]
	delete(rs.m, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	n := rand.Intn(len(rs.a))
	return rs.a[n]
}
