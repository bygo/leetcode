package main

// https://leetcode.cn/problems/minimum-genetic-mutation

// ❓ 最小基因变化

func minMutation(start string, end string, bank []string) int {
	m := make(map[string]struct{})
	for _, b := range bank {
		m[b] = struct{}{}
	}
	if start == end {
		return 0
	}

	if _, ok := m[end]; !ok {
		return -1
	}
	que := []string{start}
	cnt := 1
	base := []byte{'A', 'C', 'G', 'T'}
	for {
		queL := len(que)
		if queL == 0 {
			break
		}
		for i := 0; i < queL; i++ {
			b := []byte(que[i])
			for j := 0; j < 8; j++ {
				for k := 0; k < 4; k++ {
					if b[j] != base[k] {
						b[j], base[k] = base[k], b[j]
						s := string(b)
						if s == end {
							return cnt
						}
						if _, ok := m[s]; ok {
							que = append(que, s)
						}
						b[j], base[k] = base[k], b[j]
					}
				}
			}
		}
		cnt++
		que = que[queL:]
	}
	return -1
}
