package main

// https://leetcode-cn.com/problems/count-vowel-substrings-of-a-string

// ❓ 连续元音子串数

func countVowelSubstrings(word string) int {
	var res int
	var byteMpCnt map[byte]int

	wordL := len(word)
	for i := 0; i < wordL; i++ {
		byteMpCnt = map[byte]int{}
		for j := i; j < wordL && (word[j] == 'a' || word[j] == 'e' || word[j] == 'i' || word[j] == 'o' || word[j] == 'u'); j++ {
			byteMpCnt[word[j]]++
			if len(byteMpCnt) == 5 {
				res++
			}
		}
	}
	return res
}

func countVowelSubstringsBit(word string) int {
	var res int
	wordL := len(word)
	var state, success int32
	// 计算合法状态
	success |= 1 << int('a'-'a')
	success |= 1 << int('e'-'a')
	success |= 1 << int('i'-'a')
	success |= 1 << int('o'-'a')
	success |= 1 << int('u'-'a')

	for i := 0; i < wordL; i++ {
		// 每次清空状态
		state = 0
		for j := i; j < wordL; j++ {
			bit := word[j] - 'a'
			if success&1<<bit == 0 {
				// 非法状态
				break
			}
			// 合并状态
			state |= 1 << bit

			// 是否合法
			if state == success {
				res++
			}
		}
	}
	return res
}
