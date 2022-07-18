package main

// https://leetcode.cn/problems/maximum-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 最大二叉树

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var idxMax, numMax int
	for idx, num := range nums {
		if numMax <= num {
			idxMax = idx
			numMax = num
		}
	}

	return &TreeNode{
		Val:   numMax, // 最大值为根节点
		Left:  constructMaximumBinaryTree(nums[:idxMax]),
		Right: constructMaximumBinaryTree(nums[idxMax+1:]),
	}
}
