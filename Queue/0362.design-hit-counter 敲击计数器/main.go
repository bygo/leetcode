package main

// https://leetcode-cn.com/problems/design-hit-counter

type HitCounter struct {
	m map[int]int
}

func Constructor() HitCounter {
	return HitCounter{m: map[int]int{}}
}

func (this *HitCounter) Hit(timestamp int) {
	this.m[timestamp]++
}

func (this *HitCounter) GetHits(timestamp int) int {
	start := timestamp - 299
	if start < 0 {
		start = 0
	}
	var res int
	for i := start; i <= timestamp; i++ {
		res += this.m[i]
	}
	return res
}
