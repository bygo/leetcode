package main

// https://leetcode-cn.com/problems/implement-trie-ii-prefix-tree

// ❓ 实现 Trie II

type Trie struct {
	chMp   [26]*Trie
	cnt    int
	cntEnd int
}

func Constructor() Trie {
	return Trie{}
}

func (root *Trie) Insert(word string) {
	node := root
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &Trie{}
		}
		node = node.chMp[ch]
		node.cnt++
	}
	node.cntEnd++
}

func (root *Trie) CountWordsEqualTo(word string) int {
	node := root.search(word)
	if node == nil {
		return 0
	}
	return node.cntEnd
}

func (root *Trie) CountWordsStartingWith(prefix string) int {
	node := root.search(prefix)
	if node == nil {
		return 0
	}
	return node.cnt
}

func (root *Trie) Erase(word string) {
	node := root
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			return
		}
		node = node.chMp[ch]
		if 0 < node.cnt {
			node.cnt--
		}
	}
	if 0 < node.cntEnd {
		node.cntEnd--
	}
}

func (root *Trie) search(word string) *Trie {
	node := root
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			return nil
		}
		node = node.chMp[ch]
	}
	return node
}
