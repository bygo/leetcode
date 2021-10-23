package main

// https://leetcode-cn.com/problems/minimum-height-trees

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	var degree = make([]int, n)
	var graph = make([][]int, n)
	for _, e := range edges {
		i, j := e[0], e[1]
		degree[i]++
		degree[j]++
		graph[i] = append(graph[i], j)
		graph[j] = append(graph[j], i)
	}

	var queue = []int{}
	for i, v := range degree {
		if v == 1 {
			queue = append(queue, i)
		}
	}

	var res []int
	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		for i := 0; i < cnt; i++ {
			q := queue[i]
			degree[q]--
			for _, v := range graph[q] {
				degree[v]--
				if degree[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
		res = make([]int, cnt)
		copy(res, queue[:cnt])
		queue = queue[cnt:]
	}
	return res
}
