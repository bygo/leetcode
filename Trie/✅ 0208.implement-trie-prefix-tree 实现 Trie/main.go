package main

// https://leetcode-cn.com/problems/implement-trie-prefix-tree

// ❓ 实现 Trie

type Trie struct {
	chMp  [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{
		chMp: [26]*Trie{},
	}
}

func (t *Trie) Insert(word string) {
	node := t
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &Trie{}
		}
		node = node.chMp[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(word string) *Trie {
	node := t
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			return nil
		}
		node = node.chMp[ch]
	}
	return node
}

// 有没有单词 word

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

// 有没有前缀 prefix

func (t *Trie) StartsWith(prefix string) bool {
	node := t.SearchPrefix(prefix)
	return node != nil
}
