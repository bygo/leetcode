package main

// https://leetcode-cn.com/problems/design-search-autocomplete-system

type AutocompleteSystem struct {
	root *trie
	node *trie
	buf  []byte
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	as := AutocompleteSystem{
		root: &trie{chMp: [27]*trie{}},
	}
	root := as.root
	for i := range sentences {
		root.insert(sentences[i], times[i])
	}

	as.node = root
	return as
}

func (as *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		as.root.insert(string(as.buf), 1)
		as.buf = []byte{}
		as.node = as.root
		return nil
	}
	as.buf = append(as.buf, c)
	if c == ' ' {
		c = 0
	} else {
		c -= 96
	}

	if as.node.chMp[c] == nil {
		as.node.chMp[c] = &trie{chMp: [27]*trie{}}
	}
	var one, two, three string
	var oneCnt, twoCnt, threeCnt int

	as.node = as.node.chMp[c]
	var dfs func(node *trie)
	dfs = func(node *trie) {
		if 0 < node.cnt {
			if oneCnt < node.cnt {
				twoCnt, threeCnt = oneCnt, twoCnt
				two, three = one, two
				one = node.str
				oneCnt = node.cnt
			} else if twoCnt < node.cnt {
				threeCnt = twoCnt
				three = two
				two = node.str
				twoCnt = node.cnt
			} else if threeCnt < node.cnt {
				three = node.str
				threeCnt = node.cnt
			}
		}

		var i byte
		for ; i < 27; i++ {
			if node.chMp[i] == nil {
				continue
			}
			dfs(node.chMp[i])
		}
	}
	dfs(as.node)

	var strs []string
	if one != "" {
		strs = append(strs, one)
	}
	if two != "" {
		strs = append(strs, two)
	}
	if three != "" {
		strs = append(strs, three)
	}
	return strs
}

type trie struct {
	chMp [27]*trie
	cnt  int
	str  string
}

func (root *trie) insert(word string, cnt int) {
	node := root
	for i := range word {
		ch := word[i]
		if ch == ' ' {
			ch = 0
		} else {
			ch -= 96
		}
		if node.chMp[ch] == nil {
			node.chMp[ch] = &trie{chMp: [27]*trie{}}
		}
		node = node.chMp[ch]
	}
	node.cnt += cnt
	node.str = word
}
