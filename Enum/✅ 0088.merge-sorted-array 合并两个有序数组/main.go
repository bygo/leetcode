package main

// https://leetcode.cn/problems/merge-sorted-array

func merge(nums1 []int, m int, nums2 []int, n int) {
	m -= 1
	n -= 1
	for 0 <= n && 0 <= m {
		if nums1[m] < nums2[n] {
			nums1[m+n+1] = nums2[n]
			n--
		} else {
			nums1[m+n+1] = nums1[m]
			m--
		}
	}
	copy(nums1[:n+1], nums2[:n+1])
}
