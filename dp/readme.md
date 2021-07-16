# 0005.longest-palindromic-substring 最长回文 
```go
package main

// https://leetcode-cn.com/problems/longest-palindromic-substring

// 一维
func longestPalindrome(s string) string {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n) // f[left][right] left至right 是否回文
	}

	var left, right int
	for r := 0; r < n; r++ {
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || f[l+1][r-1]) { // 0:相等  1:相邻  2:隔一个字符
				f[l][r] = true
				if right-left < r-l {
					right = r
					left = l
				}
			}
		}
	}
	return s[left : right+1]
}

// 压缩外循环
func longestPalindrome(s string) string {
	n := len(s)
	f := make([]bool, n)

	var left, right int
	for r := 0; r < n; r++ {
		for l := 0; l < r; l++ {
			f[l] = s[l] == s[r] && (r-l <= 2 || f[l+1])
			if f[l] && right-left < r-l {
				right = r
				left = l
			}
		}
	}
	return s[left : right+1]
}

```

# 0010.regular-expression-matching .*正则匹配 
```go
package main

// https://leetcode-cn.com/problems/regular-expression-matching

// 二维
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	f := make([][]bool, m+1)
	for i := range f {
		f[i] = make([]bool, n+1)
	}

	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				if f[i][j-2] {
					f[i][j] = true
				} else if i != 0 && f[i-1][j] && (p[j-2] == s[i-1] || p[i-2] == '.') {
					f[i][j] = true
				}
			} else if i != 0 && f[i-1][j-1] && (p[j-1] == s[i-1] || p[i-1] == '.') {
				f[i][j] = true
			}
		}
	}

	return f[m][n]
}

// 递归
func isMatch(s string, p string) bool {
	return dfs(s, p)
}

func dfs(s string, p string) bool {
	m, n := len(s), len(p)
	if n == 0 {
		return m == 0
	}
	if 2 <= n && p[1] == '*' {
		return dfs(s, p[2:]) || 0 < m && (s[0] == p[0] || p[0] == '.') && dfs(s[1:], p)
	} else {
		return 0 < m && (s[0] == p[0] || p[0] == '.') && dfs(s[1:], p[1:])
	}
}

```

# 0032.longest-valid-parentheses 最长有效括号 
```go
package main

// https://leetcode-cn.com/problems/longest-valid-parentheses

// 一维
// if **() : f(n) = f(n-2) + 2
// if (*)) : f(n) = f(n-1) + 2 + f(n-f(n-1)-2)
func longestValidParentheses(s string) int {
	var res int
	var n = len(s)
	var f = make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == ')' { // 1. 当前  )
			if s[i-1] == '(' { // 2.前置  (
				if 2 <= i {
					f[i] = f[i-2] + 2
				} else {
					f[i] = 2
				}
			} else if 0 < i-f[i-1] && s[i-f[i-1]-1] == '(' {
				if 2 <= i-f[i-1] {
					f[i] = f[i-1] + f[i-f[i-1]-2] + 2
				} else {
					f[i] = f[i-1] + 2
				}
			}
			if res < f[i] {
				res = f[i]
			}
		}
	}
	return res
}

// map
func longestValidParentheses(s string) int {
	var res int
	var n = len(s)
	var f = map[int]int{}
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				f[i] = f[i-2] + 2
			} else if 0 < i-f[i-1] && s[i-f[i-1]-1] == '(' {
				f[i] = f[i-1] + f[i-f[i-1]-2] + 2
			}
			if res < f[i] {
				res = f[i]
			}
		}
	}
	return res
}

```

# 0042.trapping-rain-water 接雨水 
```go
package main

// https://leetcode-cn.com/problems/trapping-rain-water

// 前缀
// f(n) = min(leftMax,rightMax) - self
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0] // 第一个是本身
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1] // 最后一个是本身
	for i := n - 2; 0 <= i; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	var res int
	for i, h := range height {
		res += min(leftMax[i], rightMax[i]) - h // 水
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

```

# 0044.wildcard-matching *？通配符 
```go
package main

// https://leetcode-cn.com/problems/wildcard-matching

// 二维
func isMatch(s string, p string) bool {
	l1, l2 := len(s), len(p)
	f := make([][]bool, l1+1)
	for i := range f {
		f[i] = make([]bool, l2+1)
	}
	f[0][0] = true
	for j := 1; j <= l2; j++ {
		if p[j-1] != '*' {
			break
		}
		f[0][j] = true
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i-1][j] || f[i][j-1]
			} else if s[i-1] == p[j-1] || '?' == p[j-1] {
				f[i][j] = f[i-1][j-1]
			}
		}
	}
	return f[l1][l2]
}

```

# 0053.maximum-subarray 最大子序和 
```go
package main

// https://leetcode-cn.com/problems/maximum-subarray

// 前缀
// f(n) = f(n) + f(n-1)
func maxSubArray(f []int) int {
	res := f[0]
	l1 := len(f)
	for i := 1; i < l1; i++ {
		if 0 < f[i-1] { // 贡献
			f[i] += f[i-1]
		}
		if res < f[i] {
			res = f[i]
		}
	}
	return res
}

```

# 0062.unique-paths 不同路径 
```go
package main

// https://leetcode-cn.com/problems/unique-paths

// 二维
// f(i)(j) = f(i-1)(j) + f(i)(j-1)
func uniquePaths(l1, l2 int) int {
	f := make([][]int, l1)
	for i := range f {
		f[i] = make([]int, l2)
		f[i][0] = 1
	}
	for j := 0; j < l2; j++ {
		f[0][j] = 1
	}
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[l1-1][l2-1]
}

// 压缩
func uniquePaths(l1, l2 int) int {
	f := make([]int, l2)
	f[0] = 1
	for i := 0; i < l1; i++ {
		for j := 1; j < l2; j++ {
			f[j] += f[j-1]
		}
	}
	return f[l2-1]
}

```

# 0063.unique-paths-ii 不同路径 II 
```go
package main

// https://leetcode-cn.com/problems/unique-paths-ii

// 二维
// f(i)(j) = f(i-1)(j) + f(i)(j-1)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	l1 := len(obstacleGrid)
	l2 := len(obstacleGrid[0])
	f := make([][]int, l1)

	var i int
	for i < l1 { // 第一列
		if obstacleGrid[i][0] == 1 {
			break
		}
		f[i] = make([]int, l2)
		f[i][0] = 1
		i++
	}

	for i < l1 {
		f[i] = make([]int, l2)
		i++
	}

	for j := 0; j < l2; j++ { // 第一行
		if obstacleGrid[0][j] == 1 {
			break
		}
		f[0][j] = 1
	}
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if obstacleGrid[i][j] == 0 {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[l1-1][l2-1]
}

// 压缩
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	l1, l2 := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, l2)
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	f[0] = 1

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
			} else if 1 <= j {
				f[j] += f[j-1]
			}
		}
	}
	return f[l2-1]
}

```

# 0064.minimum-path-sum 最小路径和 
```go
package main

// https://leetcode-cn.com/problems/minimum-path-sum

// 二维
// f(i)(j) = f(i-1)(j) + f(i)(j-1)
func minPathSum(grid [][]int) int {
	l1, l2 := len(grid), len(grid[0])
	f := make([][]int, l1)
	f[0] = make([]int, l2)
	f[0][0] = grid[0][0]
	for i := 1; i < l1; i++ {
		f[i] = make([]int, l2)
		f[i][0] = grid[i][0] + f[i-1][0]
	}

	for i := 1; i < l2; i++ {
		f[0][i] = grid[0][i] + f[0][i-1]
	}

	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			f[i][j] = min(f[i-1][j], f[i][j-1]) + grid[i][j]
		}
	}

	return f[l1-1][l2-1]
}

// 压缩
func minPathSum(grid [][]int) int {
	l1, l2 := len(grid), len(grid[0])
	f := make([]int, l2)
	f[0] = grid[0][0]
	for j := 1; j < l2; j++ {
		f[j] = grid[0][j] + f[j-1]
	}

	for i := 1; i < l1; i++ {
		f[0] += grid[i][0]
		for j := 1; j < l2; j++ {
			f[j] = min(f[j-1], f[j]) + grid[i][j]
		}
	}

	return f[l2-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

```

# 0070.climbing-stairs 爬楼梯 
```go
package main

// https://leetcode-cn.com/problems/climbing-stairs

// 一维
// f(n) = f(n-1) + f(n-2)
func climbStairs(n int) int {
	f := make([]int, n+1)
	f[0] = 1
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// 压缩
func climbStairs(n int) int {
	x, y := 1, 1
	for i := 2; i <= n; i++ {
		y, x = x+y, y
	}
	return y
}

```

# 0072.edit-distance 编辑距离 
```go
package main

// https://leetcode-cn.com/problems/edit-distance

// 二维
// f(i)(j) = min( f(i-1)(j), f(i)(j-1), f(i-1)(j-1) )
func minDistances(word1, word2 string) int {
	l1, l2 := len(word1), len(word2)
	var f = make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		f[i] = make([]int, l2+1)
		f[i][0] = i
	}

	for i := range f[0] {
		f[0][i] = i
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			angle := f[i-1][j-1]
			if word1[i-1] == word2[j-1] {
				angle -= 1
			} // 45000
			f[i][j] = min(f[i-1][j], f[i][j-1], angle) + 1
		}
	}

	return f[l1][l2]
}

// 压缩
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	var f = make([]int, l2+1)
	for i := 0; i <= l2; i++ {
		f[i] = i
	}

	for i := 1; i <= l1; i++ {
		angle := i - 1
		f[0] = i
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				angle -= 1
			}
			f[j], angle = min(f[j-1], f[j], angle)+1, f[j]
		}
	}

	// f[i-1][j-1] 不同等于替换
	// f[i][j-1] 插入2
	// f[i-1][j] 插入1
	// 	  #  h  h  h  o
	// #  0  1  2  3  4
	// h  1  0  1  2  3
	// o  2  1  1  2  2

	//    #   s   i   t   t   e
	// #  0   1   2   3   4   5
	// k  1   1   2   3   4   5
	// i  2   2   1   2   3   4
	// t  3   3   2   1   2   3
	// t  4   4   3   2   1   2
	// e  5   5   4   3   2   1
	return f[l2]
}

func min(nums ...int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] < nums[0] {
			nums[0] = nums[i]
		}
	}
	return nums[0]
}

```

# 0087.scramble-string 扰乱字符串 
```go
package main

// https://leetcode-cn.com/problems/scramble-string

// 记忆递归
func isScramble(s1 string, s2 string) bool {
	// 记忆 减少重复计算
	var cache = make(map[string]bool)

	var dfs func(s1 string, s2 string) bool

	dfs = func(s1 string, s2 string) bool {
		var s = s1 + s2
		v, ok := cache[s]
		if ok {
			return v
		}

		if s1 == s2 {
			cache[s] = true
			return true
		}

		l := len(s1)

		var freq = map[byte]int{}
		for i := 0; i < l; i++ {
			freq[s1[i]]++
			freq[s2[i]]--
		}

		for _, v := range freq {
			if v != 0 {
				cache[s] = false
				return false
			}
		}

		//切割
		for i := 1; i < l; i++ {
			if dfs(s1[:i], s2[:i]) && dfs(s1[i:], s2[i:]) || dfs(s1[:i], s2[l-i:]) && dfs(s1[i:], s2[:l-i]) {
				cache[s] = true
				return true
			}
		}
		// 没有一个符合要求的
		cache[s] = false
		return false
	}
	return dfs(s1, s2)
}

// 三维
// f[i][j][l] 代表 字符串s1[i:i+l] s2[j:j+l]  是否和谐
func isScramble(s1 string, s2 string) bool {
	f := [30][30][31]bool{}
	l1 := len(s1)

	// 枚举长度 1 是否和谐
	for i := 0; i < l1; i++ {
		for j := 0; j < l1; j++ {
			f[i][j][1] = s1[i] == s2[j]
		}
	}

	// 枚举长度 2～l1 是否和谐
	for l := 2; l <= l1; l++ { // 长度变长
		// 枚举 A 中的起点位置
		for i := 0; i <= l1-l; i++ { //  5-2长度越长 字符串区间越小
			// 枚举 B 中的起点位置
			for j := 0; j <= l1-l; j++ { //  长度越长 字符串区间越小
				// 枚举划分位置
				for k := 1; k < l; k++ { // 1-l 长度区间 切割点
					if f[i][j][k] && f[i+k][j+k][l-k] || f[i][j+l-k][k] && f[i+k][j][l-k] {
						f[i][j][l] = true
						break
					}
				}
			}
		}
	}
	return f[0][0][l1]
}

```

# 0091.decode-ways 解码方法 
```go
package main

// https://leetcode-cn.com/problems/decode-ways

// 一维
// f(n) = f(n-1) + f(n-2)
func numDecodings(s string) int {
	l1 := len(s)
	f := make([]int, l1+1)
	f[0] = 1
	if s[0] != '0' {
		f[1] = 1
	}

	for i := 2; i <= l1; i++ {
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}

		if s[i-2] != '0' && (s[i-2]-'0')*10+s[i-1]-'0' <= 26 {
			f[i] += f[i-2]
		}
	}

	return f[l1]
}

// 压缩
func numDecodings(s string) int {
	l1 := len(s)
	x, y, z := 1, 0, 0
	if s[0] != '0' {
		y = 1
		z = 1
	}

	for i := 1; i < l1; i++ {
		if s[i] != '0' {
			z = y
		} else {
			z = 0
		}

		if s[i-1] != '0' && ((s[i-1]-'0')*10+(s[i]-'0') <= 26) {
			z += x
		}

		x, y = y, z
	}
	return z
}

```

# 0097.interleaving-string 交错字符串 
```go
package main

// https://leetcode-cn.com/problems/interleaving-string

// 二维
// f(i)(j) = f(i)(j-1) && s2[j-1] == p || f(i-1)(j) && s1[i-1] == p
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if (l1 + l2) != l3 {
		return false
	}
	f := make([][]bool, l1+1)
	for i := 0; i <= l1; i++ {
		f[i] = make([]bool, l2+1)
	}
	f[0][0] = true
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			f[i][j] = i == 0 && j == 0 || 0 < i && f[i-1][j] && s1[i-1] == s3[i+j-1] ||
				0 < j && f[i][j-1] && s2[j-1] == s3[i+j-1]
		}
	}
	return f[l1][l2]
}

// 压缩
// dp[j] 表示 当前i下
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if (l1 + l2) != l3 {
		return false
	}
	f := make([]bool, l2+1)
	f[0] = true
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if 0 < i {
				f[j] = f[j] && s1[i-1] == s3[i+j-1] // i增加后 f[j] 是否依然有效
			}

			if 0 < j { // 如果dp[j] 有效 或 可以通过增加j 使之有效
				f[j] = f[j] || f[j-1] && s2[j-1] == s3[i+j-1]
			}
		}
	}
	return f[l2]
}

```

# 0115.distinct-subsequences 不同的子序列 
```go
package main

// https://leetcode-cn.com/problems/distinct-subsequences

func numDistinct(s string, t string) int {
	l1, l2 := len(s), len(t)
	f := make([][]int, l1+1)
	for i := range f {
		f[i] = make([]int, l2+1)
		f[i][0] = 1
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			f[i][j] = f[i-1][j]   // 少一个字符 的路径数
			if s[i-1] == t[j-1] { // 如果相等 加 上次普通路径数
				f[i][j] += f[i-1][j-1]
			}
		}
	}

	return f[l1][l2]
}

```

# 0118.pascals-triangle 杨辉三角 
```go
package main

// https://leetcode-cn.com/problems/pascals-triangle

// 二维
// f(i)(j) = f(i-1)(j-1) + f(i-1)(j)
func generate(numRows int) [][]int {
	var f = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		f[i] = make([]int, i+1)
		f[i][0] = 1
		f[i][i] = 1
		for j := 1; j < i; j++ {
			f[i][j] = f[i-1][j-1] + f[i-1][j]
		}
	}

	return f
}

```

# 0119.pascals-triangle-ii 杨辉三角 II 
```go
package main

// https://leetcode-cn.com/problems/pascals-triangle-ii

// 压缩
// f(i)(j) = f(i-1)(j-1) + f(i-1)(j)
func getRow(rowIndex int) []int {
	numRows := rowIndex + 1
	var dp = make([]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[0] = 1
		dp[i] = 1
		for j := i - 1; 0 < j; j-- {
			dp[j] = dp[j-1] + dp[j]
		}
	}
	return dp
}

```

# 0120.triangle 三角形最小路径 
```go
package main

// https://leetcode-cn.com/problems/triangle

// 压缩 从顶至底
func minimumTotal(triangle [][]int) int {
	l1 := len(triangle)
	l2 := len(triangle[l1-1])
	var f = make([]int, l2)
	f[0] = triangle[0][0]

	for i := 1; i < l1; i++ {
		f[i] = f[i-1] + triangle[i][i] // 最后一个
		for j := i - 1; 0 < j; j-- {   // 中间
			f[j] = min(f[j], f[j-1]) + triangle[i][j]
		}
		f[0] = f[0] + triangle[i][0] // 第一个
	}

	for i := 1; i < l2; i++ {
		if f[i] < f[0] {
			f[0] = f[i]
		}
	}

	return f[0]
}

// 压缩 从底至顶
func minimumTotal(triangle [][]int) int {
	l1 := len(triangle)
	l2 := len(triangle[l1-1])
	var f = make([]int, l2+1)
	for i := l2 - 1; 0 <= i; i-- {
		for j := 0; j <= i; j++ {
			f[j] = min(f[j], f[j+1]) + triangle[i][j]
		}
	}

	return f[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```

# 0121.best-time-to-buy-and-sell-stock 买卖股票的最佳时机 
```go
package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock

func maxProfit(prices []int) int {
	var min, max = 1<<63 - 1, 0
	for _, p := range prices {
		if p < min {
			min = p
		}
		if max < p-min {
			max = p - min
		}
	}
	return max
}

```

# 0123.best-time-to-buy-and-sell-stock-iii 买卖股票的最佳时机 III 
```go
package main

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii

func maxProfit(prices []int) int {
	a, b := -prices[0], 0
	c, d := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		a = max(a, -prices[i])  // 最少钱购买
		b = max(b, a+prices[i]) // 最多钱卖出
		c = max(c, b-prices[i]) // 最少钱购买
		d = max(d, c+prices[i]) // 最多钱卖出
	}
	return d
}

func maxProfit(prices []int) int {
	count := 2
	n := len(prices)
	dp := make([][3][2]int, n)

	for i := 0; i < len(prices); i++ {
		for j := count; 1 <= j; j-- {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}

			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j-1][0]-prices[i], dp[i-1][j][1])
		}
	}

	return dp[len(prices)-1][count][0]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

```

# 0139.word-break 单词拆分 
```go
package main

// https://leetcode-cn.com/problems/word-break

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, w := range wordDict {
		m[w] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for r := 1; r <= n; r++ {
		for l := 0; l < r; l++ {
			if dp[l] && m[s[l:r]] {
				dp[r] = true
				break
			}
		}
	}
	return dp[n]
}

```

# 0279.perfect-squares 完全平方数 
```go
package main

// https://leetcode-cn.com/problems/perfect-squares

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = n + 1
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```

# 0303.range-sum-query-immutable 区域和检索 - 数组不可变 
```go
package main

// https://leetcode-cn.com/problems/range-sum-query-immutable

// 前缀
type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	var n = len(nums)
	var s = make([]int, n)
	s[0] = nums[0]
	for i := 1; i < n; i++ {
		s[i] = s[i-1] + nums[i]
	}
	return NumArray{nums: s}
}

func (n *NumArray) SumRange(left int, right int) int {
	if 0 == left {
		return n.nums[right]
	}
	return n.nums[right] - n.nums[left-1]
}

```

# 0322.coin-change 零钱兑换 
```go
package main

// https://leetcode-cn.com/problems/coin-change

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	dp[0] = 0
	// 每种amount 有多少种组合
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if amount < dp[amount] {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1) // 1 1  0 2
			}
		}
	}

	if amount < dp[amount] {
		return -1
	}

	return dp[amount]
}

```

# 0474.ones-and-zeroes 一和零 
```go
package main

// https://leetcode-cn.com/problems/ones-and-zeroes

func findMaxForm(strs []string, m, n int) int {
	length := len(strs)
	dp := make([][][]int, length+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i, s := range strs {
		zeros, ones := get(s)
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i+1][j][k] = dp[i][j][k]
				if zeros <= j && ones <= k {
					// 第i个
					dp[i+1][j][k] = max(dp[i+1][j][k], dp[i][j-zeros][k-ones]+1)
				}
			}
		}
	}
	return dp[length][m][n]
}

// 输入：strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3
func findMaxForm(strs []string, m, n int) int {
	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zeros, ones := get(s)

		for i := m; zeros <= i; i-- {
			for j := n; ones <= j; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func get(s string) (zeros, ones int) {
	for _, v := range s {
		if v == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

```

# 0509.fibonacci-number 斐波那契数 
```go
package main

// https://leetcode-cn.com/problems/fibonacci-number

// f(n) = f(n-1) + f(n-2)
func climbStairs(n int) int {
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// 压缩
func fib(n int) int {
	var x, y = 0, 1
	for i := 1; i <= n; i++ {
		y, x = (x+y)%1000000007, y
	}
	return x
}

```

# 0518.coin-change-2 零钱兑换 II 
```go
package main

// https://leetcode-cn.com/problems/coin-change-2

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	//多加一种硬币 ，dp[i] 可以多几种跳跃方式
	for _, coin := range coins {
		for i := 1; i <= amount; i++ {
			if coin <= i {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[amount]
}

func wrong(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	// 多加一块钱 dp[i] 可以多几种跳跃方式  like（斐波那契数）
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] += dp[i-coin]
			}
		}
	}
	return dp[amount]
}

```

# 0523.continuous-subarray-sum 连续子数组和 K的倍数 
```go
package main

// https://leetcode-cn.com/problems/continuous-subarray-sum

// 前缀余数
func checkSubarraySum(nums []int, k int) bool {
	l := len(nums)
	if l < 2 {
		return false
	}
	m := map[int]int{0: -1}
	r := 0
	for i, v := range nums {
		r = (r + v) % k
		index, ok := m[r]
		if ok {
			if 2 <= i-index {
				return true
			}
		} else {
			m[r] = i
		}
	}
	return false
}

```

# 0525.contiguous-array 连续数组 
```go
package main

// https://leetcode-cn.com/problems/contiguous-array

func findMaxLength(nums []int) int {
	m := map[int]int{0: -1}
	var counter, res int
	for i, v := range nums {
		if v == 1 {
			counter++
		} else {
			counter--
		}
		index, ok := m[counter]
		if ok {
			if res < i-index {
				res = i - index
			}
		} else {
			m[counter] = i
		}
	}
	return res
}

```

# 0877.stone-game 石头游戏 
```go
package main

// https://leetcode-cn.com/problems/stone-game

func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}

	for r := 0; r < n; r++ {
		for l := 0; l < r; l++ {
			dp[l][r] = max(piles[l]-dp[l+1][r], piles[r]-dp[l][r-1])
		}
	}

	return 0 < dp[0][n-1]
}

// 压缩
func stoneGame(piles []int) bool {
	top := len(piles) - 1
	dp := make([]int, top+1)
	dp[0] = piles[0]

	for r := 0; r < top; r++ {
		for l := 0; l < r; l++ {
			dp[l] = max(piles[l]-dp[l+1], piles[r]-dp[r-1])
		}
	}
	return 0 < max(piles[0]-dp[1], piles[top]-dp[top-1])
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

```

# 1744.can-you-eat-your-favorite-candy-on-your-favorite-day 你能在你最喜欢的那天吃到你最喜欢的糖果吗？ 
```go
package main

// https://leetcode-cn.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day

// 前缀和
func canEat(candiesCount []int, queries [][]int) []bool {
	var res []bool

	for i := 1; i < len(candiesCount); i++ {
		candiesCount[i] += candiesCount[i-1]
	}

	for _, q := range queries {
		min := q[1] + 1
		max := q[2] * min

		var left int
		var right = candiesCount[q[0]]
		if 0 == q[0] {
			left = 1
		} else {
			left = candiesCount[q[0]-1] + 1
		}

		// 最大够不到最左  最小超过最右 都为false
		//res = append(res, !(max < left || right < min))
		res = append(res, left <= max && min <= right)
	}
	return res
}

```

