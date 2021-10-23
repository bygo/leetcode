package main

// https://leetcode-cn.com/problems/permutations

func permute(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	if n == 0 {
		return res
	} else if n == 1 {
		return append(res, nums)
	}
	res = [][]int{
		{nums[0], nums[1]},
		{nums[1], nums[0]},
	}

	for k := 2; k < n; k++ {
		num := nums[k]
		resLen := len(res)
		for i := 0; i < resLen; i++ {
			for j := 0; j <= len(res[i]); j++ {
				tmp := []int{}
				if j == 0 {
					tmp = append(append(tmp, num), res[i]...)
				} else {
					t := make([]int, len(res[i][:j]))
					copy(t, res[i][:j])
					tmp = append(append(t, num), res[i][j:]...)
				}
				res = append(res, tmp)
			}
		}
		res = res[resLen:]
	}

	return res
}
