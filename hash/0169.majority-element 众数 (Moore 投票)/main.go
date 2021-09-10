package main

// https://leetcode-cn.com/problems/majority-element

func main() {
	majorityElement([]int{1, 3, 2, 2, 2, 3, 2})
}
func majorityElement(nums []int) int {
	var num, cnt int
	for i := range nums {
		if cnt == 0 {
			num = nums[i]
			cnt = 1
		} else {
			if num == nums[i] {
				cnt++
			} else if num != nums[i] {
				cnt--
			}
		}
	}
	return num
}
