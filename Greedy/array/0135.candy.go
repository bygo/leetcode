package array

// https://leetcode.cn/problems/candy

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func candy(ratings []int) int {
	rL := len(ratings)
	// TODO 左最优
	cntsLeft := make([]int, rL)
	cntsLeft[0] = 1
	for idx := 1; idx < rL; idx++ {
		if ratings[idx-1] < ratings[idx] {
			cntsLeft[idx] = cntsLeft[idx-1] + 1
		} else {
			cntsLeft[idx] = 1
		}
	}

	// TODO 右最优
	cntR := 1
	cnt := max(cntsLeft[rL-1], cntR)
	for idx := rL - 2; 0 <= idx; idx-- {
		if ratings[idx+1] < ratings[idx] {
			cntR++
		} else {
			cntR = 1
		}
		// TODO 贪心
		cnt += max(cntR, cntsLeft[idx])
	}
	return cnt
}

func candy(ratings []int) int {
	rL := len(ratings)

	// TODO
	cnt, inc, dec, pre := 1, 1, 0, 1
	for idx := 1; idx < rL; idx++ {
		// TODO 递增序列
		if ratings[idx-1] <= ratings[idx] {
			dec = 0
			// 相等为1
			if ratings[idx-1] == ratings[idx] {
				pre = 1
			} else {
				// 分配1
				pre++
			}
			// 总数
			cnt += pre
			// 前置
			inc = pre
		} else {
			// TODO 递减序列长度
			dec++
			if dec == inc { // 最后一个 不够高，也需加入序列
				dec++
			}
			cnt += dec // 每次加一轮
			pre = 1
		}
	}
	return cnt
}
