package main

// https://leetcode.cn/problems/minimum-window-substring

// ❓ 最小覆盖子串
// ⚠️ s的子串 能覆盖 t的所有字符

func minWindow(s string, t string) string {
	chMpCntS := map[byte]int{}
	chMpCntT := map[byte]int{}
	sL := len(s)
	tL := len(t)

	for i := 0; i < tL; i++ {
		ch := t[i]
		chMpCntT[ch]++
	}

	idxLeft, idxRight := 0, 0
	left := -1
	var cntCur int
	var minL = 1<<63 - 1
	for idxRight < sL {
		chRight := s[idxRight]
		if chMpCntS[chRight] < chMpCntT[chRight] {
			// 命中关键字符，并在合理范围
			cntCur++
		}
		chMpCntS[chRight]++

		for cntCur == tL {
			// 计算
			curL := idxRight - idxLeft + 1
			if curL < minL {
				minL = curL
				left = idxLeft
			}
			chLeft := s[idxLeft]
			chMpCntS[chLeft]--
			if chMpCntS[chLeft] < chMpCntT[chLeft] {
				// 减少到关键字符 并在合理范围
				cntCur--
			}
			// left 移动
			idxLeft++
		}

		// right 移动
		idxRight++
	}
	if left == -1 {
		return ""
	}

	return s[left : left+minL]
}
