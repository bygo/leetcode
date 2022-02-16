package main

// https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 将有序数组转换为二叉搜索树

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  sortedArrayToBST(nums[:len(nums)/2]), // 偶数左边多一个
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),
	}
}
