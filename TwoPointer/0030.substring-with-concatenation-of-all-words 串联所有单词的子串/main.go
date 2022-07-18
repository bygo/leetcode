package main

// substring with concatenation of all words
// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}

	var wordLen = len(words[0])
	var wordNum = len(words)
	var l1 = len(s)
	var res = make([]int, 0)
	var l2 = wordLen * wordNum
	var l, cur int
	var reset = true
	var r = l2
	var m = map[string]int{}

	for i := 0; i < wordLen; i++ {
		l = i
		cur = l
		r = l + l2
		if reset {
			reset = false
			m = map[string]int{}
			for _, w := range words {
				m[w]++
			}
		}

		for cur < r && r <= l1 {
			var index = s[cur : cur+wordLen]
			var how, ok = m[index]
			if !ok {
				if reset {
					reset = false
					m = map[string]int{}
					for _, w := range words {
						m[w]++
					}
				}
				l = cur + wordLen
				cur = l
				r = l + l2
				continue
			}
			if how <= 0 {
				m[s[l:l+wordLen]]++
				l += wordLen
				r += wordLen
				continue
			}
			reset = true
			m[index]--
			cur += wordLen
			if cur == r {
				res = append(res, l)
				m[s[l:l+wordLen]]++
				l += wordLen
				r += wordLen
			}
		}
	}
	return res
}
