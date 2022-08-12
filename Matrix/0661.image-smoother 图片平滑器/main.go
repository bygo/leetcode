package main

// https://leetcode.cn/problems/image-smoother

func imageSmoother(img [][]int) [][]int {
	rowL, colL := len(img), len(img[0])
	res := make([][]int, rowL)
	for i := range res {
		res[i] = make([]int, colL)
		for j := range res[i] {
			sum, cnt := 0, 0
			for _, row := range img[max(i-1, 0):min(i+2, rowL)] {
				for _, num := range row[max(j-1, 0):min(j+2, colL)] {
					sum += num
					cnt++
				}
			}
			res[i][j] = sum / cnt
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
