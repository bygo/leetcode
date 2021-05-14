package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

func main() {
	var md bytes.Buffer

	all, err := ioutil.ReadDir("./")
	if err == nil {
		for _, item := range all {
			sqlDir, _ := ioutil.ReadDir(fmt.Sprintf("./%s", item.Name()))
			var b = bytes.NewBufferString(fmt.Sprintf("# %s \n", item.Name()))
			for _, sqlFile := range sqlDir {
				b.WriteString("```sql\n")
				s, _ := ioutil.ReadFile(fmt.Sprintf("./%s/%s", item.Name(), sqlFile.Name()))
				b.Write(s)
				b.WriteString("\n```\n\n")
				md.WriteString(b.String())
			}
		}
	} else {
		log.Fatal(err)
	}


	ioutil.WriteFile("readme.md", md.Bytes(), 0777)
}

func longestBeautifulSubstring(word string) int {
	var m = map[rune]int{
		'a': 0, 'e': 1, 'i': 2, 'o': 3, 'u': 4,
	}
	var index, cur, res int
	for _, v := range word {
		if index == m[v] || index+1 == m[v] {
			cur++
			if v == 'u' && res < cur {
				res = cur
			}
			index = m[v]
		} else {
			index = -1
			if m[v] == 0 {
				cur = 1
			} else {
				cur = 0
			}
		}
	}
	return res
}

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	var res, cur, i, index int

	for k1 := range nums {
		i = k
		index = 0
		for k2, v2 := range nums[k1:] {
			if k2 != 0 {
				i = i - (v2-cur)*(k2+1)
			}
			if i < 0 {
				break
			}
			index++
			cur = v2
		}
		if res < index {
			res = index
		}
	}
	return res
}
