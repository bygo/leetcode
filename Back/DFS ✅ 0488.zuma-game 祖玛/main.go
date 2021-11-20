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
	var dfs func(curBoard, curHand string) int
	dfs = func(curBoard, curHand string) int {
		if len(curBoard) == 0 {
			return 0
		}
		key := [2]string{curBoard, curHand}
		_, ok := vis[key]
		if !ok {
			cur := 6
			for k := range curHand {
				// hand 相同 只插一次
				if 0 < k && curHand[k-1] == curHand[k] {
					continue
				}
				for j := 0; j < len(curBoard); j++ {
					// board 相同 只插开头
					if 0 < j && curBoard[j-1] == curHand[k] {
						continue
					}

					if j < len(curBoard) && curBoard[j] == curHand[k] ||
						0 < j && curBoard[j-1] == curBoard[j] {
						newBoard := clean(curBoard[:j] + string(curHand[k]) + curBoard[j:])
						newHand := curHand[:k] + curHand[k+1:]
						cur = min(cur, dfs(newBoard, newHand)+1)
					}
				}
			}
			vis[key] = cur
		}
		return vis[key]
	}
	res := dfs(board, hand)
	if res < 6 {
		return res
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func clean(s string) string {
	var b = []byte(s)
	var curL = len(b)
	var preL = curL + 1
	var cnt = 1
	for curL != preL {
		preL = curL
		for i := 1; i < curL; i++ {
			if b[i] == b[i-1] {
				cnt++
			} else {
				if 3 <= cnt {
					curL = preL - cnt
					b = append(b[:i-cnt], b[i:]...)
					cnt = 1
					break
				}
				cnt = 1
			}
		}
		if 3 <= cnt {
			curL = preL - cnt
			b = b[:curL]
			cnt = 1
		}
	}
	return string(b)
}
