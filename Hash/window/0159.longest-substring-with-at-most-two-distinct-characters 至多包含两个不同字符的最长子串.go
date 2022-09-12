package main

// https://leetcode.cn/problems/longest-substring-with-at-most-two-distinct-characters

// ❓ 最多包含两个不同字符的最长子串

func lengthOfLongestSubstringTwoDistinct(s string) int {
	chMpCnt := [128]int{}
	var n = len(s)
	var cntLongest, cntTyp, left, right int
	var fn = func() {
		cntCur := right - left
		if cntLongest < cntCur {
			cntLongest = cntCur
		}
	}
	for right < n {
		chRight := s[right]
		if chMpCnt[chRight] == 0 {
			// 出现新的，提前结算
			fn()
			cntTyp++
			// 超过2种
			for 2 < cntTyp {
				chLeft := s[left]
				chMpCnt[chLeft]--
				if chMpCnt[chLeft] == 0 {
					cntTyp--
				}
				left++
			}
		}
		chMpCnt[chRight]++
		right++
	}

	// 最后结算
	fn()

	return cntLongest
}
