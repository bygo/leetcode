package main

// https://leetcode-cn.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	check := func(length int) bool {
		str0, strsL := strs[0][:length], len(strs)
		for i := 1; i < strsL; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}

	lo, hi := 0, minLength
	for lo < hi {
		mid := int(uint(lo+hi+1) >> 1) // 向上尝试
		if check(mid) {
			lo = mid // 答案
		} else {
			hi = mid - 1
		}
	}
	return strs[0][:lo]
}
