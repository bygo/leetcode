package main

import "strconv"

// https://leetcode-cn.com/problems/minimum-unique-word-abbreviation

//func minAbbreviation(target string, dictionary []string) string {
//	abbrMp := map[string]struct{}{}
//	var targetL = len(target)
//
//	for _, dict := range dictionary {
//		var dictL = len(dict)
//		if targetL != dictL {
//			continue
//		}
//		var numMax = 1 << dictL
//		for subset := 0; subset < numMax; subset++ {
//			var cnt int
//			var buf []byte
//			for idx := 0; idx < dictL; idx++ {
//				if subset>>idx&1 == 1 {
//					if 0 < cnt {
//						buf = append(buf, strconv.Itoa(cnt)...)
//						cnt = 0
//					}
//					buf = append(buf, dict[idx])
//				} else {
//					cnt++
//				}
//			}
//			if 0 < cnt {
//				buf = append(buf, strconv.Itoa(cnt)...)
//			}
//			abbrMp[string(buf)] = struct{}{}
//		}
//	}
//
//	var strRes = target
//	var numMax = 1 << targetL
//	for subset := 0; subset < numMax; subset++ {
//		var cnt int
//		var buf []byte
//		for idx := 0; idx < targetL; idx++ {
//			if subset>>idx&1 == 1 {
//				if 0 < cnt {
//					buf = append(buf, strconv.Itoa(cnt)...)
//					cnt = 0
//				}
//				buf = append(buf, target[idx])
//			} else {
//				cnt++
//			}
//		}
//		if 0 < cnt {
//			buf = append(buf, strconv.Itoa(cnt)...)
//		}
//		strCur := string(buf)
//		if len(strCur) < len(strRes) {
//			_, ok := abbrMp[strCur]
//			if !ok {
//				strRes = strCur
//			}
//		}
//	}
//	return strRes
//}

func minAbbreviation(target string, dictionary []string) string {
	targetL := len(target)
	subsetBanMp := map[int]struct{}{}
	for _, dict := range dictionary {
		dictL := len(dict)
		// 不相等 永不相同
		if dictL != targetL {
			continue
		}

		var num int
		for pos := 0; pos < dictL; pos++ {
			if target[pos] == dict[pos] {
				num |= 1 << pos
			}
		}
		subset := num
		for {
			// 所有禁止的子集
			subsetBanMp[subset] = struct{}{}
			subset = (subset - 1) & num
			if subset == num {
				break
			}
		}
	}

	if len(subsetBanMp) == 0 {
		return strconv.Itoa(targetL)
	}
	var strRes = target
	bitTarget := 1<<targetL - 1
	for subset := 0; subset < bitTarget; subset++ {
		_, ok := subsetBanMp[subset]
		if ok {
			continue
		}

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
			strRes = strCur
		}
	}
	return strRes
}
