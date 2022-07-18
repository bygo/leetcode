package main

// https://leetcode.cn/problems/sentence-similarity

// ❓ 句子相似性 相似绑定
// ⚠️ great 和 fine 相似 = fine 和 great 相似
// ⚠️ fine 与 fine 相似

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	str1MpStr2MpBool := map[string]map[string]bool{}
	for _, pair := range similarPairs {
		str1 := pair[0]
		str2 := pair[1]
		if str1MpStr2MpBool[str1] == nil {
			str1MpStr2MpBool[str1] = map[string]bool{}
		}
		if str1MpStr2MpBool[str2] == nil {
			str1MpStr2MpBool[str2] = map[string]bool{}
		}
		str1MpStr2MpBool[str1][str2] = true
		str1MpStr2MpBool[str2][str1] = true
	}

	// 长度不一
	if len(sentence1) != len(sentence2) {
		return false
	}

	for i := range sentence1 {
		//if sentence1[i] == sentence2[i] || str1MpStr2MpBool[sentence1[i]][sentence2[i]] || str1MpStr2MpBool[sentence2[i]][sentence1[i]] {
		//	continue
		//} else {
		//	return true
		//}

		// 三种情况下都不相似
		if sentence1[i] != sentence2[i] &&
			!str1MpStr2MpBool[sentence1[i]][sentence2[i]] &&
			!str1MpStr2MpBool[sentence2[i]][sentence1[i]] {
			return false
		}
	}
	return true
}
