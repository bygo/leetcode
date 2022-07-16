package main

// https://leetcode-cn.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	strsL := len(strs)
	if strsL == 0 {
		return ""
	}

	check := func(lo, hi int) bool {
		for i := 1; i < strsL; i++ {
			if strs[i][lo:hi] != strs[0][lo:hi] {
				return false
			}
		}
		return true
	}
	minL := 1<<63 - 1
	for _, str := range strs {
		strL := len(str)
		if strL < minL {
			minL = strL
		}
	}

	lo, hi := 0, minL-1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if check(lo, mid+1) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return strs[0][:lo]
}
