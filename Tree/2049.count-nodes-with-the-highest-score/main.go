package main

// https://leetcode-cn.com/problems/count-nodes-with-the-highest-score

// ❓ 查找两棵二叉搜索树之和

func countHighestScoreNodes(parents []int) int {
	pL := len(parents)
	pMpChild := make([][]int, pL)
	for node := 1; node < pL; node++ {
		p := parents[node]
		pMpChild[p] = append(pMpChild[p], node)
	}

	proMax := 0
	cntRes := 0
	var dfs func(node int) int
	dfs = func(node int) int {
		var pro, cnt = 1, 0
		for _, child := range pMpChild[node] {
			// 子节点数
			cntChild := dfs(child)
			pro *= cntChild
			cnt += cntChild
		}
		if 0 < node { // 除了根节点,都乘以父节点数
			pro *= pL - cnt - 1
		}
		if pro == proMax {
			cntRes++
		} else if proMax < pro {
			proMax = pro
			cntRes = 1
		}

		// 总节点数
		return cnt + 1
	}
	dfs(0)
	return cntRes
}
