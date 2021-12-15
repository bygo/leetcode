# Distance ✅ 0888.fair-candy-swap 糖果棒交换后数量相等 
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

# Distance ✅ 1165.single-row-keyboard 机械手最少移动次数 
```go
package main

// https://leetcode-cn.com/problems/single-row-keyboard

// ❓ 输入一个单词的移动长度总和
// ⚠️ 移动长度 = abs(i - j)

func calculateTime(keyboard string, word string) int {
	// 索引统计
	var valMpIdx = [26]int{}
	for i := range keyboard {
		valMpIdx[keyboard[i]-'a'] = i
	}

	// 计算移动次数
	// 0
	var res = valMpIdx[word[0]-'a']

	// 1～l
	var l = len(word)
	var i = 1
	for i < l {
		res += abs(valMpIdx[word[i-1]-'a'], valMpIdx[word[i]-'a'])
		i++
	}

	return res
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

```

# Distance✅ 0389.find-the-difference 找不同 
```go
package main

// https://leetcode-cn.com/problems/find-the-difference

// ❓ 找出 x
// ⚠️ s = t + 随机byte(x)

func findTheDifferenceCnt(s string, t string) byte {
	// 统计
	valMpIdx := [26]int{}
	for i := range s {
		ch := s[i] - 'a'
		valMpIdx[ch]++
	}

	// map抵消
	for i := range t {
		ch := t[i] - 'a'
		if valMpIdx[ch] == 0 {
			return ch + 'a'
		}
		valMpIdx[ch]--
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

# Divide ✅ 1763.longest-nice-substring 大小写成对出现的最长子串 
```go
package main

// https://leetcode-cn.com/problems/longest-nice-substring

// ❓ 大小写成对出现的最长子串
// ⚠️ Aa Bb... 成对

func longestNiceSubstring(s string) string {
	byteMp := map[byte]struct{}{}
	for i := range s {
		byteMp[s[i]] = struct{}{}
	}

	for i := range s {
		// 找对
		ch := s[i]
		if ch <= 'Z' {
			ch += 32
		} else {
			ch -= 32
		}
		_, ok := byteMp[ch]

		if !ok {
			// 分区
			left := longestNiceSubstring(s[:i])
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

# Filter ✅ 0217.contains-duplicate 出现重复 
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

# Filter ✅ 0359.logger-rate-limiter 相同消息截流 
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

# Filter ✅ 0884.uncommon-words-from-two-sentences 两句话中的不常见单词 
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

	for _, s := range a1 {
		strMpCnt[s]++
	}

	for _, s := range a2 {
		strMpCnt[s]++
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

# Filter ✅ 1119.remove-vowels-from-a-string 删去字符串中的元音 
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

# Filter ✅ 1133.largest-unique-number 最大的唯一数 
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

# Filter ✅ 1394.find-lucky-integer-in-an-array 找出数组中的幸运数 
```go
package main

// https://leetcode-cn.com/problems/find-lucky-integer-in-an-array

// ❓ 出现次数cnt等于值val的最大值

func findLucky(arr []int) int {
	// 计数
	numMpCnt := [501]int{}
	for i := range arr {
		numMpCnt[arr[i]]++
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

# Filter ✅ 1748.sum-of-unique-elements 唯一元素的和 
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

# Filter ✅ 1935.maximum-number-of-words-you-can-type 可以输入的最大单词数 
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

# Group ✅ 0274.h-index H 指数 
```go
package main

// https://leetcode-cn.com/problems/h-index

// ❓ h指数 为 篇数cnt 等于 至少引用次数quote

func hIndex(citations []int) (h int) {
	// 后缀和：至少引用次数val == 篇幅cnt
	max := len(citations)
	quoteMpCnt := make([]int, max+1)
	for i := range citations {
		if max <= citations[i] {
			quoteMpCnt[max] ++
		} else {
			quoteMpCnt[citations[i]]++
		}
	}
	var cnt int
	// i = 0 时，为一个都没被引用
	for quote := max; 0 < quote; quote-- {
		cnt += quoteMpCnt[quote]
		if quote <= cnt {
			return quote
		}
	}
	return -1
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
	valMpComMpCnt map[string]map[string]int
}

func Constructor(dictionary []string) ValidWordAbbr {
	valMpComMpCnt := map[string]map[string]int{}
	for i := range dictionary {
		s := Compress(dictionary[i])
		if valMpComMpCnt[s] == nil {
			valMpComMpCnt[s] = map[string]int{}
		}
		valMpComMpCnt[s][dictionary[i]] ++
	}
	return ValidWordAbbr{
		valMpComMpCnt: valMpComMpCnt,
	}
}

func (vwa *ValidWordAbbr) IsUnique(word string) bool {
	s := Compress(word)
	return vwa.valMpComMpCnt[s] == nil || len(vwa.valMpComMpCnt[s]) == 1 && 0 < vwa.valMpComMpCnt[s][word]
}

// hello => h3o
func Compress(word string) string {
	l := len(word)
	return string(word[0]) + strconv.Itoa(l-2) + string(word[l-1])
}

```

# Group ✅ 0348.design-tic-tac-toe 设计井字棋 
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

# Group ✅ 0500.keyboard-row 键盘同一行的单词数 
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
	var res []string
	for _, word := range words {
		idx := 0
		l := len(word)
		for idx < l && chMpState[word[idx]] == chMpState[word[0]] {
			idx++
		}
		if idx == l {
			res = append(res, word)
		}
	}
	return res
}

```

# Group ✅ 0811.subdomain-visit-count 子域名访问次数 
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
			for idx < strL && str[idx] != '.' {
				idx++
			}
			idx++
		}
	}
	var res []string
	for host, cnt := range hostMpCnt {
		res = append(res, strconv.Itoa(cnt)+" "+host)
	}
	return res
}

```

# Group ✅ 0914.x-of-a-kind-in-a-deck-of-cards 卡牌分组 
```go
package main

// https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards

// ❓存在 int(x) 把数组分为:
// ⚠️ 每组cnt为 int(x)
// ⚠️ 每组val相同

func hasGroupsSizeX(deck []int) bool {
	// 统计元素
	numMpCnt := map[int]int{}
	for _, num := range deck {
		numMpCnt[num]++
	}

	// 最大公约数
	x := -1
	for num := range numMpCnt {
		if x == -1 {
			x = numMpCnt[num]
		} else {
			x = gcd(x, numMpCnt[num])
		}
	}
	return 2 <= x
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

// ❓ 子域名访问次数
// b.y+ig@gmail.com
// .会被忽略
// +忽略 +~@内容
func numUniqueEmails(emails []string) int {
	mailMp := map[string]struct{}{}
	for _, email := range emails {
		l1 := len(email)
		var idx int
		var cur []byte
		// 计算 b.y+ig
		for idx < l1 && email[idx] != '@' {
			if email[idx] == '+' {
				// 移动到后 @gmail.com
				for idx < l1 && email[idx] != '@' {
					idx++
				}
			} else {
				if email[idx] != '.' {
					cur = append(cur, email[idx])
				}
				idx++
			}
		}

		// 计算 @gmail.com
		cur = append(cur, email[idx:]...)
		mailMp[string(cur)] = struct{}{}
	}
	return len(mailMp)
}

```

# Group ✅ 1128.number-of-equivalent-domino-pairs 组合相同的数量 
```go
package main

// https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs

// ❓ 组合相同的数量
// ⚠️ 一个组只有2个数字

func numEquivDominoPairs(d [][]int) int {
	combMpCnt := map[int]int{}
	var res int
	for i := range d {
		if d[i][0] < d[i][1] {
			d[i][0], d[i][1] = d[i][1], d[i][0]
		}
		comb := d[i][0]*10 + d[i][1]
		res += combMpCnt[comb]
		combMpCnt[comb]++
	}
	return res
}

```

# Group ✅ 1275.find-winner-on-a-tic-tac-toe-game 找出井字棋的获胜者 
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
	l1 := len(moves)

	perMpCnt := [8]int{}

	for i := l1 - 1; 0 <= i; i -= 2 {
		row := moves[i][0]
		col := moves[i][1] + 3
		perMpCnt[row]++
		perMpCnt[col]++
		if moves[i][0] == moves[i][1] {
			perMpCnt[6]++
		}
		if moves[i][0]+moves[i][1] == 2 {
			perMpCnt[7]++
		}
	}

	for i := range perMpCnt {
		if perMpCnt[i] != 3 {
			continue
		}
		if l1%2 == 1 {
			return A
		}
		return B
	}
	if l1 < 9 {
		return Pending
	}

	return Draw
}

```

# Group ✅ 1399.count-largest-group 统计数位sum相同的组，cnt最大组的数量 
```go
package main

// https://leetcode-cn.com/problems/count-largest-group

// ❓ 统计数位sum相同的组，cnt最大组的数量

func countLargestGroup(n int) int {
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
	var cnt, res int
	for sum := range sumMpCnt {
		if cnt < sumMpCnt[sum] {
			cnt = sumMpCnt[sum]
			res = 1
		} else if sumMpCnt[sum] == cnt {
			res++
		}
	}
	return res
}

```

# Group ✅ 1512.number-of-good-pairs 好数对的数目：i < j & a[i] == a[j] 
```go
package main

// https://leetcode-cn.com/problems/number-of-good-pairs

// ❓ 好数对的数目
// ⚠️ i < j && nums[i] == nums[j]

func numIdenticalPairs(nums []int) int {
	var numMpCnt = map[int]int{}
	var res int
	for _, num := range nums {
		if 0 < numMpCnt[num] {
			res += numMpCnt[num]
		}
		numMpCnt[num]++
	}
	return res
}

```

# Group ✅ 1742.maximum-number-of-balls-in-a-box 盒子中小球的最大数量：数位和 
```go
package main

// https://leetcode-cn.com/problems/maximum-number-of-balls-in-a-box

// ❓ 盒子中小球的最大数量

func countBalls(lowLimit int, highLimit int) int {
	var sum int
	j := lowLimit
	for 0 < j {
		sum += j % 10
		j /= 10
	}
	sumMpCnt := [46]int{} // 最大9999 = 46
	var max int
	for i := lowLimit; i <= highLimit; i++ {
		sumMpCnt[sum]++
		k := i
		for k%10 == 9 { // 每次进1 减9
			sum -= 9
			k /= 10
		}
		sum++
	}

	for sum := range sumMpCnt {
		if max < sumMpCnt[sum] {
			max = sumMpCnt[sum]
		}
	}

	return max
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
	var tmp []byte
	fn := func() {
		tmpL := len(tmp)
		if 0 < tmpL {
			j := 0
			for tmp[j] == '0' && j < tmpL-1 {
				j++
			}
			strMp[string(tmp[j:])] = struct{}{}
			tmp = []byte{}
		}
	}

	for i := range word {
		if '0' <= word[i] && word[i] <= '9' {
			tmp = append(tmp, word[i])
		} else {
			fn()
		}
	}
	fn()
	return len(strMp)
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

	var res []int
	for _, num := range nums2 {
		_, ok := numMp[num]

		if ok {
			res = append(res, num)
			delete(numMp, num)
		}
	}
	return res
}

```

# Inter ✅ 0350.intersection-of-two-arrays-ii 两数组 val 交集 
```go
package main

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii

// ❓ 两数组 val 交集

func intersect(nums1 []int, nums2 []int) []int {
	numMpCnt := map[int]int{}
	for i := range nums1 {
		numMpCnt[nums1[i]]++
	}

	var res []int
	for _, num := range nums2 {
		if 0 < numMpCnt[num] {
			res = append(res, num)
			numMpCnt[num]--
		}
	}
	return res
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

	var res int
	for i := range s {
		res += chMpCnt[s[i]]
	}
	return res
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

	var res string
	var max int
	var tmp []byte
	fn := func() {
		str := string(tmp)
		if 0 < len(str) && 0 == banMpCnt[str] {
			strMpCnt[str]++
			if max < strMpCnt[str] {
				res = str
				max = strMpCnt[str]
			}
		}
		tmp = []byte{}
	}
	for i := range paragraph {
		if 'A' <= paragraph[i] && paragraph[i] <= 'Z' {
			tmp = append(tmp, paragraph[i]+32)
		} else if 'a' <= paragraph[i] && paragraph[i] <= 'z' {
			tmp = append(tmp, paragraph[i])
		} else {
			fn()
		}
	}
	fn()
	return res
}

```

# Inter ✅ 1002.find-common-characters N个数组 val 交集 
```go
package main

// https://leetcode-cn.com/problems/find-common-characters

func commonChars(words []string) {
	var res []string
	minFreq := [26]int{}
	for i := range minFreq {
		minFreq[i] = 1<<63 - 1
	}
	for _, word := range words {
		freq := [26]int{}
		for _, b := range word {
			freq[b-'a']++
		}
		for i, f := range freq[:] {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}

	// 字典序
	for i := 0; i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			res = append(res, string(rune(i+'a')))
		}
	}
	return
}

```

# Link 0356.line-reflection 直线镜像 
```go
package main

// https://leetcode-cn.com/problems/line-reflection

func isReflected(points [][]int) bool {
	m := map[int]map[int]bool{}
	max, min := -1<<63, 1<<63-1

	for _, point := range points {
		if max < point[0] {
			max = point[0]
		}
		if point[0] < min {
			min = point[0]
		}
		if m[point[0]] == nil {
			m[point[0]] = map[int]bool{}
		}
		m[point[0]][point[1]] = true
	}
	for _, point := range points {
		x := max + min - point[0]
		// x 镜像不存在 或者 不平行
		if m[x] == nil || !m[x][point[1]] {
			return false
		}
	}
	return true
}

```

# Link Trie ✅ 0720.longest-word-in-dictionary 词典中连续递接 的最长单词 
```go
package main

import "sort"

// https://leetcode-cn.com/problems/longest-word-in-dictionary

func longestWord(words []string) string {
	sort.Strings(words)
	m := map[string]struct{}{}

	var res string
	for i := range words {
		l1 := len(words[i])
		if len(words[i]) == 1 {
			m[words[i]] = struct{}{}
			if len(res) < l1 {
				res = words[i]
			}
		} else {
			_, ok := m[words[i][:l1-1]]
			if ok {
				m[words[i]] = struct{}{}
				if len(res) < l1 {
					res = words[i]
				}
			}
		}

	}
	return res
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

	var res string
	var queue = []*trie{root}
	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		res = queue[0].s
		for i := 0; i < cnt; i++ {
			for j := 0; j < 26; j++ {
				if queue[i].child[j] != nil && queue[i].child[j].s != "" {
					queue = append(queue, queue[i].child[j])
				}
			}
		}

		queue = queue[cnt:]
	}
	return res
}

```

# Link ✅ 0001.two-sum 两数之和 
```go
package main

// https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	var m = map[int]int{}
	for rawK, rawV := range nums {
		hashK, ok := m[target-rawV]
		if ok {
			return []int{hashK, rawK}
		}
		m[rawV] = rawK
	}
	return nil
}

```

# Link ✅ 0205.isomorphic-strings 同构字符串 
```go
package main

// https://leetcode-cn.com/problems/isomorphic-strings

func isIsomorphic(s string, t string) bool {
	m1 := map[byte]byte{}
	m2 := map[byte]byte{}
	l := len(s)
	for i := 0; i < l; i++ {
		ch1 := s[i]
		ch2 := t[i]
		if m1[ch1] != 0 && m1[ch1] != ch2 || m2[ch1] != 0 && m2[ch2] != ch1 {
			return false
		}
		m1[ch1] = ch2
		m2[ch2] = ch1
	}

	return true
}

```

# Link ✅ 0244.shortest-word-distance-ii 两数组相同值最小索引 II 🤚 索引队列 
```go
package main

// https://leetcode-cn.com/problems/shortest-word-distance-ii

type WordDistance struct {
	indexes map[string][]int
}

func Constructor(wordsDict []string) WordDistance {
	indexes := map[string][]int{}
	for i := range wordsDict {
		indexes[wordsDict[i]] = append(indexes[wordsDict[i]], i)
	}
	return WordDistance{
		indexes: indexes,
	}
}

func (wd *WordDistance) Shortest(word1 string, word2 string) int {
	arr1 := wd.indexes[word1]
	arr2 := wd.indexes[word2]
	i, j := len(arr1)-1, len(arr2)-1
	var min = 1<<63 - 1
	for -1 < i && -1 < j {
		var cur int
		if arr1[i] < arr2[j] {
			cur = arr2[j] - arr1[i]
			j--
		} else {
			cur = arr1[i] - arr2[j]
			i--
		}
		if cur < min {
			min = cur
		}
	}
	return min
}

```

# Link ✅ 0290.word-pattern 单词规律 
```go
package main

import "strings"

// https://leetcode-cn.com/problems/word-pattern

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	var patternSet = map[string]byte{}
	var wordsSet = map[byte]string{}
	for i := range pattern {
		p := patternSet[words[i]]
		w := wordsSet[pattern[i]]
		if p != 0 && p != pattern[i] || w != "" && w != words[i] {
			return false
		}
		patternSet[words[i]] = pattern[i]
		wordsSet[pattern[i]] = words[i]

	}

	return true
}

```

# Link ✅ 0599.minimum-index-sum-of-two-lists 两数组相同值的最小索引总和 
```go
package main

// https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists

func findRestaurant(list1 []string, list2 []string) []string {
	m := map[string]int{}
	for i := range list1 {
		m[list1[i]] = i + 1
	}

	var res []string
	var min = 1<<63 - 1
	for i := range list2 {
		if 0 < m[list2[i]] {
			cur := m[list2[i]] + i
			if cur < min {
				res = []string{list2[i]}
				min = cur
			} else if cur == min {
				res = append(res, list2[i])
			}
		}
	}
	return res
}

```

# Link ✅ 0734.sentence-similarity 句子相似性 
```go
package main

// https://leetcode-cn.com/problems/sentence-similarity

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	m := map[string]map[string]bool{}
	for _, pair := range similarPairs {
		if m[pair[0]] == nil {
			m[pair[0]] = map[string]bool{}
		}
		if m[pair[1]] == nil {
			m[pair[1]] = map[string]bool{}
		}
		m[pair[0]][pair[1]] = true
		m[pair[1]][pair[0]] = true
	}

	if len(sentence1) != len(sentence2) {
		return false
	}
	for i := range sentence1 {
		if sentence1[i] != sentence2[i] && !m[sentence1[i]][sentence2[i]] && !m[sentence2[i]][sentence1[i]] {
			return false
		}
	}
	return true
}

```

# Link ✅ 0760.find-anagram-mappings 找A在B的位置 
```go
package main

// https://leetcode-cn.com/problems/find-anagram-mappings

func anagramMappings(A []int, B []int) []int {
	var m = map[int]int{}
	for k, v := range B {
		m[v] = k
	}

	var res []int
	for _, v := range A {
		res = append(res, m[v])
	}
	return res
}

```

# Link ✅ 1346.check-if-n-and-its-double-exist 是否存在两倍数 
```go
package main

// https://leetcode-cn.com/problems/check-if-n-and-its-double-exist

func checkIfExist(arr []int) bool {
	m := map[int]int{}
	for i := range arr {
		m[arr[i]]++
	}

	for i := range m {
		if 0 < m[i*2] {
			if i != 0 || 2 <= m[i] {
				return true
			}
		}
	}
	return false
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
	l := len(changed)
	if l%2 == 1 {
		return nil
	}
	m := [100001]int{}
	var cnt int
	for i := range changed {
		m[changed[i]]++
		cnt++
	}

	var res []int
	for i := 0; i <= 50000; i++ {
		// 必须从小开始抵消
		for 0 < m[i] {
			m[i]--
			m[i*2]--
			cnt -= 2
			res = append(res, i)
			if m[i*2] < 0 {
				return nil
			}
		}
	}
	if 0 < cnt {
		return nil
	}
	return res
}

```

# Link ✅ 2013.detect-squares 检测正方形 
```go
package main

// https://leetcode-cn.com/problems/detect-squares

type DetectSquares struct {
	m map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{m: map[int]map[int]int{}}
}

func (ds *DetectSquares) Add(p []int) {
	x := p[0]
	y := p[1]
	if ds.m[x] == nil {
		ds.m[x] = map[int]int{}
	}
	ds.m[x][y] ++
}

func (ds *DetectSquares) Count(p []int) int {
	x1 := p[0]
	y1 := p[1]
	yMap := ds.m[x1] // 所有x = x1 的 y 点
	var res int
	for y2 := range yMap {
		dis := y2 - y1 // 需要的边长
		if dis != 0 {
			//  左边 右边
			ys := ds.m[x1-dis] // x1-dis 的 y 点
			if ys != nil {
				res += ys[y1] * ys[y2] * yMap[y2] // 与其他三点成为正方形
			}

			ys = ds.m[x1+dis] // x1+dis 的 y 点
			if ys != nil {
				res += ys[y1] * ys[y2] * yMap[y2] // 与其他三点成为正方形
			}
		}
	}
	return res
}

```

# Moore 投票 ✅ 0169.majority-element 众数 
```go
package main

// https://leetcode-cn.com/problems/majority-element

func majorityElement(nums []int) int {
	var num, cnt int
	for i := range nums {
		if num == nums[i] {
			cnt++
		} else {
			if cnt == 0 {
				num = nums[i]
				cnt = 1
			} else {
				cnt--
			}
		}
	}
	return num
}

```

# Moore 投票 ✅ 0229.majority-element-ii 求众数 II 
```go
package main

// https://leetcode-cn.com/problems/majority-element-ii

type candidate struct {
	num, cnt int
}

const base = 3
const csLen = base - 1

func majorityElement(nums []int) []int {
	var l = len(nums)
	var cs [csLen]candidate
	for _, num := range nums {
		var firstSlot = -1
		var k int
		for k < csLen && cs[k].num != num {
			if cs[k].cnt == 0 {
				firstSlot = k
			}
			k++
		}
		if k < csLen {
			cs[k].cnt++
		} else if k == csLen {
			if -1 < firstSlot {
				cs[firstSlot].num = num
				cs[firstSlot].cnt = 1
			} else {
				for i := range cs {
					cs[i].cnt--
				}
			}
		}
	}
	for i := range cs {
		cs[i].cnt = 0
	}
	for _, num := range nums {
		for i := range cs {
			if cs[i].num == num {
				cs[i].cnt++
				break
			}
		}
	}

	var limit = l / base
	var res []int
	for i := range cs {
		if limit < cs[i].cnt {
			res = append(res, cs[i].num)
		}
	}
	return res
}

```

# Permutation Heap ✅ 1086.high-five 前五科的均分 🤚 最高值 
```go
package main

import "sort"

// https://leetcode-cn.com/problems/high-five

func highFive(items [][]int) [][]int {
	m1 := map[int]int{}
	m2 := map[int]int{}
	slice := [][]int{}
	getSlice := func(id int) *[]int {
		_, ok := m1[id]
		if !ok {
			l := len(slice)
			m1[id] = l
			m2[l] = id
			slice = append(slice, []int{})
		}
		return &slice[m1[id]]
	}
	getId := func(index int) int {
		return m2[index]
	}
	for _, item := range items {
		s := getSlice(item[0])
		*s = append(*s, item[1])
	}

	var res [][]int
	for i := range slice {
		sort.Slice(slice[i], func(l, r int) bool {
			return slice[i][r] < slice[i][l]
		})
		var sum int
		for _, v := range slice[i][:5] {
			sum += v
		}
		sum /= 5
		res = append(res, []int{getId(i), sum})
	}

	sort.Slice(res, func(l, r int) bool {
		return res[l][0] < res[r][0]
	})
	return res
}

```

# Permutation ✅ 0128.longest-consecutive-sequence 最长递增 🤚 偏移 
```go
package main

// https://leetcode-cn.com/problems/longest-consecutive-sequence

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}

	var res int
	for num := range m {
		if !m[num-1] {
			cur := 1
			num += 1
			for m[num] {
				num++
				cur++
			}
			if res < cur {
				res = cur
			}
		}
	}
	return res
}

```

# Permutation ✅ 0242.valid-anagram 有效的字母异位词 🤚 
```go
package main

// https://leetcode-cn.com/problems/valid-anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := [26]int{}
	for i := range s {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}

	for _, cnt := range m {
		if cnt != 0 {
			return false
		}
	}
	return true
}

```

# Permutation ✅ 0266.palindrome-permutation 有效回文 
```go
package main

// https://leetcode-cn.com/problems/palindrome-permutation

func canPermutePalindrome(s string) bool {
	m := map[byte]int{}
	for i := range s {
		m[s[i]]++
	}

	var c = 0
	for _, v := range m {
		if v%2 == 1 {
			c++
		}
	}

	return c <= 1
}

```

# Permutation ✅ 0383.ransom-note  全排列 == 子排列 
```go
package main

// https://leetcode-cn.com/problems/ransom-note

func canConstruct(ransomNote string, magazine string) bool {
	m := map[byte]int{}
	for i := range magazine {
		m[magazine[i]]++
	}

	for i := range ransomNote {
		if m[ransomNote[i]] == 0 {
			return false
		}
		m[ransomNote[i]]--
	}
	return true
}

```

# Permutation ✅ 0409.longest-palindrome 最长回文串 
```go
package main

// https://leetcode-cn.com/problems/longest-palindrome

func longestPalindrome(s string) int {
	m := map[byte]int{}
	for i := range s {
		m[s[i]-'a']++
	}

	var cnt = len(s)
	var odd = 0
	for _, v := range m {
		if v%2 == 1 {
			odd++
		}
	}

	if 0 < odd {
		return cnt - odd + 1
	}
	return cnt
}

```

# Permutation ✅ 0575.distribute-candies 分糖果 🤚 最多类型 
```go
package main

// https://leetcode-cn.com/problems/distribute-candies

func distributeCandies(candyType []int) int {
	l1 := len(candyType)
	m := map[int]struct{}{}
	for i := range candyType {
		m[candyType[i]] = struct{}{} // 类型数
	}

	l2 := len(m)
	if l1/2 < l2 { // 类型数超过 一半，每样一颗
		return l1 / 2
	}
	return l2 // 小于一半 最多l2
}

```

# Permutation ✅ 0594.longest-harmonious-subsequence 最长相差 1 子排列 🤚 偏移 
```go
package main

// https://leetcode-cn.com/problems/longest-harmonious-subsequence

// 二次扫描
func findLHSTwo(nums []int) int {
	m := map[int]int{}
	for i := range nums {
		m[nums[i]]++
	}
	var res int
	for num, v := range m {
		if 0 < m[num+1] {
			cur := m[num+1] + v
			if res < cur {
				res = cur
			}
		}
	}
	return res
}

// 一次扫描
func findLHSOne(nums []int) int {
	m := map[int]int{}
	var res int
	for _, num := range nums {
		m[num]++
		if 0 < m[num+1] {
			cur := m[num+1] + m[num]
			if res < cur {
				res = cur
			}
		}
		if 0 < m[num-1] {
			cur := m[num-1] + m[num]
			if res < cur {
				res = cur
			}
		}
	}
	return res
}

```

# Permutation ✅ 0748.shortest-completing-word 最短补全词 
```go
package main

// https://leetcode-cn.com/problems/shortest-completing-word

func shortestCompletingWord(licensePlate string, words []string) string {
	m := [26]int{}
	for i := range licensePlate {
		if 'A' <= licensePlate[i] && licensePlate[i] <= 'Z' {
			m[licensePlate[i]-'A'] ++
		} else if 'a' <= licensePlate[i] && licensePlate[i] <= 'z' {
			m[licensePlate[i]-'a'] ++
		}
	}

	var res string
	var min = 1<<63 - 1
	var k int
	for i := range words {
		cur := [26]int{}
		for j := range words[i] {
			cur[words[i][j]-'a']++
		}
		for k = 0; k < 26; k++ {
			if cur[k] < m[k] {
				break
			}
		}
		if k == 26 {
			l1 := len(words[i])
			if l1 < min {
				res = words[i]
				min = len(words[i])
			}
		}
	}

	return res
}

```

# Permutation ✅ 0859.buddy-strings 交换一次后字符串相等 
```go
package main

// https://leetcode-cn.com/problems/buddy-strings

func buddyStrings(s string, goal string) bool {
	l1 := len(s)
	l2 := len(goal)
	if l1 != l2 {
		return false
	}

	if s == goal {
		m := map[byte]int{}
		for i := 0; i < l1; i++ {
			if m[s[i]] == 1 {
				return true
			}
			m[s[i]]++
		}
		return false
	}

	var first, second = -1, -1
	var cnt = [26]int{}
	for i := 0; i < l1; i++ {
		cnt[s[i]-'a']++
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}

	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}

```

# Permutation ✅ 1160.find-words-that-can-be-formed-by-characters 拼写单词 
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

# Permutation ✅ 1189.maximum-number-of-balloons “气球” 的最大数量 
```go
package main

// https://leetcode-cn.com/problems/maximum-number-of-balloons

func maxNumberOfBalloons(text string) int {
	m := map[byte]int{}
	for i := range text {
		m[text[i]]++
	}
	c := map[byte]int{
		'a': 1,
		'b': 1,
		'n': 1,
		'o': 2,
		'l': 2,
	}

	var res = 1<<63 - 1
	for b := range c {
		cur := m[b] / c[b]
		if cur < res {
			res = cur
		}
	}
	return res
}

```

# Permutation ✅ 1426.counting-elements 存在 (值+1 的后置元素) 的元素个数 🤚 偏移 
```go
package main

// https://leetcode-cn.com/problems/counting-elements

func countElements(arr []int) int {
	m := map[int]int{}
	for i := range arr {
		m[arr[i]]++
	}

	var res int
	for i := range m {
		if 0 < m[i+1] {
			res += m[i]
		}
	}
	return res
}

```

# Permutation ✅ 1460.make-two-arrays-equal-by-reversing-sub-arrays 通过翻转子数组使两个数组相等 
```go
package main

// https://leetcode-cn.com/problems/make-two-arrays-equal-by-reversing-sub-arrays

func canBeEqual(target []int, arr []int) bool {
	l1 := len(target)
	l2 := len(arr)
	if l1 != l2 {
		return false
	}
	m := map[int]int{}
	for i := 0; i < l1; i++ {
		m[target[i]]++
		m[arr[i]]--
	}

	for i := range m {
		if m[i] != 0 {
			return false
		}
	}
	return true
}

```

# Permutation ✅ LCS 02. 完成一半题目最少的题型 🤚 最少类型 
```go
package main

import "sort"

// https://leetcode-cn.com/problems/WqXACV/

func halfQuestions(questions []int) int {
	m := make([]int, 1001)
	half := len(questions) / 2
	for i := range questions {
		m[questions[i]]++
	}

	sort.Slice(m, func(i, j int) bool {
		return m[j] < m[i]
	})

	var res int
	for i := range m {
		if half <= 0 {
			return res
		}
		half -= m[i]
		res++
	}
	return res
}

```

# Pre ✅ 0930.binary-subarrays-with-sum 和相同的二元子数组 
```go
package main

// https://leetcode-cn.com/problems/binary-subarrays-with-sum

func numSubarraysWithSum(nums []int, goal int) int {
	var res, sum int
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[sum]++
		sum += num
		res += cnt[sum-goal] // 前缀和 sum[j] - sum[i] = goal
	}
	return res
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

