package main

import "strings"

// https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters

// ❓ 每一字符最少重复K次的最长子串 长度

func longestSubstring(str string, k int) int {
	var cntLongest int
	// 种类数 1～26

	for typ := 1; typ <= 26; typ++ {
		chMpCnt := [26]int{}
		cntTyp := 0      // 种类数
		cntTypValid := k // 默认全部合法
		left := 0
		for right := range str {
			// 每次加入一个字符
			chRight := str[right] - 'a'
			if chMpCnt[chRight] == 0 {
				// 加入了 新的 ch 种类
				cntTyp++
				cntTypValid--
			}
			chMpCnt[chRight]++
			if chMpCnt[chRight] == k {
				// 加入的ch种类 达到合法
				cntTypValid++
			}

			// 种类数超过当前搜索总数，移动 left 使之合法
			for typ < cntTyp {
				chLeft := str[left] - 'a'
				if chMpCnt[chLeft] == k {
					// 删除的种类 达到不合法
					cntTypValid--
				}
				chMpCnt[chLeft]--
				if chMpCnt[chLeft] == 0 {
					// 删除 ch 种类
					cntTyp--
					cntTypValid++
				}
				left++
			}

			// 种类等于 typ 且 都合法
			if cntTypValid == k {
				cntCur := right - left + 1
				if cntLongest < cntCur {
					cntLongest = cntCur
				}
			}
		}
	}
	return cntLongest
}

func longestSubstringDivide(str string, k int) int {
	var cntLongest int
	if str == "" {
		return cntLongest
	}

	// 统计
	chMpCnt := [26]int{}
	for i := range str {
		ch := str[i] - 'a'
		chMpCnt[ch]++
	}

	// 分割点
	var chSplit byte
	for ch, cnt := range chMpCnt {
		if 0 < cnt && cnt < k {
			chSplit = byte(ch + 'a')
			break
		}
	}

	// 全部合法
	if chSplit == 0 {
		return len(str)
	}

	// 分割
	for _, strSub := range strings.Split(str, string(chSplit)) {
		cntCur := longestSubstringDivide(strSub, k)
		if cntLongest < cntCur {
			cntLongest = cntCur
		}
	}
	return cntLongest
}
