package main

import (
	"runtime"
)

// https://leetcode-cn.com/problems/permutations

// 插入排序

func main() {
	//ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	//go func() {
	//	select {
	//	case <-ctx.Done():
	//		println(runtime.NumGoroutine())
	//
	//	}
	//}()
	for {
		println(runtime.NumGoroutine())
	}
}
func permute(nums []int) [][]int {
	l := len(nums)
	if l == 0 {
		return nil
	}
	var queue = [][]int{{}}
	var n = 0
	for i := range nums {
		num := nums[i]
		cnt := len(queue)
		for _, q := range queue {
			for j := 0; j <= n; j++ {
				cur := append([]int{}, q[:j]...)
				cur = append(cur, num)
				cur = append(cur, q[j:]...)
				queue = append(queue, cur)
			}
		}
		n++
		queue = queue[cnt:]
	}
	return queue
}
