package main

import (
	"sort"
	"strconv"
)

// https://leetcode-cn.com/problems/reordered-power-of-2

// ❓ 重新排序得到 2 的幂

func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0
}

func combine(n int) bool {
	nums := []byte(strconv.Itoa(n))
	numsL := len(nums)
	var cur []int
	for i := 1; i <= numsL; i++ {
		cur = append(cur, i)
	}
	cur = append(cur, n+1) // 哨兵

	var j = 0
	for j < numsL {

	}
	return false
}

func reorderedPowerOf2(n int) bool {
	nums := []byte(strconv.Itoa(n))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	numsL := len(nums)
	vis := make([]bool, numsL)
	var dfs func(cnt, num int) bool
	dfs = func(cnt, num int) bool {
		if cnt == numsL {
			return isPowerOfTwo(num)
		}
		for i := range nums {
			ch := nums[i]
			if vis[i] {
				// 使用过
				continue
			}

			if num == 0 && ch == '0' {
				// 0 开头
				continue
			}

			if 0 < i && ch == nums[i-1] && !vis[i-1] {
				//组合  位置相同且同一层 = 相同问题
				continue
			}

			vis[i] = true
			if dfs(cnt+1, num*10+int(ch-'0')) {
				return true
			}
			vis[i] = false
		}
		return false
	}
	return dfs(0, 0)
}
