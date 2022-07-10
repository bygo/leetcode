package main

import "sort"

// https://leetcode-cn.com/problems/online-election

type TopVotedCandidate struct {
	personTops []int
	times      []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	var personTops []int
	personTop := -1                    // 领先的
	personMpCnt := map[int]int{-1: -1} // 统计
	for _, person := range persons {
		personMpCnt[person]++                              // 选票
		if personMpCnt[personTop] <= personMpCnt[person] { // 是否领先，改变领先者
			personTop = person
		}
		personTops = append(personTops, personTop)
	}

	return TopVotedCandidate{personTops, times}
}

func (c *TopVotedCandidate) Q(t int) int {
	return c.personTops[sort.SearchInts(c.times, t+1)-1]
}
