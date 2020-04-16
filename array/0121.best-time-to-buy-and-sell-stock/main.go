package main

func maxProfit(prices []int) int {
	min, max := 1<<63-1, 0

	for _, v := range prices {
		if v < min {
			min = v
		}
		if max < v-min {
			max = v - min
		}
	}
	return max
}
