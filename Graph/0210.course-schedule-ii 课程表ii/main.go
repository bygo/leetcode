package main

// https://leetcode-cn.com/problems/course-schedule-ii

func findOrder(numCourses int, prerequisites [][]int) []int {
	var edges = make([][]int, numCourses)
	var visited = make([]int, numCourses)
	var valid = true

	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
	}
	var dfs func(i int)
	var res []int
	dfs = func(i int) {
		visited[i] = 1
		for _, e := range edges[i] {
			if visited[e] == 1 {
				valid = false
				return
			} else if visited[e] == 0 {
				dfs(e)
				if !valid {
					return
				}
			}
		}
		visited[i] = 2
		res = append(res, i)
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if len(res) < numCourses {
		return nil
	}
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	var edges = make([][]int, numCourses)
	var pre = make([]int, numCourses)
	var res []int
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		pre[p[0]]++
	}

	var queue = []int{}
	for i, v := range pre {
		if v == 0 {
			queue = append(queue, i)
		}
	}

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

	if len(res) < numCourses {
		return nil
	}
	return res
}
