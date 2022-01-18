package main

import "fmt"

// https://leetcode-cn.com/problems/combinations

func main() {
	combine(5, 3)
}

// 逆向
func combine(n int, k int) [][]int {
	var res [][]int
	var cur []int
	for i := 1; i <= k; i++ {
		cur = append(cur, i)
	}
	cur = append(cur, n+1) // 哨兵

	var j = 0
	for j < k {
		fmt.Printf("%+v", cur)
		res = append(res, append([]int{}, cur[:k]...))
		j = 0
		for j < k && cur[j]+1 == cur[j+1] { // 递增重置 1 2 3  6(哨兵)
			cur[j] = j + 1
			j++
		}

		// 最后一个递增 +1
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

		// 全部到达尾部 j超过k
		cur[j] ++
	}
	return res
}
