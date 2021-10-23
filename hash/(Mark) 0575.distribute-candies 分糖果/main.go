package main

// https://leetcode-cn.com/problems/distribute-candies

func distributeCandies(candyType []int) int {
	l1 := len(candyType)
	m := map[int]struct{}{}
	for i := range candyType {
		m[candyType[i]] = struct{}{} // 类型数
	}

	l2 := len(m)
	if l1/2 < l2 { // 类型数超过 一半，每样一颗
		return l1 / 2
	}
	return l2 // 小于一半 最多l2
}
