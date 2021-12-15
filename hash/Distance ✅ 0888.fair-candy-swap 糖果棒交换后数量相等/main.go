package main

// https://leetcode-cn.com/problems/fair-candy-swap

// ❓ 交换一个值后，总和相等

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	// aliceSizes map
	// aliceSizes 糖果数
	valMp := map[int]struct{}{}
	sum1 := 0
	for i := range aliceSizes {
		valMp[aliceSizes[i]] = struct{}{}
		sum1 += aliceSizes[i]
	}

	// bobSizes 糖果数
	sum2 := 0
	for i := range bobSizes {
		sum2 += bobSizes[i]
	}

	// 一增一减
	// 必为偶数
	diff := sum1 - sum2
	if diff%2 == 1 {
		// 并且因奇数向下取整，差值还会少1，虽然循环找出结果，但不是答案
		return nil
	}
	diff /= 2

	// 找差额
	for i := range bobSizes {
		aliceVal := bobSizes[i] + diff
		_, ok := valMp[aliceVal]
		if ok {
			return []int{aliceVal, bobSizes[i]}
		}
	}

	return nil
}
