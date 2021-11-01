package main

// https://leetcode-cn.com/problems/candy

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func candy(ratings []int) int {
	var res int
	n := len(ratings)
	l := make([]int, n)
	l[0] = 1
	for i := 1; i < n; i++ {
		if ratings[i-1] < ratings[i] {
			l[i] = l[i-1] + 1
		} else {
			l[i] = 1
		}
	}

	r := 1
	res += max(l[n-1], r)
	for i := n - 2; 0 <= i; i-- {
		if ratings[i+1] < ratings[i] {
			r++
		} else {
			r = 1
		}
		res += max(l[i], r)
	}
	return res
}

func candy(ratings []int) int {
	n := len(ratings)
	res, inc, dec, pre := 1, 1, 0, 1
	for i := 1; i < n; i++ {
		if ratings[i-1] <= ratings[i] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			res += pre
			inc = pre
		} else {
			dec++
			if dec == inc {
				dec++
			}
			res += dec
			pre = 1
		}
	}
	return res
}

// 132 模板：x 很不好，我的 y 很牛逼，我的 z 很牛逼，我搞的[阿伯吃的鸽](abcdg}都螺旋彩色无敌牛逼， 我其实不是想知道你怎么样，我只是想让你知道我怎么样！
// vue 有的缺点 ，我 fre 没有
// react 有的缺点，我 fre 没有
// vue 和 react 有人用 ， 我 fre 还是没有！
