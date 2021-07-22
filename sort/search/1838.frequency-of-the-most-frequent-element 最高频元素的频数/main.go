package main

// https://leetcode-cn.com/problems/frequency-of-the-most-frequent-element

func maxFrequency(nums []int, k int) int {
	var counter = make([]int, 100001)
	for _, num := range nums {
		counter[num] ++
	}

	nums = []int{}
	for num, k := range counter {
		for 0 < k {
			nums = append(nums, num)
			k--
		}
	}

	var l1 = len(nums)
	var l, r, total, res = 0, 1, 0, 1
	for r < l1 {
		total += (nums[r] - nums[r-1]) * (r - l)
		for k < total {
			total -= nums[r] - nums[l]
			l++
		}
		if res < r-l+1 {
			res = r - l + 1
		}
		r++
	}
	return res
}
