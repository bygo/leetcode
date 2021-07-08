package main

// Link: https://leetcode-cn.com/problems/binary-subarrays-with-sum

// 前缀
func numSubarraysWithSum(nums []int, goal int) (res int) {
	cnt := map[int]int{}
	var sum int
	for _, num := range nums {
		cnt[sum]++
		sum += num
		res += cnt[sum-goal]
	}
	return res
}

// two ptr 1,0,1,0,1
func numSubarraysWithSum(nums []int, goal int) (res int) {
	var l1, l2, sum1, sum2 int
	for right, num := range nums {
		sum1 += num
		for l1 <= right && goal < sum1 {
			sum1 -= nums[l1]
			l1++
		}

		sum2 += num
		for l2 <= right && goal <= sum2 {
			sum2 -= nums[l2]
			l2++
		}
		res += l2 - l1
	}
	return res
}
