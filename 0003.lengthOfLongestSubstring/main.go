package main

import "leetcode"

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

 */

func lengthOfLongestSubstring(s string) int {
	var hashMap = map[rune]bool{}
	var length int
	var sor int
	for _, v := range s {
		for hashMap[v] {
			delete(hashMap, rune(s[sor]))
			sor++
		}
		hashMap[v] = true
		if len(hashMap) > length {
			length = len(hashMap)
		}
	}
	return length
}

func main() {
	leetcode.D(func() interface{} {
		return lengthOfLongestSubstring("aabaab!bb")
	})
}

/**
思路：
1.HashMap 存储当前子串
2.出现重复时，不断去除最左字符，直至HashMap不再重复
 */
