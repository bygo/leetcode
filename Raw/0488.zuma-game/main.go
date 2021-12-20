package main

import (
	"sort"
)

func findMinStep(board string, hand string) int {
	handBuf := []byte(hand)
	sort.Slice(handBuf, func(i, j int) bool {
		return handBuf[i] < handBuf[j]
	})
	hand = string(handBuf)
	vis := map[[2]string]int{}

	var dfs func(board, hand string) int
	dfs = func(board, hand string) int {
		if board == "" {
			return 0
		}
		key := [2]string{board, hand}
		_, ok := vis[key]
		if !ok {
			depCur := 6
			for idxHand := range hand {
				if 0 < idxHand && hand[idxHand-1] == hand[idxHand] {
					continue
				}
				for idxBoard := range board {
					if 0 < idxBoard && board[idxBoard-1] == hand[idxHand] {
						continue
					}

					if board[idxBoard] == hand[idxHand] || 0 < idxBoard && board[idxBoard-1] == board[idxBoard] {
						newBoard := clean(board[:idxBoard] + string(hand[idxHand]) + board[idxBoard:])
						newHand := hand[:idxHand] + hand[idxHand+1:]
						dep := dfs(newBoard, newHand) + 1
						if dep < depCur {
							depCur = dep
						}
					}
				}
			}
			vis[key] = depCur
		}
		return vis[key]
	}

	depMin := dfs(board, hand)
	if depMin == 6 {
		return -1
	}
	return depMin
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
					cur = append(cur[:i-cnt], cur[i:]...)
					curL = preL - cnt
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
