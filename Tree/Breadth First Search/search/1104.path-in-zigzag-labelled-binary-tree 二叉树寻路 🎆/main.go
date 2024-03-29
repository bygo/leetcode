package main

// https://leetcode.cn/problems/path-in-zigzag-labelled-binary-tree

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

	for dep, num := range nums {
		if dep%2 == 0 {
			continue
		}
		// 需要还原成补数的 end - num + start
		start := 1 << (numsL - dep - 1)
		end := 1<<(numsL-dep) - 1
		nums[dep] = end - num + start
	}

	// 正序
	lo, hi := 0, numsL-1
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}

	return nums
}
