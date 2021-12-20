package main

import (
	"sort"
)

// https://leetcode-cn.com/problems/heaters

// ❓ houses 能被 heaters 覆盖
// ⚠️ 均无序

func findRadius(houses, heaters []int) int {
	// 二分
	sort.Ints(heaters)
	heatersL := len(heaters)
	var distJust int
	for _, house := range houses {
		// 右边取暖器
		right := sort.SearchInts(heaters, house+1) // left < right  heaters[left] < house+1 <= heaters[right]
		var distMin = 1<<63 - 1
		if right < heatersL {
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

// 双指针

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findRadius(houses, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	var idxHeaters int
	var distJust int
	for _, house := range houses {
		distCur := abs(house - heaters[idxHeaters])
		for idxHeaters+1 < len(heaters) {
			distNext := abs(house - heaters[idxHeaters+1])
			if distNext <= distCur {
				idxHeaters++
				distCur = distNext
			} else {
				break
			}
		}

		if distJust < distCur {
			distJust = distCur
		}
	}
	return distJust
}
