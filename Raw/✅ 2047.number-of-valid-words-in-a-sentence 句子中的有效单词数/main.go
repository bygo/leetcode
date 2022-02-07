package main

// https://leetcode-cn.com/problems/number-of-valid-words-in-a-sentence

func main() {

}
func countValidWords(sentence string) int {
	var cnt int
	var idx int
	var sL = len(sentence)
	for idx < sL {
		var state int

		// 起始位
		for idx < sL && 'a' <= sentence[idx] && sentence[idx] <= 'z' {
			state |= 1
			idx++
		}

		if idx < sL && sentence[idx] == '-' {
			state |= 2
			idx++
		}

		for idx < sL && 'a' <= sentence[idx] && sentence[idx] <= 'z' {
			state |= 4
			idx++
		}

		if idx < sL && (sentence[idx] == ',' || sentence[idx] == '.' || sentence[idx] == '!') {
			state |= 8
			idx++
		}

		if idx == sL || sentence[idx] == ' ' {
			if state == 1 || state == 7 || state == 8 || state == 9 || state == 15 {
				cnt++
			}
			idx++
		} else {
			// 合法起始位
			for idx < sL && sentence[idx] != ' ' {
				idx++
			}
		}
		for idx < sL && sentence[idx] == ' ' {
			idx++
		}

	}
	return cnt
}
