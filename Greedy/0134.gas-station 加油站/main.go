package main

// https://leetcode.cn/problems/gas-station

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	i := 0
	for i < l {
		var j = 0
		var cnt = 0
		for j < l {
			index := (i + j) % l
			cnt += gas[index] - cost[index]
			if cnt < 0 {
				break
			}
			j++
		}

		if j == l {
			return i
		}
		i += j + 1
	}
	return -1
}
