package main

// substring with concatenation of all words
// https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}

	wordLen := len(words[0])
	wordNum := len(words)
	sLen := len(s)
	res := make([]int, 0)
	l := wordLen * wordNum
	var left, cursor int
	reset := true
	right := l
	var m map[string]int8
	for i := 0; i < wordLen; i++ {
		left = i
		cursor = left
		right = left + l
		if reset {
			reset = false
			m = map[string]int8{}
			for _, w := range words {
				m[w]++
			}
		}

		for cursor < right && right <= sLen {
			index := s[cursor : cursor+wordLen]
			how, ok := m[index]
			if !ok {
				if reset {
					reset = false
					m = map[string]int8{}
					for _, w := range words {
						m[w]++
					}
				}
				left = cursor + wordLen
				cursor = left
				right = left + l
				continue
			}
			if how <= 0 {
				m[s[left:left+wordLen]]++
				left += wordLen
				right += wordLen
				continue
			}
			reset = true
			m[index]--
			cursor += wordLen
			if cursor == right {
				res = append(res, left)
				m[s[left:left+wordLen]]++
				left += wordLen
				right += wordLen
			}
		}
	}
	return res
}

/**
思路:四指针i left cursor right
*/
