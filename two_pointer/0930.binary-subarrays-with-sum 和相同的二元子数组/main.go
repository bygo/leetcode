package main

// https://leetcode-cn.com/problems/binary-subarrays-with-sum

func numSubarraysWithSum(nums []int, goal int) (res int) {
	var l1, l2, sum1, sum2 int
	for r, num := range nums {
		sum1 += num
		for l1 <= r && goal < sum1 {
			sum1 -= nums[l1]
			l1++
		}

		sum2 += num
		for l2 <= r && goal <= sum2 {
			sum2 -= nums[l2]
			l2++
		}
		res += l2 - l1
	}
	return res
}
