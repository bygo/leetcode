package main

import "strconv"

// https://leetcode-cn.com/problems/minimum-unique-word-abbreviation

// ❓ 最短独占单词缩写

func minAbbreviation(target string, dictionary []string) string {
	targetL := len(target)
	banMp := map[int]struct{}{}
	for _, word := range dictionary {
		wordL := len(word)
		// 不相等 永不相同
		if wordL != targetL {
			continue
		}

		var subset int
		for idx := 0; idx < wordL; idx++ {
			if target[idx] == word[idx] { // 相同的位 必须加入 禁止子集
				subset += 1 << idx
			}
		}
		sub := subset
		for {
			// 所有禁止的子集
			banMp[sub] = struct{}{}
			if sub == 0 {
				break
			}
			sub = (sub - 1) & subset
			//if subset == num {
			//	break
			//}
		}
	}

	if len(banMp) == 0 {
		return strconv.Itoa(targetL)
	}

	var strRes = target
	subsetMax := 1<<targetL - 1                     // =subsetMax 为本身 已计算
	for subset := 1; subset < subsetMax; subset++ { // =0 为全缩略，肯定不是答案，
		_, ok := banMp[subset]
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
	targetL := len(target)

	for _, word := range dictionary {
		var wordL = len(word)
		if targetL != wordL {
			continue
		}
		var subsetMax = 1<<wordL - 1
		for subset := 1; subset < subsetMax; subset++ {
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
			subsetBanMp[string(buf)] = struct{}{}
		}
	}
	if len(subsetBanMp) == 0 {
		return strconv.Itoa(targetL)
	}
	var strRes = target
	var subsetMax = 1<<targetL - 1
	for subset := 0; subset < subsetMax; subset++ {
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
