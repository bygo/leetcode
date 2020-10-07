package main

func sortColors(nums []int) {
	var l int
	r := len(nums) - 1
	for k := 0; k <= r; k++ {
		if nums[k] == 0 {
			nums[l], nums[k] = nums[k], nums[l]
			l++
		} else if nums[k] == 2 {
			nums[k], nums[r] = nums[r], nums[k]
			r--
			k-- //当前位置继续循环
		}
	}
}
