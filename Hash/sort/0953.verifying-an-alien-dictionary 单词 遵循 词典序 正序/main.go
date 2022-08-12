package main

// https://leetcode.cn/problems/verifying-an-alien-dictionary

// ❓ 单词 遵循 词典序 正序

func isAlienSorted(words []string, order string) bool {
	// 词典序
	chMpIdx := [26]int{}
	for i := range order {
		ch := order[i] - 'a'
		chMpIdx[ch] = i
	}

	var strPre string
	var strPreL int

	for _, strCur := range words {
		strCurL := len(strCur)
		idx := 0
		for idx < strPreL && idx < strCurL {
			chPre := strPre[idx] - 'a'
			chCur := strCur[idx] - 'a'
			idxPre := chMpIdx[chPre]
			idxCur := chMpIdx[chCur]
			if idxPre < idxCur {
				// 合法
				break
			} else if idxCur < idxPre {
				// 不合法
				return false
			} else if idxPre == idxCur {
				// 相等 对比下一个
				idx++
			}
		}
		if idx == strCurL && strCurL < strPreL {
			return false
		}
		strPre = strCur
		strPreL = strCurL
	}
	return true
}
