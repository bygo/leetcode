package main

// https://leetcode.cn/problems/group-shifted-strings

// ❓ 移位字符串分组

func groupStrings(strings []string) [][]string {
	strMpID := map[string]int{}
	var IDsStrings [][]string
	getID := func(str string) int {
		_, ok := strMpID[str]
		if !ok {
			strMpID[str] = len(IDsStrings)
			IDsStrings = append(IDsStrings, []string{})
		}
		return strMpID[str]
	}

	for _, str := range strings {
		idxStr := 1
		strL := len(str)
		strBuf := make([]byte, strL)
		for idxStr < strL {
			ch := (str[idxStr] - str[0] + 26) % 26 // +26以防变负
			strBuf = append(strBuf, ch)
			idxStr++
		}
		id := getID(string(strBuf))
		IDsStrings[id] = append(IDsStrings[id], str)
	}
	return IDsStrings
}
