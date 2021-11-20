package main

import "strings"

// https://leetcode-cn.com/problems/uncommon-words-from-two-sentences

func uncommonFromSentences(s1 string, s2 string) []string {
	a1 := strings.Split(s1, " ")
	a2 := strings.Split(s2, " ")

	m := map[string]int{}

	for i := range a1 {
		m[a1[i]]++
	}

	for i := range a2 {
		m[a2[i]]++
	}

	var res []string
	for k := range m {
		if m[k] == 1 {
			res = append(res, k)
		}
	}
	return res
}
