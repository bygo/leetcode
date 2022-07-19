package main

// https://leetcode.cn/problems/implement-strstr/

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := getNext(needle)
	var i, j int
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] { //对比
			j++
			i++
		} else {
			j = next[j] //跳
		}
	}

	if j == len(needle) {
		return i - j
	}
	return -1
}

func getNext(p string) []int {
	length := len(p)
	next := make([]int, length+1)
	next[0] = -1
	l := -1
	r := 0
	for r < length {
		if l == -1 || p[l] == p[r] {
			l++
			r++
			next[r] = l
		} else {
			l = next[l]
		}
	}
	return next
}
