package main

import (
	"sort"
)

// https://leetcode.cn/problems/zuma-game

// DFS
func findMinStep(board string, hand string) int {
	bytesHand := []byte(hand)
	sort.Slice(bytesHand, func(i, j int) bool {
		return bytesHand[i] < bytesHand[j]
	})
	hand = string(bytesHand)
	used := map[[2]string]int{}
	// Bottom up, cache minimum path to avoid duplication, ring

	var dfs func(board, hand string) int
	dfs = func(board, hand string) int {
		if board == "" {
			return 0
		}
		key := [2]string{board, hand}
		_, ok := used[key]
		if ok {
			return used[key]
		}
		depMin := 6 // 1 <= hand <= 5
		for i := range hand {
			if 0 < i && hand[i-1] == hand[i] {
				// 1.Hand is inserted only once, because the path is the same
				continue
			}
			for j := range board {
				if 0 < j && board[j-1] == hand[i] {
					// 2.The first head, same as board, because the path is the same
					continue
				}

				// a.next to board ✅
				// b.the middle of the same board ✅
				// c.between different boards ❌
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
		used[key] = depMin
		return used[key]
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

// BFS
func findMinStep(board string, hand string) int {
	bytesHand := []byte(hand)
	sort.Slice(bytesHand, func(i, j int) bool {
		return bytesHand[i] < bytesHand[j]
	})
	hand = string(bytesHand)
	used := map[[2]string]bool{}
	// Up bottom, cache minimum path to avoid duplication
	// Because the 'hand' will decrease, it won't form a ring

	key := [2]string{board, hand}
	used[key] = true
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
				// 1.Hand is inserted only once, because the path is the same
				if 0 < i && handCur[i-1] == handCur[i] {
					continue
				}
				for j := range boardCur {
					// 2.The first head, same as board, because the path is the same
					if 0 < j && boardCur[j-1] == handCur[i] {
						continue
					}

					// a.next to board ✅
					// b.the middle of the same board ✅
					// c.between different boards ❌
					if boardCur[j] == handCur[i] || 0 < j && boardCur[j-1] == boardCur[j] {
						boardRaw := boardCur[:j] + string(handCur[i]) + boardCur[j:]
						boardNew := clean(boardRaw)
						if len(boardNew) == 0 {
							return dep
						}
						handNew := handCur[:i] + handCur[i+1:]
						str := [2]string{boardNew, handNew}
						if !used[str] {
							que = append(que, str)
						}
						used[str] = true
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
