package main

// https://leetcode.cn/problems/wildcard-matching

func isMatch(s string, p string) bool {
	idxS := len(s) - 1
	idxP := len(p) - 1
	// TODO 后缀普通匹配
	for 0 <= idxS && 0 <= idxP && p[idxP] != '*' {
		if !charMatch(s[idxS], p[idxP]) {
			return false
		}
		s = s[:idxS]
		p = p[:idxP]
		idxS--
		idxP--
	}
	if idxP == -1 {
		return idxS == -1
	}
	idxS, idxP = 0, 0
	sL := len(s)
	pL := len(p)
	sRecord, pRecord := -1, -1
	for idxS < sL && pRecord < pL {
		if p[idxP] == '*' { // *匹配，先跳过记录位置，下次匹配失败时，可从*重新开始
			idxP++
			sRecord, pRecord = idxS, idxP
		} else if charMatch(s[idxS], p[idxP]) { // 普通匹配成功
			idxS++
			idxP++
		} else if -1 < sRecord && sRecord+1 < sL { // 有记录, 直接贪心(也像回溯)，获得 sRecord++
			sRecord++
			idxS, idxP = sRecord, pRecord
		} else {
			return false
		}
	}
	// 剩下必须全为*
	for idx := idxP; idx < pL; idx++ {
		if p[idx] != '*' {
			return false
		}
	}
	return true
}

func charMatch(s, p byte) bool {
	return s == p || p == '?'
}
