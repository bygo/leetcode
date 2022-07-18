package main

// https://leetcode.cn/problems/number-of-matching-subsequences

func numMatchingSubseq(s string, words []string) int {
	chMpNodes := [26][]*Node{}

	for _, word := range words {
		ch := word[0] - 'a'
		chMpNodes[ch] = append(chMpNodes[ch], &Node{word, 0})
	}
	sL := len(s)
	var cnt int
	for i := 0; i < sL; i++ {
		ch := s[i] - 'a'
		if 0 < len(chMpNodes[ch]) {
			nodes := chMpNodes[ch]
			chMpNodes[ch] = nil
			for _, node := range nodes {
				if node.idx == len(node.word)-1 {
					cnt++
				} else {
					node.idx++
					chCur := node.word[node.idx] - 'a'
					chMpNodes[chCur] = append(chMpNodes[chCur], node)
				}
			}

		}
	}
	return cnt
}

type Node struct {
	word string
	idx  int
}
