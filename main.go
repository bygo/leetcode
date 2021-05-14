package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	Stat       Stat       `json:"stat"`
	Difficulty Difficulty `json:"difficulty"`
	File       string
	Algorithm  string
}

type Difficulty struct {
	Level int `json:"level"`
}

func getMinSwaps(num string, k int) int {
	var arr = []int{}
	for _, v := range num {
		arr = append(arr, int(v-48))
	}
	tail := len(arr) - 1
	i := 0
	j := 0
	for i < k {
		for j < tail {
			if arr[tail-j-1] < arr[tail-j] {
				arr[tail-j], arr[tail-j-1] = arr[tail-j-1], arr[tail-j]
				i++
				break
			}
			j++
		}
	}

	return i
}

type Stat struct {
	QuestionId          int    `json:"question_id"`
	QuestionTitle       string `json:"question__title"`
	QuestionTitleSlug   string `json:"question__title_slug"`
	QuestionHide        bool   `json:"question__hide"`
	TotalAcs            int    `json:"total_acs"`
	TotalSubmitted      int    `json:"total_submitted"`
	TotalColumnArticles int    `json:"total_column_articles"`
	FrontendQuestionId  string `json:"frontend_question_id"`
	IsNewQuestion       int    `json:"is_new_question"`
}

type Profile struct {
	Username        string    `json:"user_name"`
	NumSolved       int       `json:"num_solved"`
	NumTotal        int       `json:"num_total"`
	AcEasy          int       `json:"ac_easy"`
	AcMedium        int       `json:"ac_medium"`
	AcHard          int       `json:"ac_hard"`
	StatStatusPairs []Problem `json:"stat_status_pairs"`
}

var problems = map[string][]Problem{}

//数据结构 排序，并且存在时生成
var SolutionOrder = []string{
	"array", "linked_list", "math", "stack", "string", "tree",
}
var readme string
var profile Profile
var currentClassName string
var c string

func main() {
	will()

	a := flag.String("p", "r", "generate readme.md | start a new problem")
	t := flag.String("t", "", "type")
	flag.Parse()
	if *a == "r" {
		read()
		output()
	} else {
		//id, err := strconv.Atoi(*a)
		//check(err)
		start(*a, *t)
	}

	did()
}

func will() {
	var body []byte
	body, err := ioutil.ReadFile("all.json")
	if err != nil {
		println("Update all.json")
		resp, err := http.Get(strings.Join([]string{
			"https://leetcode-cn.com/api/problems/all/",
		}, ""))
		check(err)

		body, _ = ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		ioutil.WriteFile("all.json", body, os.ModePerm)
	}
	json.Unmarshal(body, &profile)
}

func read() {
	class, err := ioutil.ReadDir("./")
	check(err)
	for _, c := range class {
		if c.IsDir() && c.Name()[0] != '.' && valueOf(SolutionOrder, c.Name()) {
			currentClassName = c.Name()
			problems[currentClassName] = []Problem{}
			getSolutions(c, currentClassName)
		}
	}
}

func getSolutions(dir os.FileInfo, path string) {
	if strings.Contains(dir.Name(), "_test") {
		return
	}
	if dir.IsDir() {
		dirs, err := ioutil.ReadDir(path)
		check(err)
		for _, d := range dirs {
			getSolutions(d, path+"/"+d.Name())
		}
		return
	}

	desc := strings.Split(path, "/")
	l := len(desc)
	//fmt.Printf("%+v", desc)
	algorithm := strings.Replace(strings.Join(desc[2:l-1], "."), "_", " ", -1)
	title := strings.Split(desc[1], ".")
	//id, _ := strconv.Atoi(title[0])
	p := find(title[0])
	p.File = "https://github.com/temporaries/leetcode/tree/master/" + path
	p.Algorithm = algorithm
	problems[currentClassName] = append(problems[currentClassName], p)
}

func getTemplates(dir os.FileInfo, path string) {
	if strings.Contains(dir.Name(), "_test") {
		return
	}
	if dir.IsDir() {
		dirs, err := ioutil.ReadDir(path)
		check(err)
		for _, d := range dirs {
			getTemplates(d, path+"/"+d.Name())
		}
		return
	}
}

func find(id string) Problem {
	id = strings.TrimLeft(id, "0")
	left, right := 0, len(profile.StatStatusPairs)
	for left < right {
		//i, _ := strconv.Atoi(profile.StatStatusPairs[left].Stat.FrontendQuestionId)
		if profile.StatStatusPairs[left].Stat.FrontendQuestionId == id {
			return profile.StatStatusPairs[left]
		}
		left++
	}
	//for left < right {
	//	mid := (left + right) / 2
	//	i, _ := strconv.Atoi(profile.StatStatusPairs[mid].Stat.FrontendQuestionId)
	//	if id < i {
	//		left = mid
	//	} else if i < id {
	//		right = mid
	//	} else {
	//		return profile.StatStatusPairs[mid]
	//	}
	//}

	panic(id + " not exists")
}

func output() {
	stubReadme, err := ioutil.ReadFile("./readmeStub.md")
	difficulty := []string{
		"Easy", "Medium", "Hard",
	}
	check(err)

	var directoryIndex string
	for _, index := range SolutionOrder {
		className := normalizeClassTitle(index)
		directoryIndex += fmt.Sprintf("- [%s](#%s)\n\r", className, className)
		problems := problems[index]
		stubClass, err := ioutil.ReadFile("./class.stub")
		check(err)
		class := strings.Replace(string(stubClass), "@DummyClass", className, 1)
		for k, problem := range problems {
			questionId := fmt.Sprintf("%04s", problem.Stat.FrontendQuestionId)
			questionTitle := problem.Stat.QuestionTitle
			questionTitleSlug := problem.Stat.QuestionTitleSlug
			questionAcceptance := fmt.Sprintf("%.1f%s", float64(problem.Stat.TotalAcs)*100/float64(problem.Stat.TotalSubmitted), "%")
			questionDifficulty := difficulty[problem.Difficulty.Level-1]
			if k != 0 && problems[k-1].Stat.QuestionId == problem.Stat.QuestionId {
				questionId, questionTitle, questionDifficulty, questionAcceptance, questionTitleSlug = "", "", "", "", ""
			}
			class += fmt.Sprintf("\n| %s | [%-32s](https://leetcode-cn.com/problems/%s) | %s | %s | [Go](%s)| %s",
				questionId,
				questionTitle,
				questionTitleSlug,
				questionAcceptance,
				questionDifficulty,
				problem.File,
				problem.Algorithm,
			)
		}
		readme += class
	}
	readme = strings.Replace(string(stubReadme), "@DummyTable", readme, 1)
	readme = strings.Replace(readme, "@DummyIndex", directoryIndex, 1)
	ioutil.WriteFile("readme.md", []byte(readme), os.ModePerm)
}

func normalizeClassTitle(title string) string {
	words := strings.Split(title, "_")
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	s := strings.Join(words, "")
	return s
}

type Stub struct {
	Stub string
	File string
}

func start(id string, t string) {
	p := find(id)

	problemStub, err := ioutil.ReadFile("./problem.stub")
	v, ok := map[string]Stub{
		"t": {"./tree.stub", "main.go"},
		"l": {"./link.stub", "main.go"},
		"s": {"./sql.stub", "1.sql"},
	}[t]
	if !ok {
		return
	}

	problemStub, err = ioutil.ReadFile(v.Stub)

	check(err)
	i, _ := strconv.Atoi(p.Stat.FrontendQuestionId)
	path := fmt.Sprintf("%04d.%s", i, p.Stat.QuestionTitleSlug)

	problemStub = bytes.Replace(problemStub, []byte("DumpTitle"), []byte(p.Stat.QuestionTitle), 1)
	problemStub = bytes.Replace(problemStub, []byte("DumpLink"), []byte(fmt.Sprintf("https://leetcode-cn.com/problems/%s", p.Stat.QuestionTitleSlug)), 1)

	os.Mkdir(path, os.ModePerm)

	ioutil.WriteFile(fmt.Sprintf("%s/%s", path, v.File), problemStub, os.ModePerm)
}

//func parse(s []string) map[string]string {
//	args := make(map[string]string)
//	for _, v := range s {
//		arg := strings.Split(v, "=")
//		args[arg[0]] = arg[1]
//	}
//	return args
//}

func did() {
	os.Remove("leetcode")
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func valueOf(str []string, value string) bool {
	for _, v := range str {
		if v == value {
			return true
		}
	}
	return false
}
