package main

import "time"

// https://leetcode-cn.com/problems/longest-nice-substring

// ❓ 大小写成对出现的最长子串
// ⚠️ Aa Bb... 成对

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		longestNiceSubstring("jceWzKNUrLjceWzKNUrLQvxRyljvkUwjceWzKNUrLQvxRyljvkUwCjceWzKNUrLQvxRyljvkUwCUymYuyAMAJLVaFfQfnUymjceWzKNUrLQvxRyljvkUwCUymYuyAMAJLVaFfQfnYuyAMAJLVaFfQfnCUymYuyAMAJLVaFfQfnQvxRyljceWzjceWzKNUrLQvxRyljvkUwCUymYuyAMAJLVaFfQfnKNUrLQvxRyljvkUwCUymYuyAMAJLVaFfQfnjvkUwCUymYuyAMAJLVaFfQfn")
	}
	end := time.Since(start)
	println(end.Seconds())
}

func longestNiceSubstring1(s string) string {
	chMp := map[byte]struct{}{}
	for i := range s {
		chMp[s[i]] = struct{}{}
	}

	for i := range s {
		// 找对
		ch := s[i]
		if ch <= 'Z' {
			ch += 32
		} else {
			ch -= 32
		}
		_, ok := chMp[ch]

		if !ok {
			// 分区
			left := longestNiceSubstring(s[:i])
			right := longestNiceSubstring(s[i+1:])
			if len(left) < len(right) {
				return right
			}
			return left
		}
	}
	return s
}

func longestNiceSubstring(s string) string {
	chMp := map[byte]struct{}{}
	for i := range s {
		chMp[s[i]] = struct{}{}
	}

	for i := range s {
		ch := s[i]
		if ch <= 'Z' {
			ch += 32
		} else if 'a' <= ch {
			ch -= 32
		}
		_, ok := chMp[ch]
		if !ok {
			left := longestNiceSubstring(s[:i])
			if len(s[i+1:]) <= len(left) {
				return left
			}
			right := longestNiceSubstring(s[i+1:])
			if len(left) < len(right) {
				return right
			}
			return left
		}
	}
	return s
}
