package main

// https://leetcode-cn.com/problems/single-row-keyboard

// ❓ 输入一个单词的移动长度总和
// ⚠️ 移动长度 = abs(i - j)

func calculateTime(keyboard string, word string) int {
	// 索引统计
	var chMpIdx = [26]int{}
	for idx := range keyboard {
		chMpIdx[keyboard[idx]-'a'] = idx
	}

	// 计算移动次数
	// 0
	var res = chMpIdx[word[0]-'a']

	// 1～l
	var l = len(word)
	var i = 1
	for i < l {
		chPrev := word[i-1] - 'a'
		chNext := word[i] - 'a'
		res += abs(chMpIdx[chPrev] - chMpIdx[chNext])
		i++
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
