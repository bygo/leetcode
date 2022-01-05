package main

// https://leetcode-cn.com/problems/prefix-and-suffix-search

type WordFilter struct {
	root *trie
}

func Constructor(words []string) WordFilter {
	wf := WordFilter{
		root: &trie{chMpTrie: map[[2]byte]*trie{}},
	}

	for idx, word := range words {
		wf.root.insert(word, idx)
	}
	return wf
}

func (wf *WordFilter) F(prefix string, suffix string) int {
	return wf.root.search(prefix, suffix)
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */

type trie struct {
	chMpTrie map[[2]byte]*trie
	idx      int
}

func (t *trie) insert(word string, idx int) {
	node := t
	wordL := len(word)
	for i := 0; i < wordL; i++ {
		nodeTmp := node
		for j := i; j < wordL; j++ {
			ch := [2]byte{word[j] - '`', 0}
			if nodeTmp.chMpTrie[ch] == nil {
				nodeTmp.chMpTrie[ch] = &trie{chMpTrie: map[[2]byte]*trie{}}
			}
			nodeTmp = nodeTmp.chMpTrie[ch]
			nodeTmp.idx = idx
		}
		nodeTmp = node
		for j := wordL - 1 - i; 0 <= j; j-- {
			ch := [2]byte{0, word[j] - '`'}
			if nodeTmp.chMpTrie[ch] == nil {
				nodeTmp.chMpTrie[ch] = &trie{chMpTrie: map[[2]byte]*trie{}}
			}
			nodeTmp = nodeTmp.chMpTrie[ch]
			nodeTmp.idx = idx
		}

		ch := [2]byte{word[i] - '`', word[wordL-1-i] - '`'}
		if node.chMpTrie[ch] == nil {
			node.chMpTrie[ch] = &trie{chMpTrie: map[[2]byte]*trie{}}
		}
		node = node.chMpTrie[ch]
		node.idx = idx
	}
}

func (t *trie) search(prefix, suffix string) int {
	node := t
	prefixL := len(prefix)
	suffixL := len(suffix)

	var i int
	for i < prefixL || i < suffixL {
		var ch [2]byte
		if i < prefixL {
			ch[0] = prefix[i]
		}

		if i < suffixL {
			ch[1] = suffix[suffixL-i-1]
		}
		node = node.chMpTrie[ch]
		if node == nil {
			return -1
		}
		i++
	}
	return node.idx
}
