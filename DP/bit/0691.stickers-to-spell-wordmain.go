package bit

// https://leetcode.cn/problems/stickers-to-spell-word

func minStickers(stickers []string, target string) int {
	tL := len(target)
	f := make([]int, 1<<tL)
	for i := range f {
		f[i] = -1
	}
	f[0] = 0
	var dfs func(int) int
	dfs = func(subset int) int {
		if f[subset] != -1 {
			return f[subset]
		}
		f[subset] = tL + 1
		for _, sticker := range stickers {
			sub := subset
			cnt := [26]int{}
			for _, ch := range sticker {
				cnt[ch-'a']++
			}
			for idx, ch := range target {
				if subset>>idx&1 == 0 {
					continue
				}

				if cnt[ch-'a'] == 0 { // 木有
					continue
				}
				cnt[ch-'a']--
				sub ^= 1 << idx
			}
			if sub < subset {
				f[subset] = min(f[subset], dfs(sub)+1)
			}
		}
		return f[subset]
	}
	cnt := dfs(1<<tL - 1)
	if cnt <= tL {
		return cnt
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
