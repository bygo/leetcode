package main

import "strconv"

// https://leetcode-cn.com/problems/minimum-unique-word-abbreviation

// ❓ 最短独占单词缩写

func minAbbreviation(target string, dictionary []string) string {
	targetL := len(target)
	subsetBanMp := map[int]struct{}{}
	for _, word := range dictionary {
		wordL := len(word)
		// 不相等 永不相同
		if wordL != targetL {
			continue
		}

		var subset int
		for idx := 0; idx < wordL; idx++ {
			if target[idx] == word[idx] { // 相同
				subset += 1 << idx
			}
		}
		sub := subset
		for {
			// 所有禁止的子集
			subsetBanMp[sub] = struct{}{}
			if sub == 0 {
				break
			}
			sub = (sub - 1) & subset
			//if subset == num {
			//	break
			//}
		}
	}

	if len(subsetBanMp) == 0 {
		return strconv.Itoa(targetL)
	}
	var strRes = target
	subsetMax := 1<<targetL - 1 // 本身已计算
	for subset := 0; subset < subsetMax; subset++ {
		_, ok := subsetBanMp[subset]
		if ok {
			continue
		}

		var cnt int
		var buf []byte

		for idx := 0; idx < targetL; idx++ {
			if subset>>idx&1 == 1 {
				if 0 < cnt {
					buf = append(buf, strconv.Itoa(cnt)...)
					cnt = 0
				}
				buf = append(buf, target[idx])
			} else {
				// 未出现的压缩
				cnt++
			}
		}
		if 0 < cnt {
			buf = append(buf, strconv.Itoa(cnt)...)
		}

		strCur := string(buf)
		if len(strCur) < len(strRes) {
			strRes = strCur
		}
	}
	return strRes
}

// ban 字符串

func minAbbreviation(target string, dictionary []string) string {
	subsetBanMp := map[string]struct{}{}
	var targetL = len(target)

	for _, dict := range dictionary {
		var dictL = len(dict)
		if targetL != dictL {
			continue
		}
		var numMax = 1 << dictL
		for subset := 0; subset < numMax; subset++ {
			var cnt int
			var buf []byte
			for idx := 0; idx < dictL; idx++ {
				if subset>>idx&1 == 1 {
					if 0 < cnt {
						buf = append(buf, strconv.Itoa(cnt)...)
						cnt = 0
					}
					buf = append(buf, dict[idx])
				} else {
					cnt++
				}
			}
			if 0 < cnt {
				buf = append(buf, strconv.Itoa(cnt)...)
			}
			subsetBanMp[string(buf)] = struct{}{}
		}
	}
	if len(subsetBanMp) == 0 {
		return strconv.Itoa(targetL)
	}
	var strRes = target
	var numMax = 1<<targetL - 1
	for subset := 0; subset < numMax; subset++ {
		var cnt int
		var buf []byte
		for pos := 0; pos < targetL; pos++ {
			if subset>>pos&1 == 1 {
				if 0 < cnt {
					buf = append(buf, strconv.Itoa(cnt)...)
					cnt = 0
				}
				buf = append(buf, target[pos])
			} else {
				cnt++
			}
		}
		if 0 < cnt {
			buf = append(buf, strconv.Itoa(cnt)...)
		}
		strCur := string(buf)
		if len(strCur) < len(strRes) {
			_, ok := subsetBanMp[strCur]
			if !ok {
				strRes = strCur
			}
		}
	}
	return strRes
}
