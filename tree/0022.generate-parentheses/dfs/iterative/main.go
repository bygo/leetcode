package iterative

//闭合数
func generateParenthesis(n int) []string {
	return closed(n)
}

func closed(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		res = append(res, "")
	} else {
		for i := 0; i < n; i++ {
			leftCombination := closed(i)
			rightCombination := closed(n - 1 - i)
			for _, left := range leftCombination {
				for _, right := range rightCombination {
					res = append(res, "("+left+")"+right)
				}
			}
			// （）
			//  0 = ""
			//  1 = "" ()
			//  2 = "" ()()  (())
			//  3 = "" (()())  ()(())  (())() ((()))
		}
	}
	return res
}
