package main

// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted

func twoSum(nums []int, target int) []int {
	numsL := len(nums)
	for idx := 0; idx < numsL; idx++ {
		lo, hi := idx+1, numsL // idx+1 开始查找
		cur := target - nums[idx]
		for lo < hi {
			mid := int(uint(lo+hi) >> 1)
			if nums[mid] < cur {
				lo = mid + 1
			} else if cur < nums[mid] {
				hi = mid
			} else if nums[mid] == cur {
				return []int{idx + 1, mid + 1}
			}
		}
	}
	return []int{-1, -1}
}

// 双指针
func twoSum_(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum == target {
			return []int{lo + 1, hi + 1}
		} else if sum < target {
			lo++
		} else if target < sum {
			hi--
		}
	}
	return []int{-1, -1}
}
