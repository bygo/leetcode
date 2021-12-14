package main

import "strconv"

// 求公牛母牛数
// 公牛: idx & val 交集
// 母牛: val 交集
// https://leetcode-cn.com/problems/bulls-and-cows

func getHint(secret string, guess string) string {
	m1 := [10]int{}
	m2 := [10]int{}
	var bull, cow int
	for i := range guess {
		if guess[i] == secret[i] {
			bull++
		} else {
			m1[secret[i]-'0']++
			m2[guess[i]-'0']++
		}
	}

	for i := 0; i < 10; i++ {
		cow += min(m1[i], m2[i])
	}
	return "A" + strconv.Itoa(bull) + "B" + strconv.Itoa(cow)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
