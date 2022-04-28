package main

// https://leetcode-cn.com/problems/h-index

// ❓ h指数 为 篇数cnt 等于 至少引用次数quote

func hIndex(citations []int) (h int) {
	quoteMax := len(citations)
	// 计数
	quoteMpCnt := make([]int, quoteMax+1)
	for _, quote := range citations {
		if quoteMax <= quote {
			quoteMpCnt[quoteMax] ++
		} else {
			quoteMpCnt[quote]++
		}
	}
	var cnt int

	// 从大到小，quote = 0 时，为一个都没被引用
	for quote := quoteMax; 0 <= quote; quote-- {
		cnt += quoteMpCnt[quote] // 大于等于 quote 的篇数
		if quote <= cnt {        // 后缀和 == 当前篇幅cnt
			return quote
		}
	}
	return -1
}
