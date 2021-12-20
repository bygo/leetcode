package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Title: Find Nearest Right Node in Binary Tree

func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var queue = []*TreeNode{root}

	for {
		counter := len(queue)
		if counter == 0 {
			break
		}

		for k, q := range queue[:counter] {
			if q == u {
				if k == counter-1 {
					return nil
				} else {
					return queue[k+1]
				}
			}
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[counter:]
	}
	return nil
}
