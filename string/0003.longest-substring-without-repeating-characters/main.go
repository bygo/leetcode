package main

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

func lengthOfLongestSubstring1(s string) int {
	var left int
	var leftMap = map[rune]int{}
	var length int
	var currentLength int
	for right, v := range s {
		if leftMap[v] > left {
			left = leftMap[v]
		}
		leftMap[v] = right + 1
		currentLength = leftMap[v] - left
		if currentLength > length {
			length = currentLength
		}
	}
	return length
}

func lengthOfLongestSubstring2(s string) int {
	var left int
	var length int
	var currentLength int
	for right, v := range s {
		for i := left; i < right; i++ {
			if rune(s[i]) == v {
				left = i + 1
				break
			}
		}
		currentLength = right - left + 1
		if currentLength > length {
			length = currentLength
		}
	}
	return length
}

/**
思路1：
1.HashMap 存储 Index
2.尝试从 HashMap 获取 Index 设为Left

思路2：
1.Left 存储当前子串左边界
2.出现重复时，找出子串里重复值的Index，设为Left
 */
