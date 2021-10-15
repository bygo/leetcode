package main

func findString(words []string, s string) int {
	l, r := 0, len(words)
	for l < r {
		mid := l + (r-l)/2
		for l < mid && words[mid] == "" {
			mid--
		}
		cur := words[mid]
		if cur == s {
			return mid
		} else if cur < s {
			l = mid + 1
		} else if s < cur {
			r = mid
		}
	}
	return -1
}
