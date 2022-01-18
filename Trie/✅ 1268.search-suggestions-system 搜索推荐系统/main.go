package main

import "sort"

// https://leetcode-cn.com/problems/search-suggestions-system

// ❓ 搜索推荐系统

type Trie struct {
	chMp [26]*Trie
	idx  []int
}

func (t *Trie) Insert(idx int, word string) {
	node := t
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &Trie{}
		}
		node = node.chMp[ch]
		node.idx = append(node.idx, idx)
	}
}

func (t *Trie) Search(word string) []int {
	node := t
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			return nil
		}
		node = node.chMp[ch]
	}
	return node.idx
}

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	root := &Trie{}
	for idx, product := range products {
		root.Insert(idx, product)
	}

	var combsProducts = make([][]string, len(searchWord))
	for i := range searchWord {
		ch := searchWord[i] - 'a'
		if root.chMp[ch] == nil {
			break
		}
		root = root.chMp[ch]
		if 3 < len(root.idx) {
			root.idx = root.idx[:3]
		}
		for _, idx := range root.idx {
			combsProducts[i] = append(combsProducts[i], products[idx])
		}
	}
	return combsProducts
}
