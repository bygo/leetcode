package main

// https://leetcode-cn.com/problems/binary-tree-coloring-game

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树着色游戏

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	half := n / 2
	var dfs func(node *TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, false
		}
		cntL, bL := dfs(node.Left)
		if bL {
			return 0, bL
		}
		cntR, bR := dfs(node.Right)
		if bR {
			return 0, bR
		}
		if node.Val == x {
			// 剩下的 还有一半，y必赢
			// 或者选中小于等于一半，x必输
			if half < cntL || half < cntR || cntL+cntR+1 <= half {
				return 0, true
			}
		}
		return cntL + cntR + 1, false
	}
	_, b := dfs(root)
	return b
}
