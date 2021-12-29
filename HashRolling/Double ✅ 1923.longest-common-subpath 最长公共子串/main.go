package main

// https://leetcode-cn.com/problems/longest-common-subpath

// ❓ 最长公共子串

const (
	base1 = 133020331
	base2 = 13331
	mod1  = 1000000007
	mod2  = 1000000009
)

func longestCommonSubpath(n int, paths [][]int) int {
	// 生成 前缀树
	pathsL := len(paths)
	preHashs := make([][][2]int, pathsL)
	var minL = 1<<63 - 1
	for i := 0; i < pathsL; i++ {
		path := paths[i]
		pathL := len(path)
		if pathL < minL {
			minL = pathL
		}
		preHashs[i] = make([][2]int, pathL+1)
		preHash := preHashs[i]
		for j := 0; j < pathL; j++ {
			bitLow := int(path[j])
			preHash[j+1][0] = (preHash[j][0]*base1 + bitLow) % mod1
			preHash[j+1][1] = (preHash[j][1]*base2 + bitLow) % mod2
		}
	}

	preMul1 := make([]int, minL+1)
	preMul2 := make([]int, minL+1)

	preMul1[0] = 1
	preMul2[0] = 1

	for i := 1; i <= minL; i++ {
		preMul1[i] = preMul1[i-1] * base1 % mod1
		preMul2[i] = preMul2[i-1] * base2 % mod2
	}

	check := func(curL int) bool {
		// nums1 hash值计算
		mul1 := preMul1[curL-1]
		mul2 := preMul2[curL-1]

		preHash := preHashs[0]
		hashCur := preHash[curL]
		hashMpBool := map[[2]int]bool{hashCur: true}
		pathCur := paths[0]
		pathCurL := len(pathCur)
		for idx := 0; idx < pathCurL-curL; idx++ {
			bitHigh := pathCur[idx]
			bitLow := pathCur[idx+curL]
			hashCur[0] = (hashCur[0] - bitHigh*mul1%mod1 + mod1) % mod1 // 减最高
			hashCur[0] = hashCur[0] * base1 % mod1                      // 提升
			hashCur[0] = hashCur[0] + bitLow                            // 加最低
			hashCur[0] %= mod1

			hashCur[1] = (hashCur[1] - bitHigh*mul2%mod2 + mod2) % mod2 // 减最高
			hashCur[1] = hashCur[1] * base2 % mod2                      // 提升
			hashCur[1] = hashCur[1] + bitLow                            // 加最低
			hashCur[1] %= mod2
			hashMpBool[hashCur] = true
		}

		for i := 1; i < pathsL; i++ {
			preHash = preHashs[i]
			hashCur = preHash[curL]
			hashMpBoolCur := map[[2]int]bool{hashCur: true}
			pathCur = paths[i]
			pathCurL = len(pathCur)
			for idx := 0; idx < pathCurL-curL; idx++ {
				bitHigh := pathCur[idx]
				bitLow := pathCur[idx+curL]
				hashCur[0] = (hashCur[0] - bitHigh*mul1%mod1 + mod1) % mod1 // 减最高
				hashCur[0] = hashCur[0] * base1 % mod1                      // 提升
				hashCur[0] = hashCur[0] + bitLow                            // 加最低
				hashCur[0] %= mod1

				hashCur[1] = (hashCur[1] - bitHigh*mul2%mod2 + mod2) % mod2 // 减最高
				hashCur[1] = hashCur[1] * base2 % mod2                      // 提升
				hashCur[1] = hashCur[1] + bitLow                            // 加最低
				hashCur[1] %= mod2
				hashMpBoolCur[hashCur] = true
			}

			for hash := range hashMpBool {
				if !hashMpBoolCur[hash] {
					delete(hashMpBool, hash)
				}
			}
		}
		return 0 < len(hashMpBool)
	}

	// 搜索 合法长度
	left, right := 1, minL+1
	for left < right {
		mid := int(uint(left+right) >> 1)
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}
