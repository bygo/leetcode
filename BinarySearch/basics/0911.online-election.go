package main

// https://leetcode.cn/problems/online-election

type TopVotedCandidate struct {
	tops  []int
	times []int
	tL    int
}

func (c *TopVotedCandidate) Q(t int) int { // TODO
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
	personLead := -1
	personMpCnt := map[int]int{personLead: 1}
	for _, person := range persons {
		personMpCnt[person]++
		if personMpCnt[personLead] <= personMpCnt[person] {
			personLead = person
		}
		tops = append(tops, personLead)
	}
	return TopVotedCandidate{tops, times, tL}
}
