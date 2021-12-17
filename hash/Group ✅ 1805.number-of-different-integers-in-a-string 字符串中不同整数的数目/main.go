package main

// https://leetcode-cn.com/problems/number-of-different-integers-in-a-string

// ❓ 字符串中不同整数的数目
// ⚠️ 0 01 001

func numDifferentIntegers(word string) int {
	strMp := map[string]struct{}{}
	var numBuf []byte

	// 整数结算
	cntNum := func() {
		tmpL := len(numBuf)
		if 0 < tmpL {
			j := 0
			// 保存至少一个零
			for numBuf[j] == '0' && j < tmpL-1 {
				j++
			}
			strMp[string(numBuf[j:])] = struct{}{}
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
