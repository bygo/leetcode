package main

// https://leetcode.cn/problems/sparse-array-search-lcci/

func findString(words []string, str string) int {
	lo, hi := 0, len(words)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		for lo < mid && words[mid] == "" { // Move `mid` to the far left
			mid--
		}

		if words[mid] == str {
			return mid
		} else if words[mid] < str { // Hidden case `words[mid] == ""`
			lo = mid + 1
		} else if str < words[mid] {
			hi = mid
		}
	}
	return -1
}

func findString(words []string, str string) int {
	lo, hi := -1, len(words)-1
	for lo < hi {
		mid := int(uint(lo+hi+1) >> 1)
		for mid < hi && words[mid] == "" { // Move `mid` to the far right
			mid++
		}

		if words[mid] == str {
			return mid
		} else if str < words[mid] || words[mid] == "" {
			hi = mid - 1
		} else if words[mid] < str {
			lo = mid
		}
	}
	return -1
}
