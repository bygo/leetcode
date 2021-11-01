package main

// https://leetcode-cn.com/problems/most-common-word

func mostCommonWord(paragraph string, banned []string) string {
	m1 := map[string]int{}
	m2 := map[string]int{}
	banned = append(banned, "#")
	for i := range banned {
		cur := []byte{}
		for j := range banned[i] {
			if 'A' <= banned[i][j] && banned[i][j] <= 'Z' {
				cur = append(cur, banned[i][j]+32)
			} else if 'a' <= banned[i][j] && banned[i][j] <= 'z' {
				cur = append(cur, banned[i][j])
			}
		}
		m1[string(cur)]++
	}

	res := ""
	max := 0
	cur := []byte{}
	f := func() {
		s := string(cur)
		if 0 < len(cur) {
			if 0 == m1[s] {
				m2[s]++
				if max < m2[s] {
					res = s
					max = m2[s]
				}
			}
		}
		cur = []byte{}
	}
	for i := range paragraph {
		if 'A' <= paragraph[i] && paragraph[i] <= 'Z' {
			cur = append(cur, paragraph[i]+32)
		} else if 'a' <= paragraph[i] && paragraph[i] <= 'z' {
			cur = append(cur, paragraph[i])
		} else {
			f()
		}
	}
	f()
	return res
}
