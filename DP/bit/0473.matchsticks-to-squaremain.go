package bit

// https://leetcode.cn/problems/matchsticks-to-square

func makesquare(matchsticks []int) bool {
	total := 0
	for _, matchstick := range matchsticks {
		total += matchstick
	}
	if total%4 != 0 {
		return false
	}
	cnt := total / 4

	mL := len(matchsticks)
	numMax := 1 << mL
	f := make([]int, numMax)
	for i := 1; i < numMax; i++ {
		f[i] = -1
	}

	for subset := 1; subset < numMax; subset++ {
		for idx, num := range matchsticks {
			if subset>>idx&1 == 0 {
				continue
			}
			sub := subset ^ (1 << idx)
			if 0 <= f[sub] && f[sub]+num <= cnt { // TODO 寻找路径，只要能填入，就是合法，不管从而何来
				f[subset] = (f[sub] + num) % cnt
				break
			}
			// 5 6 4 = 9
		}
	}
	return f[numMax-1] == 0
}
