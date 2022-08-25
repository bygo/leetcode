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
	cntsLeft := make([]int, rL)
	cntsLeft[0] = 1
	for idx := 1; idx < rL; idx++ {
		if ratings[idx-1] < ratings[idx] {
			cntsLeft[idx] = cntsLeft[idx-1] + 1
		} else {
			cntsLeft[idx] = 1
		}
	}

	cntR := 1
	cnt := max(cntsLeft[rL-1], cntR)
	for idx := rL - 2; 0 <= idx; idx-- {
		if ratings[idx+1] < ratings[idx] {
			cntR++
		} else {
			cntR = 1
		}
		cnt += max(cntR, cntsLeft[idx])
	}
	return cnt
}

func candy(ratings []int) int {
	rL := len(ratings)
	cnt, inc, dec, pre := 1, 1, 0, 1
	for idx := 1; idx < rL; idx++ {
		if ratings[idx-1] <= ratings[idx] {
			dec = 0
			if ratings[idx] == ratings[idx-1] {
				pre = 1
			} else {
				pre++
			}
			cnt += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			cnt += dec
			pre = 1
		}
	}
	return cnt
}
