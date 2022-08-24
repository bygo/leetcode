package main

import "math/bits"

// https://leetcode.cn/problems/number-of-valid-words-for-each-puzzle

func findNumOfValidWords(words []string, puzzles []string) []int {
	setMpCnt := map[int]int{}
	for _, word := range words {
		var subset int
		for _, ch := range word {
			ch -= 'a'
			subset |= 1 << ch
		}
		// The longest puzzle len(puzzles[i]) == 7
		if bits.OnesCount(uint(subset)) <= 7 {
			setMpCnt[subset]++
		}
	}

	var cntsPuzzle = make([]int, len(puzzles))
	for idx, puzzle := range puzzles {
		first := 1 << (puzzle[0] - 'a')

		var subset int
		for _, ch := range puzzle[1:] {
			subset |= 1 << (ch - 'a')
		}
		sub := subset
		for {
			cntsPuzzle[idx] += setMpCnt[sub|first]
			sub = (sub - 1) & subset
			if sub == subset {
				break
			}
		}

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
	return cntsPuzzle
}
