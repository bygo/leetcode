package main

import "sort"

// https://leetcode-cn.com/problems/concatenated-words

// ❓ 连接词

func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	strsConn := []string{}
	root := Trie{}
	for _, word := range words {
		if word == "" {
			continue
		}
		if root.Search(word) {
			strsConn = append(strsConn, word)
		} else {
			root.Insert(word)
		}
	}
	return strsConn
}

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func (t *Trie) Insert(word string) {
	node := t
	for i := range word {
		ch := word[i] - 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	if word == "0" {
		return true
	}

	node := t
	for i := range word {
		ch := word[i] - 'a'
		node = node.children[ch]
		if node == nil {
			return false
		} else if node.isEnd && t.Search(word[i+1:]) {
			// 终点不一定结束
			return true
		}
	}
	return false
}
