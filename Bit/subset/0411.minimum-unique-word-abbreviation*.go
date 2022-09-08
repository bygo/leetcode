package main

import "strconv"

// https://leetcode.com/problems/minimum-unique-word-abbreviation

func minAbbreviation(target string, dictionary []string) string {
	tL := len(target)
	banMp := map[int]bool{}
	for _, word := range dictionary {
		wL := len(word)
		if wL != tL {
			continue
		}

		var subset int
		for idx := 0; idx < wL; idx++ {
			if target[idx] == word[idx] {
				subset += 1 << idx
			}
		}
		sub := subset
		for {
			// TODO
			banMp[sub] = true
			if sub == 0 {
				break
			}
			sub = (sub - 1) & subset
			//if sub == subsetMax {
			//	break
			//}
		}
	}

	if len(banMp) == 0 {
		return strconv.Itoa(tL)
	}

	var bufRes = []byte(target)
	subsetMax := 1<<tL - 1                          // = subsetMax is computed for itself
	for subset := 1; subset < subsetMax; subset++ { // = 0 is the full abbreviation, which is definitely not the answer
		if banMp[subset] {
			continue
		}

		var cnt int
		var buf []byte

		for idx := 0; idx < tL; idx++ {
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

		if len(buf) < len(bufRes) {
			bufRes = buf
		}
	}
	return string(bufRes)
}
