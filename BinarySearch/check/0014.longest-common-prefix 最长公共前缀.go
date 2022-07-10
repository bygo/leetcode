package main

// https://leetcode-cn.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	check := func(length int) bool {
		str0, strL := strs[0][:length], len(strs)
		for i := 1; i < strL; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}

	minLength := len(strs[0])
	for _, str := range strs {
		strL := len(str)
		if strL < minLength {
			minLength = strL
		}
	}

	lo, hi := 0, minLength+1
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return strs[0][:lo-1]
}
