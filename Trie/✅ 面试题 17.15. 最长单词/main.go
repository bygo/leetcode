package main

// https://leetcode-cn.com/problems/longest-word-lcci/

// ❓ 恢复空格

type Trie struct {
	chMp  [26]*Trie
	isEnd bool
}

func (root *Trie) Insert(word string) {
	node := root
	for i := range word {
		ch := word[i] - 'a'
		if node.chMp[ch] == nil {
			node.chMp[ch] = &Trie{}
		}
		node = node.chMp[ch]
	}
	node.isEnd = true
}

func longestWord(words []string) string {
	root := &Trie{}
	for _, word := range words {
		root.Insert(word)
	}
	strMax := ""
	var strCur string
	var nodeCur *Trie
	var dfs func(idx, cnt int) bool
	dfs = func(idx, cnt int) bool {
		if idx == len(strCur) {
			// 组合2次以上，且刚好结束
			return 2 <= cnt && nodeCur == root
		}
		ch := strCur[idx] - 'a'
		if nodeCur.chMp[ch] == nil {
			return false
		}
		nodeCur = nodeCur.chMp[ch]
		if nodeCur.isEnd {
			nodeTmp := nodeCur
			nodeCur = root
			if dfs(idx+1, cnt+1) {
				return true
			}
			// 回溯
			nodeCur = nodeTmp
		}
		return dfs(idx+1, cnt)
	}

	for _, word := range words {
		strCur = word
		nodeCur = root
		if len(strMax) < len(strCur) || len(strMax) == len(strCur) && word < strMax {
			if dfs(0, 0) {
				strMax = strCur
			}
		}
	}
	return strMax
}
