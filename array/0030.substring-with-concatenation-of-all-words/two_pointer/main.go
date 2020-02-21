package main

/**
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

 
示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]
 */

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