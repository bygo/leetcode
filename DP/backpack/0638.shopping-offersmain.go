package backpack

// https://leetcode.cn/problems/shopping-offers

func shoppingOffers(price []int, specials [][]int, needs []int) int {
	pL := len(price)
	nL := len(needs)
	// 过滤
	specialsFilter := [][]int{}
	for _, special := range specials {
		totalCnt, totalPrice := 0, 0
		for i, cnt := range special[:pL] {
			totalCnt += cnt
			totalPrice += cnt * price[i]
		}
		if 0 < totalCnt && special[pL] < totalPrice {
			specialsFilter = append(specialsFilter, special)
		}
	}

	f := map[string]int{}
	var dfs func([]byte) int
	dfs = func(curNeeds []byte) (minPrice int) {
		if res, ok := f[string(curNeeds)]; ok {
			return res
		}
		for i, p := range price {
			minPrice += int(curNeeds[i]) * p // 原价价格
		}
		needsNext := make([]byte, pL)
		for _, special := range specialsFilter {
			var idx int
			for idx < nL {
				if curNeeds[idx] < byte(special[idx]) { // 超出需要
					break
				}
				needsNext[idx] = curNeeds[idx] - byte(special[idx])
				idx++
			}
			if idx == nL {
				minPrice = min(minPrice, dfs(needsNext)+special[pL])
			}
		}
		f[string(curNeeds)] = minPrice
		return
	}

	needsCur := make([]byte, pL)
	for i, need := range needs {
		needsCur[i] = byte(need)
	}
	return dfs(needsCur)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
