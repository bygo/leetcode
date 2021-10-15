package main

// https://leetcode-cn.com/problems/kth-missing-positive-number

func findKthPositive(arr []int, k int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)/2
		cur := arr[mid] - mid - 1
		if cur < k {
			l = mid + 1
		} else if k <= cur {
			r = mid
		}
	}
	return k + l
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numColor(root *TreeNode) int {
	m := map[int]struct{}{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			dfs(root.Right)
			m[root.Val] = struct{}{}
		}
	}
	dfs(root)
	return len(m)
}

func bicycleYard(position []int, terrain [][]int, obstacle [][]int) [][]int {
	var graph = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	l1 := len(terrain)
	l2 := len(terrain[0])

	var visit = make([][][102]bool, l1)
	for i := range visit {
		visit[i] = make([][102]bool, l2)
	}

	x := position[0]
	y := position[1]
	visit[x][y][1] = true
	var dfs func(x, y, speed int)
	dfs = func(x, y, speed int) {
		for _, v := range graph {
			nx, ny := x+v[0], y+v[1]
			if 0 <= nx && nx < l1 && 0 <= ny && ny < l2 {
				h1 := terrain[x][y]
				h2 := terrain[nx][ny]
				o2 := obstacle[nx][ny]
				s := speed + h1 - h2 - o2
				if 0 < s && !visit[nx][ny][s] {
					visit[nx][ny][s] = true
					dfs(nx, ny, s)
				}
			}
		}
	}
	dfs(x, y, 1)
	var res [][]int
	visit[x][y][1] = false
	for i, row := range visit {
		for j, v := range row {
			if v[1] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
