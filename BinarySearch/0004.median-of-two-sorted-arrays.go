package main

// https://leetcode.cn/problems/median-of-two-sorted-arrays/

// 整体二分法

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l2 < l1 {
		nums1, nums2, l1, l2 = nums2, nums1, l2, l1 // TODO l1短
	}

	half := (l1 + l2 + 1) >> 1 // 向上取整,保证左边数量永大于右边
	lo, hi := 0, l1
	for lo < hi {
		m1 := int(uint(lo+hi+1) >> 1)
		m2 := half - m1
		if m1 < l1 && nums1[m1] < nums2[m2-1] {
			lo = m1 + 1
		} else if 0 < m1 && nums2[m2] < nums1[m1-1] {
			hi = m1
		} else {
			var leftMax int
			if m1 == 0 {
				leftMax = nums2[m2]
			} else if m2 == 0 {
				leftMax = nums1[m1]
			} else {
				leftMax = max(nums2[m2], nums1[m1])
			}
			if (l1+l2)%2 == 1 {
				return float64(leftMax)
			}

			var rightMin int
			if m1 == l1 {
				rightMin = nums2[m2]
			} else if m2 == l2 {
				rightMin = nums1[m1]
			} else {
				rightMin = min(nums1[m1], nums2[m2])
			}

			return float64(leftMax+rightMin) / 2
		}
	}
	return 0
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// k数二分法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)

	if l%2 == 1 {
		return float64(getKthElement(nums1, nums2, l/2+1))
	} else {
		return float64(getKthElement(nums1, nums2, l/2)+getKthElement(nums1, nums2, l/2+1)) / 2
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	var idx1, idx2 int // todo 偏移1
	numsL1, numsL2 := len(nums1), len(nums2)
	for {
		if idx1 == numsL1 {
			return nums2[idx2-1+k]
		}
		if idx2 == numsL2 {
			return nums1[idx1-1+k]
		}
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}
		half := k / 2
		idxJump1 := min(idx1+half, numsL1) - 1
		idxJump2 := min(idx2+half, numsL2) - 1
		if nums1[idxJump1] < nums2[idxJump2] {
			k -= idxJump1 - (idx1 - 1) // TODO 跳跃了几个位置
			idx1 = idxJump1 + 1        // TODO 下一个位置
		} else {
			k -= idxJump2 - (idx2 - 1)
			idx2 = idxJump2 + 1
		}
	}
}
