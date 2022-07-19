package main

// https://leetcode.cn/problems/number-of-valid-words-for-each-puzzle

type trie struct {
	chMp [26]*trie
	cnt  int
}

func (t *trie) insert(word string) {
	// 排序并去重
	chMp := [26]bool{}
	for i := range word {
		ch := word[i] - 'a'
		chMp[ch] = true
	}

	node := t
	for ch, b := range chMp {
		if b {
			if node.chMp[ch] == nil {
				node.chMp[ch] = &trie{}
			}
			node = node.chMp[ch]
		}
	}
	node.cnt++
}

func findNumOfValidWords(words []string, puzzles []string) []int {
	root := &trie{}
	for _, word := range words {
		root.insert(word)
	}

	numsPuzzle := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		first := puzzle[0] - 'a'
		chMp := [26]bool{}
		for j := range puzzle {
			ch := puzzle[j] - 'a'
			chMp[ch] = true
		}

		var bufPuzzle []byte
		for ch, b := range chMp {
			if b {
				bufPuzzle = append(bufPuzzle, byte(ch))
			}
		}
		// 所有子集
		var dfs func(idx int, node *trie) int
		dfs = func(idx int, node *trie) int {
			if node == nil {
				return 0
			}

			// 结算
			if idx == len(bufPuzzle) {
				return node.cnt
			}

			// 选择
			ch := bufPuzzle[idx]
			cnt := dfs(idx+1, node.chMp[ch])

			// 不选择
			if bufPuzzle[idx] != first {
				// 当 idx 不为首字母时，可以不选择第 idx 个字母
				cnt += dfs(idx+1, node)
			}

			return cnt
		}

		numsPuzzle[i] = dfs(0, root)
	}

	return numsPuzzle
}
