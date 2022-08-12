package main

import "sort"

// https://leetcode.cn/problems/concatenated-words

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
	chMp  [26]*Trie
	isEnd bool
}

func (root *Trie) Insert(word string) {
	node := root
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &Trie{}
		}
		node = node.chMp[ch]
	}
	node.isEnd = true
}

func (root *Trie) Search(word string) bool {
	if word == "" {
		return true
	}

	node := root
	for i := range word {
		ch := word[i] - 'a'
		node = node.chMp[ch]
		if node == nil {
			return false
		} else if node.isEnd && root.Search(word[i+1:]) { // 进入下一次寻找
			// 终点不一定结束
			return true
		}
	}
	return false
}
