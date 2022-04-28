package main

// https://leetcode-cn.com/problems/maximum-number-of-words-you-can-type

// ❓ 可以输入的最大单词数
// ⚠️ hello word , w

func canBeTypedWords(text string, brokenLetters string) int {
	// 损坏键 计数
	var chMpDis = map[byte]bool{}
	for i := range brokenLetters {
		ch := brokenLetters[i]
		chMpDis[ch] = true
	}
	var cntInput int

	var idx int
	var textL = len(text)
	for idx < textL {
		// 损坏键
		for idx < textL && text[idx] != ' ' && !chMpDis[text[idx]] {
			idx++
		}

		// 最后 或者 ' '
		if idx == textL || text[idx] == ' ' {
			cntInput++
		} else {
			// 出现损坏，移动到下一个单词
			for idx < textL && text[idx] != ' ' {
				idx++
			}
		}
		idx++
	}
	return cntInput
}
