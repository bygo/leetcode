package main

// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string

func findAnagrams(s string, p string) []int {
	l1 := len(s)
	l2 := len(p)
	l := 0
	r := l2 - 1
	var res []int
	if l1 < l2 {
		return res
	}
	m1 := [26]int{}
	for i := range p {
		m1[p[i]-'a']++
	}

	m2 := [26]int{}
	for l < r {
		m2[s[l]-'a']++
		l++
	}
	l = 0
	for r < l1 {
		m2[s[r]-'a']++
		if m1 == m2 {
			res = append(res, l)
		}
		m2[s[l]-'a']--
		l++
		r++
	}

	return res
}
