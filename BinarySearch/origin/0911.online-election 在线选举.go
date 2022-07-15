package main

// https://leetcode-cn.com/problems/online-election

type TopVotedCandidate struct {
	tops  []int
	times []int
	tL    int
}

func (c *TopVotedCandidate) Q(t int) int {
	lo, hi := 0, c.tL
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if c.times[mid] <= t {
			lo = mid + 1
		} else if t < c.times[mid] {
			hi = mid
		}
	}
	return c.tops[lo-1]
	//return c.tops[sort.SearchInts(c.times, t+1)-1]
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	tL := len(times)
	var tops []int
	top := -1
	pMpCnt := map[int]int{-1: 01}
	for _, p := range persons {
		pMpCnt[p]++
		if pMpCnt[top] <= pMpCnt[p] {
			top = p
		}
		tops = append(tops, top)
	}
	return TopVotedCandidate{tops, times, tL}
}
