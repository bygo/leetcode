package main

// https://leetcode-cn.com/problems/distribute-candies

func distributeCandies(candyType []int) int {
	l1 := len(candyType)
	m := map[int]int{}
	for i := range candyType {
		m[candyType[i]]++
	}

	l2 := len(m)
	if l1/2 < l2 {
		return l1 / 2
	}
	return l2
}
