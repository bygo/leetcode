package main

// https://leetcode-cn.com/problems/design-add-and-search-words-data-structure

type Trie struct {
	child [26]*Trie
	isEnd bool
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.child[ch] == nil {
			node.child[ch] = &Trie{}
		}
		node = node.child[ch]
	}
	node.isEnd = true
}

type WordDictionary struct {
	root *Trie
}

func Constructor() WordDictionary {
	return WordDictionary{&Trie{}}
}

func (d *WordDictionary) AddWord(word string) {
	d.root.Insert(word)
}

func (d *WordDictionary) Search(word string) bool {
	var dfs func(int, *Trie) bool
	l := len(word)
	dfs = func(index int, node *Trie) bool {
		if index == l {
			return node.isEnd
		}
		ch := word[index]
		if ch != '.' {
			child := node.child[ch-'a']
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for i := range node.child {
				if node.child[i] != nil && dfs(index+1, node.child[i]) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, d.root)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
