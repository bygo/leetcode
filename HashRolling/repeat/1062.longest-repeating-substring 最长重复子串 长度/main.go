package main

// https://leetcode.cn/problems/longest-repeating-substring

// ❓ 最长重复子串 的长度
const base = 131

func longestRepeatingSubstring(s string) int {
	sL := len(s)
	var longestL int
	var preHash = make([]uint, sL+1)
	var preMul = make([]uint, sL+1)
	preMul[0] = 1
	for i := 1; i <= sL; i++ {
		bitLo := uint(s[i-1])
		preHash[i] = preHash[i-1]*base + bitLo
		preMul[i] = preMul[i-1] * base
	}

	check := func(curL int) bool {
		// 初始化
		var hash = preHash[curL]
		var mul = preMul[curL-1]

		hashMpState := map[uint]int8{hash: 1}

		for i := 0; i < sL-curL; i++ {
			bitHi := uint(s[i])
			bitLo := uint(s[i+curL])
			hash = hash - bitHi*mul // 减最高
			hash = hash * base      // 提升
			hash = hash + bitLo     // 加最低

			if hashMpState[hash] < 2 {
				hashMpState[hash]++
				if hashMpState[hash] == 2 {
					longestL = curL // 下一个
					return true
				}
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
	return longestL
}
