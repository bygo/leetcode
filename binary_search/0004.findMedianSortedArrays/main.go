package main

/**
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5

解释：假设数组合并后，偶数取中位两个数平均，奇数取中位数
 */

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 {
		nums1, nums2, l1, l2 = nums2, nums1, l2, l1 //保证num1最短，只要计算出num1的k1，就可计算出num2的k2
	}
	var min, max, k1, k2 int
	half := (l1 + l2 + 1) / 2 // 向上取整,保证左边数量永大于右边
	max = l1
	for min <= max {
		k1 = (min + max) / 2 // 二分
		k2 = half - k1       //k2永远等于 half-k1
		if k1 < max && nums2[k2-1] > nums1[k1] { //如果 右2 大于 左1，二分范围min为 k1+1
			min = k1 + 1
		} else if k1 > min && nums1[k1-1] > nums2[k2] { //如果 右1 大于 左2,二分范围max 为k1-1
			max = k1 - 1
		} else { // 右1右1 全部大于左1 左2，计算值
			var leftMax int
			if k1 == 0 {
				leftMax = nums2[k2-1]
			} else if k2 == 0 {
				leftMax = nums1[k1-1]
			} else {
				leftMax = maxInt(nums2[k2-1], nums1[k1-1])
			}
			if (l1+l2)%2 == 1 {
				return float64(leftMax)
			}

			var rightMin int
			if k1 == l1 {
				rightMin = nums2[k2]
			} else if k2 == l2 {
				rightMin = nums1[k1]
			} else {
				rightMin = minInt(nums1[k1], nums2[k2])
			}

			return float64(leftMax+rightMin) / 2
		}
	}
	return 0
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
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
