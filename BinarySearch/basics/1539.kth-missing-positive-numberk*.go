package main

// https://leetcode.cn/problems/kth-missing-positive-number

// general
func findKthPositive(arr []int, k int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		cur := arr[mid] - mid - 1
		// TODO 索引差
		if cur < k {
			lo = mid + 1
		} else if k <= cur {
			hi = mid
		}
	}
	return k + lo
}

// offset
func findKthPositive(arr []int, k int) int {
	lo, hi := -1, len(arr)-1
	for lo < hi {
		mid := int(uint(lo+hi+1) >> 1)
		cur := arr[mid] - mid - 1
		if cur < k {
			lo = mid
		} else if k <= cur {
			hi = mid - 1
		}
	}
	return k + lo + 1
}

// equal
func findKthPositive(arr []int, k int) int {
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid := int(uint(lo+hi) >> 1)
		cur := arr[mid] - mid - 1
		if cur < k {
			lo = mid + 1
		} else if k <= cur {
			hi = mid - 1
		}
	}
	return k + lo
}
