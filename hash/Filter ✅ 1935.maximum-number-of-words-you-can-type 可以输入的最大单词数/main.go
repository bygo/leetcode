package main

// https://leetcode-cn.com/problems/maximum-number-of-words-you-can-type

// ❓可以输入的最大单词数

func canBeTypedWords(text string, brokenLetters string) int {
	// 损坏键统计
	var broMpDis = map[byte]bool{}
	for i := range brokenLetters {
		broMpDis[brokenLetters[i]] = true
	}
	var res int
	var normal bool

	var i int
	var l = len(text)
	for i < l {
		// 单词组成
		normal = true
		for i < l && text[i] != ' ' {
			// 损坏键
			if broMpDis[text[i]] {
				normal = false
				// 走到下一个单词
				for i < l && text[i] != ' ' {
					i++
				}
			} else {
				i++
			}
		}

		if normal {
			res++
		}
		i++
	}
	return res
}
