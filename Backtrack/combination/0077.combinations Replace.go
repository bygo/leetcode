package main

// https://leetcode.cn/problems/combinations

// reverse thinking
func combine(numMax int, needL int) [][]int {
	var combNums [][]int
	var nums []int
	for i := 1; i <= needL; i++ {
		nums = append(nums, i) // first combination
	}
	nums = append(nums, numMax+1) // guard

	var idx = 0
	for idx < needL {
		combNums = append(combNums, append([]int{}, nums[:needL]...))
		idx = 0
		for idx < needL && nums[idx]+1 == nums[idx+1] { // increment , guard 6
			nums[idx] = idx + 1 // reset idx(1,2,3)
			idx++
		}
		nums[idx]++

		// all arrive at the tail `idx` over k
	}
	return combNums
}

// 1 2 3  6

// 1 2 4  6
// 1 3 4  6
// 2 3 4  6

// 1 2 5  6
// 1 3 5  6
// 2 3 5  6
// 1 4 5  6
// 2 4 5  6
// 3 4 5  6
