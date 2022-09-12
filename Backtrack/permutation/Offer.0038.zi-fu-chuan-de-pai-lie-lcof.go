package main

import "sort"

// https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof/

func permutation(s string) []string {
	var strs []string
	raw := []byte(s)
	sort.Slice(raw, func(i, j int) bool { return raw[i] < raw[j] })
	rawL := len(raw)
	perm := make([]byte, rawL)
	used := make([]bool, rawL)
	var dfs func(int)
	dfs = func(idx int) {
		if idx == rawL {
			strs = append(strs, string(perm))
			return
		}
		for i := 0; i < rawL; i++ {
			if used[i] {
				continue
			}
			if 0 < i && !used[i-1] && raw[i-1] == raw[i] {
				continue
			}
			used[i] = true
			perm[idx] = raw[i]
			dfs(idx + 1)
			used[i] = false
		}
	}
	dfs(0)
	return strs
}

// The next bigger number
func permutation(s string) []string {
	var strs []string
	t := []byte(s)
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	for {
		strs = append(strs, string(t))
		if !nextPermutation(t) {
			break
		}
	}
	return strs
}

func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func nextPermutation(nums []byte) bool {
	numsL := len(nums)
	idx := numsL - 2
	for 0 <= idx && nums[idx+1] <= nums[idx] { // TODO 寻找更小的数，后置单调递增
		idx--
	}
	if idx < 0 {
		return false
	}

	// TODO 找到第一个比目标大的数
	// 15432
	// 21345

	// 123
	// 132

	j := numsL - 1
	for 0 <= j && nums[j] <= nums[idx] {
		j--
	}
	// 交换
	nums[idx], nums[j] = nums[j], nums[idx]

	// 后置全部重置为最小位
	reverse(nums[idx+1:])
	return true
}

// 1 2 3 4
// 1 2 4 3
// 1 3 2 4
//
