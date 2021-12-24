package main

// https://leetcode-cn.com/problems/maximum-product-of-the-length-of-two-palindromic-substrings

// ❓ 两个奇数回文子字符串长度的最大乘积

const base = 131

func maxProduct(s string) int64 {
	sL := len(s)
	buf := []byte(s)
	var preHashLeft = make([]uint, sL+1)
	var preHashRight = make([]uint, sL+1)
	var preMul = make([]uint, sL+1)
	preMul[0] = 1
	for i := 1; i <= sL; i++ {
		bitLowLeft := uint(s[i-1])
		bitLowRight := uint(s[sL-i])
		preHashLeft[i] = preHashLeft[i-1]*base + bitLowLeft
		preHashRight[i] = preHashRight[i-1]*base + bitLowRight
		preMul[i] = preMul[i-1] * base
	}

	check := func(curL int, buf []byte) bool {
		if curL%2 == 0 {
			return false
		}
		bufL := len(buf)
		hashLeft := preHashLeft[curL]
		hashRight := preHashRight[bufL] - preHashRight[bufL-curL]*preMul[curL]
		mul := preMul[curL-1]
		if hashLeft == hashRight {
			return true
		}
		var bitLow, bitHigh uint
		for i := 0; i < bufL; i++ {
			bitHigh = uint(buf[i])
			bitLow = uint(buf[i+curL])
			hashLeft = hashLeft - bitHigh*mul
			hashLeft *= mul
			hashLeft += bitLow

			hashRight = hashRight - bitHigh
			hashRight += bitLow * mul
			if hashLeft == hashRight {
				return true
			}
		}
		return false
	}

	var numMax int
	for i := 0; i < sL; i++ {
		var left, right = 1, sL + 1
		for left < right {
			mid := int(uint(left+right) >> 1)
			if check(mid, buf[i:]) {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	return int64(numMax)
}
