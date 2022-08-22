package main

// https://leetcode.cn/problems/kth-missing-positive-number

func findKthPositive(arr []int, k int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		cur := arr[mid] - mid - 1
		// 索引差
		if cur < k {
			lo = mid + 1
		} else if k <= cur {
			hi = mid
		}
	}
	return k + lo
}
