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
	vis := map[[2]string]bool{}
	// 置顶向下 缓存最小路径 避免环重复
	// 因为hand的会减少，所以不会形成环

	key := [2]string{board, hand}
	vis[key] = true
	que := [][2]string{key}
	var dep int
	for {
		queL := len(que)
		if queL == 0 {
			break
		}

		dep++
		for _, q := range que {
			boardCur := q[0]
			handCur := q[1]
			for i := range handCur {
				// 1.hand 相同 只插一次,因为路径相同
				if 0 < i && handCur[i-1] == handCur[i] {
					continue
				}
				for j := range boardCur {
					// 2.与board 相同 只插第一个头部,因为路径相同
					if 0 < j && boardCur[j-1] == handCur[i] {
						continue
					}

					// a.直接插入board旁边 ✅
					// b.直接插入相同board中间 ✅
					// c.直接插入不同board中间 ❌
					if boardCur[j] == handCur[i] || 0 < j && boardCur[j-1] == boardCur[j] {
						boardRaw := boardCur[:j] + string(handCur[i]) + boardCur[j:]
						boardNew := clean(boardRaw)
						handNew := handCur[:i] + handCur[i+1:]
						if len(boardNew) == 0 {
							return dep
						}
						str := [2]string{boardNew, handNew}
						if !vis[str] {
							que = append(que, str)
						}
						vis[str] = true
					}
				}
			}
		}
		que = que[queL:]
	}
	return -1
}

func clean(board string) string {
	curL := len(board)
	preL := curL + 1
	cur := []byte(board)
	var cnt int
	for preL != curL {
		preL = curL
		cnt = 1
		for idx := 1; idx < curL; idx++ {
			if cur[idx-1] == cur[idx] {
				cnt++
			} else {
				// 超过三个相同,消除
				if 3 <= cnt {
					copy(cur[idx-cnt:], cur[idx:])
					curL = preL - cnt
					cur = cur[:curL]
					break
				}
				cnt = 1
			}
		}
		// 超过三个相同,消除
		if 3 <= cnt {
			curL = preL - cnt
			cur = cur[:curL]
		}
	}
	return string(cur)
}
