package main

// https://leetcode.cn/problems/number-of-ways-to-reconstruct-a-tree

// ❓ 重构一棵树的方案数

func checkWays(pairs [][]int) int {
	edges := map[int]map[int]bool{}

	for _, pair := range pairs {
		if edges[pair[0]] == nil {
			edges[pair[0]] = map[int]bool{}
		}
		edges[pair[0]][pair[1]] = true

		if edges[pair[1]] == nil {
			edges[pair[1]] = map[int]bool{}
		}
		edges[pair[1]][pair[0]] = true
	}
	graphL := len(edges)

	root := 0
	for num, numsNeighbour := range edges {
		if len(numsNeighbour) == graphL-1 {
			root = num
			break
		}
	}
	// 没有找根节点
	if root == 0 {
		return 0
	}

	var cnt = 1

	// 校验
	for num, numsNeighbour := range edges {
		// 根节点忽略
		if num == root {
			continue
		}

		parent := 0

		degreeCur := len(numsNeighbour)
		degreeParent := 1<<63 - 1
		for numNeighbour := range numsNeighbour {
			degree := len(edges[numNeighbour])
			if degreeCur <= degree && degree < degreeParent {
				parent = numNeighbour
				degreeParent = degree
			}
		}
		if parent == 0 {
			return 0
		}
		for numNeighbour := range numsNeighbour {
			if numNeighbour != parent && !edges[parent][numNeighbour] {
				return 0
			}
		}

		if degreeParent == degreeCur {
			cnt = 2
		}
	}
	return cnt
}
