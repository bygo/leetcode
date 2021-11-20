package main

import "sort"

// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/

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
			// 限定为组合 相似路径剪枝
			if vis[j] || 0 != j && !vis[j-1] && raw[j-1] == raw[j] {
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
func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func nextPermutation(nums []byte) bool {
	n := len(nums)
	i := n - 2
	for 0 <= i && nums[i+1] <= nums[i] { // 后面那个比前面大
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for 0 <= j && nums[j] <= nums[i] {
		j--
	}
	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
	return true
}

func permutationNext(s string) (ans []string) {
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	for {
		ans = append(ans, string(t))
		if !nextPermutation(t) {
			break
		}
	}
	return
}
