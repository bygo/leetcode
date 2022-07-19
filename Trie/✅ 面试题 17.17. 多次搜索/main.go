package main

// https://leetcode.cn/problems/multi-search-lcci/

// ❓ 面试题 17.17. 多次搜索

type Trie struct {
	chMp  [26]*Trie
	idxes []int
}

func (t *Trie) Insert(word string, idx int) {
	node := t
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &Trie{}
		}
		node = node.chMp[ch]
	}
	node.idxes = append(node.idxes, idx)
}

func (t *Trie) Search(word string) []int {
	node := t
	idxes := []int{}
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			return idxes
		}
		node = node.chMp[ch]
		idxes = append(idxes, node.idxes...)
	}
	return idxes
}

func multiSearch(big string, smalls []string) [][]int {
	root := &Trie{}
	for idx, small := range smalls {
		root.Insert(small, idx)
	}

	var idxBig int
	bigL := len(big)
	smallsL := len(smalls)
	var combIdxes = make([][]int, smallsL)
	for idxBig < bigL {
		idxes := root.Search(big[idxBig:])
		for _, idx := range idxes {
			combIdxes[idx] = append(combIdxes[idx], idxBig)
		}
		idxBig++
	}
	return combIdxes
}
