package Stack

// https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop

func finalPrices(prices []int) []int {
	pL := len(prices)
	nums := make([]int, pL)
	stack := []int{0}
	for idx := pL - 1; 0 <= idx; idx-- {
		p := prices[idx]
		for 1 < len(stack) && p < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		nums[idx] = p - stack[len(stack)-1]
		stack = append(stack, p)
	}
	return nums
}
