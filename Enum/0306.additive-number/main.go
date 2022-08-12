package main

// https://leetcode.cn/problems/additive-number

func stringAdd(x, y string) string {
	var numStr = []byte{}
	xTop := len(x) - 1
	yTop := len(y) - 1

	carry, numCur := 0, 0
	for 0 <= xTop || 0 <= yTop || carry != 0 {
		numCur = carry
		if 0 <= xTop {
			numCur += int(x[xTop] - '0')
			xTop--
		}
		if 0 <= yTop {
			numCur += int(y[yTop] - '0')
			yTop--
		}
		carry = numCur / 10
		numCur %= 10
		numStr = append(numStr, byte(numCur)+'0')
	}

	var left, right = 0, len(numStr) - 1
	for left < right {
		numStr[left], numStr[right] = numStr[right], numStr[left]
		left++
		right--
	}
	return string(numStr)
}

// 滚动验证
func valid(num string, secondStart, secondEnd int) bool {
	numL := len(num)
	firstStart, firstEnd := 0, secondStart-1
	for secondEnd < numL {
		first := num[firstStart : firstEnd+1]
		second := num[secondStart : secondEnd+1]

		third := stringAdd(first, second)
		thirdStart := secondEnd + 1
		thirdEnd := secondEnd + len(third)
		// 超过 | 不等于
		if numL <= thirdEnd || num[thirdStart:thirdEnd+1] != third {
			break
		}
		// 到达终点
		if thirdEnd == numL-1 {
			return true
		}
		firstStart, firstEnd = secondStart, secondEnd
		secondStart, secondEnd = thirdStart, thirdEnd
	}
	return false
}

func isAdditiveNumber(num string) bool {
	numTop := len(num) - 1
	// 枚举起始位
	for secondStart := 1; secondStart < numTop; secondStart++ {
		// 枚举结束位
		for secondEnd := secondStart; secondEnd < numTop; secondEnd++ {
			// 0 开始的数字，只能使用 0
			if num[secondStart] == '0' && secondStart != secondEnd {
				break
			}
			if valid(num, secondStart, secondEnd) {
				return true
			}
		}
	}
	return false
}
