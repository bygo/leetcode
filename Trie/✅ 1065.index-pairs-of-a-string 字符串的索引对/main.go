package main

// https://leetcode.cn/problems/index-pairs-of-a-string
// ❓ 字符串的索引对

func indexPairs(text string, words []string) [][]int {
	root := &Trie{chMp: map[byte]*Trie{}}
	for _, word := range words {
		root.insert(word)
	}
	textL := len(text)
	var combsIdxes [][]int
	for i := 0; i < textL; i++ {
		deps := root.search(text[i:])
		for _, dep := range deps {
			combsIdxes = append(combsIdxes, []int{i, i + dep})
		}
	}
	return combsIdxes
}

type Trie struct {
	chMp  map[byte]*Trie
	isEnd bool
}

func (t *Trie) insert(word string) {
	node := t
	for i := range word {
		ch := word[i]
		if node.chMp[ch] == nil {
			node.chMp[ch] = &Trie{chMp: map[byte]*Trie{}}
		}
		node = node.chMp[ch]
	}
	node.isEnd = true
}

func (t *Trie) search(word string) []int {
	node := t
	var dep = 0
	var deps []int
	for i := range word {
		ch := word[i]
		if node.chMp[ch] == nil {
			return deps
		}
		node = node.chMp[ch]
		if node.isEnd {
			deps = append(deps, dep)
		}
		dep++
	}
	return deps
}
