package main

// https://leetcode-cn.com/problems/combinations

func main() {
	combine(5, 3)
}

// 逆向
func combine(n int, k int) [][]int {
	var combNums [][]int
	var cur []int
	for i := 1; i <= k; i++ {
		cur = append(cur, i)
	}
	cur = append(cur, n+1) // 哨兵

	var cntIncr = 0
	for cntIncr < k {
		combNums = append(combNums, append([]int{}, cur[:k]...))
		cntIncr = 0
		for cntIncr < k && cur[cntIncr]+1 == cur[cntIncr+1] { // 递增 判断  6(哨兵)
			cur[cntIncr] = cntIncr + 1 // 重置 1 2 3
			cntIncr++
		}

		// 最后一个递增 +1
		// 1 2 3  6

		// 1 2 4  6
		// 1 3 4  6
		// 2 3 4  6

		// 1 2 5  6
		// 1 3 5  6
		// 2 3 5  6
		// 1 4 5  6
		// 2 4 5  6
		// 3 4 5  6

		// 全部到达尾部 idx超过k
		cur[cntIncr]++
	}
	return combNums
}
