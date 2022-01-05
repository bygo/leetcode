package main

// https://leetcode-cn.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	lo, hi := 0, minLength
	for lo < hi {
		mid := (hi-lo+1)/2 + lo
		if isCommonPrefix(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return strs[0][:lo]
}
