package main

import (
	"strconv"
)

// https://leetcode-cn.com/problems/word-abbreviation

// ❓ 单词缩写

type pair struct {
	word string
	idx  int
}

func wordsAbbreviation(words []string) []string {
	wordsL := len(words)
	var wordsStd = make([]string, wordsL)
	var lenMpWords = map[[3]int][]pair{}
	for idx, word := range words {
		wordL := len(word)
		key := [3]int{int(word[0]), wordL, int(word[wordL-1])}
		lenMpWords[key] = append(lenMpWords[key], pair{word, idx})
	}
	for _, words := range lenMpWords {
		root := &Trie{}
		for _, p := range words {
			root.Insert(p.word)
		}
		for _, p := range words {
			word := p.word
			wordTop := len(word) - 1
			idxPre := root.Search(word)
			idxCom := wordTop - idxPre - 1

			if 1 < idxCom {
				wordCom := []byte(word[:idxPre+1])
				wordCom = append(wordCom, strconv.Itoa(idxCom)...)
				wordCom = append(wordCom, word[wordTop])
				wordsStd[p.idx] = string(wordCom)
			} else {
				wordsStd[p.idx] = word
			}
		}
	}

	return wordsStd
}

type Trie struct {
	chMp [26]*Trie
	idx  int
	cnt  int
}

func (t *Trie) Insert(word string) {
	node := t
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &Trie{}
		}
		node = node.chMp[ch]
		node.cnt++
	}
}

func (t *Trie) Search(word string) int {
	node := t
	wordL := len(word)
	for i := range word {
		ch := word[i] - 'a'
		node = node.chMp[ch]
		if node == nil || node.cnt == 1 {
			return i
		}
	}
	return wordL
}
