package main

// https://leetcode-cn.com/problems/maximum-number-of-words-you-can-type

// ❓ 可以输入的最大单词数
// ⚠️ hello word , w

func canBeTypedWords(text string, brokenLetters string) int {
	// 损坏键 计数
	var chMpDis = map[byte]bool{}
	for i := range brokenLetters {
		chMpDis[brokenLetters[i]] = true
	}
	var res int

	var idx int
	var l = len(text)
	for idx < l {
		// 损坏键
		for idx < l && text[idx] != ' ' && !chMpDis[text[idx]] {
			idx++
		}

		if idx == l || text[idx] == ' ' {
			res++
		} else {
			for idx < l && text[idx] != ' ' {
				idx++
			}
		}
		idx++
	}
	return res
}
