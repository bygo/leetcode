package main

import "strconv"

// https://leetcode-cn.com/problems/bulls-and-cows

func getHint(secret string, guess string) string {
	m1 := [10]int{}
	m2 := [10]int{}

	var res []byte
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

	res = append(res, strconv.Itoa(bull)...)
	res = append(res, 'A')
	res = append(res, strconv.Itoa(cow)...)
	res = append(res, 'B')
	return string(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
