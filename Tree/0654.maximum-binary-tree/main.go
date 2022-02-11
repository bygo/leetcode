package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var maxK, maxV = -1, 0
	for k, v := range nums {
		if maxV <= v {
			maxK = k
			maxV = v
		}
	}

	if maxK == -1 {
		return nil
	}

	return &TreeNode{
		Val:   maxV,
		Left:  constructMaximumBinaryTree(nums[:maxK]),
		Right: constructMaximumBinaryTree(nums[maxK+1:]),
	}
}
