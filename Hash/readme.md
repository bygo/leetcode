# Cnt ✅ 0389.find-the-difference 找不同 
```go
package main

// https://leetcode-c 	n.com/problems/find-the-difference

// ❓ 找出 x
// ⚠️ s = t + 随机byte(x)

func findTheDifferenceCnt(s string, t string) byte {
	// 统计
	chMpCnt := [26]int{}
	for i := range s {
		ch := s[i] - 'a'
		chMpCnt[ch]++
	}

	// map抵消
	for i := range t {
		ch := t[i] - 'a'
		if chMpCnt[ch] == 0 {
			return t[i]
		}
		chMpCnt[ch]--
	}
	return 0
}

func findTheDifferenceDistance(s string, t string) byte {
	// sum抵消
	var sum = 0
	for i := range s {
		sum = sum - int(s[i]) + int(t[i])
	}
	return byte(sum + int(t[len(t)-1]))
}

func findTheDifferenceXor(s string, t string) byte {
	// 异或抵消
	var sum byte
	for i := range s {
		sum ^= s[i] ^ t[i]
	}
	return sum ^ t[len(t)-1]
}

```

# Cnt ✅ 0811.subdomain-visit-count 子域名访问次数 
```go
package main

import (
	"strconv"
)

// https://leetcode-cn.com/problems/subdomain-visit-count

// ❓ 子域名访问次数
// ⚠️ 10 www.by.com

func subdomainVisits(cpdomains []string) []string {
	var hostMpCnt = map[string]int{}
	for _, str := range cpdomains {
		// 出现次数
		var cnt int
		var idx int
		var strL = len(str)
		for str[idx] != ' ' {
			cnt = cnt*10 + int(str[idx]-'0')
			idx++
		}
		idx++

		// 域名
		for idx < strL {
			hostMpCnt[str[idx:]] += cnt
			for idx < strL && str[idx] != '.' { // 移动到 .
				idx++
			}
			idx++ // 去除 .
		}
	}
	var cntMp []string
	for host, cnt := range hostMpCnt {
		cntMp = append(cntMp, strconv.Itoa(cnt)+" "+host)
	}
	return cntMp
}

```

# Cnt ✅ 0888.fair-candy-swap 糖果棒交换后数量相等 
```go
package main

// https://leetcode-cn.com/problems/fair-candy-swap

// ❓ 交换一个值后，总和相等

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	// aliceSizes map
	// aliceSizes 糖果数
	valMp := map[int]struct{}{}
	sum1 := 0
	for i := range aliceSizes {
		valMp[aliceSizes[i]] = struct{}{}
		sum1 += aliceSizes[i]
	}

	// bobSizes 糖果数
	sum2 := 0
	for i := range bobSizes {
		sum2 += bobSizes[i]
	}

	// 一增一减
	// 必为偶数
	diff := sum1 - sum2
	if diff%2 == 1 {
		// 并且因奇数向下取整，差值还会少1，虽然循环找出结果，但不是答案
		return nil
	}
	diff /= 2

	// 找差额
	for i := range bobSizes {
		aliceVal := bobSizes[i] + diff
		_, ok := valMp[aliceVal]
		if ok {
			return []int{aliceVal, bobSizes[i]}
		}
	}

	return nil
}

```

# Comb Trie ✅ 0720.longest-word-in-dictionary 词典中连续递接 的最长单词 
```go
package main

import "sort"

// https://leetcode-cn.com/problems/longest-word-in-dictionary

// ❓词典中连续递接 的最长单词

func longestWord(words []string) string {
	sort.Strings(words)
	strMp := map[string]struct{}{}

	var longestStr string
	for _, str := range words {
		strL := len(str)
		if len(str) == 1 {
			// 单个字符 直接加入
			strMp[str] = struct{}{}
			// 比答案更长
			if len(longestStr) == 0 {
				longestStr = str
			}
		} else {
			// 多个字符 判断是否有前缀
			_, ok := strMp[str[:strL-1]]
			if ok {
				strMp[str] = struct{}{}
				// 比答案更长
				if len(longestStr) < strL {
					longestStr = str
				}
			}
		}

	}
	return longestStr
}

type trie struct {
	child [26]*trie
	s     string
}

func longestWordTrie(words []string) string {
	root := &trie{}
	for _, word := range words {
		n := root
		for j := range word {
			ch := word[j] - 'a'
			if n.child[ch] == nil {
				n.child[ch] = &trie{}
			}
			n = n.child[ch]
		}
		n.s = word
	}

	var longestStr string
	var queue = []*trie{root}
	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		longestStr = queue[0].s
		for i := 0; i < cnt; i++ {
			for j := 0; j < 26; j++ {
				if queue[i].child[j] != nil && queue[i].child[j].s != "" {
					queue = append(queue, queue[i].child[j])
				}
			}
		}

		queue = queue[cnt:]
	}
	return longestStr
}

```

# Comb ✅ 0001.two-sum 两数之和 
```go
package main

// https://leetcode-cn.com/problems/two-sum/

// ❓ 两数之和

func twoSum(nums []int, target int) []int {
	var numMpIdx = map[int]int{}
	for idxRaw, numRaw := range nums {
		numHash := target - numRaw
		idxHash, ok := numMpIdx[numHash]
		if ok {
			return []int{idxHash, idxRaw}
		}
		numMpIdx[numRaw] = idxRaw
	}
	return nil
}

```

# Comb ✅ 0348.design-tic-tac-toe 设计井字棋 
```go
package main

// https://leetcode-cn.com/problems/design-tic-tac-toe

// ❓设计无限大井字棋 n不止3

var symbols = map[int]int{
	1: 1,
	2: -1,
}

type TicTacToe struct {
	rowMpCnt []int
	colMpCnt []int
	angMpCnt [2]int
	n        int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		rowMpCnt: make([]int, n),
		colMpCnt: make([]int, n),
		angMpCnt: [2]int{},
		n:        n,
	}
}

func (t *TicTacToe) Move(row int, col int, player int) int {
	symbol := symbols[player]
	t.rowMpCnt[row] += symbol
	if t.win(t.rowMpCnt[row]) {
		return player
	}
	t.colMpCnt[col] += symbol
	if t.win(t.colMpCnt[col]) {
		return player
	}

	if col == row {
		t.angMpCnt[0] += symbol
		if t.win(t.angMpCnt[0]) {
			return player
		}
	}
	if col+row == t.n-1 {
		t.angMpCnt[1] += symbol
		if t.win(t.angMpCnt[1]) {
			return player
		}
	}
	return 0
}

func (t *TicTacToe) win(i int) bool {
	return i == t.n || i == -t.n
}

```

# Comb ✅ 0500.keyboard-row 键盘同一行的单词数 
```go
package main

// https://leetcode-cn.com/problems/keyboard-row

// ❓ 在键盘同一行的单词数

var chMpState = map[byte]uint8{}

func init() {
	for _, ch := range []byte("qwertyuiopQWERTYUIOP") {
		chMpState[ch] = 1
	}
	for _, ch := range []byte("asdfghjklASDFGHJKL") {
		chMpState[ch] = 2
	}
	for _, ch := range []byte("zxcvbnmZXCVBNM") {
		chMpState[ch] = 3
	}
}

func findWords(words []string) []string {
	var wordSame []string
	for _, word := range words {
		wordL := len(word)
		chZero := word[0]
		idx := 1
		for idx < wordL {
			chCur := word[idx]
			if chMpState[chCur] != chMpState[chZero] {
				break
			}
			idx++
		}
		if idx == wordL {
			wordSame = append(wordSame, word)
		}
	}
	return wordSame
}

```

# Comb ✅ 1128.number-of-equivalent-domino-pairs 组合相同的数量 
```go
package main

// https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs

// ❓ 组合相同的数量
// ⚠️ 一个组只有2个数字

func numEquivDominoPairs(d [][]int) int {
	combMpCnt := map[int]int{}
	var cnt int
	for i := range d {
		if d[i][0] < d[i][1] {
			d[i][0], d[i][1] = d[i][1], d[i][0]
		}
		comb := d[i][0]*10 + d[i][1]
		cnt += combMpCnt[comb]
		combMpCnt[comb]++
	}
	return cnt
}

```

# Comb ✅ 1275.find-winner-on-a-tic-tac-toe-game 找出井字棋的获胜者 
```go
package main

// https://leetcode-cn.com/problems/find-winner-on-a-tic-tac-toe-game

// ❓ 3x3 井字棋

const (
	A       = "A"
	B       = "B"
	Pending = "Pending"
	Draw    = "Draw"
)

func tictactoe(moves [][]int) string {
	movesL := len(moves)

	combMpCnt := [8]int{}

	// 倒序
	for i := movesL - 1; 0 <= i; i -= 2 {
		row := moves[i][0]
		col := moves[i][1] + 3
		combMpCnt[row]++
		combMpCnt[col]++
		if moves[i][0] == moves[i][1] {
			combMpCnt[6]++
		}
		if moves[i][0]+moves[i][1] == 2 {
			combMpCnt[7]++
		}
	}

	for i := range combMpCnt {
		if combMpCnt[i] != 3 {
			continue
		}
		// A先手
		if movesL%2 == 1 {
			return A
		}
		return B
	}
	if movesL < 9 {
		return Pending
	}

	return Draw
}

```

# Comb ✅ 1512.number-of-good-pairs 好数对的数目：i < j & a[i] == a[j] 
```go
package main

// https://leetcode-cn.com/problems/number-of-good-pairs

// ❓ 好数对的数目
// ⚠️ i < j && nums[i] == nums[j]

func numIdenticalPairs(nums []int) int {
	var numMpCnt = map[int]int{}
	var cntGood int
	for _, num := range nums {
		if 0 < numMpCnt[num] {
			cntGood += numMpCnt[num]
		}
		numMpCnt[num]++
	}
	return cntGood
}

```

# Group ✅ 0288.unique-word-abbreviation 单词是唯一缩写 
```go
package main

import (
	"strconv"
)

// https://leetcode-cn.com/problems/unique-word-abbreviation

// ❓ 单词是唯一的缩写

type ValidWordAbbr struct {
	comMpStrMpCnt map[string]map[string]int
}

func Constructor(dictionary []string) ValidWordAbbr {
	comMpStrMpCnt := map[string]map[string]int{}
	for _, str := range dictionary {
		com := Compress(str)
		if comMpStrMpCnt[com] == nil {
			comMpStrMpCnt[com] = map[string]int{}
		}
		comMpStrMpCnt[com][str] ++
	}
	return ValidWordAbbr{
		comMpStrMpCnt: comMpStrMpCnt,
	}
}

func (vwa *ValidWordAbbr) IsUnique(str string) bool {
	com := Compress(str)
	return vwa.comMpStrMpCnt[com] == nil || len(vwa.comMpStrMpCnt[com]) == 1 && 0 < vwa.comMpStrMpCnt[com][str]
}

// hello => h3o

func Compress(str string) string {
	wordL := len(str)
	return string(str[0]) + strconv.Itoa(wordL-2) + string(str[wordL-1])
}

```

# Group ✅ 0914.x-of-a-kind-in-a-deck-of-cards 卡牌分组 
```go
package main

// https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards

// ❓存在 int(x) 把数组分为:
// ⚠️ 每组cnt为 int(g)
// ⚠️ 每组val相同

func hasGroupsSizeX(deck []int) bool {
	// 统计元素
	numMpCnt := map[int]int{}
	for _, num := range deck {
		numMpCnt[num]++
	}

	// 最大公约数
	g := -1
	for num := range numMpCnt {
		if g == -1 {
			g = numMpCnt[num]
		} else {
			g = gcd(g, numMpCnt[num])
		}
	}
	return 2 <= g
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}

```

# Group ✅ 0929.unique-email-addresses 不同的邮件地址数 
```go
package main

// https://leetcode-cn.com/problems/unique-email-addresses

// ❓ 不同的邮件地址数
// b.y+ig@gmail.com
// .会被忽略
// +忽略 +~@内容
func numUniqueEmails(emails []string) int {
	mailMp := map[string]struct{}{}
	for _, email := range emails {
		emailL := len(email)
		var idx int
		var emailTmp []byte
		// 处理 b.y+ig
		for idx < emailL && email[idx] != '@' {
			if email[idx] == '+' {
				// 移动到后 @gmail.com
				for idx < emailL && email[idx] != '@' {
					idx++
				}
			} else {
				if email[idx] != '.' {
					emailTmp = append(emailTmp, email[idx])
				}
				idx++
			}
		}

		// 处理 @gmail.com
		emailTmp = append(emailTmp, email[idx:]...)

		mailMp[string(emailTmp)] = struct{}{}
	}
	return len(mailMp)
}

```

# Group ✅ 1399.count-largest-group 数位和相同为一组，含有最大数量的有多少组 
```go
package main

// https://leetcode-cn.com/problems/count-largest-group

// ❓ 数位和相同为一组，含有最大数量的有多少组

func countLargestGroup(n int) int {
	// 统计
	sumMpCnt := map[int]int{}
	for i := 1; i <= n; i++ {
		j := i
		sum := 0
		for 0 < j {
			sum += j % 10
			j /= 10
		}
		sumMpCnt[sum] ++
	}
	// 最大数量
	var cnt, cntLargest int
	for sum := range sumMpCnt {
		if cnt < sumMpCnt[sum] {
			cnt = sumMpCnt[sum]
			cntLargest = 1
		} else if sumMpCnt[sum] == cnt {
			cntLargest++
		}
	}
	return cntLargest
}

```

# Group ✅ 1742.maximum-number-of-balls-in-a-box 盒子中小球的最大数量：数位和 
```go
package main

// https://leetcode-cn.com/problems/maximum-number-of-balls-in-a-box

// ❓ 盒子中小球的最大数量

func countBalls(lowLimit int, hiLimit int) int {
	// 数位和
	var sum int
	j := lowLimit
	for 0 < j {
		sum += j % 10
		j /= 10
	}
	sumMpCnt := [46]int{} // 最大9999 = 46
	var cntMax int
	for i := lowLimit; i <= hiLimit; i++ {
		sumMpCnt[sum]++
		k := i
		for k%10 == 9 { // 每次进1 减9
			sum -= 9
			k /= 10
		}
		sum++
	}

	// 计算最大数量
	for sum := range sumMpCnt {
		if cntMax < sumMpCnt[sum] {
			cntMax = sumMpCnt[sum]
		}
	}

	return cntMax
}

```

# Group ✅ 1805.number-of-different-integers-in-a-string 字符串中不同整数的数目 
```go
package main

// https://leetcode-cn.com/problems/number-of-different-integers-in-a-string

// ❓ 字符串中不同整数的数目
// ⚠️ 0 01 001

func numDifferentIntegers(word string) int {
	strMp := map[string]struct{}{}
	var numBuf []byte

	// 整数结算
	cntNum := func() {
		tmpL := len(numBuf)
		if 0 < tmpL {
			j := 0
			// 保存至少一个零
			for numBuf[j] == '0' && j < tmpL-1 {
				j++
			}
			strMp[string(numBuf[j:])] = struct{}{}
			numBuf = []byte{}
		}
	}

	for i := range word {
		if '0' <= word[i] && word[i] <= '9' {
			numBuf = append(numBuf, word[i])
		} else {
			cntNum()
		}
	}
	cntNum()
	return len(strMp)
}

```

# Idx ✅ 0244.shortest-word-distance-ii 两单词距离最小索引 
```go
package main

// https://leetcode-cn.com/problems/shortest-word-distance-ii

// ❓ 两单词距离最小索引

type WordDistance struct {
	strMpIdxes map[string][]int
}

func Constructor(wordsDict []string) WordDistance {
	strMpIdxes := map[string][]int{}
	for i, str := range wordsDict {
		strMpIdxes[str] = append(strMpIdxes[str], i)
	}
	return WordDistance{
		strMpIdxes: strMpIdxes,
	}
}

func (wd *WordDistance) Shortest(word1 string, word2 string) int {
	idxes1 := wd.strMpIdxes[word1]
	idxes2 := wd.strMpIdxes[word2]
	i, j := len(idxes1)-1, len(idxes2)-1
	var distMin = 1<<63 - 1
	var distTmp int
	for -1 < i && -1 < j {
		if idxes1[i] < idxes2[j] {
			// 2逼近1
			distTmp = idxes2[j] - idxes1[i]
			j--
		} else {
			// 1逼近2
			distTmp = idxes1[i] - idxes2[j]
			i--
		}

		// 距离
		if distTmp < distMin {
			distMin = distTmp
		}
	}
	return distMin
}

```

# Idx ✅ 0599.minimum-index-sum-of-two-lists 两数组相同值的最小索引集合 
```go
package main

// https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists

// ❓ 两数组相同值的最小索引集合
// ⚠️ 没有重复

func findRestaurant(list1 []string, list2 []string) []string {
	strMpIdx := map[string]int{}
	for i, str := range list1 {
		strMpIdx[str] = i + 1
	}

	var stringsDist []string
	var distMin = 1<<63 - 1
	for i := range list2 {
		if 0 < strMpIdx[list2[i]] {
			// 存在时计算
			dist := strMpIdx[list2[i]] + i

			if dist < distMin {
				// 比当前还小
				stringsDist = []string{list2[i]}
				distMin = dist
			} else if dist == distMin {
				// 累加
				stringsDist = append(stringsDist, list2[i])
			}
		}
	}
	return stringsDist
}

```

# Idx ✅ 0760.find-anagram-mappings 在B找A的位置 
```go
package main

// https://leetcode-cn.com/problems/find-anagram-mappings

// ❓ 在B找A的位置

func anagramMappings(A []int, B []int) []int {
	// B 的位置
	var numMpIdx = map[int]int{}
	for idx, num := range B {
		numMpIdx[num] = idx
	}

	var idxesA []int
	for _, num := range A {
		idxesA = append(idxesA, numMpIdx[num])
	}
	return idxesA
}

```

# Idx ✅ 1165.single-row-keyboard 机械手最少移动次数 
```go
package main

// https://leetcode-cn.com/problems/single-row-keyboard

// ❓ 输入一个单词的移动长度总和
// ⚠️ 移动长度 = abs(i - j)

func calculateTime(keyboard string, word string) int {
	// 索引统计
	var chMpIdx = [26]int{}
	for idx := range keyboard {
		ch := keyboard[idx] - 'a'
		chMpIdx[ch] = idx
	}

	// 计算移动次数
	// 0
	ch0 := word[0] - 'a'
	var cntDist = chMpIdx[ch0]

	// 1～l
	for i := 1; i < len(word); i++ {
		chPrev := word[i-1] - 'a'
		chNext := word[i] - 'a'
		cntDist += abs(chMpIdx[chPrev] - chMpIdx[chNext])
		i++
	}

	return cntDist
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

```

# Inter ✅ 0299.bulls-and-cows 两数组 val idx 交集数 & val 交集数 
```go
package main

import "strconv"

// https://leetcode-cn.com/problems/bulls-and-cows

// ❓ 公牛母牛数
// 公牛: idx & val 交集
// 母牛: val 交集

func getHint(secret string, guess string) string {
	secretMpCnt := [10]int{}
	guessMpCNt := [10]int{}
	var bull, cow int
	for i := range guess {
		if guess[i] == secret[i] {
			bull++
		} else {
			secretMpCnt[secret[i]-'0']++
			guessMpCNt[guess[i]-'0']++
		}
	}

	for num := 0; num < 10; num++ {
		cow += min(secretMpCnt[num], guessMpCNt[num])
	}
	return "A" + strconv.Itoa(bull) + "B" + strconv.Itoa(cow)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```

# Inter ✅ 0349.intersection-of-two-arrays 两数组 val 交集 & 取唯一 
```go
package main

// https://leetcode-cn.com/problems/intersection-of-two-arrays

// ❓ 相对交集 val 取唯一值

func intersection(nums1 []int, nums2 []int) []int {
	var numMp = map[int]struct{}{}
	for _, num := range nums1 {
		numMp[num] = struct{}{}
	}

	var numsInter []int
	for _, num := range nums2 {
		_, ok := numMp[num]
		if ok {
			numsInter = append(numsInter, num)
			delete(numMp, num)
		}
	}
	return numsInter
}

```

# Inter ✅ 0350.intersection-of-two-arrays-ii 两数组 val 交集 
```go
package main

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii

// ❓ 两数组 val 交集

func intersect(nums1 []int, nums2 []int) []int {
	numMpCnt := map[int]int{}
	for _, num := range nums1 {
		numMpCnt[num]++
	}

	var numsInter []int
	for _, num := range nums2 {
		if 0 < numMpCnt[num] {
			numsInter = append(numsInter, num)
			numMpCnt[num]--
		}
	}
	return numsInter
}

```

# Inter ✅ 0771.jewels-and-stones 两数组 val 交集数 
```go
package main

// https://leetcode-cn.com/problems/jewels-and-stones/

// ❓ 两数组 val 交集数
func numJewelsInStones(j string, s string) int {
	var chMpCnt = map[byte]int{}
	for i := range j {
		chMpCnt[j[i]] = 1
	}

	var cntInter int
	for i := range s {
		cntInter += chMpCnt[s[i]]
	}
	return cntInter
}

```

# Inter ✅ 0819.most-common-word 两数组 val 最大交集数 
```go
package main

// https://leetcode-cn.com/problems/most-common-word

// ❓ 两数组 val 最大交集数

func mostCommonWord(paragraph string, banned []string) string {
	banMpCnt := map[string]int{}
	strMpCnt := map[string]int{}
	for i := range banned {
		banMpCnt[banned[i]] = 1
	}

	var strInter string
	var strTmp []byte
	var cntMax int

	fn := func() {
		str := string(strTmp)
		// 不在禁用列表
		if 0 < len(str) && 0 == banMpCnt[str] {
			strMpCnt[str]++
			if cntMax < strMpCnt[str] {
				strInter = str
				cntMax = strMpCnt[str]
			}
		}
		strTmp = []byte{}
	}
	for i := range paragraph {
		if 'A' <= paragraph[i] && paragraph[i] <= 'Z' {
			strTmp = append(strTmp, paragraph[i]+32)
		} else if 'a' <= paragraph[i] && paragraph[i] <= 'z' {
			strTmp = append(strTmp, paragraph[i])
		} else {
			fn()
		}
	}
	fn()
	return strInter
}

```

# Inter ✅ 1002.find-common-characters N个字符串 公共字符 
```go
package main

// https://leetcode-cn.com/problems/find-common-characters

// ❓ N个字符串 公共字符

func commonChars(words []string) []string {
	// 占位
	var chsCommon []string
	chMpCnt := [26]int{}
	for i := range chMpCnt {
		chMpCnt[i] = 1<<63 - 1
	}

	for _, word := range words {
		// 计数
		chMpCntTmp := [26]int{}
		for _, ch := range word {
			chMpCntTmp[ch-'a']++
		}
		// 计算最小值
		for ch, cntTmp := range chMpCntTmp {
			if cntTmp < chMpCnt[ch] {
				chMpCnt[ch] = cntTmp
			}
		}
	}

	// 字典序
	for i := 0; i < 26; i++ {
		for j := 0; j < chMpCnt[i]; j++ {
			chsCommon = append(chsCommon, string(byte(i+'a')))
		}
	}
	return chsCommon
}

```

# Link 0356.line-reflection 直线镜像 
```go
package main

// https://leetcode-cn.com/problems/line-reflection

func isReflected(points [][]int) bool {
	xMpYMpBool := map[int]map[int]bool{}
	max, min := -1<<63, 1<<63-1

	// 找出最大范围 并统计
	for _, point := range points {
		x := point[0]
		y := point[1]
		if max < x {
			max = x
		}
		if x < min {
			min = x
		}
		if xMpYMpBool[x] == nil {
			xMpYMpBool[x] = map[int]bool{}
		}
		xMpYMpBool[x][y] = true
	}

	// 使所有平行，必须找出中垂线

	for _, point := range points {
		x := max + min - point[0]
		y := point[1]
		// x 镜像不存在 或者 不平行
		if xMpYMpBool[x] == nil || !xMpYMpBool[x][y] {
			return false
		}
	}
	return true
}

```

# Link ✅ 2006.count-number-of-pairs-with-absolute-difference-k 差的绝对值为 K 的数对数目 
```go
package main

// https://leetcode-cn.com/problems/count-number-of-pairs-with-absolute-difference-k

func countKDifference(nums []int, k int) int {
	m := map[int]int{}
	var res int
	for _, num := range nums {
		res += m[num-k] + m[num+k]
		m[num]++
	}
	return res
}

```

# Link ✅ 2007.find-original-array-from-doubled-array 从双倍数组中还原原数组 
```go
package main

// https://leetcode-cn.com/problems/find-original-array-from-doubled-array

func findOriginalArray(changed []int) []int {
	if len(changed)%2 == 1 {
		return nil
	}
	// 计数
	numMpCnt := [100001]int{}
	var cnt int
	for _, num := range changed {
		numMpCnt[num]++
		cnt++
	}

	var numsOri []int
	for num := 0; num <= 50000; num++ {
		// 必须从小开始抵消
		for 0 < numMpCnt[num] {
			numMpCnt[num]--
			numMpCnt[num*2]--
			cnt -= 2
			numsOri = append(numsOri, num)
			if numMpCnt[num*2] < 0 {
				return nil
			}
		}
	}
	if 0 < cnt {
		return nil
	}
	return numsOri
}

```

# Link ✅ 2013.detect-squares 检测正方形 
```go
package main

// https://leetcode-cn.com/problems/detect-squares

// ❓ 输入的点，能形成几个正方形

type DetectSquares struct {
	xMpYMpCnt map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{xMpYMpCnt: map[int]map[int]int{}}
}

func (ds *DetectSquares) Add(p []int) {
	x := p[0]
	y := p[1]
	if ds.xMpYMpCnt[x] == nil {
		ds.xMpYMpCnt[x] = map[int]int{}
	}
	ds.xMpYMpCnt[x][y] ++
}

func (ds *DetectSquares) Count(p []int) int {
	x := p[0]
	y := p[1]
	yMpCnt := ds.xMpYMpCnt[x] // 所有x = x1 的 y 点
	var cntSquare int
	for yCur := range yMpCnt {
		dist := yCur - y // 需要的边长
		if dist != 0 {
			// 左边 右边
			yCurMpCnt := ds.xMpYMpCnt[x-dist] // x1-dis 的 y 点
			if yCurMpCnt != nil {
				cntSquare += yCurMpCnt[y] * yCurMpCnt[yCur] * yMpCnt[yCur] // 与其他三点成为正方形
			}

			yCurMpCnt = ds.xMpYMpCnt[x+dist] // x1+dis 的 y 点
			if yCurMpCnt != nil {
				cntSquare += yCurMpCnt[y] * yCurMpCnt[yCur] * yMpCnt[yCur] // 与其他三点成为正方形
			}
		}
	}
	return cntSquare
}

```

# Moore 投票 ✅ 0169.majority-element 众数 
```go
package main

// https://leetcode-cn.com/problems/majority-element

// ❓ 超过一半的数

func majorityElement(nums []int) int {
	var numMode, cnt int
	for _, num := range nums {
		if numMode == num {
			cnt++
		} else {
			if cnt == 0 {
				numMode = num
				cnt = 1
			} else {
				cnt--
			}
		}
	}
	return numMode
}

```

# Moore 投票 ✅ 0229.majority-element-ii 求众数 II 
```go
package main

// https://leetcode-cn.com/problems/majority-element-ii

// ❓ 超过1/n的数

type Elem struct {
	num, cnt int
}

const base = 3
const elemL = base - 1

func majorityElement(nums []int) []int {
	var l = len(nums)
	var elems [elemL]Elem
	for _, num := range nums {
		var slotInsert = -1
		var idx int
		// 找出第一个插槽
		for idx < elemL {
			if elems[idx].num == num {
				// 同值插槽
				break
			} else if elems[idx].cnt == 0 {
				// 第一个空槽
				slotInsert = idx
			}
			idx++
		}

		if idx < elemL {
			// 同值插槽
			elems[idx].cnt++
		} else if idx == elemL {
			if -1 < slotInsert {
				// 有空槽 替换
				elems[slotInsert].num = num
				elems[slotInsert].cnt = 1
			} else {
				// 全部自减
				for i := range elems {
					elems[i].cnt--
				}
			}
		}
	}

	// 置零
	for i := range elems {
		elems[i].cnt = 0
	}

	// 因为抵消，需要重新计数
	for _, num := range nums {
		for i := range elems {
			if elems[i].num == num {
				elems[i].cnt++
				break
			}
		}
	}

	// 多个答案
	var limit = l / base
	var numsMode []int
	for i := range elems {
		if limit < elems[i].cnt {
			numsMode = append(numsMode, elems[i].num)
		}
	}
	return numsMode
}

```

# Order ✅ 0205.isomorphic-strings 同构字符串 
```go
package main

// https://leetcode-cn.com/problems/isomorphic-strings

func isIsomorphic(s string, t string) bool {
	ch1MpCh2 := map[byte]byte{}
	ch2MpCh1 := map[byte]byte{}
	sL := len(s)
	for i := 0; i < sL; i++ {
		ch1 := s[i]
		ch2 := t[i]



		if ch1MpCh2[ch1] != 0 && ch1MpCh2[ch1] != ch2 || ch2MpCh1[ch2] != 0 && ch2MpCh1[ch2] != ch1 {
			return false
		}
		ch1MpCh2[ch1] = ch2
		ch2MpCh1[ch2] = ch1
	}

	return true
}

```

# Order ✅ 0290.word-pattern  句子与单词 双向绑定 
```go
package main

import "strings"

// https://leetcode-cn.com/problems/word-pattern

// ❓ 句子与单词 双向绑定

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	var strMpCh = map[string]byte{}
	var chMpStr = map[byte]string{}
	for i := range pattern {
		ch := strMpCh[words[i]]
		str := chMpStr[pattern[i]]
		if ch != 0 && ch != pattern[i] || str != "" && str != words[i] {
			return false
		}
		strMpCh[words[i]] = pattern[i]
		chMpStr[pattern[i]] = words[i]

	}
	return true
}

```

# Order ✅ 0734.sentence-similarity 句子相似性 双向绑定 
```go
package main

// https://leetcode-cn.com/problems/sentence-similarity

// ❓ 句子相似性 双向绑定
// ⚠️ great 和 fine 相似 相当 fine 和 great 相似

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

	if len(sentence1) != len(sentence2) {
		return false
	}
	for i := range sentence1 {
		if sentence1[i] != sentence2[i] && !str1MpStr2MpBool[sentence1[i]][sentence2[i]] && !str1MpStr2MpBool[sentence2[i]][sentence1[i]] {
			return false
		}
	}
	return true
}

```

# Perm ✅ 0128.longest-consecutive-sequence 最长递增 
```go
package main

// https://leetcode-cn.com/problems/longest-consecutive-sequence

func longestConsecutive(nums []int) int {
	numMpBool := map[int]bool{}
	for _, num := range nums {
		numMpBool[num] = true
	}

	var cntLongest int
	for num := range numMpBool {
		if !numMpBool[num-1] {
			cnt := 1
			num += 1
			for numMpBool[num] {
				num++
				cnt++
			}
			if cntLongest < cnt {
				cntLongest = cnt
			}
		}
	}
	return cntLongest
}

```

# Perm ✅ 0242.valid-anagram 有效的异位词 
```go
package main

// https://leetcode-cn.com/problems/valid-anagram

// ❓ 有效的异位词
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chMpCnt := [26]int{}
	for i := range s {
		chMpCnt[s[i]-'a']++
		chMpCnt[t[i]-'a']--
	}

	for _, cnt := range chMpCnt {
		if cnt != 0 {
			return false
		}
	}
	return true
}

// ❌ 通过字母code总和 求差，比如 bc 与 ad 相等， 却不是异位词

```

# Perm ✅ 0266.palindrome-permutation 有效回文 
```go
package main

// https://leetcode-cn.com/problems/palindrome-permutation

// ❓ 有效回文

func canPermutePalindrome(s string) bool {
	chMpCnt := map[byte]int{}
	for i := range s {
		chMpCnt[s[i]]++
	}

	// 奇数
	var odd = 0
	for _, cnt := range chMpCnt {
		if cnt%2 == 1 {
			odd++
		}
	}

	return odd <= 1
}

```

# Perm ✅ 0383.ransom-note  全排列 == 子排列 
```go
package main

// https://leetcode-cn.com/problems/ransom-note

// ❓ 从magazine组合成ransomNote

func canConstruct(ransomNote string, magazine string) bool {
	chMpCnt := map[byte]int{}
	for i := range magazine {
		chMpCnt[magazine[i]]++
	}

	for i := range ransomNote {
		if chMpCnt[ransomNote[i]] == 0 {
			return false
		}
		chMpCnt[ransomNote[i]]--
	}
	return true
}

```

# Perm ✅ 0409.longest-palindrome 最长回文串 
```go
package main

// https://leetcode-cn.com/problems/longest-palindrome

// ❓ 能排列为最长的回文串

func longestPalindrome(s string) int {
	chMpCnt := map[byte]int{}
	for i := range s {
		chMpCnt[s[i]-'a']++
	}

	var cntLongest = len(s)
	var odd = 0
	for _, cnt := range chMpCnt {
		if cnt%2 == 1 {
			odd++
		}
	}

	if 0 < odd {
		return cntLongest - odd + 1
	}
	return cntLongest
}

```

# Perm ✅ 0575.distribute-candies 得到最多类型的糖果 
```go
package main

// https://leetcode-cn.com/problems/distribute-candies

// ❓ 得到最多类型的糖果
func distributeCandies(candy []int) int {
	candyL := len(candy)
	typMp := map[int]struct{}{}
	for _, typ := range candy {
		typMp[typ] = struct{}{} // 类型数
	}

	typMpL := len(typMp)
	if candyL/2 < typMpL { // 类型数超过 一半，每样一颗
		return candyL / 2
	}
	return typMpL // 小于一半 最多typMpL
}

```

# Perm ✅ 0594.longest-harmonious-subsequence 最长相差 1 子序列 
```go
package main

// https://leetcode-cn.com/problems/longest-harmonious-subsequence

// ❓ 最长相差1 子序列
// 二次扫描
func findLHSTwo(nums []int) int {
	numMpCnt := map[int]int{}
	for _, num := range nums {
		numMpCnt[num]++
	}

	var cntLongest int
	for num, cnt := range numMpCnt {
		if 0 < numMpCnt[num+1] {
			cntCur := numMpCnt[num+1] + cnt
			if cntLongest < cntCur {
				cntLongest = cntCur
			}
		}
	}
	return cntLongest
}

// 一次扫描
func findLHSOne(nums []int) int {
	numMpCnt := map[int]int{}
	var cntLongest int
	for _, num := range nums {
		numMpCnt[num]++
		if 0 < numMpCnt[num+1] {
			cntCur := numMpCnt[num+1] + numMpCnt[num]
			if cntLongest < cntCur {
				cntLongest = cntCur
			}
		}

		if 0 < numMpCnt[num-1] {
			cntCur := numMpCnt[num-1] + numMpCnt[num]
			if cntLongest < cntCur {
				cntLongest = cntCur
			}
		}
	}
	return cntLongest
}

```

# Perm ✅ 0748.shortest-completing-word 最短补全词 
```go
package main

// https://leetcode-cn.com/problems/shortest-completing-word

// ❓ 最短补全词
// ⚠️ step 能覆盖 s1T2p 所有字母 stp
// ⚠️ 不区分大小写

func shortestCompletingWord(licensePlate string, words []string) string {
	chMpCnt := [26]int{}
	// 计算次数
	for i := range licensePlate {
		if 'A' <= licensePlate[i] && licensePlate[i] <= 'Z' {
			chMpCnt[licensePlate[i]-'A'] ++
		} else if 'a' <= licensePlate[i] && licensePlate[i] <= 'z' {
			chMpCnt[licensePlate[i]-'a'] ++
		}
	}

	var strShortest string
	var strShortestL = 1<<63 - 1
	var ch int
	for i := range words {
		// 计算当前次数
		chMpCntCur := [26]int{}
		for j := range words[i] {
			chMpCntCur[words[i][j]-'a']++
		}
		for ch = 0; ch < 26; ch++ {
			if chMpCntCur[ch] < chMpCnt[ch] {
				// 少于就不合法
				break
			}
		}

		if ch == 26 {
			// 合法判断
			strCurL := len(words[i])
			if strCurL < strShortestL {
				strShortest = words[i]
				strShortestL = len(words[i])
			}
		}
	}

	return strShortest
}

```

# Perm ✅ 0859.buddy-strings 交换一次后字符串相等 
```go
package main

// https://leetcode-cn.com/problems/buddy-strings

// ❓ 交换一次后字符串相等

func buddyStrings(str string, goal string) bool {
	strL := len(str)
	goalL := len(goal)
	if strL != goalL {
		return false
	}

	var slotFirst, slotSecond = -1, -1
	var chMpCnt = [26]int{}
	for i := 0; i < strL; i++ {
		ch := str[i] - 'a'
		chMpCnt[ch]++
		if str[i] != goal[i] {
			if slotFirst == -1 {
				slotFirst = i
			} else if slotSecond == -1 {
				slotSecond = i
			} else {
				return false
			}
		}
	}

	isSame := func() bool {
		for _, cnt := range chMpCnt {
			if 2 <= cnt {
				return true
			}
		}
		return false
	}

	return slotSecond != -1 && str[slotFirst] == goal[slotSecond] && str[slotSecond] == goal[slotFirst] || slotFirst == -1 && isSame()
}

```

# Perm ✅ 1086.high-five 每个学生 最高的五科 成绩的 平均分 
```go
package main

import "sort"

// https://leetcode-cn.com/problems/high-five

// ❓ 每个学生 最高的五科 成绩的 平均分

func highFive(items [][]int) [][]int {
	idMpIdx := map[int]int{}
	idxMpId := map[int]int{}
	grades := [][]int{}

	// 成绩池
	getGrades := func(id int) *[]int {
		_, ok := idMpIdx[id]
		if !ok {
			l := len(grades)
			idMpIdx[id] = l
			idxMpId[l] = id
			grades = append(grades, []int{})
		}
		return &grades[idMpIdx[id]]
	}

	// 获取id
	getId := func(index int) int {
		return idxMpId[index]
	}

	for _, item := range items {
		// or head
		grades := getGrades(item[0])
		*grades = append(*grades, item[1])
	}

	var students [][]int
	for i := range grades {
		sort.Slice(grades[i], func(l, r int) bool {
			return grades[i][r] < grades[i][l]
		})
		// 前五颗总分
		var sum int
		for _, v := range grades[i][:5] {
			sum += v
		}
		// 平均分
		avg := sum / 5
		students = append(students, []int{getId(i), avg})
	}

	// 按 ID 正序
	sort.Slice(students, func(l, r int) bool {
		return students[l][0] < students[r][0]
	})
	return students
}

```

# Perm ✅ 1160.find-words-that-can-be-formed-by-characters 拼写单词 
```go
package main

// https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters

func countCharacters(words []string, chars string) int {
	m := [26]int{}
	for i := range chars {
		m[chars[i]-'a']++
	}

	var res int
	for i := range words {
		cur := [26]int{}
		word := words[i]
		for j := range words[i] {
			ch := word[j] - 'a'
			cur[ch]++
		}
		k := 0
		for k < 26 {
			if m[k] < cur[k] {
				break
			}
			k++
		}
		if k == 26 {
			res += len(words[i])
		}
	}
	return res
}

```

# Perm ✅ 1189.maximum-number-of-balloons 气球的最大数量 
```go
package main

// https://leetcode-cn.com/problems/maximum-number-of-balloons

// ❓ 气球的最大数量

func maxNumberOfBalloons(text string) int {
	chMpCnt := map[byte]int{}
	for i := range text {
		chMpCnt[text[i]]++
	}

	chMpCntBalloon := map[byte]int{
		'a': 1,
		'b': 1,
		'n': 1,
		'o': 2,
		'l': 2,
	}

	var cntMax = 1<<63 - 1
	for ch := range chMpCntBalloon {
		cnt := chMpCnt[ch] / chMpCntBalloon[ch]
		if cnt < cntMax {
			cntMax = cnt
		}
	}
	return cntMax
}

```

# Perm ✅ 1426.counting-elements 存在后置元素的元素个数 
```go
package main

// https://leetcode-cn.com/problems/counting-elements

// 存在后置元素的元素个数
func countElements(arr []int) int {
	numMpCnt := map[int]int{}
	for _, num := range arr {
		numMpCnt[num]++
	}

	var cnt int
	for num := range numMpCnt {
		if 0 < numMpCnt[num+1] {
			cnt += numMpCnt[num]
		}
	}
	return cnt
}
```

# Perm ✅ 1460.make-two-arrays-equal-by-reversing-sub-arrays 无限翻转子数组 使两个数组相等 
```go
package main

// https://leetcode-cn.com/problems/make-two-arrays-equal-by-reversing-sub-arrays

// ❓ 无限翻转子数组 使两个数组相等

func canBeEqual(target []int, arr []int) bool {
	targetL := len(target)
	arrL := len(arr)
	if targetL != arrL {
		return false
	}
	numMpCnt := map[int]int{}
	for i := 0; i < targetL; i++ {
		numMpCnt[target[i]]++
		numMpCnt[arr[i]]--
	}

	for i := range numMpCnt {
		if numMpCnt[i] != 0 {
			return false
		}
	}
	return true
}

```

# Perm ✅ 1763.longest-nice-substring 大小写成对出现的最长子串 
```go
package main

// https://leetcode-cn.com/problems/longest-nice-substring

// ❓ 大小写成对出现的最长子串
// ⚠️ Aa Bb... 成对

func longestNiceSubstring(s string) string {
	chMp := map[byte]struct{}{}
	for i := range s {
		chMp[s[i]] = struct{}{}
	}

	for i := range s {
		ch := s[i]
		if ch <= 'Z' {
			ch += 32
		} else if 'a' <= ch {
			ch -= 32
		}
		_, ok := chMp[ch]
		if !ok {
			left := longestNiceSubstring(s[:i])
			if len(s[i+1:]) <= len(left) {
				return left
			}
			right := longestNiceSubstring(s[i+1:])
			if len(left) < len(right) {
				return right
			}
			return left
		}
	}
	return s
}

```

# Perm ✅ LCS 02. 完成一半题目最少的题型 
```go
package main

import "sort"

// https://leetcode-cn.com/problems/WqXACV/

// ❓ 完成一半题目最少题型
// ⚠️ 2*N 道题目
func halfQuestions(questions []int) int {
	// 类型计数
	typMpCnt := make([]int, 1001)
	half := len(questions) / 2
	for i := range questions {
		typMpCnt[questions[i]]++
	}

	// 类型排序
	sort.Slice(typMpCnt, func(i, j int) bool {
		return typMpCnt[j] < typMpCnt[i]
	})

	// 最少类型
	var cntTyp int
	for i := range typMpCnt {
		if half <= 0 {
			return cntTyp
		}
		half -= typMpCnt[i]
		cntTyp++
	}
	return 10086
}

```

# Pre ✅ 0930.binary-subarrays-with-sum 和为 goal 的二元子数组 
```go
package main

// https://leetcode-cn.com/problems/binary-subarrays-with-sum

// ❓ 和为 goal 的二元子数组
// ⚠️ 1 0 0 1 0 0 0 0 1
//  0 1 1 1 2 2 2 2 2 3

func numSubarraysWithSum(nums []int, goal int) int {
	var cnt, sum int
	sumMpCnt := map[int]int{}
	for _, num := range nums {
		sumMpCnt[sum]++ // 0
		sum += num
		cnt += sumMpCnt[sum-goal] // 前缀和 sum[j] - goal = sum[i]  包含自身的组合
	}
	return cnt
}

```

# Pre ✅ 1711.count-good-meals 大餐计数 
```go
package main

// https://leetcode-cn.com/problems/count-good-meals

func countPairs(deliciousness []int) (res int) {
	var preMax, sumMax int
	cnt := map[int]int{}
	for _, delicious := range deliciousness {
		if preMax < delicious {
			preMax = delicious
			sumMax = preMax << 1
		}
		for sum := 1; sum <= sumMax; sum <<= 1 {
			res += cnt[sum-delicious]
		}
		cnt[delicious] += 1
	}
	return res % (1e9 + 7)
}

```

# Queue ✅ 0362.design-hit-counter 敲击计数器 
```go
package main

// https://leetcode-cn.com/problems/design-hit-counter

type HitCounter struct {
	m []Counter
}

type Counter struct {
	cnt       int
	timestamp int
}

func Constructor() HitCounter {
	return HitCounter{m: []Counter{}}
}

func (hc *HitCounter) Hit(timestamp int) {
	n := len(hc.m) - 1
	if -1 < n && hc.m[n].timestamp == timestamp {
		hc.m[n].cnt++
	} else {
		hc.m = append(hc.m, Counter{cnt: 1, timestamp: timestamp})
	}
}

func (hc *HitCounter) GetHits(timestamp int) int {
	var res int
	var t int
	var start = timestamp - 299
	for i := range hc.m {
		if hc.m[i].timestamp < start {
			t++
		} else {
			res += hc.m[i].cnt
		}
	}
	hc.m = hc.m[t:]
	return res
}

```

# Scope ✅ 0041.first-missing-positive 缺失的首个正数 🤚 剔除 标记 筛选 
```go
package main

// https://leetcode-cn.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	l1 := len(nums)
	// 剔除
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = l1 + 1
		}
	}

	// 记录
	nums = append(nums, l1+1)
	for _, num := range nums {
		num = abs(num)

		if num <= l1 {
			nums[num] = - abs(nums[num])
		}
	}

	// 筛选
	for i := 1; i <= l1; i++ {
		if 0 < nums[i] {
			return i
		}
	}
	return l1 + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

```

# Score ✅ 0387.first-unique-character-in-a-string 字符串中的第一个唯一字符 🤚 标记索引 剔除 
```go
package main

// https://leetcode-cn.com/problems/first-unique-character-in-a-string

func firstUniqCharBit(s string) int {
	l1 := len(s)
	m := [26]int{}
	for i := range m {
		m[i] = l1
	}
	for i := 0; i < l1; i++ {
		ch := s[i] - 'a'
		if m[ch] == l1 {
			m[ch] = i
		} else {
			m[ch] = l1 + 1
		}
	}

	var res = l1
	for i := range m {
		if m[i] < res {
			res = m[i]
		}
	}
	if res < l1 {
		return res
	}
	return -1
}

func firstUniqCharQueue(s string) int {
	type pair struct {
		ch  byte
		pos int
	}
	l1 := len(s)
	m := [26]int{}
	for i := range m {
		m[i] = l1
	}

	var queue []pair
	for i := range s {
		ch := s[i] - 'a'
		if m[ch] == l1 {
			m[ch] = i
			queue = append(queue, pair{
				ch:  ch,
				pos: i,
			})
		} else {
			m[ch] = l1 + 1
			//for 0 < len(queue) && m[queue[0].ch] == l1+1 { // 队列优化map
			//	queue = queue[1:]
			//}
		}
	}
	for 0 < len(queue) && m[queue[0].ch] == l1+1 { // 队列优化map
		queue = queue[1:]
	}

	if 0 < len(queue) {
		return queue[0].pos
	}

	return -1
}

```

# Score ✅ 0442.find-all-duplicates-in-an-array 等值数组中 重复的数 🤚 偏移 标记 
```go
package main

// https://leetcode-cn.com/problems/find-all-duplicates-in-an-array

func findDuplicates(nums []int) []int {
	n := len(nums) + 1
	var res []int
	for i := range nums {
		index := (nums[i] - 1) % n // todo: n % n - 1 == -1
		if nums[index]/n == 1 {
			res = append(res, index+1)
		}
		nums[index] += n // todo: min = n + 1
	}
	return res
}

```

# Score ✅ 0448.find-all-numbers-disappeared-in-an-array 等值数组中 消失的数 🤚 偏移 标记 筛选 
```go
package main

// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array

func findDisappearedNumbers(nums []int) []int {
	var res []int
	n := len(nums)
	for _, num := range nums {
		index := (num - 1) % n
		nums[index] += n
	}

	for i, num := range nums {
		if num <= n { // 代表数字不存在
			res = append(res, i+1)
		}
	}
	return res
}

```

# Score ✅ 0604.design-compressed-string-iterator 迭代压缩字符串 
```go
package main

// https://leetcode-cn.com/problems/design-compressed-string-iterator

type StringIterator struct {
	q []pair
}

type pair struct {
	ch  byte
	cnt int
}

func Constructor(compressedString string) StringIterator {
	q := []pair{}
	l1 := len(compressedString)
	cnt := 0
	var cur byte
	var i = 0

	f := func() {
		if 0 < cnt {
			q = append(q, pair{
				ch:  cur,
				cnt: cnt,
			})
		}
	}
	for i < l1 {
		if '0' <= compressedString[i] && compressedString[i] <= '9' {
			cnt = cnt*10 + int(compressedString[i]-'0')
		} else {
			f()
			cur = compressedString[i]
			cnt = 0
		}
		i++
	}
	f()
	return StringIterator{
		q: q,
	}
}

func (sI *StringIterator) Next() (res byte) {
	if !sI.HasNext() {
		return ' '
	}
	sI.q[0].cnt--
	res = sI.q[0].ch
	if sI.q[0].cnt == 0 {
		sI.q = sI.q[1:]
	}
	return
}

func (sI *StringIterator) HasNext() bool {
	return 0 < len(sI.q)
}

```

# Score ✅ 1636.sort-array-by-increasing-frequency 按照频率将数组升序排序 🤚 权重 
```go
package main

import "sort"

// https://leetcode-cn.com/problems/sort-array-by-increasing-frequency

func frequencySort(nums []int) []int {
	m := map[int]int{}
	for i := range nums {
		m[nums[i]]++
	}

	for i := range nums {
		nums[i] = 201*m[nums[i]] - nums[i] + 100
	}

	sort.Ints(nums)

	for i := range nums {
		nums[i] = -(nums[i] % 201) + 100
	}
	return nums
}

```

# Score ✅ 1796.second-largest-digit-in-a-string 字符串中第二大的数字 
```go
package main

// https://leetcode-cn.com/problems/second-largest-digit-in-a-string

func secondHighest(s string) int {
	var first, second = -1, -1
	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			num := int(s[i] - '0')
			if first < num {
				second = first
				first = num
			} else if num < first && second < num {
				second = num
			}
		}
	}
	return second
}

```

# Score ✅ 1941.check-if-all-characters-have-equal-number-of-occurrences 检查是否所有字符出现次数相同 
```go
package main

// https://leetcode-cn.com/problems/check-if-all-characters-have-equal-number-of-occurrences

func areOccurrencesEqual(s string) bool {
	m := map[byte]int{}
	base := s[0]
	for i := range s {
		m[s[i]] ++
	}
	for k := range m {
		if m[k] != m[base] {
			return false
		}
	}
	return true
}

```

# Score ✅ 2032.two-out-of-three 至少在两个数组中出现的值 
```go
package main

// https://leetcode-cn.com/problems/two-out-of-three

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m := map[int]int{}
	var res []int

	for i := range nums1 {
		if m[nums1[i]] == 0 {
			m[nums1[i]] = 1
		}
	}

	for i := range nums2 {
		if m[nums2[i]] < 2 {
			if m[nums2[i]] == 1 {
				res = append(res, nums2[i])
			}
			m[nums2[i]] += 2
		}
	}

	for i := range nums3 {
		if m[nums3[i]] == 1 || m[nums3[i]] == 2 {
			res = append(res, nums3[i])
			m[nums3[i]] += 4
		}
	}
	return res
}

```

# Sort ✅ 0451.sort-characters-by-frequency 根据字符出现频率排序 
```go
package main

import (
	"bytes"
	"sort"
)

// https://leetcode-cn.com/problems/sort-characters-by-frequency

func frequencySortBucket(s string) string {
	var freq = map[byte]int{}
	var max = 0
	for i := range s {
		freq[s[i]]++
		if max < freq[s[i]] {
			max = freq[s[i]]
		}
	}

	var buckets = make([][]byte, max+1)
	for ch, c := range freq {
		buckets[c] = append(buckets[c], ch)
	}

	var res = make([]byte, 0, len(s))
	for i := max; 0 < i; i-- {
		for _, ch := range buckets[i] {
			res = append(res, bytes.Repeat([]byte{ch}, i)...)
		}
	}
	return string(res)
}

func frequencySortQuick(s string) string {
	var freq = map[byte]int{}
	for i := range s {
		freq[s[i]]++
	}

	type pair struct {
		cnt  int
		char byte
	}

	var pairs []pair
	for i, v := range freq {
		pairs = append(pairs, pair{
			cnt:  v,
			char: i,
		})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[j].cnt < pairs[i].cnt
	})

	var res []byte
	for _, p := range pairs {
		res = append(res, bytes.Repeat([]byte{p.char}, p.cnt)...)
	}
	return string(res)
}

```

# Sort ✅ 0953.verifying-an-alien-dictionary 单词 遵循 词典序 
```go
package main

// https://leetcode-cn.com/problems/verifying-an-alien-dictionary

func isAlienSorted(words []string, order string) bool {
	m := [26]int{}
	for i, v := range order {
		m[int(v-'a')] = i
	}

	var l, r string
	ll := 0
	for i := range words {
		r = words[i]
		rl := len(r)
		j := 0
		for j < ll && j < rl {
			p1 := m[l[j]-'a']
			p2 := m[r[j]-'a']
			if p1 < p2 {
				break
			} else if p2 < p1 {
				return false
			} else if p1 == p2 {
				j++
			}
		}
		if j == rl && rl < ll {
			return false
		}
		l = r
		ll = rl
	}
	return true
}

```

# Sort ✅ 1122.relative-sort-array 数组A 遵循 数组B 排序 
```go
package main

// https://leetcode-cn.com/problems/relative-sort-array

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := [1001]int{}
	for i := range arr1 {
		m[arr1[i]]++
	}

	res := []int{}
	for i := range arr2 {
		for 0 < m[arr2[i]] {
			res = append(res, arr2[i])
			m[arr2[i]]--
		}
	}

	for i := range m {
		for 0 < m[i] {
			res = append(res, i)
			m[i]--
		}
	}
	return res
}

```

# Sort ✅ 1331.rank-transform-of-an-array 数组序号转换 
```go
package main

import "sort"

// https://leetcode-cn.com/problems/rank-transform-of-an-array

func arrayRankTransform(arr []int) []int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)

	m := map[int]int{}
	cnt := 0
	for i := range sorted {
		if m[sorted[i]] == 0 {
			cnt++
		}
		m[sorted[i]] = cnt
	}

	for i := range arr {
		arr[i] = m[arr[i]]
	}
	return arr
}

```

# Sort ✅ 1365.how-many-numbers-are-smaller-than-the-current-number 有多少小于当前数字的数字 
```go
package main

// https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/

func smallerNumbersThanCurrent(nums []int) []int {
	var pre [101]int
	res := make([]int, len(nums))
	for i := range nums {
		pre[nums[i]]++
	}

	for i := 1; i < 101; i++ {
		pre[i] += pre[i-1]
	}
	for i, num := range nums {
		if num != 0 {
			res[i] = pre[num-1]
		}
	}
	return res
}

```

# Sort ✅ 1370.increasing-decreasing-string 上升下降字符串 
```go
package main

// https://leetcode-cn.com/problems/increasing-decreasing-string

func sortString(s string) string {
	m := [26]int{}
	for i := range s {
		m[s[i]-'a']++
	}
	var res []byte
	l1 := len(s)
	for len(res) < l1 {
		for i := 0; i <= 25; i++ {
			if 0 < m[i] {
				res = append(res, byte(i+'a'))
				m[i]--
			}
		}

		for i := 25; 0 <= i; i-- {
			if 0 < m[i] {
				res = append(res, byte(i+'a'))
				m[i]--
			}
		}
	}
	return string(res)
}

```

# Sort ✅ 1640.check-array-formation-through-concatenation 能否连接形成数组 
```go
package main

// https://leetcode-cn.com/problems/check-array-formation-through-concatenation

func canFormArray(arr []int, pieces [][]int) bool {
	m := map[int]int{}
	for i := range arr {
		m[arr[i]] = i + 1
	}

	for _, piece := range pieces {
		top := len(piece) - 1
		for j := range piece {
			if m[piece[j]] == 0 || j < top && m[piece[j]]+1 != m[piece[j+1]] {
				return false
			}
		}
	}
	return true
}

```

# Sort ❌ 1833.maximum-ice-cream-bars 雪糕的最大数量 
```go
package main

// https://leetcode-cn.com/problems/maximum-ice-cream-bars

func maxIceCream(costs []int, coins int) int {
	var res int
	var max int
	for _, c := range costs { // 查找最大值
		if max < c {
			max = c
		}
	}

	var freq = make([]int, max+1)
	for _, c := range costs { // 计数 出现的次数
		freq[c]++
	}

	for i := 1; i <= max && i <= coins; i++ {
		var c = min(freq[i], coins/i) // 最多能买到的雪糕
		res += c
		coins -= i * c
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```

# Vector ✅ 0049.group-anagrams 字母异位词分组 
```go
package main

// https://leetcode-cn.com/problems/group-anagrams

func groupAnagrams(strs []string) [][]string {
	m := map[[26]int][]string{}
	for i := range strs {
		var vector [26]int
		for j := range strs[i] {
			vector[strs[i][j]-'a']++
		}
		m[vector] = append(m[vector], strs[i])
	}
	var res [][]string
	for i := range m {
		res = append(res, m[i])
	}
	return res
}

```

# Vector ✅ 0249.group-shifted-strings 移位字符串分组 
```go
package main

// https://leetcode-cn.com/problems/group-shifted-strings

func groupStrings(strings []string) [][]string {
	m := map[string]int{}
	res := [][]string{}
	getId := func(s string) int {
		_, ok := m[s]
		if !ok {
			m[s] = len(res)
			res = append(res, []string{})
		}
		return m[s]
	}

	for _, char := range strings {
		j := 1
		l := len(char)
		b := make([]byte, l)
		for j < l {
			b = append(b, (char[j]-char[0]+26)%26) // 以防变负
			j++
		}
		id := getId(string(b))
		res[id] = append(res[id], char)
	}
	return res
}

```

# Where ✅ 0217.contains-duplicate 出现重复 
```go
package main

// https://leetcode-cn.com/problems/contains-duplicate

// ❓ 出现重复

func containsDuplicate(nums []int) bool {
	var valMpCnt = map[int]int{}
	for _, val := range nums {
		if 0 < valMpCnt[val] {
			return true
		}
		valMpCnt[val] += 1
	}
	return false
}

```

# Where ✅ 0274.h-index H 指数 
```go
package main

// https://leetcode-cn.com/problems/h-index

// ❓ h指数 为 篇数cnt 等于 至少引用次数quote

func hIndex(citations []int) (h int) {
	// 后缀和：至少引用次数val == 篇幅cnt
	quoteMax := len(citations)
	quoteMpCnt := make([]int, quoteMax+1)
	for i := range citations {
		if quoteMax <= citations[i] {
			quoteMpCnt[quoteMax] ++
		} else {
			quoteMpCnt[citations[i]]++
		}
	}
	var cnt int
	// i = 0 时，为一个都没被引用
	for quote := quoteMax; 0 < quote; quote-- {
		cnt += quoteMpCnt[quote]
		if quote <= cnt {
			return quote
		}
	}
	return -1
}

```

# Where ✅ 0359.logger-rate-limiter 相同消息截流 
```go
package main

// https://leetcode-cn.com/problems/logger-rate-limiter

// ❓相同消息截流

type Logger struct {
	msgMpTime map[string]int
}

func Constructor() Logger {
	return Logger{msgMpTime: make(map[string]int)}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	// 还在截流期
	if timestamp < l.msgMpTime[message] {
		return false
	}
	// 10秒
	l.msgMpTime[message] = timestamp + 10
	return true
}

```

# Where ✅ 0884.uncommon-words-from-two-sentences 两句话中的不常见单词 
```go
package main

import "strings"

// https://leetcode-cn.com/problems/uncommon-words-from-two-sentences

// ❓ 在两个数组只出现一次的单词

func uncommonFromSentences(s1 string, s2 string) []string {
	// 计数
	a1 := strings.Split(s1, " ")
	a2 := strings.Split(s2, " ")

	strMpCnt := map[string]int{}

	for _, str := range a1 {
		strMpCnt[str]++
	}

	for _, str := range a2 {
		strMpCnt[str]++
	}

	// 是否一次
	var res []string
	for str := range strMpCnt {
		if strMpCnt[str] == 1 {
			res = append(res, str)
		}
	}
	return res
}

```

# Where ✅ 1119.remove-vowels-from-a-string 删去字符串中的元音 
```go
package main

// https://leetcode-cn.com/problems/remove-vowels-from-a-string

// ❓ 移除元音字符

func removeVowels(s string) string {
	l := len(s)
	var res = make([]byte, 0, l)
	for i := range s {
		// 元音跳过
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			l--
		} else {
			res = append(res, s[i])
		}
	}
	return string(res[:l])
}

```

# Where ✅ 1133.largest-unique-number 最大的唯一数 
```go
package main

// https://leetcode-cn.com/problems/largest-unique-number

// ❓ 最大的唯一值

func largestUniqueNumber(nums []int) int {
	// 计数
	numMpCnt := [1001]int{}
	for _, num := range nums {
		numMpCnt[num]++
	}

	// 倒序
	for num := 1000; 0 <= num; num-- {
		if numMpCnt[num] == 1 {
			return num
		}
	}
	return -1
}

```

# Where ✅ 1346.check-if-n-and-its-double-exist 是否存在两倍数 
```go
package main

// https://leetcode-cn.com/problems/check-if-n-and-its-double-exist

// ❓是否存在两倍数

func checkIfExist(arr []int) bool {
	numMpCnt := map[int]int{}
	for _, num := range arr {
		numMpCnt[num]++
	}

	for num := range numMpCnt {
		if 0 < numMpCnt[num*2] {
			if num != 0 || 2 <= numMpCnt[num] {
				return true
			}
		}
	}
	return false
}

```

# Where ✅ 1394.find-lucky-integer-in-an-array 找出数组中的幸运数 
```go
package main

// https://leetcode-cn.com/problems/find-lucky-integer-in-an-array

// ❓ 出现次数cnt等于值val的最大值

func findLucky(arr []int) int {
	// 计数
	numMpCnt := [501]int{}
	for _, num := range arr {
		numMpCnt[num]++
	}

	// 倒序
	for num := 500; 0 < num; num-- {
		if numMpCnt[num] == num {
			return num
		}
	}
	return -1
}

```

# Where ✅ 1748.sum-of-unique-elements 唯一元素的和 
```go
package main

// https://leetcode-cn.com/problems/sum-of-unique-elements

// ❓ 所有唯一元素的和

func sumOfUnique(nums []int) int {
	var numMpCnt = map[int]int{}
	var res int
	for _, num := range nums {
		if numMpCnt[num] == 0 {
			res += num
			numMpCnt[num] = 1
		} else if numMpCnt[num] == 1 {
			res -= num
			numMpCnt[num] = 2
		}
	}
	return res
}

```

# Where ✅ 1935.maximum-number-of-words-you-can-type 可以输入的最大单词数 
```go
package main

// https://leetcode-cn.com/problems/maximum-number-of-words-you-can-type

// ❓ 可以输入的最大单词数
// ⚠️ hello word , w

func canBeTypedWords(text string, brokenLetters string) int {
	// 损坏键 计数
	var chMpDis = map[byte]bool{}
	for i := range brokenLetters {
		chMpDis[brokenLetters[i]] = true
	}
	var res int

	var idx int
	var l = len(text)
	for idx < l {
		// 损坏键
		for idx < l && text[idx] != ' ' && !chMpDis[text[idx]] {
			idx++
		}

		if idx == l || text[idx] == ' ' {
			res++
		} else {
			for idx < l && text[idx] != ' ' {
				idx++
			}
		}
		idx++
	}
	return res
}

```

# Where ✅ 2062.count-vowel-substrings-of-a-string 连续元音子串数 
```go
package main

// https://leetcode-cn.com/problems/count-vowel-substrings-of-a-string

// ❓ 连续元音子串数

func countVowelSubstrings(word string) int {
	var res int
	var byteMpCnt map[byte]int

	wordL := len(word)
	for i := 0; i < wordL; i++ {
		byteMpCnt = map[byte]int{}
		for j := i; j < wordL && (word[j] == 'a' || word[j] == 'e' || word[j] == 'i' || word[j] == 'o' || word[j] == 'u'); j++ {
			byteMpCnt[word[j]]++
			if len(byteMpCnt) == 5 {
				res++
			}
		}
	}
	return res
}

func countVowelSubstringsBit(word string) int {
	var res int
	wordL := len(word)
	var state, success int32
	// 计算合法状态
	success |= 1 << int('a'-'a')
	success |= 1 << int('e'-'a')
	success |= 1 << int('i'-'a')
	success |= 1 << int('o'-'a')
	success |= 1 << int('u'-'a')

	for i := 0; i < wordL; i++ {
		// 每次清空状态
		state = 0
		for j := i; j < wordL; j++ {
			bit := word[j] - 'a'
			if success&1<<bit == 0 {
				// 非法状态
				break
			}
			// 合并状态
			state |= 1 << bit

			// 是否合法
			if state == success {
				res++
			}
		}
	}
	return res
}

```

# Window Divide ❌ 0395.longest-substring-with-at-least-k-repeating-characters 至少包含 K 个重复字符的最长子串 
```go
package main

import "strings"

// https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters

func longestSubstring(s string, k int) int {
	var res int
	for typ := 1; typ <= 26; typ++ {
		cnt := [26]int{}
		typCnt := 0
		typValid := k
		l := 0
		for r := range s {
			ch := s[r] - 'a'
			if cnt[ch] == 0 {
				typCnt++
				typValid--
			}
			cnt[ch]++
			if cnt[ch] == k {
				typValid++
			}

			for typ < typCnt {
				ch := s[l] - 'a'
				if cnt[ch] == k {
					typValid--
				}
				cnt[ch]--
				if cnt[ch] == 0 {
					typCnt--
					typValid++
				}
				l++
			}

			if typValid == k {
				cur := r - l + 1
				if res < cur {
					res = cur
				}
			}
		}
	}
	return res
}

func longestSubstringDivide(s string, k int) int {
	var res int
	if s == "" {
		return res
	}

	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	var split byte
	for i, c := range cnt {
		if 0 < c && c < k {
			split = 'a' + byte(i)
			break
		}
	}
	if split == 0 {
		return len(s)
	}

	for _, subStr := range strings.Split(s, string(split)) {
		cur := longestSubstringDivide(subStr, k)
		if res < cur {
			res = cur
		}
	}
	return res
}

```

# Window Vector ✅ 0438.find-all-anagrams-in-a-string 找到字符串中所有字母异位词 
```go
package main

// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string

func findAnagrams(s string, p string) []int {
	l1 := len(s)
	l2 := len(p)
	l := 0
	r := l2 - 1
	var res []int
	if l1 < l2 {
		return res
	}
	m1 := [26]int{}
	for i := range p {
		m1[p[i]-'a']++
	}

	m2 := [26]int{}
	for l < r {
		m2[s[l]-'a']++
		l++
	}
	l = 0
	for r < l1 {
		m2[s[r]-'a']++
		if m1 == m2 {
			res = append(res, l)
		}
		m2[s[l]-'a']--
		l++
		r++
	}

	return res
}

```

# Window ✅ 0003.longest-substring-without-repeating-characters 无重复的最长子串 
```go
package main

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	var res, cur, l int
	var m = map[byte]int{}
	for i := range s {
		if l < m[s[i]] {
			l = m[s[i]] // 缩
		}
		cur = i - l + 1
		if res < cur {
			res = cur
		}
		m[s[i]] = i + 1
		// 出现位置 + 1
		// A.否则 重复[a]bca 0 < 0 不成立
		// B.否则 cur = i - l + 1 不成立
		// C.也可 l = m[s[i]] + 1 ，但 A 不成立
	}
	return res
}

```

# Window ✅ 0159.longest-substring-with-at-most-two-distinct-characters 至多包含两个不同字符的最长子串 
```go
package main

// https://leetcode-cn.com/problems/longest-substring-with-at-most-two-distinct-characters

func lengthOfLongestSubstringTwoDistinct(s string) int {
	m := [128]int{}
	var n = len(s)
	var max, cnt, l, r int
	var f = func() {
		cur := r - l
		if max < cur {
			max = cur
		}
	}
	for r < n {
		if m[s[r]] == 0 {
			f()
			cnt++
			for 2 < cnt {
				m[s[l]]--
				if m[s[l]] == 0 {
					cnt--
				}
				l++
			}
		}
		m[s[r]]++
		r++
	}
	f()

	return max
}

```

# Window ✅ 0219.contains-duplicate-ii 两个重复元素距离小于等于K 
```go
package main

// https://leetcode-cn.com/problems/contains-duplicate-ii

func containsNearbyDuplicate(nums []int, k int) bool {
	var m = map[int]int{}
	for r, num := range nums {
		l, ok := m[num]
		if ok && r-l <= k {
			return true
		}
		m[num] = r
	}
	return false
}

```

# Window ✅ 0246.strobogrammatic-number 中心对称数 
```go
package main

// https://leetcode-cn.com/problems/strobogrammatic-number

func isStrobogrammatic(num string) bool {
	m := map[byte]byte{'6': '9', '9': '6', '8': '8', '0': '0', '1': '1'}

	l, r := 0, len(num)-1
	for l <= r {
		if m[num[l]] != num[r] || m[num[l]] == 0 {
			return false
		}
		l++
		r--
	}
	return true
}

```

# Window ✅ 0697.degree-of-an-array 度相同的最短连续子数组 
```go
package main

// https://leetcode-cn.com/problems/degree-of-an-array

type e struct {
	cnt int
	l   int
	r   int
}

// 最大的度必定横跨首尾
func findShortestSubArray(nums []int) int {
	m := map[int]*e{}
	for i := range nums {
		_, ok := m[nums[i]]
		if ok {
			m[nums[i]].r = i
			m[nums[i]].cnt++
		} else {
			m[nums[i]] = &e{
				cnt: 1,
				l:   i,
				r:   i,
			}
		}
	}

	var cnt int
	var res int
	for _, e := range m {
		if cnt < e.cnt {
			cnt = e.cnt
			res = e.r - e.l + 1
		} else if cnt == e.cnt {
			cur := e.r - e.l + 1
			if cur < res {
				res = cur
			}
		}
	}
	return res
}

```

# Window ✅ 1624.largest-substring-between-two-equal-characters 两个相同字符之间的最长子字符串 
```go
package main

// https://leetcode-cn.com/problems/largest-substring-between-two-equal-characters

func maxLengthBetweenEqualCharacters(s string) int {
	m := map[byte]int{}
	var max = -1

	for r := range s {
		l, ok := m[s[r]]
		if !ok {
			m[s[r]] = r + 1
		} else {
			cur := r - l
			if max < cur {
				max = cur
			}
		}
	}
	return max
}

```

# Window ❌ 0325.maximum-size-subarray-sum-equals-k 和等于 k 的最长子数组长度 
```go
package main

// https://leetcode-cn.com/problems/maximum-size-subarray-sum-equals-k

func maxSubArrayLen(nums []int, k int) int {
	m := map[int]int{0: -1}
	var pre, max, cur int
	for r := range nums {
		pre += nums[r]
		l, ok := m[pre-k]
		if ok {
			cur = r - l
			if max < cur {
				max = cur
			}
		}

		_, ok = m[pre]
		if !ok {
			m[pre] = r
		}
	}
	return max
}

```

# Window ❌ 0340.longest-substring-with-at-most-k-distinct-characters 至多包含 K 个不同字符的最长子串 
```go
package main

// https://leetcode-cn.com/problems/longest-substring-with-at-most-k-distinct-characters

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	m := map[byte]int{}
	var cur, max, cnt, l int
	for r := range s {
		if m[s[r]] == 0 {
			cnt++
		}
		m[s[r]]++
		for k < cnt {
			m[s[l]]--
			if m[s[l]] == 0 {
				cnt--
			}
			l++
		}
		cur = r - l + 1
		if max < cur {
			max = cur
		}
	}
	return max
}

```

# Window ❌ 0961.n-repeated-element-in-size-2n-array 2N 中重复 N 次的元素 
```go
package main

// https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array

func repeatedNTimes(nums []int) int {
	l1 := len(nums)
	for j := 1; j <= 2; j++ {
		for i := 0; i < l1-j; i++ {
			if nums[i] == nums[i+j] {
				return nums[i]
			}
		}
	}

	return -1
}

```

