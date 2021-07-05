package main

//Title: Find Center of Star Graph
// Link: https://leetcode-cn.com/problems/find-center-of-star-graph

//func findCenter(edges [][]int) int {
//	if edges[0][0] == edges[1][0] {
//		return edges[0][0]
//	} else if edges[0][0] == edges[1][1] {
//		return edges[0][0]
//	} else {
//		return edges[0][1]
//	}
//}


func findCenter(edges [][]int) int {
	m := map[int]int{}
	for _, e := range edges {
		for _, v := range e {
			m[v]++
			if 2 == m[v] {
				return v
			}
		}
	}
	return 0
}