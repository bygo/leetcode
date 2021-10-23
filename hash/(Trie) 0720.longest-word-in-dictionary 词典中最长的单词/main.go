package main

import "sort"

// https://leetcode-cn.com/problems/longest-word-in-dictionary

func longestWord(words []string) string {
	sort.Strings(words)
	m := map[string]struct{}{}

	var cnt int
	var res string
	for i := range words {
		l1 := len(words[i])
		if len(words[i]) == 1 {
			m[words[i]] = struct{}{}
			if cnt < l1 {
				cnt = l1
				res = words[i]
			}
		} else {
			_, ok := m[words[i][:l1-1]]
			if ok {
				m[words[i]] = struct{}{}
				if cnt < l1 {
					cnt = l1
					res = words[i]
				}
			}
		}

	}
	return res
}