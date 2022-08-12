package main

// https://leetcode.cn/problems/scramble-string

// 三维
// f[i][j][l] 代表 字符串s1[i:i+l] s2[j:j+l]  是否和谐
func isScramble(s1 string, s2 string) bool {
	f := [30][30][31]bool{}
	sL := len(s1)

	// 枚举长度 1 是否和谐
	for i := 0; i < sL; i++ {
		for j := 0; j < sL; j++ {
			f[i][j][1] = s1[i] == s2[j]
		}
	}

	// 枚举长度 2～l1 是否和谐
	for curL := 2; curL <= sL; curL++ { // 长度变长
		// 枚举 A 中的起点位置
		for i := 0; i <= sL-curL; i++ { //  5-2长度越长 字符串区间越小
			// 枚举 B 中的起点位置
			for j := 0; j <= sL-curL; j++ { //  长度越长 字符串区间越小
				// 枚举划分位置
				for k := 1; k < curL; k++ { // 1-l 长度区间 切割点
					if f[i][j][k] && f[i+k][j+k][curL-k] || f[i][j+curL-k][k] && f[i+k][j][curL-k] {
						f[i][j][curL] = true
						break
					}
				}
			}
		}
	}
	return f[0][0][sL]
}
