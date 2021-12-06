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
		for j < k && cur[j]+1 == cur[j+1] { // 递增重置
			println()
			cur[j] = j + 1
			j++
		}
		cur[j] ++
	}
	return res
}
