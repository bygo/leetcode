# Cnt Diff Xor ✅ 0389.find-the-difference 找不同

```go
package main

// https://leetcode-cn.com/problems/find-the-difference

func findTheDifferenceCnt(s string, t string) byte {
	m := [26]int{}
	for i := range s {
		ch := s[i] - 'a'
		m[ch]++
	}

	for i := range t {
		ch := t[i] - 'a'
		if m[ch] == 0 {
			return ch + 'a'
		}
		m[ch]--
	}
	return 0
}

func findTheDifferenceDistance(s string, t string) byte {
	var sum = 0
	for i := range s {
		sum = sum - int(s[i]) + int(t[i])
	}
	return byte(sum + int(t[len(t)-1]))
}

func findTheDifferenceXor(s string, t string) byte {
	var sum byte
	for i := range s {
		sum ^= s[i] ^ t[i]
	}
	return sum ^ t[len(t)-1]
}

```

# Cnt Heap ❌ 1086.high-five 前五科的均分

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

# Cnt ✅ 0359.logger-rate-limiter 限流器

```go
package main

// https://leetcode-cn.com/problems/logger-rate-limiter

type Logger struct {
	m map[string]int
}

func Constructor() Logger {
	return Logger{m: make(map[string]int)}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if timestamp < l.m[message] {
		return false
	}
	l.m[message] = timestamp + 10
	return true
}

```

# Cnt ✅ 0604.design-compressed-string-iterator 迭代压缩字符串

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

/**
 * Your StringIterator object will be instantiated and called as such:
 * obj := Constructor(compressedString);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

```

# Cnt ✅ 0888.fair-candy-swap 糖果棒交换后数量相等

```go
package main

// https://leetcode-cn.com/problems/fair-candy-swap

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	m := map[int]struct{}{}
	sum1 := 0
	for i := range aliceSizes {
		m[aliceSizes[i]] = struct{}{}
		sum1 += aliceSizes[i]
	}

	sum2 := 0
	for i := range bobSizes {
		sum2 += bobSizes[i]
	}

	diff := (sum1 - sum2) / 2

	for i := range bobSizes {
		_, ok := m[bobSizes[i]+diff]
		if ok {
			return []int{bobSizes[i] + diff, bobSizes[i]}
		}
	}
	return nil
}

```

# Cnt ✅ 1941.check-if-all-characters-have-equal-number-of-occurrences 检查是否所有字符出现次数相同

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

# Cnt ✅ LCS 02. 完成一半题目最少的题型

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

# Diff ✅ 0299.bulls-and-cows 公牛奶牛

```go
package main

import "strconv"

// https://leetcode-cn.com/problems/bulls-and-cows

func getHint(secret string, guess string) string {
	m1 := [10]int{}
	m2 := [10]int{}

	var res []byte
	var bull, cow int
	for i := range guess {
		if guess[i] == secret[i] {
			bull++
		} else {
			m1[secret[i]-'0']++
			m2[guess[i]-'0']++
		}
	}

	for i := 0; i < 10; i++ {
		cow += min(m1[i], m2[i])
	}

	res = append(res, strconv.Itoa(bull)...)
	res = append(res, 'A')
	res = append(res, strconv.Itoa(cow)...)
	res = append(res, 'B')
	return string(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

```

# Diff ✅ 0349.intersection-of-two-arrays 两个数组的交集

```go
package main

// https://leetcode-cn.com/problems/intersection-of-two-arrays

func intersection(nums1 []int, nums2 []int) []int {
	var m = map[int]struct{}{}
	for _, v := range nums1 {
		m[v] = struct{}{}
	}

	var res []int
	for _, v := range nums2 {
		_, ok := m[v]
		if ok {
			res = append(res, v)
			delete(m, v)
		}
	}
	return res
}

```

# Diff ✅ 0350.intersection-of-two-arrays-ii 两个数组的交集 II

```go
package main

// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for i := range nums1 {
		m[nums1[i]]++
	}

	var res []int
	for i := range nums2 {
		if 0 < m[nums2[i]] {
			res = append(res, nums2[i])
			m[nums2[i]]--
		}
	}
	return res
}

```

# Diff ✅ 0771.jewels-and-stones 找J在S中的数量

```go
package main

// https://leetcode-cn.com/problems/jewels-and-stones/

func numJewelsInStones(J string, S string) int {
	var m = map[byte]int{}
	for i := range J {
		m[J[i]] = 1
	}

	var res int
	for i := range S {
		res += m[S[i]]
	}
	return res
}

```

# Diff ✅ 0819.most-common-word 最常见的单词

```go
package main

// https://leetcode-cn.com/problems/most-common-word

func mostCommonWord(paragraph string, banned []string) string {
	m1 := map[string]int{}
	m2 := map[string]int{}
	banned = append(banned, "#")
	for i := range banned {
		m1[banned[i]]++
	}

	res := ""
	max := 0
	cur := []byte{}
	f := func() {
		s := string(cur)
		if 0 < len(cur) {
			if 0 == m1[s] {
				m2[s]++
				if max < m2[s] {
					res = s
					max = m2[s]
				}
			}
		}
		cur = []byte{}
	}
	for i := range paragraph {
		if 'A' <= paragraph[i] && paragraph[i] <= 'Z' {
			cur = append(cur, paragraph[i]+32)
		} else if 'a' <= paragraph[i] && paragraph[i] <= 'z' {
			cur = append(cur, paragraph[i])
		} else {
			f()
		}
	}
	f()
	return res
}

```

# Diff ✅ 1002.find-common-characters 查找公共字符

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
	for i := 0; i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			res = append(res, string(i+'a'))
		}
	}
	return
}

```

# Distance ✅ 1165.single-row-keyboard 机械手最少移动次数

```go
package main

// https://leetcode-cn.com/problems/single-row-keyboard

func calculateTime(keyboard string, word string) int {
	var m = [26]int{}
	for i := range keyboard {
		m[keyboard[i]-'a'] = i
	}

	var res = m[word[0]-'a']
	var l = len(word)
	var i = 1
	for i < l {
		res += abs(m[word[i-1]-'a'], m[word[i]-'a'])
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

# Divide ✅ 1119.remove-vowels-from-a-string 删去字符串中的元音

```go
package main

// https://leetcode-cn.com/problems/remove-vowels-from-a-string

func removeVowels(s string) string {
	l := len(s)
	var res = make([]byte, 0, l)
	for i := range s {
		if s[i] != 'a' && s[i] != 'e' && s[i] != 'i' && s[i] != 'o' && s[i] != 'u' {
			res = append(res, s[i])
		} else {
			l--
		}
	}
	return string(res[:l])
}

```

# Divide ✅ 1763.longest-nice-substring 大小写都出现的最长子串

```go
package main

// https://leetcode-cn.com/problems/longest-nice-substring

func longestNiceSubstring(s string) string {
	var dfs func(sub string) string
	dfs = func(sub string) string {
		var m = map[byte]int{}
		for i := range sub {
			if sub[i] <= 'Z' && (m[sub[i]+32] == 0 || m[sub[i]+32] == 1) {
				m[sub[i]+32] += 2
			} else if 'a' <= sub[i] && (m[sub[i]] == 0 || m[sub[i]] == 2) {
				m[sub[i]] += 1
			}
		}

		var left int
		var res string
		for j := range sub {
			ch := sub[j]
			if ch <= 'Z' {
				ch += 32
			}
			if m[ch] != 3 {
				cur := dfs(sub[left:j])
				if len(res) < len(cur) {
					res = cur
				}
				left = j + 1
			}
		}
		if 0 < left {
			cur := dfs(sub[left:])
			if len(res) < len(cur) {
				res = cur
			}
			return res
		}
		return sub
	}
	return dfs(s)
}

```

# Filter ✅ 0217.contains-duplicate 出现重复

```go
package main

//
// https://leetcode-cn.com/problems/contains-duplicate

func containsDuplicate(nums []int) bool {
	var m = map[int]int{}
	for _, num := range nums {
		if 0 < m[num] {
			return true
		}
		m[num] += 1
	}
	return false
}

```

# Filter ✅ 0884.uncommon-words-from-two-sentences 两句话中的不常见单词

```go
package main

import "strings"

// https://leetcode-cn.com/problems/uncommon-words-from-two-sentences

func uncommonFromSentences(s1 string, s2 string) []string {
	a1 := strings.Split(s1, " ")
	a2 := strings.Split(s2, " ")

	m := map[string]int{}

	for i := range a1 {
		m[a1[i]]++
	}

	for i := range a2 {
		m[a2[i]]++
	}

	var res []string
	for k := range m {
		if m[k] == 1 {
			res = append(res, k)
		}
	}
	return res
}

```

# Filter ✅ 1133.largest-unique-number 最大唯一数

```go
package main

// https://leetcode-cn.com/problems/largest-unique-number

func largestUniqueNumber(nums []int) int {
	m := [1001]int{}
	for i := range nums {
		m[nums[i]]++
	}

	for i := 1000; 0 <= i; i-- {
		if m[i] == 1 {
			return i
		}
	}
	return -1
}

```

# Filter ✅ 1394.find-lucky-integer-in-an-array 找出数组中的幸运数

```go
package main

// https://leetcode-cn.com/problems/find-lucky-integer-in-an-array

func findLucky(arr []int) int {
	m := [501]int{}
	for i := range arr {
		m[arr[i]]++
	}

	for i := 500; 0 < i; i-- {
		if m[i] == i {
			return i
		}
	}
	return -1
}

```

# Filter ✅ 1748.sum-of-unique-elements 唯一元素的和

```go
package main

// https://leetcode-cn.com/problems/sum-of-unique-elements

func sumOfUnique(nums []int) int {
	var m = map[int]int{}
	var res int
	for _, v := range nums {
		if m[v] == 0 {
			res += v
			m[v] = 1
		} else if m[v] == 1 {
			res -= v
			m[v] = 2
		}
	}
	return res
}

```

# Filter ✅ 1935.maximum-number-of-words-you-can-type 可以输入的最大单词数

```go
package main

// https://leetcode-cn.com/problems/maximum-number-of-words-you-can-type

func canBeTypedWords(text string, brokenLetters string) int {
	var m = map[byte]bool{}
	for i := range brokenLetters {
		m[brokenLetters[i]] = true
	}
	var res int
	var normal = false

	var i int
	var l = len(text)
	for i < l {
		normal = true
		for i < l && text[i] != ' ' {
			if m[text[i]] {
				normal = false
				for i < l && text[i] != ' ' {
					i++
				}
			} else {
				i++
			}
		}
		if normal {
			res++
		}
		i++
	}
	return res
}

```

# Group GCD ✅ 0914.x-of-a-kind-in-a-deck-of-cards 卡牌分组

```go
package main

// https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards

func hasGroupsSizeX(deck []int) bool {
	if len(deck) <= 1 {
		return false
	}

	m := map[int]int{}
	for i := range deck {
		m[deck[i]]++
	}

	g := -1
	for i := range m {
		if 0 < m[i] {
			if g == -1 {
				g = m[i]
			} else {
				g = gcd(g, m[i])
			}
		}
	}
	return 2 <= g
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	} else {
		return gcd(y%x, x)
	}
}

```

# Group ✅ 0288.unique-word-abbreviation 单词是唯一缩写

```go
package main

import "strconv"

// https://leetcode-cn.com/problems/unique-word-abbreviation

type ValidWordAbbr struct {
	m map[string]map[string]int
}

func Constructor(dictionary []string) ValidWordAbbr {
	m := map[string]map[string]int{}
	for i := range dictionary {
		s := Abbr(dictionary[i])
		if m[s] == nil {
			m[s] = map[string]int{}
		}
		m[s][dictionary[i]] ++
	}
	return ValidWordAbbr{
		m: m,
	}
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	s := Abbr(word)
	return this.m[s] == nil || len(this.m[s]) == 1 && 0 < this.m[s][word]
}

func Abbr(word string) string {
	l := len(word)
	b := []byte{word[0]}
	b = append(b, strconv.Itoa(l-2)...)
	b = append(b, word[l-1])
	return string(b)
}

/**
 * Your ValidWordAbbr object will be instantiated and called as such:
 * obj := Constructor(dictionary);
 * param_1 := obj.IsUnique(word);
 */

```

# Group ✅ 0348.design-tic-tac-toe 设计井字棋

```go
package main

// https://leetcode-cn.com/problems/design-tic-tac-toe

type TicTacToe struct {
	row []int
	col []int
	ang [2]int
	n   int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		row: make([]int, n),
		col: make([]int, n),
		ang: [2]int{},
		n:   n,
	}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	symbol := 1
	if player == 2 {
		symbol = -1
	}
	this.row[row] += symbol
	if this.win(this.row[row]) {
		return player
	}
	this.col[col] += symbol
	if this.win(this.col[col]) {
		return player
	}

	if col-row == 0 {
		this.ang[0] += symbol
		if this.win(this.ang[0]) {
			return player
		}
	}
	if col+row == this.n-1 {
		this.ang[1] += symbol
		if this.win(this.ang[1]) {
			return player
		}
	}
	return 0
}

func (this *TicTacToe) win(i int) bool {
	return i == this.n || i == -this.n
}

```

# Group ✅ 0500.keyboard-row 只返回键盘同一行的字母

```go
package main

// https://leetcode-cn.com/problems/keyboard-row

var m = map[byte]uint{}

func init() {
	for _, v := range []byte("qwertyuiopQWERTYUIOP") {
		m[v] = 1
	}
	for _, v := range []byte("asdfghjklASDFGHJKL") {
		m[v] = 2
	}
	for _, v := range []byte("zxcvbnmZXCVBNM") {
		m[v] = 3
	}
}

func findWords(words []string) []string {
	var res []string
	for _, word := range words {
		i := 0
		l := len(word)
		for i < l && m[word[i]] == m[word[0]] {
			i++
		}
		if i == l {
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
	"bytes"
	"strconv"
)

// https://leetcode-cn.com/problems/subdomain-visit-count

func subdomainVisits(cpdomains []string) []string {
	var m = map[string]int{}
	for _, c := range cpdomains {
		var num int
		var i int
		var n = len(c)
		for c[i] != ' ' {
			num = num*10 + int(c[i]-'0')
			i++
		}
		i += 1

		var words [][]byte
		cur := []byte{}
		for i < n {
			if c[i] == '.' {
				words = append(words, cur)
				cur = []byte{}
			} else {
				cur = append(cur, c[i])
			}
			i++
		}
		if 0 < len(cur) {
			words = append(words, cur)
		}

		l := len(words)
		for j := 0; j < l; j++ {
			w := bytes.Join(words[j:], []byte{'.'})
			m[string(w)] += num
		}
	}
	var res []string
	for k, v := range m {
		var cur []byte
		cur = append(cur, strconv.Itoa(v)...)
		cur = append(cur, ' ')
		cur = append(cur, k...)
		res = append(res, string(cur))
	}
	return res
}

```

# Group ✅ 0929.unique-email-addresses 不同的邮件地址数

```go
package main

// https://leetcode-cn.com/problems/unique-email-addresses

func numUniqueEmails(emails []string) int {
	m := map[string]struct{}{}
	for _, email := range emails {
		l1 := len(email)
		var i int
		var cur []byte
		for i < l1 && email[i] != '@' {
			if email[i] == '+' {
				for i < l1 && email[i] != '@' {
					i++
				}
			} else {
				if email[i] != '.' {
					cur = append(cur, email[i])
				}
				i++
			}
		}
		for i < l1 {
			cur = append(cur, email[i])
			i++
		}
		m[string(cur)] = struct{}{}
	}
	return len(m)
}

```

# Group ✅ 1128.number-of-equivalent-domino-pairs 等价多米诺骨牌对的数量

```go
package main

// https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs

func numEquivDominoPairs(d [][]int) int {
	m := map[int]int{}
	var res int
	for i := range d {
		if d[i][0] < d[i][1] {
			d[i][0], d[i][1] = d[i][1], d[i][0]
		}
		v := d[i][0]*10 + d[i][1]
		res += m[v]
		m[v]++
	}
	return res
}

```

# Group ✅ 1275.find-winner-on-a-tic-tac-toe-game 找出井字棋的获胜者

```go
package main

// https://leetcode-cn.com/problems/find-winner-on-a-tic-tac-toe-game

func tictactoe(moves [][]int) string {
	top := len(moves) - 1

	m := [8]int{}

	for i := top; 0 <= i; i -= 2 {
		m[moves[i][0]]++
		m[moves[i][1]+3]++
		if moves[i][0] == moves[i][1] {
			m[6]++
		}
		if moves[i][0]+moves[i][1] == 2 {
			m[7]++
		}
	}

	l1 := len(moves)
	var res string
	if l1%2 == 0 {
		res = "B"
	} else {
		res = "A"
	}
	for i := range m {
		if m[i] == 3 {
			return res
		}
	}
	if l1 < 9 {
		return "Pending"
	}

	return "Draw"
}

```

# Group ✅ 1399.count-largest-group 统计最大组的数目

```go
package main

// https://leetcode-cn.com/problems/count-largest-group

func countLargestGroup(n int) int {
	m := map[int]int{}
	for i := 1; i <= n; i++ {
		j := i
		sum := 0
		for 0 < j {
			sum += j % 10
			j /= 10
		}
		m[sum] ++
	}
	cnt := 0
	max := 0
	for i := range m {
		if max < m[i] {
			max = m[i]
			cnt = 1
		} else if m[i] == max {
			cnt++
		}
	}
	return cnt
}

```

# Group ✅ 1512.number-of-good-pairs 好数对的数目 i < j && a[i] == a[j]

```go
package main

// https://leetcode-cn.com/problems/number-of-good-pairs

func numIdenticalPairs(nums []int) int {
	var m = map[int]int{}
	var res int
	for _, v := range nums {
		if 0 < m[v] {
			res += m[v]
		}
		m[v]++
	}
	return res
}

```

# Group ✅ 1742.maximum-number-of-balls-in-a-box 盒子中小球的最大数量

```go
package main

// https://leetcode-cn.com/problems/maximum-number-of-balls-in-a-box

func countBalls(lowLimit int, highLimit int) int {
	var sum int
	j := lowLimit
	for 0 < j {
		sum += j % 10
		j /= 10
	}
	m := [46]int{}
	var max int
	for i := lowLimit; i <= highLimit; i++ {
		m[sum]++
		k := i
		for k%10 == 9 { // 每次进1 减9
			sum -= 9
			k /= 10
		}
		sum++
	}

	for i := range m {
		if max < m[i] {
			max = m[i]
		}
	}

	return max
}

```

# Group ✅ 1805.number-of-different-integers-in-a-string 字符串中不同整数的数目

```go
package main

// https://leetcode-cn.com/problems/number-of-different-integers-in-a-string

func numDifferentIntegers(word string) int {
	m := map[string]struct{}{}
	var cur []byte
	f := func() {
		l1 := len(cur)
		if 0 < l1 {
			j := 0
			for cur[j] == '0' && j < l1-1 {
				j++
			}
			m[string(cur[j:])] = struct{}{}
			cur = []byte{}
		}
	}
	for i := range word {
		if '0' <= word[i] && word[i] <= '9' {
			cur = append(cur, word[i])
		} else {
			f()
		}
	}
	f()
	return len(m)
}

```

# Group ✅ 2006.count-number-of-pairs-with-absolute-difference-k 差的绝对值为 K 的数对数目

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

# Group ✅ 2032.two-out-of-three 至少在两个数组中出现的值

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

# Link ✅ 0953.verifying-an-alien-dictionary 验证外星语词典

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

# Link ✅ 1122.relative-sort-array 数组的相对排序

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

# Link ✅ 1331.rank-transform-of-an-array 数组序号转换

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

# Link ✅ 1346.check-if-n-and-its-double-exist 检查整数及其两倍数是否存在

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

# Link ✅ 1370.increasing-decreasing-string 上升下降字符串

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

# Link ❌ 0244.shortest-word-distance-ii 最短单词距离 II

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

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */

```

# Link ❌ 0599.minimum-index-sum-of-two-lists 相同值的最小索引总和

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

# Link ❌ 1640.check-array-formation-through-concatenation 能否连接形成数组

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

# Link ❌ 2007.find-original-array-from-doubled-array 从双倍数组中还原原数组

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

# Link ❌ 2013.detect-squares 检测正方形

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

# Moore 投票 ❌ 0229.majority-element-ii 求众数 II

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

# Permutation ✅ 0383.ransom-note 全排列 == 子排列

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

# Pre ❌ 1711.count-good-meals 大餐计数

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

# Score ❌ 0387.first-unique-character-in-a-string 字符串中的第一个唯一字符 🤚 标记索引 剔除

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
			b = append(b, (char[j]-char[0]+26)%26) // 以方变负
			j++
		}
		id := getId(string(b))
		res[id] = append(res[id], char)
	}
	return res
}

```

# Vector ✅ 0438.find-all-anagrams-in-a-string 找到字符串中所有字母异位词

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

# Window ✅ 0961.n-repeated-element-in-size-2n-array 2N 中重复 N 次的元素

```go
package main

// https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array

func repeatedNTimes(nums []int) int {
	l1 := len(nums)
	for j := 1; j <= 3; j++ {
		for i := 0; i < l1-j; i++ {
			if nums[i] == nums[i+j] {
				return nums[i]
			}
		}
	}

	return -1
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

