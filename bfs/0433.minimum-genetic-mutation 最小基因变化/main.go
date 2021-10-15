package main

// https://leetcode-cn.com/problems/minimum-genetic-mutation

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
	queue := []string{start}
	res := 1
	base := []byte{'A', 'C', 'G', 'T'}
	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		for i := 0; i < cnt; i++ {
			b := []byte(queue[i])
			for j := 0; j < 8; j++ {
				for k := 0; k < 4; k++ {
					if b[j] != base[k] {
						b[j], base[k] = base[k], b[j]
						s := string(b)
						if s == end {
							return res
						}
						if _, ok := m[s]; ok {
							queue = append(queue, s)
						}
						b[j], base[k] = base[k], b[j]
					}
				}
			}
		}
		res++
		queue = queue[cnt:]
	}
	return -1
}
