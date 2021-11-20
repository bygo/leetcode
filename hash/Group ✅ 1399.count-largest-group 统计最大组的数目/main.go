package main

// https://leetcode-cn.com/problems/count-largest-group

func countLargestGroup(n int) int {
	m := map[int]int{}
	for i := 1; i <= n; i++ {
		j := i
		sum := 0
		for 0 < j {
			sum += j % 10
			j /= 10
		}
		m[sum] ++
	}
	var cur, cnt int
	for i := range m {
		if cur < m[i] {
			cur = m[i]
			cnt = 1
		} else if m[i] == cur {
			cnt++
		}
	}
	return cnt
}
