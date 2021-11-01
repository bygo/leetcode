package main

// https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

func firstUniqChar(s string) byte {
	l1 := len(s)
	m := [26]int{}
	for i := range m {
		m[i] = l1
	}

	for i := range s {
		k := s[i] - 'a'
		if m[k] == l1 {
			m[k] = i
		} else {
			m[k] = l1 + 1
		}
	}

	var res = l1
	for i := range m {
		if m[i] < res {
			res = m[i]
		}
	}
	if res == l1 {
		return ' '
	}
	return s[res]
}
