package main

// https://leetcode.cn/problems/pyramid-transition-matrix

func pyramidTransition(bottom string, allowed []string) bool {
	strMpBuf := map[string][]byte{}
	for _, v := range allowed {
		strMpBuf[v[:2]] = append(strMpBuf[v[:2]], v[2])
	}

	var dfs func(buf []byte, idx int) bool
	dfs = func(buf []byte, idx int) bool {
		bufL := len(buf)
		if bufL == 1 {
			return true
		}
		if bufL == idx+1 {
			return dfs(buf[:bufL-1], 0) // TODO
		}
		str := string(buf[idx : idx+2])
		for _, ch := range strMpBuf[str] { // TODO
			buf[idx] = ch
			if dfs(buf, idx+1) {
				return true
			}
		}
		return false
	}

	return dfs([]byte(bottom), 0)
}
