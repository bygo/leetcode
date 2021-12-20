package main

import "math"

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n))) // 完全平方数 约数为奇数，灯泡亮，其他为成对出现
}
