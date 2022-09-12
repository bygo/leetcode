package string

// https://leetcode.cn/problems/longest-chunked-palindrome-decomposition

func longestDecomposition(text string) int {
	textL := len(text)
	if textL <= 1 {
		return textL
	}

	// TODO 贪心
	for i := 1; i < textL; i++ {
		if text[:i] == text[textL-i:] {
			return 2 + longestDecomposition(text[i:textL-i])
		}
	}
	return 1
}
