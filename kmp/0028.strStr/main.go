package main

/**
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

 */

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

/**
思路1:kmp

todo：听说sunday更稳定，二刷再来写sunday
 */
