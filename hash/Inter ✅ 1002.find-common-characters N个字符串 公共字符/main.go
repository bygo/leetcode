package main

// https://leetcode-cn.com/problems/find-common-characters

// ❓ N个字符串 公共字符

func commonChars(words []string) []string {
	// 占位
	var chsCommon []string
	chMpCnt := [26]int{}
	for i := range chMpCnt {
		chMpCnt[i] = 1<<63 - 1
	}

	for _, word := range words {
		// 计数
		chMpCntTmp := [26]int{}
		for _, ch := range word {
			chMpCntTmp[ch-'a']++
		}
		// 计算最小值
		for ch, cntTmp := range chMpCntTmp {
			if cntTmp < chMpCnt[ch] {
				chMpCnt[ch] = cntTmp
			}
		}
	}

	// 字典序
	for i := 0; i < 26; i++ {
		for j := 0; j < chMpCnt[i]; j++ {
			chsCommon = append(chsCommon, string(byte(i+'a')))
		}
	}
	return chsCommon
}
