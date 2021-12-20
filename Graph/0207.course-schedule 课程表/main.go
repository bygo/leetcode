package main

// https://leetcode-cn.com/problems/course-schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	var edges = make([][]int, numCourses)
	var visited = make([]int, numCourses)
	var res = true
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
	}

	var dfs func(i int)
	dfs = func(i int) {
		visited[i] = 1
		for _, e := range edges[i] {
			if visited[e] == 1 {
				res = false
				return
			} else if visited[e] == 0 {
				dfs(e)
				if !res {
					return
				}
			}
		}
		visited[i] = 2
	}

	for i := 0; i < numCourses && res; i++ {
		dfs(i)
	}
	return res
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	var edges = make([][]int, numCourses)
	var pre = make([]int, numCourses)
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		pre[p[0]]++
	}

	var queue = []int{}

	for i, p := range pre {
		if p == 0 {
			queue = append(queue, i)
		}
	}
	var res []int
	for 0 < len(queue) {
		q := queue[0]
		queue = queue[1:]
		res = append(res, q)
		for _, v := range edges[q] {
			pre[v]--
			if pre[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return len(res) == numCourses
}
