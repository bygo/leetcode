package main

import "sort"

func findMinStep(board string, hand string) int {
	h := []byte(hand)
	sort.Slice(h, func(i, j int) bool {
		return h[i] < h[j]
	})
	hand = string(h)
	vis := map[[2]string]int{}

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
				if 0 < i && hand[i-1] == hand[i] {
					continue
				}
				for j := range board {
					if 0 < j && board[j-1] == hand[i] {
						continue
					}

					if board[j] == hand[i] || 0 < j && board[j-1] == board[j] {
						newBoard := clean(board[:j] + string(hand[i]) + board[j:])
						newHand := hand[:i] + hand[i+1:]
						cur := dfs(newBoard, newHand) + 1
						if cur < min {
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
