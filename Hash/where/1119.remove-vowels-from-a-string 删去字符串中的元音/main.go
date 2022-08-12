package main

// https://leetcode.cn/problems/remove-vowels-from-a-string

// ❓ 移除元音字符

func removeVowels(str string) string {
	//strL := len(str)
	var strBuf []byte
	for i := range str {
		// 元音跳过
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' {
			continue
		} else {
			strBuf = append(strBuf, str[i])
		}
	}
	return string(strBuf)
}
