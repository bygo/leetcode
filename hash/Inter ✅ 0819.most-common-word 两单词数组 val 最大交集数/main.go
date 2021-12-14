package main

// https://leetcode-cn.com/problems/most-common-word

func mostCommonWord(paragraph string, banned []string) string {
	ban := map[string]int{}
	m2 := map[string]int{}
	banned = append(banned, "#")
	for i := range banned {
		ban[banned[i]]++
	}

	var res string
	var max int
	var cur []byte
	f := func() {
		str := string(cur)
		if 0 < len(str) && 0 == ban[str] {
			m2[str]++
			if max < m2[str] {
				res = str
				max = m2[str]
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
