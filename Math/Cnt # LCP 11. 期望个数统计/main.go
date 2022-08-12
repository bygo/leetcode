package main

// https://leetcode.cn/problems/qi-wang-ge-shu-tong-ji/

func expectNumber(scores []int) int {
	m := map[int]struct{}{}
	for i := range scores {
		m[scores[i]] = struct{}{}
	}
	return len(m)
}
