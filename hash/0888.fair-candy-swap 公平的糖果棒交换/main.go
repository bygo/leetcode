package main

// https://leetcode-cn.com/problems/fair-candy-swap

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	m := map[int]struct{}{}
	sum1 := 0
	for i := range aliceSizes {
		m[aliceSizes[i]] = struct{}{}
		sum1 += aliceSizes[i]
	}

	sum2 := 0
	for i := range bobSizes {
		sum2 += bobSizes[i]
	}

	diff := (sum1 - sum2) / 2

	for i := range bobSizes {
		_, ok := m[bobSizes[i]+diff]
		if ok {
			return []int{bobSizes[i] + diff, bobSizes[i]}
		}
	}
	return nil
}
