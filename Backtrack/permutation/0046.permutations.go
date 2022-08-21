package main

import "strings"

// https://leetcode.cn/problems/permutations

// DFS
func permute(raw []int) [][]int {
	var permNums [][]int
	rawL := len(raw)
	nums := make([]int, rawL)
	used := make([]bool, rawL)
	var dfs func(int)
	dfs = func(idx int) {
		if idx == rawL {
			permNums = append(permNums, append([]int{}, nums...))
			return
		}
		for i := 0; i < rawL; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			nums[idx] = raw[i]
			dfs(idx + 1)
			used[i] = false
		}
	}
	dfs(0)
	return permNums
}

// Insert
func permute(raw []int) [][]int {
	numsL := len(raw)
	if numsL == 0 {
		return nil
	}
	var que = [][]int{{}}
	var numL = 0
	for _, num := range raw {
		queL := len(que)
		for _, q := range que {
			for idx := 0; idx <= numL; idx++ {
				nums := append([]int{}, q[:idx]...)
				nums = append(nums, num)
				nums = append(nums, q[idx:]...)

				que = append(que, nums)
			}
		}
		// 1
		// 1 2
		// 2 1
		// 3 1 2
		// 1 3 2
		// 1 2 3

		// 3 2 1
		// 2 3 1
		// 2 1 3
		numL++
		que = que[queL:]
	}
	return que
}

// like ZuMa
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
			cur = append(cur, int(str[j])-10) // offset value
		}
		permNums = append(permNums, append([]int{}, cur...))
	}
	return permNums
}

func build(nums []int) string {
	b := strings.Builder{}
	for _, v := range nums {
		b.WriteByte(byte(v + 10)) // offset value
	}
	return b.String()
}
