package main

// https://leetcode.cn/problems/number-of-different-integers-in-a-string

// ❓ 字符串中不同整数的数目
// ⚠️ 0 01 001

func numDifferentIntegers(word string) int {
	strMp := map[string]struct{}{}
	var numBuf []byte

	// 整数结算
	cntNum := func() {
		numBufL := len(numBuf)
		if 0 < numBufL {
			idx := 0
			// 保存至少一个零
			for numBuf[idx] == '0' && idx < numBufL-1 {
				idx++
			}
			str := string(numBuf[idx:])
			strMp[str] = struct{}{}
			numBuf = []byte{}
		}
	}

	for i := range word {
		if '0' <= word[i] && word[i] <= '9' {
			numBuf = append(numBuf, word[i])
		} else {
			cntNum()
		}
	}
	cntNum()
	return len(strMp)
}
