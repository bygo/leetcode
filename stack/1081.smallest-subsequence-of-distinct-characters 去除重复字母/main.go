package main

// https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters

func smallestSubsequence(s string) string {
	var m = map[byte]int{}
	for i := range s {
		m[s[i]]++
	}

	var is = map[byte]bool{}
	var res = []byte{}
	for i := range s {
		if !is[s[i]] {
			top := len(res) - 1
			for -1 < top && s[i] < res[top] {
				if m[res[top]] == 0 {
					break
				}
				is[res[top]] = false
				res = res[:top]
				top--
			}
			res = append(res, s[i])
			is[s[i]] = true
		}
		m[s[i]]--
	}

	return string(res)
}
