package main

// https://leetcode-cn.com/problems/distinct-echo-substrings

// ❓ 不同的循环子字符串 总数

const base uint = 131

func distinctEchoSubstrings(text string) int {
	textL := len(text)
	preHash := make([]uint, textL+1)
	preMul := make([]uint, textL+1)
	preMul[0] = 1

	// 初始化
	for curL := 1; curL <= textL; curL++ {
		bitLow := uint(text[curL-1])
		preHash[curL] = preHash[curL-1]*base + bitLow
		preMul[curL] = preMul[curL-1] * base
	}

	// 穷举
	hash := func(idx, curL int) uint {
		// right 为开区间
		return preHash[idx+curL] - preHash[idx]*preMul[curL] // 减去多乘  curL 阶
	}

	hashMpBool := map[uint]bool{}

	for idx1 := 0; idx1 < textL; idx1++ { // 第一串起始位
		idx2 := idx1 + 1 // 第二串起始位
		for {
			curL := idx2 - idx1 // 距离等于长度
			if textL < idx2+curL {
				break
			}
			hashLeft := hash(idx1, curL) // 前置hash
			// 不存在时
			if !hashMpBool[hashLeft] {
				hashRight := hash(idx2, curL)
				// 相等，构成循环字符串
				if hashLeft == hashRight {
					hashMpBool[hashLeft] = true
				}
			}
			idx2++
		}
	}
	return len(hashMpBool)
}
