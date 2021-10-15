package main

// https://leetcode-cn.com/problems/word-search-ii

type Trie struct {
	children [26]*Trie
	word     string
}

func (t *Trie) Add(word string) {
	node := t
	for i := range word {
		ch := word[i] - 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.word = word
}

var graph = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func findWords(board [][]byte, words []string) []string {
	t := &Trie{}
	for _, word := range words {
		t.Add(word)
	}
	l1, l2 := len(board), len(board[0])
	m := map[string]bool{}

	var dfs func(node *Trie, x, y int)
	dfs = func(node *Trie, x, y int) {
		ch := board[x][y]
		node = node.children[ch-'a']
		if node == nil {
			return
		}
		if node.word != "" {
			m[node.word] = true
		}

		board[x][y] = '#'
		for _, dir := range graph {
			nx, ny := x+dir[0], y+dir[1]
			if 0 <= nx && nx < l1 && 0 <= ny && ny < l2 && board[nx][ny] != '#' {
				dfs(node, nx, ny)
			}
		}
		board[x][y] = ch
	}
	for x, row := range board {
		for y := range row {
			dfs(t, x, y)
		}
	}
	var res []string
	for i := range m {
		res = append(res, i)
	}
	return res
}
