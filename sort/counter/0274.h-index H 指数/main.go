package main

// https://leetcode-cn.com/problems/h-index

func hIndex(citations []int) (h int) {
	l1 := len(citations)
	counter := make([]int, l1+1)
	for _, citation := range citations {
		if l1 <= citation {
			counter[l1]++
		} else {
			counter[citation]++
		}
	}
	var total int
	for i := l1; 0 <= i; i-- {
		total += counter[i]
		if i <= total {
			return i
		}
	}
	return 0
}
