package main

func lengthOfLongestSubstring(s string) int {
	var left int
	var leftMap = map[rune]int{}
	var length int
	var currentLength int
	for right, v := range s {
		if left < leftMap[v] {
			left = leftMap[v]
		}
		leftMap[v] = right + 1
		currentLength = leftMap[v] - left
		if length < currentLength {
			length = currentLength
		}
	}
	return length
}

func lengthOfLongestSubstring(s string) int {
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
		if length < currentLength {
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
