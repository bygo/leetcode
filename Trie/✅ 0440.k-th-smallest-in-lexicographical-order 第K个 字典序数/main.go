package main

// https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order

// ❓ 第K个 字典序数

func findKthNumber(n int, k int) int {
	cnt := 1 // 第一个
	num := 1 // 第一个
	for cnt < k {
		numLeft := num      // 当前
		numRight := num + 1 // 下一个
		cntCur := 0         // 统计

		for numLeft <= n {
			// 查出范围
			cntCur += numRight - numLeft // 统计个数 不包含right 节点 1~2  1-9
			numLeft *= 10
			numRight = min(numRight*10, n+1) // 最大到 n+1  不包含right
		}

		if k < cntCur+cnt {
			// 超过范围
			num *= 10 // 进入子树
			cnt++     // 当前节点
		} else {
			num++         // 进入right 节点
			cnt += cntCur // 当前节点所有节点数 并且统计元素数量
		}
	}
	return num
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
