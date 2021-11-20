package main

// https://leetcode-cn.com/problems/minimum-window-substring

func minWindow(s string, t string) string {
	sm, tm := map[byte]int{}, map[byte]int{}
	l1 := len(s)
	l2 := len(t)

	for i := 0; i < l2; i++ {
		tm[t[i]]++
	}

	l, r := 0, 0
	left := -1
	var cnt int
	var min = 1<<63 - 1
	for r < l1 {
		// 拓边
		if sm[s[r]] < tm[s[r]] {
			cnt++
		}
		sm[s[r]]++

		for cnt == l2 { // 收边
			cur := r - l + 1
			if cur < min {
				min = cur
				left = l
			}
			sm[s[l]]--
			if sm[s[l]] < tm[s[l]] {
				cnt--
			}
			l++
		}
		r++
	}
	if left == -1 {
		return ""
	}

	return s[left : left+min]
}
