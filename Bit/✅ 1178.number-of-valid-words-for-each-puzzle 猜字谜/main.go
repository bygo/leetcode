package main

import "math/bits"

// https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle

// ❓ 猜字谜

func findNumOfValidWords(words []string, puzzles []string) []int {
	numMpCnt := map[int]int{}
	for _, word := range words {
		var num int
		for _, ch := range word {
			ch -= 'a'
			num |= 1 << ch
		}
		if bits.OnesCount(uint(num)) <= 7 {
			numMpCnt[num]++
		}
	}

	var numsPuzzle = make([]int, len(puzzles))
	for i, word := range puzzles {
		ch := word[0] - 'a'
		first := 1 << ch
		var num int
		for _, ch := range word[1:] {
			ch -= 'a'
			num |= 1 << ch
		}
		subset := num
		for {
			numsPuzzle[i] += numMpCnt[subset|first]
			// 递减时，依次由从低位到高位去除1，
			// 每次将挑出一个1形成新子集，
			// 到达高位时，高位无法再借1，整体将损失一个1，继续从`新的整体`挑出1
			// 到达-1时, -1 & num = num
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
	return numsPuzzle
}
