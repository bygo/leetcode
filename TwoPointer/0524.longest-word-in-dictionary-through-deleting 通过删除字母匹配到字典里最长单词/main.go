package main

// https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting

func findLongestWord(s string, dictionary []string) string {
	var res string
	for _, d := range dictionary {
		i := 0
		for j := range s {
			if d[i] == s[j] {
				i++
				if i == len(d) {
					if len(res) < i || len(res) == i && d < res {
						res = d
					}
					break
				}
			}
		}
	}
	return res
}
