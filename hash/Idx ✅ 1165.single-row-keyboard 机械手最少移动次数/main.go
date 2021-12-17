package main

// https://leetcode-cn.com/problems/single-row-keyboard

// ❓ 输入一个单词的移动长度总和
// ⚠️ 移动长度 = abs(i - j)

func calculateTime(keyboard string, word string) int {
	// 索引统计
	var chMpIdx = [26]int{}
	for idx := range keyboard {
		ch := keyboard[idx] - 'a'
		chMpIdx[ch] = idx
	}

	// 计算移动次数
	// 0
	ch0 := word[0] - 'a'
	var cntDist = chMpIdx[ch0]

	// 1～l
	for i := 1; i < len(word); i++ {
		chPrev := word[i-1] - 'a'
		chNext := word[i] - 'a'
		cntDist += abs(chMpIdx[chPrev] - chMpIdx[chNext])
		i++
	}

	return cntDist
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
