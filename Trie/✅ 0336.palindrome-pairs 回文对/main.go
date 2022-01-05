package main

import "sort"

// https://leetcode-cn.com/problems/palindrome-pairs

func palindromePairs(words []string) [][]int {
	prefix := Trie{}
	suffix := Trie{}
	wordsL := len(words)
	// 插入所有值

	wordsStd := make([]string, wordsL)
	wordsMirror := make([]string, wordsL)
	wordMpID := map[string]int{}

	for i, word := range words {
		wordsStd[i] = word
		wordMpID[word] = i
		wordTop := len(word) - 1
		var wordMirrorBuf []byte
		for j := wordTop; 0 <= j; j-- {
			wordMirrorBuf = append(wordMirrorBuf, word[j])
		}
		wordsMirror[i] = string(wordMirrorBuf)
	}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	var pairsNums [][]int

	for i := 0; i < wordsL; i++ {
		word := words[i]
		idx := wordMpID[word]

		wordMirror := wordsMirror[idx]
		wordL := len(word)

		idxesPrefix := prefix.Search(wordMirror)

		for _, idxPrefix := range idxesPrefix {
			if idxPrefix == idx {
				continue
			}
			wordPrefixL := len(wordsStd[idxPrefix])
			if valid(word[:wordL-wordPrefixL]) {
				pairsNums = append(pairsNums, []int{idxPrefix, idx})
			}
		}

		idxesSuffix := suffix.Search(word)
		for _, idxSuffix := range idxesSuffix {
			if idxSuffix == idx {
				continue
			}
			wordSuffixL := len(wordsStd[idxSuffix])
			if valid(word[wordSuffixL:]) {
				pairsNums = append(pairsNums, []int{idx, idxSuffix})
			}
		}

		prefix.Insert(word, idx+1)
		suffix.Insert(wordMirror, idx+1)
	}

	return pairsNums
}

type Trie struct {
	chMp     [26]*Trie
	idxWords int
}

func (root *Trie) Insert(word string, idxWords int) {
	node := root
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &Trie{}
		}
		node = node.chMp[ch]
	}
	node.idxWords = idxWords
}

func (root *Trie) Search(word string) (idxesShort []int) {
	wordL := len(word)
	node := root

	if 0 < node.idxWords {
		idxesShort = append(idxesShort, node.idxWords-1)
	}

	for i := 0; i < wordL; i++ {
		ch := word[i] - 'a'
		node = node.chMp[ch]
		if node == nil {
			return
		}
		if 0 < node.idxWords {
			idxesShort = append(idxesShort, node.idxWords-1)
		}
	}
	return
}

func valid(word string) bool {
	wordL := len(word)
	mid := wordL / 2
	for i := 0; i < mid; i++ {
		if word[i] != word[wordL-i-1] {
			return false
		}
	}

	return true
}
