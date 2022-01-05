package main

// https://leetcode-cn.com/problems/add-bold-tag-in-string

// ❓ 给字符串添加加粗标签

type trie struct {
	chMp  map[byte]*trie
	isEnd bool
}

func addBoldTag(s string, words []string) string {
	node := &trie{chMp: map[byte]*trie{}}
	for _, word := range words {
		node.insert(word)
	}

	sL := len(s)
	var mask = make([]bool, sL)
	var idxMax int
	for i := 0; i < sL; i++ {
		idxMaxCur := i + node.search(s[i:])
		if idxMax < idxMaxCur {
			for j := i; j < idxMaxCur; j++ {
				mask[j] = true
			}
			idxMax = idxMaxCur
		}
	}
	var bufBold []byte
	var i int
	for i < sL {
		if mask[i] {
			bufBold = append(bufBold, "<b>"...)
			for i < sL && mask[i] {
				bufBold = append(bufBold, s[i])
				i++
			}
			bufBold = append(bufBold, "</b>"...)
		} else {
			bufBold = append(bufBold, s[i])
			i++
		}
	}
	return string(bufBold)
}

func (t *trie) insert(s string) {
	node := t
	for i := range s {
		ch := s[i]
		if node.chMp[ch] == nil {
			node.chMp[ch] = &trie{chMp: map[byte]*trie{}}
		}
		node = node.chMp[ch]
	}
	node.isEnd = true
}

func (t *trie) search(s string) int {
	node := t
	var cntMax int
	var cnt int
	for i := range s {
		ch := s[i]
		if node.chMp[ch] == nil {
			return cntMax
		}
		node = node.chMp[ch]
		cnt++
		if node.isEnd {
			cntMax = cnt
		}
	}
	return cntMax
}
