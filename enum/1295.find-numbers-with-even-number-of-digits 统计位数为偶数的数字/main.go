package main

func findNumbers(nums []int) int {
	var res int
	for _, v := range nums {
		cur := 0
		for 0 < v {
			v /= 10
			cur++
		}
		if cur%2 == 0 {
			res++
		}
	}
	return res
}
