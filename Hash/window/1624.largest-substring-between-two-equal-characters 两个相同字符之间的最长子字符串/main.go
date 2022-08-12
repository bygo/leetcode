package main

// https://leetcode.cn/problems/largest-substring-between-two-equal-characters

// ❓ 两个相同字符之间的最长子字符串

func maxLengthBetweenEqualCharacters(str string) int {
	chMpIdx := map[byte]int{}
	var maxL = -1

	for idxCur := range str {
		chCur := str[idxCur]
		idxPre, ok := chMpIdx[chCur]
		if ok {
			// 存在时计算
			curL := idxCur - idxPre + 1
			if maxL < curL {
				maxL = curL
			}
		} else {
			// 不存在存储
			chMpIdx[chCur] = idxCur
		}
	}
	return maxL
}
