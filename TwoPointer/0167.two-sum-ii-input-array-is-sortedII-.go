package TwoPointer

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted

// two pointer
func twoSum(nums []int, target int) []int {
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
