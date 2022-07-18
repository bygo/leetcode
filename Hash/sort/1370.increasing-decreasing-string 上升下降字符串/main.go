package main

// https://leetcode.cn/problems/increasing-decreasing-string

// ❓ 上升下降字符串

func sortString(s string) string {
	chMpCnt := [26]int{}
	for i := range s {
		ch := s[i] - 'a'
		chMpCnt[ch]++
	}
	var strBuf []byte
	sL := len(s)
	for len(strBuf) < sL {
		// 上升
		for ch := 0; ch <= 25; ch++ {
			if 0 < chMpCnt[ch] {
				strBuf = append(strBuf, byte(ch+'a'))
				chMpCnt[ch]--
			}
		}

		// 下降
		for ch := 25; 0 <= ch; ch-- {
			if 0 < chMpCnt[ch] {
				strBuf = append(strBuf, byte(ch+'a'))
				chMpCnt[ch]--
			}
		}
	}
	return string(strBuf)
}
