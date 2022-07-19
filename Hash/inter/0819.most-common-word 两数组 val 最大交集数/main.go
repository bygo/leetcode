package main

// https://leetcode.cn/problems/most-common-word

// ❓ 两数组 val 最大交集数

func mostCommonWord(paragraph string, banned []string) string {
	banMpCnt := map[string]int{}
	strMpCnt := map[string]int{}
	for i := range banned {
		banMpCnt[banned[i]] = 1
	}

	var strInter string
	var strTmp []byte
	var cntMax int

	fn := func() {
		str := string(strTmp)
		// 不在禁用列表
		if 0 < len(str) && 0 == banMpCnt[str] {
			strMpCnt[str]++
			if cntMax < strMpCnt[str] {
				strInter = str
				cntMax = strMpCnt[str]
			}
		}
		strTmp = []byte{}
	}
	for i := range paragraph {
		if 'A' <= paragraph[i] && paragraph[i] <= 'Z' {
			strTmp = append(strTmp, paragraph[i]+32)
		} else if 'a' <= paragraph[i] && paragraph[i] <= 'z' {
			strTmp = append(strTmp, paragraph[i])
		} else {
			fn()
		}
	}
	fn()
	return strInter
}
