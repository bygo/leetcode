package main

// 稀疏数组搜索
func findString(words []string, s string) int {
	lo, hi := 0, len(words)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		for lo < mid && words[mid] == "" {
			mid--
		}
		str := words[mid]
		if str == s {
			return mid
		} else if str < s {
			lo = mid + 1
		} else if s < str {
			hi = mid
		}
	}
	return -1
}
