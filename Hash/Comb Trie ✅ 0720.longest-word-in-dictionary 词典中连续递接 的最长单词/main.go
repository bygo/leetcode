package main

import "sort"

// https://leetcode-cn.com/problems/longest-word-in-dictionary

// ❓词典中连续递接 的最长单词

func longestWord(words []string) string {
	sort.Strings(words)
	strMp := map[string]struct{}{}

	var longestStr string
	for _, str := range words {
		strL := len(str)
		if len(str) == 1 {
			// 单个字符 直接加入
			strMp[str] = struct{}{}
			// 比答案更长
			if len(longestStr) == 0 {
				longestStr = str
			}
		} else {
			// 多个字符 判断是否有前缀
			_, ok := strMp[str[:strL-1]]
			if ok {
				strMp[str] = struct{}{}
				// 比答案更长
				if len(longestStr) < strL {
					longestStr = str
				}
			}
		}

	}
	return longestStr
}

type trie struct {
	child [26]*trie
	s     string
}

func longestWordTrie(words []string) string {
	root := &trie{}
	for _, word := range words {
		n := root
		for j := range word {
			ch := word[j] - 'a'
			if n.child[ch] == nil {
				n.child[ch] = &trie{}
			}
			n = n.child[ch]
		}
		n.s = word
	}

	var longestStr string
	var queue = []*trie{root}
	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		longestStr = queue[0].s
		for i := 0; i < cnt; i++ {
			for j := 0; j < 26; j++ {
				if queue[i].child[j] != nil && queue[i].child[j].s != "" {
					queue = append(queue, queue[i].child[j])
				}
			}
		}

		queue = queue[cnt:]
	}
	return longestStr
}
