package main

// https://leetcode.cn/problems/gas-station

func canCompleteCircuit(gas []int, cost []int) int {
	gasL := len(gas)
	idx := 0
	for idx < gasL {
		var steps = 0
		var cnt = 0
		for steps < gasL {
			idxCur := (idx + steps) % gasL    // TODO 环
			cnt += gas[idxCur] - cost[idxCur] // 油量
			if cnt < 0 {
				break
			}
			steps++
		}

		if steps == gasL {
			return idx
		}

		// TODO 剪枝
		idx += steps + 1
	}
	return -1
}
