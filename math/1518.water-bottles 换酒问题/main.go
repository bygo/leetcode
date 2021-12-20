package main

// https://leetcode-cn.com/problems/water-bottles

func numWaterBottles(numBottles int, numExchange int) int {
	var numNew int
	var cnt = numBottles // 先全喝
	for numExchange <= numBottles {
		numNew = numBottles / numExchange // 兑换
		cnt += numNew                     // 新增

		numBottles -= numExchange * numNew // 减去兑换的
		numBottles += numNew               // 剩余的瓶子 加上新增的
	}
	return cnt
}

func numWaterBottles(numBottles int, numExchange int) int {
	if numBottles < numExchange {
		return numBottles
	}
	return (numBottles-numExchange)/(numExchange-1) + 1 + numBottles
}
