package main

import "sort"

// https://leetcode-cn.com/problems/online-election

type TopVotedCandidate struct {
	tops, times []int
}

func Constructor(persons, times []int) TopVotedCandidate {
	var tops []int
	top := -1                  // 领先的
	cnt := map[int]int{-1: -1} // 统计
	for _, p := range persons {
		cnt[p]++                // 选票
		if cnt[top] <= cnt[p] { //
			top = p
		}
		tops = append(tops, top)
	}

	return TopVotedCandidate{tops, times}
}

func (c *TopVotedCandidate) Q(t int) int {
	return c.tops[sort.SearchInts(c.times, t+1)-1] // t+1 = right
}
