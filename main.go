package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Problem struct {
	No    string
	Title string
	File  string
}

var container = map[string][]Problem{}
var readme string

func main() {
	read()
	output()
	did()
}

func read() {
	classes, err := ioutil.ReadDir("./")
	check(err)

	for _, class := range classes {
		if class.IsDir() && class.Name()[0] != '.' {
			cName := class.Name()
			container[cName] = []Problem{}

			problems, err := ioutil.ReadDir(cName)
			check(err)

			for _, problem := range problems {
				desc := strings.Split(problem.Name(), ".")
				container[cName] = append(container[cName], Problem{
					No:    desc[0],
					Title: desc[1],
					File:  "https://github.com/temporaries/leetcode/tree/master/" + class.Name() + "/" + problem.Name() + "/main.go",
				})
			}
		}
	}
}

func output() {
	stubReadme, err := ioutil.ReadFile("./layout.md")
	check(err)
	for index, problems := range container {
		stubClass, err := ioutil.ReadFile("./class.md")
		check(err)
		class := strings.Replace(string(stubClass), "@DummyClass", normalizeClassTitle(index), 1)
		for _, problem := range problems {
			class += fmt.Sprintf("\n| %s | [%-32s](https://leetcode-cn.com/problems/%s) | [Go](%s)|",
				problem.No, normalizeProblemTitle(problem.Title), normalizeProblemUri(problem.Title), problem.File)
		}
		readme += class
	}

	readme = strings.Replace(string(stubReadme), "@DummyTable", readme, 1)
	ioutil.WriteFile("readme.md", []byte(readme), os.ModePerm)
}

func normalizeClassTitle(title string) string {
	words := strings.Split(title, "_")
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	s := strings.Join(words, " ")
	return s
}

func normalizeProblemUri(t string) string {
	l := len(t)
	title := append(make([]byte, l*2))
	copy(title[0:l], t)
	for i := 0; i < l; i++ {
		if 64 < title[i] && title[i] < 91 {
			title[i] += 'a' - 'A'
			copy(title[i+1:l+1], title[i:l])
			l++
			title[i] = '-'
			i++
		}
	}
	return strings.Trim(string(title[0:l]), "")
}
func normalizeProblemTitle(t string) string {
	l := len(t)
	title := append(make([]byte, l*2))
	copy(title[0:l], t)
	for i := 0; i < l; i++ {
		if 64 < title[i] && title[i] < 91 {
			copy(title[i+1:l+1], title[i:l])
			l++
			title[i] = ' '
			i++
		}
	}
	title[0] -= 'a' - 'A'
	return strings.Trim(string(title[0:l]), "")
}

func did() {
	os.Remove("leetcode")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
