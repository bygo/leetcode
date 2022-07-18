package main

// https://leetcode.cn/problems/shortest-completing-word

// ❓ 最短补全词
// ⚠️ step 能覆盖 s1T2p 所有字母 stp
// ⚠️ 不区分大小写

func shortestCompletingWord(licensePlate string, words []string) string {
	chMpCnt := [26]int{}
	// 计算次数
	for i := range licensePlate {
		if 'A' <= licensePlate[i] && licensePlate[i] <= 'Z' {
			chMpCnt[licensePlate[i]-'A']++
		} else if 'a' <= licensePlate[i] && licensePlate[i] <= 'z' {
			chMpCnt[licensePlate[i]-'a']++
		}
	}

	var strShortest string
	var strShortestL = 1<<63 - 1
	var ch int
	for i := range words {
		// 计算当前次数
		chMpCntCur := [26]int{}
		for j := range words[i] {
			chMpCntCur[words[i][j]-'a']++
		}
		for ch = 0; ch < 26; ch++ {
			if chMpCntCur[ch] < chMpCnt[ch] {
				// 少于就不合法
				break
			}
		}

		if ch == 26 {
			// 合法判断
			strCurL := len(words[i])
			if strCurL < strShortestL {
				strShortest = words[i]
				strShortestL = len(words[i])
			}
		}
	}

	return strShortest
}
