package main

import "sort"

// https://leetcode-cn.com/problems/escape-a-large-maze

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 离散化 a，返回的哈希表中的键值对分别为 a 中的原始值及其离散化后的对应值
func discrete(a []int) (map[int]int, int) {
	sort.Ints(a)

	id := 0
	if 0 < a[0] {
		id = 1
	}
	mapping := map[int]int{a[0]: id}
	pre := a[0]
	for _, v := range a[1:] {
		if v != pre {
			if v == pre+1 {
				id++
			} else {
				id += 2
			}
			mapping[v] = id
			pre = v
		}
	}

	const boundary int = 1e6
	if a[len(a)-1] != boundary-1 {
		id++
	}

	return mapping, id
}

func isEscapePossible(block [][]int, source, target []int) bool {
	n := len(block)
	if n < 2 {
		return true
	}
	rows := []int{source[0], target[0]}
	cols := []int{source[1], target[1]}
	for _, b := range block {
		rows = append(rows, b[0])
		cols = append(cols, b[1])
	}

	// 离散化行列坐标
	rMapping, rBound := discrete(rows)
	cMapping, cBound := discrete(cols)

	grid := make([][]bool, rBound+1)
	for i := range grid {
		grid[i] = make([]bool, cBound+1)
	}
	for _, b := range block {
		grid[rMapping[b[0]]][cMapping[b[1]]] = true
	}

	sx, sy := rMapping[source[0]], cMapping[source[1]]
	tx, ty := rMapping[target[0]], cMapping[target[1]]
	grid[sx][sy] = true
	q := []pair{{sx, sy}}
	for 0 < len(q) {
		p := q[0]
		q = q[1:]
		for _, d := range dirs {
			x, y := p.x+d.x, p.y+d.y
			if 0 <= x && x <= rBound && 0 <= y && y <= cBound && !grid[x][y] {
				if x == tx && y == ty {
					return true
				}
				grid[x][y] = true
				q = append(q, pair{x, y})
			}
		}
	}
	return false
}
