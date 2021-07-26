package main

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l2 < l1 {
		nums1, nums2, l1, l2 = nums2, nums1, l2, l1 //保证num1最短？ 只要计算出num1的k1，就可计算出num2的k2
	}
	var left, right, m1, m2, half int
	half = (l1 + l2 + 1) / 2 // 向上取整,保证左边数量永大于右边
	right = l1
	for left <= right {
		m1 = (left + right) / 2                    // 二分
		m2 = half - m1                             //k2永远等于 half-m1
		if m1 < right && nums1[m1] < nums2[m2-1] { //如果 右2 大于 左1，二分范围min为 m1+1
			left = m1 + 1
		} else if left < m1 && nums2[m2] < nums1[m1-1] { //如果 右1 大于 左2,二分范围max 为m1-1
			right = m1 - 1
		} else { // 右1右2 全部大于左1 左2，计算值
			var leftMax int
			if m1 == 0 {
				leftMax = nums2[m2-1]
			} else if m2 == 0 {
				leftMax = nums1[m1-1]
			} else {
				leftMax = max(nums2[m2-1], nums1[m1-1])
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

/**
思路：
1.计算出两个数组中位数切口half，并且保证 left1+left2 = right1+right2 (+1)
2.保证right1,right2里的所有元素全部大于left1,left2的所有元素
  如果条件 left1.max > right2.min 或者 left2.max > right1.min 成立时，不断二分，重新计算 num1 切口，直至不再满足条件

备注：
left1  + right1 = num1
left2  + right2 = num2
*/
