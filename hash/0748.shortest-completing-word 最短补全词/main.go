package main

// https://leetcode-cn.com/problems/shortest-completing-word

func shortestCompletingWord(licensePlate string, words []string) string {
	m := [26]int{}
	for i := range licensePlate {
		if 'A' <= licensePlate[i] && licensePlate[i] <= 'Z' {
			m[licensePlate[i]-'A'] ++
		} else if 'a' <= licensePlate[i] && licensePlate[i] <= 'z' {
			m[licensePlate[i]-'a'] ++
		}
	}

	var res string
	var min = 1<<63 - 1
	var k int
	for i := range words {
		cur := [26]int{}
		for j := range words[i] {
			cur[words[i][j]-'a']++
		}
		for k = 0; k < 26; k++ {
			if cur[k] < m[k] {
				break
			}
		}
		if k == 26 {
			l1 := len(words[i])
			if l1 < min {
				res = words[i]
				min = len(words[i])
			}
		}
	}

	return res
}
