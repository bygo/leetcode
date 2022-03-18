package main

// https://leetcode-cn.com/problems/leaf-similar-trees/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 叶子相似树

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var nums []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		if node.Left == nil && node.Right == nil {
			nums = append(nums, node.Val)
		}
	}
	dfs(root1)
	nums1 := append([]int{}, nums...)
	nums = []int{}
	dfs(root2)
	if len(nums1) != len(nums) {
		return false
	}

	for idx := range nums1 {
		if nums[idx] != nums1[idx] {
			return false
		}
	}
	return true
}
