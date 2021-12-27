package main

// https://leetcode-cn.com/problems/occurrences-after-bigram

// ❓ first second 紧接后的词

func findOcurrences(text string, first string, second string) []string {
	textL := len(text)
	var bufCur []byte
	var strPre string
	var strCur string
	var strsOrder []string
	fn := func() {
		if strPre == first && strCur == second {
			defer func() {
				strsOrder = append(strsOrder, strCur)
			}()
		}
		strPre = strCur
		strCur = string(bufCur)
		bufCur = []byte{}
	}
	for i := 0; i < textL; i++ {
		if text[i] == ' ' {
			fn()
		} else {
			bufCur = append(bufCur, text[i])
		}
	}
	fn()
	return strsOrder
}
