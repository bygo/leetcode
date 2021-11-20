package main

import "sort"

// https://leetcode-cn.com/problems/longest-word-in-dictionary

func longestWord(words []string) string {
	sort.Strings(words)
	m := map[string]struct{}{}

	var res string
	for i := range words {
		l1 := len(words[i])
		if len(words[i]) == 1 {
			m[words[i]] = struct{}{}
			if len(res) < l1 {
				res = words[i]
			}
		} else {
			_, ok := m[words[i][:l1-1]]
			if ok {
				m[words[i]] = struct{}{}
				if len(res) < l1 {
					res = words[i]
				}
			}
		}

	}
	return res
}

type trie struct {
	child [26]*trie
	s     string
}

func longestWordTrie(words []string) string {
	root := &trie{}
	for _, word := range words {
		n := root
		for j := range word {
			ch := word[j] - 'a'
			if n.child[ch] == nil {
				n.child[ch] = &trie{}
			}
			n = n.child[ch]
		}
		n.s = word
	}

	var res string
	var queue = []*trie{root}
	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		res = queue[0].s
		for i := 0; i < cnt; i++ {
			for j := 0; j < 26; j++ {
				if queue[i].child[j] != nil && queue[i].child[j].s != "" {
					queue = append(queue, queue[i].child[j])
				}
			}
		}

		queue = queue[cnt:]
	}
	return res
}
