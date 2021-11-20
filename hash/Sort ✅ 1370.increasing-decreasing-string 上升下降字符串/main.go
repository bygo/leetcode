package main

// https://leetcode-cn.com/problems/increasing-decreasing-string

func sortString(s string) string {
	m := [26]int{}
	for i := range s {
		m[s[i]-'a']++
	}
	var res []byte
	l1 := len(s)
	for len(res) < l1 {
		for i := 0; i <= 25; i++ {
			if 0 < m[i] {
				res = append(res, byte(i+'a'))
				m[i]--
			}
		}

		for i := 25; 0 <= i; i-- {
			if 0 < m[i] {
				res = append(res, byte(i+'a'))
				m[i]--
			}
		}
	}
	return string(res)
}
