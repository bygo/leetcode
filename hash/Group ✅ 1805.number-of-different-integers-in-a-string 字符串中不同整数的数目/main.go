package main

// https://leetcode-cn.com/problems/number-of-different-integers-in-a-string

// ❓ 字符串中不同整数的数目
// ⚠️ 0 01 001

func numDifferentIntegers(word string) int {
	strMp := map[string]struct{}{}
	var tmp []byte
	fn := func() {
		tmpL := len(tmp)
		if 0 < tmpL {
			j := 0
			for tmp[j] == '0' && j < tmpL-1 {
				j++
			}
			strMp[string(tmp[j:])] = struct{}{}
			tmp = []byte{}
		}
	}

	for i := range word {
		if '0' <= word[i] && word[i] <= '9' {
			tmp = append(tmp, word[i])
		} else {
			fn()
		}
	}
	fn()
	return len(strMp)
}
