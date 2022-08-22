package Array

// https://leetcode.cn/problems/shifting-letters-ii

func shiftingLetters(s string, shifts [][]int) string {
	sL := len(s)
	diff := make([]int, sL+1)
	for _, shift := range shifts {
		shift[2] = shift[2]<<1 - 1
		diff[shift[0]] += shift[2]
		diff[shift[1]+1] -= shift[2]
	}

	buf := []byte(s)
	var offset int
	for i := range buf {
		offset = (diff[i]+offset)%26 + 26
		buf[i] = (buf[i]-'a'+byte(offset))%26 + 'a'
	}
	return string(buf)
}
