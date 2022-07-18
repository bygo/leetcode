package main

import "sort"

// https://leetcode.cn/problems/hand-of-straights

// ❓ 一手顺子

func isNStraightHand(hand []int, groupSize int) bool {
	handL := len(hand)
	if handL%groupSize != 0 {
		return false
	}

	numMpCnt := map[int]int{}
	for i := 0; i < handL; i++ {
		num := hand[i]
		numMpCnt[num]++
	}

	sort.Ints(hand)
	for i := 0; i < handL; i++ {
		num := hand[i]
		if numMpCnt[num] == 0 {
			continue
		}
		for j := 0; j < groupSize; j++ {
			if numMpCnt[num] == 0 {
				return false
			}
			numMpCnt[num]--
			num++
		}
	}
	return true
}
