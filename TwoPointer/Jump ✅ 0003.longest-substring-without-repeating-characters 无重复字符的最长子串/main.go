package main

// https://leetcode.cn/problems/longest-substring-without-repeating-characters

// :34
func lengthOfLongestSubstring(s string) int {
	var idxLeft, curL, maxL int
	for idxRight := range s {
		for i := idxLeft; i < idxRight; i++ {
			if s[idxRight] == s[i] {
				// 出现相等 跳到下一个
				idxLeft = i + 1
				break
			}
		}
		curL = idxRight - idxLeft + 1
		if maxL < curL {
			maxL = curL
		}
	}
	return maxL
}
