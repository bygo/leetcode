package main

// https://leetcode-cn.com/problems/h-index

// ❓h指数 == (篇数cnt == 至少引用次数val)

func hIndex(citations []int) (h int) {
	// 后缀和：至少引用次数val == 篇幅cnt
	max := len(citations)
	quoteMpCnt := make([]int, max+1)
	for i := range citations {
		if max <= citations[i] {
			quoteMpCnt[max] ++
		} else {
			quoteMpCnt[citations[i]]++
		}
	}
	var cnt int
	// i = 0 时，为一个都没被引用
	for quote := max; 0 < quote; quote-- {
		cnt += quoteMpCnt[quote]
		if quote <= cnt {
			return quote
		}
	}
	return -1
}
