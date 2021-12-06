package main

import (
	"strings"
)

// https://leetcode-cn.com/problems/permutations

func permute(nums []int) [][]int {
	var queue [][2]string
	n := len(nums)
	if n == 0 {
		return nil
	}

	var board, hand = "", build(nums)
	key := [2]string{"", hand}
	queue = [][2]string{key}
	vis := map[[2]string]bool{}
	vis[key] = true
	var newHand, newBoard = hand, ""
	for 0 < len(newHand) {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		for i := 0; i < cnt; i++ {
			q := queue[i]
			board = q[0]
			hand = q[1]
			for j := range hand {
				newBoard = board + string(hand[j])
				newHand = hand[:j] + hand[j+1:]
				key = [2]string{newBoard, newHand}
				if !vis[key] {
					queue = append(queue, key)
				}
				vis[key] = true
			}
		}
		queue = queue[cnt:]
	}
	var res [][]int
	for i := range queue {
		var cur []int
		for j := range queue[i][0] {
			cur = append(cur, int(queue[i][0][j])-10)
		}
		res = append(res, append([]int{}, cur...))
	}
	return res
}

func build(nums []int) string {
	b := strings.Builder{}
	for _, v := range nums {
		b.WriteByte(byte(v + 10))
	}
	return b.String()
}
