package main

// https://leetcode-cn.com/problems/h-index

func hIndex(citations []int) (h int) {
	l1 := len(citations)
	freq := make([]int, l1+1)
	for i := range citations {
		if l1 <= citations[i] { // 最多l1 篇
			freq[l1]++
		} else {
			freq[citations[i]]++
		}
	}
	var cnt int
	for i := l1; 0 <= i; i-- {
		cnt += freq[i]
		if i <= cnt {
			return i
		}
	}
	return 0
}
