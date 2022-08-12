package main

func decompressRLElist(nums []int) []int {
	var res []int
	l := len(nums)
	i := 0
	for i < l {
		for 0 < nums[i] {
			nums[i] -= 1
			res = append(res, nums[i+1])
		}
		i += 2
	}
	return res
}
