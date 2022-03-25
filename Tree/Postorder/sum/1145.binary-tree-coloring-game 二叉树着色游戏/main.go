package main

// https://leetcode-cn.com/problems/binary-tree-coloring-game

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ❓ 二叉树着色游戏
// ⚠️ 是否能拿到中心点

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
		cntCur := cntL + cntR + 1
		if node.Val == x {
			// 涂子节点，左右任意 还有一半，x必赢

			// 涂父节点：包含当前节点，小于等于一半（奇数向下取整）x必赢
			if half < cntL || half < cntR || cntCur <= half {
				return 0, true
			}
		}
		return cntCur, false
	}
	_, b := dfs(root)
	return b
}
