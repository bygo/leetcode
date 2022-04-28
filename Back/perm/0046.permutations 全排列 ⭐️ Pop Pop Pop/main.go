package main

import (
	"strings"
)

// https://leetcode-cn.com/problem	s/permutations

func permute(nums []int) [][]int {
	var que [][2]string
	numsL := len(nums)
	if numsL == 0 {
		return nil
	}

	var board, hand = "", build(nums)
	key := [2]string{"", hand}
	que = [][2]string{key}
	var handNew, boardNew = hand, ""
	for {
		handL := len(handNew)
		if handL == 0 {
			break
		}
		queL := len(que)
		for _, q := range que {
			board = q[0]
			hand = q[1]
			for j := range hand {
				boardNew = board + string(hand[j])
				handNew = hand[:j] + hand[j+1:]
				key = [2]string{boardNew, handNew}
				que = append(que, key)
			}
		}
		que = que[queL:]
	}
	var permNums [][]int
	for i := range que {
		var cur []int
		str := que[i][0]
		for j := range str {
			cur = append(cur, int(str[j])-10)
		}
		permNums = append(permNums, append([]int{}, cur...))
	}
	return permNums
}

func build(nums []int) string {
	b := strings.Builder{}
	for _, v := range nums {
		b.WriteByte(byte(v + 10))
	}
	return b.String()
}
