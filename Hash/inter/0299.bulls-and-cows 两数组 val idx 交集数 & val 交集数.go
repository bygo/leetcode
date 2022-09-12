package main

import "strconv"

// https://leetcode.cn/problems/bulls-and-cows

// ❓ 公牛母牛数
// 公牛: idx & val 交集
// 母牛: val 交集

func getHint(secret string, guess string) string {
	secretMpCnt := [10]int{}
	guessMpCNt := [10]int{}
	var bull, cow int
	for i := range guess {
		if guess[i] == secret[i] {
			bull++
		} else {
			secretMpCnt[secret[i]-'0']++
			guessMpCNt[guess[i]-'0']++
		}
	}

	for num := 0; num < 10; num++ {
		cow += min(secretMpCnt[num], guessMpCNt[num])
	}
	return "A" + strconv.Itoa(bull) + "B" + strconv.Itoa(cow)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
