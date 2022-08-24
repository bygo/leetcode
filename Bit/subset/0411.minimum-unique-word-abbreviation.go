package main

import "strconv"

// https://leetcode-cn.com/problems/minimum-unique-word-abbreviation

func minAbbreviation(target string, dictionary []string) string {
	targetL := len(target)
	banMp := map[int]bool{}
	for _, word := range dictionary {
		wordL := len(word)
		if wordL != targetL {
			continue
		}

		var subsetMax int
		for idx := 0; idx < wordL; idx++ {
			if target[idx] == word[idx] {
				subsetMax += 1 << idx
			}
		}
		sub := subsetMax
		for {
			banMp[sub] = true
			if sub == 0 {
				break
			}
			sub = (sub - 1) & subsetMax
			//if subset == num {
			//	break
			//}
		}
	}

	if len(banMp) == 0 {
		return strconv.Itoa(targetL)
	}

	var strRes = target
	subsetMax := 1<<targetL - 1                     // =subsetMax is computed for itself
	for subset := 1; subset < subsetMax; subset++ { // =0 is the full abbreviation, which is definitely not the answer
		if banMp[subset] {
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
	subsetBanMp := map[string]bool{}
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
			subsetBanMp[string(buf)] = true
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
		if len(strRes) <= len(strCur) {
			continue
		}
		if subsetBanMp[strCur] {
			continue
		}
		strRes = strCur
	}
	return strRes
}
