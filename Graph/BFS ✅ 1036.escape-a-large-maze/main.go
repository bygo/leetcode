package main

// https://leetcode-cn.com/problems/escape-a-large-maze

var dirs = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func isEscapePossible(block [][]int, source, target []int) bool {
	const (
		blocked = -1
		valid   = 0
		found   = 1
	)

	blockL := len(block)
	if blockL < 2 {
		return true
	}

	blockMpBool := map[[2]int]bool{}
	for _, b := range block {
		blockMpBool[[2]int{b[0], b[1]}] = true
	}

	check := func(s, t []int) int {
		sx, sy := s[0], s[1]
		tx, ty := t[0], t[1]
		// 最大包围圈
		cntMax := blockL * (blockL - 1) / 2

		que := [][2]int{{sx, sy}}
		visMpBool := map[[2]int]bool{{sx, sy}: true}
		for point := range blockMpBool {
			visMpBool[point] = true
		}
		for 0 < cntMax {
			cnt := len(que)
			if cnt == 0 {
				break
			}
			for _, point := range que {
				for _, dir := range dirs {
					x, y := point[0]+dir[0], point[1]+dir[1]
					pointNew := [2]int{x, y}
					if 0 <= x && x < 1e6 && 0 <= y && y < 1e6 && !visMpBool[pointNew] {
						if x == tx && y == ty {
							return found
						}
						cntMax--
						visMpBool[pointNew] = true
						que = append(que, pointNew)
					}
				}
			}
			que = que[cnt:]
		}

		// 没有突破包围圈
		if 0 < cntMax {
			return blocked
		}
		return valid
	}

	sourceToTarget := check(source, target)
	if sourceToTarget == found {
		return true
	}
	targetToSource := check(target, source)
	if targetToSource == sourceToTarget && targetToSource != blocked {
		return true
	}
	return false
}
