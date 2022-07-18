package main

// https://leetcode.cn/problems/count-nodes-with-the-highest-score

// ❓ 统计最高分的节点数目
// ⚠️ 除去当前节点，所有子树,数量的乘积

func countHighestScoreNodes(parents []int) int {
	pL := len(parents)
	pMpChild := make([][]int, pL)
	for nodeCur := 1; nodeCur < pL; nodeCur++ {
		parent := parents[nodeCur]
		pMpChild[parent] = append(pMpChild[parent], nodeCur)
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
