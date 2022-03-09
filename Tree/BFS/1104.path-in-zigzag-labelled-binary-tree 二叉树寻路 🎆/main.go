package main

// https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树寻路
// ⚠️ 偶层倒序

// 📚 偶数都在左节点,奇数都在右节点，/2 时 必定找到父节点
// 📚 补数 = end - num + start

func pathInZigZagTree(label int) []int {
	var nums []int
	// 倒序 14 7 3 1
	for 0 < label {
		nums = append(nums, label)
		label >>= 1
	}

	numsL := len(nums)
	// 正序
	lo, hi := 0, numsL-1
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}

	// 奇数开始，偶数需要反排
	// 偶数开始，奇数需要反排
	need := numsL % 2
	// 0 偶数行 正确
	// 1 奇数行 正确
	for dep, num := range nums {
		// 0 => 1 偏移
		if (dep+1)%2 == need {
			continue
		}
		// 需要还原成补数的 end - num + start
		start := 1 << dep
		end := 1<<(dep+1) - 1
		nums[dep] = end - num + start
	}

	return nums
}
