package _080_remove_duplicates_from_sorted_array_ii_删除有序数组中的重复项_II

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii

func removeDuplicates(nums []int) int {
	l1 := len(nums)
	if l1 <= 2 {
		return l1
	}
	var slow, fast = 2, 2

	for fast < l1 {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
