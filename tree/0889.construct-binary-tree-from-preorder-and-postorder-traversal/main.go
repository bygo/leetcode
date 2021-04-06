package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	if len(pre) == 1 {
		return &TreeNode{Val: pre[0]}
	}
	for i := range post {
		if post[i] == pre[1] {
			return &TreeNode{
				Val:   pre[0],
				Left:  constructFromPrePost(pre[1:i+2], post[:i+1]),
				Right: constructFromPrePost(pre[i+2:], post[i+1:len(post)-1]),
			}

		}
	}
	return nil
}