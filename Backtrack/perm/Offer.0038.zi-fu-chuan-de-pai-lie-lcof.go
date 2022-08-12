package main

import "sort"

// https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof/

// 排列

func permutation(s string) []string {
	var res []string
	raw := []byte(s)
	sort.Slice(raw, func(i, j int) bool { return raw[i] < raw[j] })
	l := len(raw)
	perm := make([]byte, l)
	vis := make([]bool, l)
	var dfs func(int)
	dfs = func(i int) {
		if i == l {
			res = append(res, string(perm))
			return
		}
		for j := 0; j < l; j++ {
			// 已使用
			// 相似路径剪枝 限定为组合
			if vis[j] || 0 < j && !vis[j-1] && raw[j-1] == raw[j] {
				continue
			}
			vis[j] = true
			perm[i] = raw[j]
			dfs(i + 1)
			vis[j] = false
		}
	}
	dfs(0)
	return res
}

// 下一个更大的数
func permutationNext(s string) []string {
	var res []string
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	for {
		res = append(res, string(t))
		if !nextPermutation(t) {
			break
		}
	}
	return res
}

func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func nextPermutation(nums []byte) bool {
	numsL := len(nums)
	idx := numsL - 2
	for 0 <= idx && nums[idx+1] <= nums[idx] { // 寻找比较小的前置位
		idx--
	}
	if idx < 0 {
		return false
	}

	// 找到比较小的前置位

	// 寻找第一个比目标值大的值
	j := numsL - 1
	for 0 <= j && nums[j] <= nums[idx] {
		j--
	}
	// 交换
	nums[idx], nums[j] = nums[j], nums[idx]

	// 后置位全部重置 回归最小位
	reverse(nums[idx+1:])
	return true
}

// 1 2 3 4
// 1 2 4 3
// 1 3 2 4
//
