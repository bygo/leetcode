package main

// https://leetcode.cn/problems/strings-differ-by-one-character

// ❓ 只有一个不同字符的字符串

func main() {
	println(differByOne([]string{"abcd", "acbd", "aacd"}))
}

const base = 131

func differByOne(dict []string) bool {
	dictL := len(dict)
	if dictL == 0 {
		return false
	}
	hashMp := map[uint]struct{}{}
	strMpHash := map[string]uint{}
	strL := len(dict[0])
	mulPre := make([]uint, strL+1)
	mulPre[0] = 1
	for i := 1; i <= strL; i++ {
		mulPre[i] = mulPre[i-1] * base
	}

	// 如果为true的答案多，提前计算，增加匹配概率
	for _, str := range dict {
		var hash uint
		for i := 0; i < strL; i++ {
			hash = hash*base + uint(str[i])
		}
		strMpHash[str] = hash
		hashMp[hash] = struct{}{}
	}

	for _, str := range dict {
		// 寻找相似
		for i := 0; i < strL; i++ {
			hashTmp := strMpHash[str] - uint(str[i])*mulPre[strL-i-1]
			var j byte = 'a'
			for ; j <= 'z'; j++ {
				if j == str[i] {
					continue
				}
				hashCur := hashTmp + uint(j)*mulPre[strL-i-1]
				_, ok := hashMp[hashCur]
				if ok {
					return true
				}
			}
		}
	}
	return false
}
