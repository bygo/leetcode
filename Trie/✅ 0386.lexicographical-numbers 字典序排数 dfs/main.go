package main

// https://leetcode-cn.com/problems/lexicographical-numbers

// ❓ 字典序排数
// ⚠️ 乘以10进入下层 除以10返回上层

func lexicalOrder(n int) []int {
	var numsOrder []int
	var dfs func(num int)
	dfs = func(num int) {
		if n < num {
			return
		}
		numsOrder = append(numsOrder, num)
		for i := 0; i < 10; i++ {
			dfs(num*10 + i)

		}
	}
	for i := 1; i < 10; i++ {
		dfs(i)
	}
	return numsOrder
}
