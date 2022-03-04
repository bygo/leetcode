package main

// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

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
		if root.Right == nil {
			nums = append(nums, root.Val) //后序遍历
			root = root.Left              //移动
		} else {
			min = root.Right //寻找右树最小节点
			for min.Left != nil && min.Left != root.Left {
				min = min.Left
			}

			if min.Left == nil {
				nums = append(nums, root.Val) //后序遍历
				min.Left = root.Left
				root = root.Right
			} else {
				root = root.Left //跳跃
				min.Left = nil
			}
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
