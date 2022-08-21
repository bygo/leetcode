package main

import "sort"

// https://leetcode.cn/problems/zi-fu-chuan-de-pai-lie-lcof/

func permutation(s string) []string {
	var res []string
	raw := []byte(s)
	sort.Slice(raw, func(i, j int) bool { return raw[i] < raw[j] })
	l := len(raw)
	perm := make([]byte, l)
	used := make([]bool, l)
	var dfs func(int)
	dfs = func(i int) {
		if i == l {
			res = append(res, string(perm))
			return
		}
		for j := 0; j < l; j++ {
			if used[j] {
				continue
			}
			if 0 < j && !used[j-1] && raw[j-1] == raw[j] {
				continue
			}
			used[j] = true
			perm[i] = raw[j]
			dfs(i + 1)
			used[j] = false
		}
	}
	dfs(0)
	return res
}

// The next bigger number
func permutation(s string) []string {
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
	for 0 <= idx && nums[idx+1] <= nums[idx] { // Find for smaller number
		idx--
	}
	if idx < 0 {
		return false
	}

	// Found the smaller number

	// Find the first value that is larger than the target value
	j := numsL - 1
	for 0 <= j && nums[j] <= nums[idx] {
		j--
	}
	// Exchange
	nums[idx], nums[j] = nums[j], nums[idx]

	// Post bits are all reset back to the minimum bit
	reverse(nums[idx+1:])
	return true
}

// 1 2 3 4
// 1 2 4 3
// 1 3 2 4
//
