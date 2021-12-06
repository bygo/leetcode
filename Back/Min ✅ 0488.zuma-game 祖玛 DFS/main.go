package main

import (
	"sort"
)

// https://leetcode-cn.com/problems/zuma-game

func findMinStep(board string, hand string) int {
	h := []byte(hand)
	sort.Slice(h, func(i, j int) bool {
		return h[i] < h[j]
	})
	hand = string(h)
	vis := map[[2]string]int{}
	// 置底向上 缓存最小路径 避免重复，
	// 出现相同时 减少递归

	var dfs func(board, hand string) int
	dfs = func(board, hand string) int {
		if board == "" {
			return 0
		}
		key := [2]string{board, hand}
		_, ok := vis[key]
		if !ok {
			min := 6
			for i := range hand {
				if 0 < i && hand[i-1] == hand[i] { // 相同 忽略
					continue
				}
				for j := range board {
					if 0 < j && board[j-1] == hand[i] {
						// 2.与board 相同 只插第一个头部
						continue
					}

					// a.直接插入board旁边 ✅
					// b.直接插入相同board中间 ✅
					// c.直接插入不同board中间 ❌
					if board[j] == hand[i] || 0 < j && board[j-1] == board[j] {
						newBoard := clean(board[:j] + string(hand[i]) + board[j:])
						newHand := hand[:i] + hand[i+1:]
						cur := dfs(newBoard, newHand) + 1
						if cur < min { // 最小路径
							min = cur
						}
					}
				}
			}
			vis[key] = min
		}
		return vis[key]
	}
	res := dfs(board, hand)
	if res == 6 {
		return -1
	}
	return res
}

func clean(board string) string {
	curL := len(board)
	preL := curL + 1
	cur := []byte(board)
	var cnt = 1
	for preL != curL {
		preL = curL
		for i := 1; i < curL; i++ {
			if cur[i-1] == cur[i] {
				cnt++
			} else {
				if 3 <= cnt {
					copy(cur[i-cnt:], cur[i:])
					curL = preL - cnt
					cur = cur[:curL]
					cnt = 1
					break
				}
				cnt = 1
			}
		}
		if 3 <= cnt {
			curL = preL - cnt
			cur = cur[:curL]
			cnt = 1
		}
	}
	return string(cur)
}
