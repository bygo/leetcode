package main

// https://leetcode-cn.com/problems/pancake-sorting

// ❓ 煎饼排序

func pancakeSort(arr []int) []int {
	var nums []int

	for top := len(arr) - 1; 0 < top; top-- {
		idx := 0
		// 寻找最大值
		for i, num := range arr[:top+1] {
			if arr[idx] < num {
				idx = i
			}
		}
		// 最大值刚好在指定位置
		if idx == top {
			continue
		}
		// 将最大值反转到第一位
		lo, hi := 0, idx
		for lo < hi {
			arr[lo], arr[hi] = arr[hi], arr[lo]
			lo++
			hi--
		}

		// 将最大值反转到最后一位
		lo, hi = 0, top
		for lo < hi {
			arr[lo], arr[hi] = arr[hi], arr[lo]
			lo++
			hi--
		}
		// 2次反转的k值
		nums = append(nums, idx+1, top+1)
	}
	return nums
}
