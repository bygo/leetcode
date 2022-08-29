package Array

// https://leetcode.cn/problems/shuffle-the-array

func shuffle(nums []int, n int) []int {
	var res = make([]int, 0, n*2)
	for idx, num := range nums[:n] {
		res = append(res, num)
		res = append(res, nums[n+idx])

	}
	return res
}
