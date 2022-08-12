package main

// https://leetcode.cn/problems/binary-tree-postorder-traversal/

// ❓ 二叉树的后序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var nums []int
	var min *TreeNode
	for root != nil {
		nums = append(nums, root.Val) //后序遍历
		if root.Right == nil {
			root = root.Left //移动
		} else {
			min = root.Right //寻找右树最小节点
			for min.Left != nil && min.Left != root {
				min = min.Left
			}

			//后序指针处理
			min.Left = root.Left
			root, root.Right = root.Right, nil
		}
	}

	lo, hi := 0, len(nums)-1
	for lo < hi {
		nums[lo], nums[hi] = nums[hi], nums[lo]
		lo++
		hi--
	}
	return nums
}
