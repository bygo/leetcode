package main

import "sort"

func breakfastNumber(staple []int, drinks []int, x int) int {
	l1 := len(staple)
	l2 := len(drinks)

	sort.Ints(staple)
	sort.Ints(drinks)

	var res int
	var i, j = 0, l2 - 1
	for i < l1 && 0 <= j {
		if staple[i]+drinks[j] <= x {
			res += j + 1
			i++
		} else {
			j--
		}
	}
	return res % 1000000007
}

func breakfastNumber(staple []int, drinks []int, x int) int {
	sum := make([]int, x+1)

	for _, v := range staple {
		if v < x {
			sum[v]++
		}
	}

	for i := 1; i <= x; i++ {
		sum[i] += sum[i-1]
	}

	var res int
	for _, v := range drinks {
		if v < x {
			res += sum[x-v]
		}
	}

	return res % 1000000007
}
