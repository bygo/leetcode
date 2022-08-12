package main

// https://leetcode.cn/problems/surrounded-regions

func solveDFS(board [][]byte) {
	l1 := len(board)
	l2 := len(board[0])
	if l1 == 0 || l2 == 0 {
		return
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || l1 <= x || y < 0 || l2 <= y || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'
		// 搜索
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}

	for i := 0; i < l1; i++ {
		dfs(i, 0)
		dfs(i, l2-1)
	}

	for i := 1; i < l2-1; i++ {
		dfs(0, i)
		dfs(l1-1, i)
	}

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func solveBFS(board [][]byte) {
	l1 := len(board)
	l2 := len(board[0])
	if l1 == 0 || l2 == 0 {
		return
	}
	queue := [][]int{}
	// 所有边
	for i := 0; i < l1; i++ {
		if board[i][0] == 'O' {
			queue = append(queue, []int{i, 0})
			board[i][0] = 'A'
		}
		if board[i][l2-1] == 'O' {
			queue = append(queue, []int{i, l2 - 1})
			board[i][l2-1] = 'A'
		}
	}
	for i := 1; i < l2-1; i++ {
		if board[0][i] == 'O' {
			queue = append(queue, []int{0, i})
			board[0][i] = 'A'
		}
		if board[l1-1][i] == 'O' {
			queue = append(queue, []int{l1 - 1, i})
			board[l1-1][i] = 'A'
		}
	}

	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		for i := range queue {
			cur := queue[i]
			for _, dir := range [4][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				nx := cur[0] + dir[0]
				ny := cur[1] + dir[1]
				if nx < 0 || l1 <= nx || ny < 0 || l2 <= ny || board[nx][ny] != 'O' {
					continue
				}
				queue = append(queue, []int{nx, ny})
				board[nx][ny] = 'A'
			}
		}

		queue = queue[cnt:]
	}
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
