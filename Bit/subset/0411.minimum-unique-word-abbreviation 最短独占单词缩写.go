package main

import "strconv"

// https://leetcode.cn/problems/minimum-unique-word-abbreviation

// ❓ 最短独占单词缩写

func minAbbreviation(target string, dictionary []string) string {
	// 禁止子集
	targetL := len(target)
	subsetMp := map[int]bool{}
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
			subsetMp[sub] = true
			if sub == 0 {
				break
			}
			sub = (sub - 1) & subset
			//if subset == num {
			//	break
			//}
		}
	}

	if len(subsetMp) == 0 {
		// 如果有长度相同，0必定在即合理
		return strconv.Itoa(targetL)
	}

	// 计算 target 所有子集
	var strRes = target
	subsetMax := 1<<targetL - 1 // 本身已计算
	for subset := 1; subset < subsetMax; subset++ {
		if subsetMp[subset] {
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

		if len(buf) < len(strRes) {
			strRes = string(buf)
		}
	}
	return strRes
}

// ban 字符串

func minAbbreviation(target string, dictionary []string) string {
	subsetMp := map[string]bool{}
	var targetL = len(target)

	for _, word := range dictionary {
		var wordL = len(word)
		if targetL != wordL {
			continue
		}
		var subsetMax = 1 << wordL
		for subset := 0; subset < subsetMax; subset++ {
			var cnt int
			var buf []byte
			for idx := 0; idx < wordL; idx++ {
				if subset>>idx&1 == 1 {
					if 0 < cnt {
						buf = append(buf, strconv.Itoa(cnt)...)
						cnt = 0
					}
					buf = append(buf, word[idx])
				} else {
					cnt++
				}
			}
			if 0 < cnt {
				buf = append(buf, strconv.Itoa(cnt)...)
			}
			subsetMp[string(buf)] = true
		}
	}
	if len(subsetMp) == 0 {
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
			if subsetMp[strCur] {
				strRes = strCur
			}
		}
	}
	return strRes
}
