package main

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

// ❓ 无重复的最长子串

func lengthOfLongestSubstring(s string) int {
	var longestL, curL, left int
	var chMpIdx = map[byte]int{}
	for i := range s {
		ch := s[i]
		if left < chMpIdx[ch] {
			left = chMpIdx[ch] // 移动到 idx+1
		}
		curL = i - left + 1
		if longestL < curL {
			longestL = curL
		}
		chMpIdx[ch] = i + 1

		// 出现位置 + 1
		// A.否则 重复[a]bca 0 < 0 不成立
		// B.否则 curL = i - l + 1 不成立
		// C.也可 curL = chMpIdx[ch] + 1 ，但 A 不成立
	}
	return longestL
}
