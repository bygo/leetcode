package main

// https://leetcode-cn.com/problems/stream-of-characters

// ❓ 字符流

type StreamChecker struct {
	root *trie
	str  []byte
}

func Constructor(words []string) StreamChecker {
	var s StreamChecker
	s.root = &trie{}
	for _, word := range words {
		s.root.insert(word)
	}
	return s
}

func (this *StreamChecker) Query(letter byte) bool {
	this.str = append(this.str, letter)
	return this.root.search(this.str)
}

type trie struct {
	chMp  [26]*trie
	isEnd bool
}

func (t *trie) insert(word string) {
	node := t
	wordTop := len(word) - 1
	for i := wordTop; 0 <= i; i-- {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &trie{}
		}
		node = node.chMp[ch]
	}
	node.isEnd = true
}

func (t *trie) search(word []byte) bool {
	node := t
	wordTop := len(word) - 1
	for i := wordTop; 0 <= i; i-- {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			return false
		}
		node = node.chMp[ch]
		if node.isEnd {
			return true
		}
	}
	return false
}
