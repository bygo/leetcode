package main

// https://leetcode-cn.com/problems/single-row-keyboard

// ❓ 输入一个单词的移动长度总和
// ⚠️ 移动长度 = abs(i - j)

func calculateTime(keyboard string, word string) int {
	// 索引统计
	var valMpIdx = [26]int{}
	for i := range keyboard {
		valMpIdx[keyboard[i]-'a'] = i
	}

	// 计算移动次数
	// 0
	var res = valMpIdx[word[0]-'a']

	// 1～l
	var l = len(word)
	var i = 1
	for i < l {
		res += abs(valMpIdx[word[i-1]-'a'], valMpIdx[word[i]-'a'])
		i++
	}

	return res
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
