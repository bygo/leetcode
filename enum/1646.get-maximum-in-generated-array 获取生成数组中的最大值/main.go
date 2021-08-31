package _646_get_maximum_in_generated_array_获取生成数组中的最大值

// https://leetcode-cn.com/problems/get-maximum-in-generated-array

func getMaximumGenerated(n int) int {
	var res int
	if n == 0 {
		return res
	}
	nums := make([]int, n+1)
	nums[1] = 1
	for i := 2; i <= n; i++ {
		nums[i] = nums[i/2] + i%2*nums[i/2+1]
	}
	for _, v := range nums {
		if res < v {
			res = v
		}
	}
	return res
}
