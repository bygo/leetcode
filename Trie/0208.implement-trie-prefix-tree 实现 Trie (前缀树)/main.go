package main

// https://leetcode-cn.com/problems/implement-trie-prefix-tree

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{
		children: [26]*Trie{},
	}
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

func (t *Trie) SearchPrefix(word string) *Trie {
	node := t
	for i := range word {
		ch := word[i] - 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.SearchPrefix(prefix)
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
