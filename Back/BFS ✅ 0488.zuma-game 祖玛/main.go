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
	var queue = [][2]string{{board, hand}}

	vis := map[[2]string]bool{}
	vis[[2]string{board, hand}] = true
	var dep int
	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		dep++

		for i := 0; i < cnt; i++ {
			q := queue[i]
			curBoard := q[0]
			curHand := q[1]
			for k := range curHand {
				// 1.hand 相同 只插一次
				if 0 < k && curHand[k-1] == curHand[k] {
					continue
				}
				l1 := len(curBoard)
				for j := 0; j < l1; j++ {
					// 2.board 相同 只插第一个头部
					if 0 < j && curBoard[j-1] == curHand[k] {
						continue
					}

					// a.直接插入board旁边 ✅
					// b.直接插入相同board中间 ✅
					// c.直接插入不同board中间 ❌
					if curBoard[j] == curHand[k] ||
						0 < j && j < len(curBoard) &&
							curBoard[j-1] == curBoard[j] {
						newBoard := clean(curBoard[:j] + string(curHand[k]) + curBoard[j:])
						newHand := curHand[:k] + curHand[k+1:]
						if len(newBoard) == 0 {
							return dep
						}
						str := [2]string{newBoard, newHand}
						if !vis[str] {
							queue = append(queue, str)
						}
						vis[str] = true
					}
				}
			}
		}
		queue = queue[cnt:]
	}
	return -1
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
