package main

// https://leetcode-cn.com/problems/h-index

func hIndex(citations []int) (h int) {
	l1 := len(citations)
	freq := make([]int, l1+1)
	for _, citation := range citations {
		if l1 <= citation {
			freq[l1]++
		} else {
			freq[citation]++
		}
	}
	var tot int
	for i := l1; 0 <= i; i-- {
		tot += freq[i]
		if i <= tot {
			return i
		}
	}
	return 0
}
