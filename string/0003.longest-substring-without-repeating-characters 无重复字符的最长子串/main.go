package main

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	var l int
	var lM = map[rune]int{}
	var res int
	var cur int
	for r, v := range s {
		if l < lM[v] {
			l = lM[v]
		}
		lM[v] = r + 1
		cur = lM[v] - l
		if res < cur {
			res = cur
		}
	}
	return res
}

func lengthOfLongestSubstring(s string) int {
	var l int
	var res int
	var cur int
	for r, v := range s {
		for i := l; i < r; i++ {
			if rune(s[i]) == v {
				l = i + 1
				break
			}
		}
		cur = r - l + 1
		if res < cur {
			res = cur
		}
	}
	return res
}

/**
思路1：
1.HashMap 存储 Index
2.尝试从 HashMap 获取 Index 设为Left

思路2：
1.Left 存储当前子串左边界
2.出现重复时，找出子串里重复值的Index，设为Left
*/
