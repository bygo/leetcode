package main

// https://leetcode-cn.com/problems/jump-game-iii

// ❓ 跳跃游戏 III

func canReach(arr []int, start int) bool {
	var l1 = len(arr) - 1
	var queue = []int{start}
	var visited = map[int]bool{}
	for 0 < len(queue) {
		var cnt = len(queue)
		for _, idx := range queue[:cnt] {
			if arr[idx] == 0 {
				return true
			}
			var l = idx - arr[idx]
			var r = idx + arr[idx]
			if 0 <= l && !visited[l] {
				queue = append(queue, l)
				visited[l] = true
			}

			if r <= l1 && !visited[r] {
				queue = append(queue, r)
				visited[r] = true
			}
		}
		queue = queue[cnt:]
	}

	return false
}
