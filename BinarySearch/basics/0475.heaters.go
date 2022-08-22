package main

import (
	"sort"
)

// https://leetcode.cn/problems/heaters

func findRadius(houses, heaters []int) int {
	// 二分
	sort.Ints(heaters)
	hL := len(heaters)
	var distJust int
	for _, house := range houses {
		// 右边取暖器
		right := sort.SearchInts(heaters, house+1)
		var distMin = 1<<63 - 1
		if right < hL {
			distCur := heaters[right] - house
			if distCur < distMin {
				distMin = distCur
			}
		}

		// 左边取暖器
		left := right - 1
		if 0 <= left {
			distCur := house - heaters[left]
			if distCur < distMin {
				distMin = distCur
			}
		}
		// 当前最小距离 是否大于 合适的距离
		if distJust < distMin {
			distJust = distMin
		}
	}
	return distJust
}

func findRadius(houses, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var idx int
	var distJust int
	for _, house := range houses {
		distCur := abs(house - heaters[idx]) // 当前距离
		for idx+1 < len(heaters) {           // 是否还有下一个取暖器
			distNext := abs(house - heaters[idx+1])
			if distCur < distNext { // 必须更近
				break
			}
			idx++
			distCur = distNext
		}

		if distJust < distCur {
			distJust = distCur
		}
	}
	return distJust
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
