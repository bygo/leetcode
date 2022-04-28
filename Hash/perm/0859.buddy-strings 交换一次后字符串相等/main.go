package main

// https://leetcode-cn.com/problems/buddy-strings

// ❓ 交换一次后字符串相等

func buddyStrings(str string, goal string) bool {
	strL := len(str)
	goalL := len(goal)
	if strL != goalL {
		return false
	}

	var slotFirst, slotSecond = -1, -1
	var chMpCnt = [26]int{}
	for i := 0; i < strL; i++ {
		ch := str[i] - 'a'
		chMpCnt[ch]++
		if str[i] != goal[i] {
			if slotFirst == -1 {
				slotFirst = i
			} else if slotSecond == -1 {
				slotSecond = i
			} else {
				return false
			}
		}
	}

	isSame := func() bool {
		for _, cnt := range chMpCnt {
			if 2 <= cnt {
				return true
			}
		}
		return false
	}

	return slotSecond != -1 && str[slotFirst] == goal[slotSecond] && str[slotSecond] == goal[slotFirst] || slotFirst == -1 && isSame()
}
