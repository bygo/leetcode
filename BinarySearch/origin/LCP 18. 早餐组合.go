package main

import "sort"

// https://leetcode-cn.com/problems/2vYnGI/

func breakfastNumber(staple []int, drinks []int, x int) int {
	sL := len(staple)
	sort.Ints(staple)
	find := func(target int) int {
		lo, hi := 0, sL
		for lo < hi {
			mid := int(uint(lo+hi) >> 1)
			if staple[mid] <= target {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		return lo
	}

	var cnt int
	for _, num := range drinks {
		if num < x {
			cnt += find(x - num)
		}
	}
	return cnt % 1000000007
}

// two pointer
func breakfastNumber(staple []int, drinks []int, x int) int {
	sL := len(staple)
	dL := len(drinks)

	sort.Ints(staple)
	sort.Ints(drinks)

	var cnt int
	lo, hi := 0, dL-1
	for lo < sL && 0 <= hi {
		if staple[lo]+drinks[hi] <= x {
			cnt = cnt + hi + 1
			lo++
		} else {
			hi--
		}
	}
	return cnt % 1000000007
}

// hash pre
func breakfastNumber(staple []int, drinks []int, x int) int {
	sum := make([]int, x+1)

	for _, num := range staple {
		if num < x {
			sum[num]++
		}
	}

	// 前缀
	for i := 1; i <= x; i++ {
		sum[i] += sum[i-1]
	}

	var cnt int
	for _, num := range drinks {
		if num < x {
			cnt += sum[x-num]
		}
	}
	return cnt % 1000000007
}
