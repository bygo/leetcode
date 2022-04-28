package main

// https://leetcode-cn.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(lo int, hi int) string
	lcp = func(lo int, hi int) string {
		if lo == hi {
			return strs[lo]
		}
		mid := int(uint(lo+hi) >> 1)
		strLeft, strRight := lcp(lo, mid), lcp(mid+1, hi)
		minL := min(len(strLeft), len(strRight))
		for i := 0; i < minL; i++ {
			if strLeft[i] != strRight[i] {
				return strLeft[:i]
			}
		}
		return strLeft[:minL]
	}
	return lcp(0, len(strs)-1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
