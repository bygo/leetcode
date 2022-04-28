package main

// https://leetcode-cn.com/problems/longest-substring-with-at-most-k-distinct-characters

// ❓ 至多包含 K 个不同字符的最长子串

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	chMpCnt := map[byte]int{}
	var curL, maxL, cntTyp, idxLeft int
	for idxRight := range s {
		chRight := s[idxRight]
		// 临界
		if chMpCnt[chRight] == 0 {
			cntTyp++
		}
		chMpCnt[chRight]++
		for k < cntTyp {
			chLeft := s[idxLeft]
			chMpCnt[chLeft]--
			// 临界
			if chMpCnt[chLeft] == 0 {
				cntTyp--
			}
			idxLeft++
		}
		curL = idxRight - idxLeft + 1
		if maxL < curL {
			maxL = curL
		}
	}
	return maxL
}
