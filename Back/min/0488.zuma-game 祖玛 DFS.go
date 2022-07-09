package main

import (
	"sort"
)

// https://leetcode-cn.com/problems/zuma-game

func findMinStep(board string, hand string) int {
	bytesHand := []byte(hand)
	sort.Slice(bytesHand, func(i, j int) bool {
		return bytesHand[i] < bytesHand[j]
	})
	hand = string(bytesHand)
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
		if ok {
			return vis[key]
		}
		depMin := 6 // 1 <= hand <= 5
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
					depCur := dfs(newBoard, newHand) + 1
					if depCur < depMin { // 最小路径
						depMin = depCur
					}
				}
			}
		}
		vis[key] = depMin
		return vis[key]
	}
	dep := dfs(board, hand)
	if dep == 6 {
		return -1
	}
	return dep
}

func clean(board string) string {
	curL := len(board)
	preL := curL + 1

	buf := []byte(board)
	for curL != preL {
		preL = curL
		cnt := 1
		for idx := 1; idx < curL; idx++ {
			if buf[idx-1] == buf[idx] {
				cnt++
			} else {
				if 3 <= cnt {
					copy(buf[idx-cnt:], buf[idx:])
					idx = idx - cnt
					curL = curL - cnt
				}
				cnt = 1
			}

		}

		if 3 <= cnt {
			curL = curL - cnt
			buf = buf[:curL]
		}
	}
	return string(buf[:curL])
}
