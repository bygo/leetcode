package add

// https://leetcode.cn/problems/partitioning-into-minimum-number-of-deci-binary-numbers

// 123 => 111 011 001 => 3

func minPartitions(str string) int {
	var max byte = 0
	for i := range str {
		if max < str[i] {
			max = str[i]
		}
	}
	return int(max - '0')
}
