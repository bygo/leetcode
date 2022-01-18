package main

// ❓ 单词的压缩编码
// https://leetcode-cn.com/problems/short-encoding-of-words

func minimumLengthEncoding(words []string) int {
	root := &Trie{chMp: map[byte]*Trie{}}
	var cnt int
	for _, word := range words {
		root.Insert(word)
	}

	var dep int
	var dfs func(node *Trie)
	dfs = func(node *Trie) {
		if len(node.chMp) == 0 {
			cnt += dep + 1
			return
		}

		for _, n := range node.chMp {
			dep++
			dfs(n)
			dep--
		}
	}

	dfs(root)
	return cnt
}

type Trie struct {
	chMp map[byte]*Trie
}

func (t *Trie) Insert(word string) {
	node := t
	wordL := len(word)
	idx := wordL - 1
	for 0 <= idx {
		ch := word[idx]
		if node.chMp[ch] == nil {
			node.chMp[ch] = &Trie{chMp: map[byte]*Trie{}}
		}
		node = node.chMp[ch]
		idx--
	}
}
