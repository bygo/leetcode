package main

// https://leetcode-cn.com/problems/combinations

// 逆向
func combine(n int, k int) [][]int {
	var combNums [][]int
	var nums []int
	for i := 1; i <= k; i++ {
		nums = append(nums, i)
	}
	nums = append(nums, n+1) // 哨兵

	var idx = 0
	for idx < k {
		combNums = append(combNums, append([]int{}, nums[:k]...))
		idx = 0
		for idx < k && nums[idx]+1 == nums[idx+1] { // 递增 判断  6(哨兵)
			nums[idx] = idx + 1 // 重置 1 2 3
			idx++
		}
		nums[idx]++

		// 全部到达尾部 idx超过k
	}
	return combNums
}

// 顺序的最后一个+1 ，碰到哨兵结束
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
