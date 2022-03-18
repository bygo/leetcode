package main

// https://leetcode-cn.com/problems/two-sum-bsts

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 查找两棵二叉搜索树之和

func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	var stack1 []*TreeNode
	var stack2 []*TreeNode

	for root1 != nil || root2 != nil {
		// 入栈
		for root1 != nil {
			stack1 = append(stack1, root1)
			root1 = root1.Left
		}

		for root2 != nil {
			stack2 = append(stack2, root2)
			root2 = root2.Right
		}

		// 出栈
		top1 := len(stack1) - 1
		top2 := len(stack2) - 1
		for 0 <= top1 && 0 <= top2 {
			sumCur := stack1[top1].Val + stack2[top2].Val
			if sumCur == target {
				return true
			} else if sumCur < target {
				root1 = stack1[top1].Right
				stack1 = stack1[:top1]
				top1--
				if root1 != nil {
					break
				}
			} else if target < sumCur {
				root2 = stack2[top2].Left
				stack2 = stack2[:top2]
				top2--
				if root2 != nil {
					break
				}
			}
		}
	}
	return false
}
