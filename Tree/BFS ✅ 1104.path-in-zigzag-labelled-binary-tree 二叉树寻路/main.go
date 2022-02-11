package main

// https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树寻路

func pathInZigZagTree(label int) []int {
	var nums []int
	// 14 7 3 1
	for 0 < label {
		nums = append(nums, label)
		label >>= 1
	}

	lo, hi := 0, len(nums)-1
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}

	y := len(nums) % 2
	for idx, num := range nums {
		if idx%2 == y {
			// 需要还原成补数的
			nums[idx] = 1<<(idx+1) + 1<<idx - num - 1
		}
	}

	return nums
}
