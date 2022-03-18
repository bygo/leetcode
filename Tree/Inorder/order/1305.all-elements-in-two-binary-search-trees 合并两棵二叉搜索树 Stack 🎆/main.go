package main

// https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 两棵二叉搜索树中的所有元素

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var stack1 []*TreeNode
	var stack2 []*TreeNode
	var nums []int
	for root1 != nil || root2 != nil {
		for root1 != nil {
			stack1 = append(stack1, root1)
			root1 = root1.Left
		}

		for root2 != nil {
			stack2 = append(stack2, root2)
			root2 = root2.Left
		}

		top1 := len(stack1) - 1
		top2 := len(stack2) - 1
		for 0 <= top1 || 0 <= top2 {
			if 0 <= top1 && 0 <= top2 {
				if stack1[top1].Val <= stack2[top2].Val {
					nums = append(nums, stack1[top1].Val)
					root1 = stack1[top1].Right
					stack1 = stack1[:top1]
					if root1 != nil {
						break
					}
					top1--
				} else {
					nums = append(nums, stack2[top2].Val)
					root2 = stack2[top2].Right
					stack2 = stack2[:top2]
					if root2 != nil {
						break
					}
					top2--
				}
			} else if 0 <= top1 {
				nums = append(nums, stack1[top1].Val)
				root1 = stack1[top1].Right
				stack1 = stack1[:top1]
				if root1 != nil {
					break
				}
				top1--
			} else if 0 <= top2 {
				nums = append(nums, stack2[top2].Val)
				root2 = stack2[top2].Right
				stack2 = stack2[:top2]
				if root2 != nil {
					break
				}
				top2--
			}
		}
	}
	return nums
}
