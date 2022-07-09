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
						if len(boardNew) == 0 {
							return dep
						}
						handNew := handCur[:i] + handCur[i+1:]
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
