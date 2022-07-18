package main

import "container/list"

// https://leetcode.cn/problems/design-hashmap

const base = 1024

type e struct {
	key, value int
}
type MyHashMap struct {
	data []list.List
}

func Constructor() MyHashMap {
	return MyHashMap{data: make([]list.List, base)}
}

func (m *MyHashMap) Put(key int, value int) {
	h := m.hash(key)
	for cur := m.data[h].Front(); cur != nil; cur = cur.Next() {
		if cur.Value.(*e).key == key {
			cur.Value = &e{key, value}
			return
		}
	}
	m.data[h].PushBack(&e{key, value})
}

func (m *MyHashMap) Get(key int) int {
	h := m.hash(key)
	for cur := m.data[h].Front(); cur != nil; cur = cur.Next() {
		if cur.Value.(*e).key == key {
			return cur.Value.(*e).value
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := m.hash(key)
	for cur := m.data[h].Front(); cur != nil; cur = cur.Next() {
		if cur.Value.(*e).key == key {
			m.data[h].Remove(cur)
		}
	}
}

func (m *MyHashMap) hash(key int) int {
	return key % base
}
