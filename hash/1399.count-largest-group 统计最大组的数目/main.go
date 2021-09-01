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
	cnt := 0
	max := 0
	for i := range m {
		if max < m[i] {
			max = m[i]
			cnt = 1
		} else if m[i] == max {
			cnt++
		}
	}
	return cnt
}
