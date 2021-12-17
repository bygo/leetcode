package main

// https://leetcode-cn.com/problems/h-index

// ❓ h指数 为 篇数cnt 等于 至少引用次数quote

func hIndex(citations []int) (h int) {
	// 后缀和：至少引用次数val == 篇幅cnt
	quoteMax := len(citations)
	quoteMpCnt := make([]int, quoteMax+1)
	for i := range citations {
		if quoteMax <= citations[i] {
			quoteMpCnt[quoteMax] ++
		} else {
			quoteMpCnt[citations[i]]++
		}
	}
	var cnt int
	// i = 0 时，为一个都没被引用
	for quote := quoteMax; 0 < quote; quote-- {
		cnt += quoteMpCnt[quote]
		if quote <= cnt {
			return quote
		}
	}
	return -1
}
