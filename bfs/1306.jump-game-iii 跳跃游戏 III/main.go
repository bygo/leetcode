package main

// Q: can i `jump` to `position' which the value is 0 ?
// https://leetcode-cn.com/problems/jump-game-iii

// A: bfs
func canReach(arr []int, start int) bool {
	var l1 = len(arr) - 1
	var queue = []int{start}
	var visited = map[int]bool{}
	for 0 < len(queue) {
		count := len(queue)
		for _, index := range queue {
			if arr[index] == 0 {
				return true
			}
			left := index - arr[index]
			right := index + arr[index]
			if 0 <= left && !visited[left] {
				queue = append(queue, left)
				visited[left] = true
			}
			if right <= l1 && !visited[right] {
				queue = append(queue, right)
				visited[right] = true
			}
		}
		queue = queue[count:]
	}

	return false
}
