package main

// https://leetcode.cn/problems/design-add-and-search-words-data-structure

// ❓ 添加与搜索单词

type WordDictionary struct {
	Trie
}

type Trie struct {
	chMp  [26]*Trie
	isEnd bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (root *Trie) AddWord(word string) {
	node := root
	for _, ch := range word {
		ch -= 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &Trie{}
		}
		node = node.chMp[ch]
	}
	node.isEnd = true
}

func (root *Trie) Search(word string) bool {
	wordL := len(word)

	var dfs func(int, *Trie) bool
	dfs = func(idx int, node *Trie) bool {
		if idx == wordL {
			return node.isEnd
		}
		ch := word[idx]
		if ch == '.' {
			// 模糊匹配
			for i := range node.chMp {
				if node.chMp[i] != nil && dfs(idx+1, node.chMp[i]) {
					return true
				}
			}
		} else if 'a' <= ch && ch <= 'z' {
			// 精准匹配
			node = node.chMp[ch-'a']
			if node != nil && dfs(idx+1, node) {
				return true
			}
		}
		return false
	}
	return dfs(0, root)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
