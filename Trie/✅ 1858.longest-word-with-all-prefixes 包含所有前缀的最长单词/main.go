package main

// https://leetcode-cn.com/problems/longest-word-with-all-prefixes

// ❓ 包含所有前缀的最长单词

type trie struct {
	chMp [26]*trie
	str  string
}

func (root *trie) insert(word string) {
	node := root
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &trie{}
		}
		node = node.chMp[ch]
	}
	node.str = word
}

func longestWord(words []string) string {
	root := &trie{}
	for _, word := range words {
		root.insert(word)
	}
	var strLongest string
	var dfs func(node *trie)
	dfs = func(node *trie) {
		if len(strLongest) < len(node.str) {
			strLongest = node.str
		}

		for _, n := range node.chMp {
			if n != nil && 0 < len(n.str) {
				dfs(n)
			}
		}
	}
	dfs(root)
	return strLongest
}
