package main

/**
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/

func nextPermutation(nums []int) {
	left := len(nums) - 2
	if left < 0 {
		return
	}
	for 0 < left && nums[left] >= nums[left+1] { //找小数
		left--
	}

	right := len(nums) - 1
	for left < right && nums[left] >= nums[right] { //找比小数大的数
		right--
	}
	if left != right { //找到交换
		nums[left], nums[right] = nums[right], nums[left]
		left++
	}
	right = len(nums) - 1
	for left < right { //升序
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

/**
思路:理解什么是比当前数字还大的数，然后堆逻辑
例1：123654->124356
例1：123->132

*/
