package main

// https://leetcode-cn.com/problems/self-crossing

// 外卷 可以进入 内卷
func isSelfCrossing(distance []int) bool {
	l := len(distance)
	if 5 <= l && distance[3] == distance[1] && distance[2]-distance[0] <= distance[4] {
		return true
	}
	for i := 3; i < len(distance); i++ {
		// 第 1 类路径交叉的情况
		if distance[i-2] <= distance[i] && distance[i-1] <= distance[i-3] { // 长度一定要比上次长，位置 靠外
			return true
		}

		// 内卷失败
		if 5 <= i && distance[i-1] <= distance[i-3] &&
			distance[i-2]-distance[i-4] <= distance[i] &&
			distance[i-3]-distance[i-5] <= distance[i-1] &&
			distance[i-4] < distance[i-2] {
			return true
		}
	}
	return false
}
