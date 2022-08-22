package TwoPointer

// https://leetcode.cn/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence

func isPrefixOfWord(sentence string, searchWord string) int {
	sL := len(sentence)
	wL := len(searchWord)
	cnt := 0
	for idx := 0; idx < sL; idx++ {
		idxLeft := idx
		for idx < sL && sentence[idx] != ' ' {
			idx++
		}
		cnt++
		idxRight := idxLeft + wL
		if idxRight <= sL && sentence[idxLeft:idxRight] == searchWord {
			return cnt
		}
	}
	return -1
}
