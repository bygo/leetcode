package main

// https://leetcode-cn.com/problems/most-common-word

// ❓ 两数组 val 最大交集数

func mostCommonWord(paragraph string, banned []string) string {
	banMpCnt := map[string]int{}
	strMpCnt := map[string]int{}
	for i := range banned {
		banMpCnt[banned[i]] = 1
	}

	var res string
	var max int
	var tmp []byte
	fn := func() {
		str := string(tmp)
		if 0 < len(str) && 0 == banMpCnt[str] {
			strMpCnt[str]++
			if max < strMpCnt[str] {
				res = str
				max = strMpCnt[str]
			}
		}
		tmp = []byte{}
	}
	for i := range paragraph {
		if 'A' <= paragraph[i] && paragraph[i] <= 'Z' {
			tmp = append(tmp, paragraph[i]+32)
		} else if 'a' <= paragraph[i] && paragraph[i] <= 'z' {
			tmp = append(tmp, paragraph[i])
		} else {
			fn()
		}
	}
	fn()
	return res
}
