package main

// https://leetcode-cn.com/problems/longest-duplicate-substring

// ❓ 最长重复子串

const base = 131

func longestDupSubstring(s string) string {
	sL := len(s)
	var strLongest string
	var preHash = make([]uint, sL+1)
	var preMul = make([]uint, sL+1)
	preMul[0] = 1
	for i := 1; i <= sL; i++ {
		bitLow := uint(s[i-1])
		preHash[i] = preHash[i-1]*base + bitLow
		preMul[i] = preMul[i-1] * base
	}

	check := func(curL int) bool {
		// 初始化
		hash := preHash[curL]
		mul := preMul[curL-1]

		hashMpState := map[uint]int8{hash: 1}

		// 当前最高位系数
		for i := curL; i < sL; i++ {
			bitHigh := uint(s[i-curL])
			bitLow := uint(s[i])
			hash = hash - bitHigh*mul // 减最高
			hash = hash * base        // 提升
			hash = hash + bitLow      // 加最低
			hashMpState[hash]++

			if hashMpState[hash] == 2 {
				strLongest = s[i-curL+1 : i+1] // 下一个
				return true
			}
		}
		return false
	}

	// 二分
	left, right := 1, sL+1
	for left < right {
		mid := int((uint(left + right)) >> 1)
		if check(mid) { // 合法时继续寻找更大的值
			left = mid + 1
		} else {
			right = mid
		}
	}
	return strLongest
}
