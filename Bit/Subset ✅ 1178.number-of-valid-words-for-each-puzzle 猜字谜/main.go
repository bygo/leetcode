package main

import "math/bits"

// https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle

// ❓ 猜字谜
// ⚠️ word 包含 第一个 puzzle 第一个字母
// ⚠️ word 是 puzzle 的子集

func findNumOfValidWords(words []string, puzzles []string) []int {
	// 生成谜底 hash
	numMpCnt := map[int]int{}
	for _, word := range words {
		var num int
		for _, ch := range word {
			ch -= 'a'
			num |= 1 << ch
		}
		// 谜题最长 len(puzzles[i]) == 7
		if bits.OnesCount(uint(num)) <= 7 {
			numMpCnt[num]++
		}
	}

	// 每个谜题 `puzzles` 的 谜底数 `word`
	var numsPuzzles = make([]int, len(puzzles))
	for idx, puzzle := range puzzles {
		first := 1 << (puzzle[0] - 'a')

		// puzzles 全集
		var num int
		for _, ch := range puzzle[1:] {
			num |= 1 << (ch - 'a')
		}
		// puzzles 所有子集
		subset := num
		for {
			// first 必须存在
			numsPuzzles[idx] += numMpCnt[subset|first]
			// 递减时，依次由从低位到高位去除1，
			// 每次将挑出一个1形成新子集，
			// 到达高位时，高位无法再借1，整体将损失一个1，继续从`新的整体`挑出1
			// 到达 0 - 1 时, -1 & num = num
			subset = (subset - 1) & num
			if subset == num {
				break
			}
		}

		// 穷举
		//ch := word[0] - 'a'
		//first := 1 << ch
		//for numChoose := 0; numChoose < 1<<6; numChoose++ {
		//	var num int
		//	for j := 0; j < 6; j++ {
		//		if numChoose>>j&1 == 1 {
		//			num |= 1 << (word[j+1] - 'a')
		//		}
		//	}
		//	numsPuzzle[i] += numMpCnt[num|first]
		//}
	}
	return numsPuzzles
}
