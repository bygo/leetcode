package main

// https://leetcode-cn.com/problems/permutations

// 插入排序

func permute(nums []int) [][]int {
	numsL := len(nums)
	if numsL == 0 {
		return nil
	}
	var que = [][]int{{}}
	var n = 0
	for _, num := range nums {
		queL := len(que)
		for _, q := range que {
			for j := 0; j <= n; j++ {
				numCur := append([]int{}, q[:j]...)
				numCur = append(numCur, num)
				numCur = append(numCur, q[j:]...)

				que = append(que, numCur)
			}
		}
		// 1
		// 1 2
		// 2 1
		// 3 1 2
		// 1 3 2
		// 1 2 3

		// 3 2 1
		// 2 3 1
		// 2 1 3
		n++
		que = que[queL:]
	}
	return que
}
