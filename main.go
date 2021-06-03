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

const (
	ApiProblemsAll     = "https://leetcode-cn.com/api/problems/all/"
	AllJsonFile        = "all.json"
	ALlJsonUpdatedText = "update all.json"

	StubPrefix = "stub/"
	RealPrefix = ""

	ReadmeClassStub = "class.stub"
	ReadmeMdStub    = "readmeStub.md"
	ReadmeName      = "readme.md"

	DummyClass = "@DummyClass"
	DummyTable = "@DummyTable"
	DummyIndex = "@DummyIndex"
	DummyLink  = "@DummyLink"
	DummyTitle = "@DummyTitle"

	Easy   = "Easy"
	Medium = "Medium"
	Hard   = "Hard"

	DefaultProblemStub = "default"
)

var profile Profile
var currentClassName string

var problems = map[string][]Problem{}

var solutionOrder = []string{
	"array", "bit", "dp", "hash", "linked_list", "sql", "math", "stack", "string", "tree",
}

var problemStubs = map[string]*Stub{
	"default": {"problem.stub", "main.go"},
	"tree":    {"tree.stub", "main.go"},
	"link":    {"link.stub", "main.go"},
	"sql":     {"sql.stub", "1.sql"},
}
var dummyClassBuf = []byte(DummyClass)
var dummyTableBuf = []byte(DummyTable)
var dummyIndexBuf = []byte(DummyIndex)
var dummyLinkBuf = []byte(DummyLink)
var dummyLinkTitle = []byte(DummyTitle)

func main() {
	body, err := ioutil.ReadFile(AllJsonFile)
	if err == os.ErrNotExist {
		resp, err := http.Get(ApiProblemsAll)
		check(err)
		defer func() {
			err = resp.Body.Close()
			check(err)
			println(ALlJsonUpdatedText)
		}()

		body, err = ioutil.ReadAll(resp.Body)
	}
	check(err)

	check(ioutil.WriteFile(AllJsonFile, body, os.ModePerm))

	check(json.Unmarshal(body, &profile))

	problemID := flag.String("p", "", "problem id")
	problemTyp := flag.String("t", "", "problem type")
	flag.Parse()
	if *problemID == "" {
		buildReadme()
		return
	}
	p := find(*problemID)
	v, ok := problemStubs[*problemTyp]
	if !ok {
		v = problemStubs[DefaultProblemStub]
	}

	problemStub, err := ioutil.ReadFile(StubPrefix + v.Stub)
	check(err)

	i, err := strconv.Atoi(p.Stat.FrontendQuestionId)
	check(err)

	path := fmt.Sprintf("%04d.%s", i, p.Stat.QuestionTitleSlug)

	problemStub = bytes.Replace(problemStub, dummyLinkTitle, []byte(p.Stat.QuestionTitle), 1)
	problemStub = bytes.Replace(problemStub, dummyLinkBuf, []byte(fmt.Sprintf("https://leetcode-cn.com/problems/%s", p.Stat.QuestionTitleSlug)), 1)

	check(os.Mkdir(path, os.ModePerm))

	check(ioutil.WriteFile(fmt.Sprintf("%s/%s", path, RealPrefix+v.Real), problemStub, os.ModePerm))
}

func getSolutions(dir os.FileInfo, path string) {
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
	if l == 2 {
		return
	}
	algorithm := strings.Replace(strings.Join(desc[2:l-1], "."), "_", " ", -1)
	title := strings.Split(desc[1], ".")
	problem := find(title[0])
	problem.File = "https://github.com/bygo/leetcode/tree/master/" + strings.Replace(path, " ", "+", -1)
	problem.Algorithm = algorithm
	problems[currentClassName] = append(problems[currentClassName], problem)
}

func find(id string) Problem {
	left, right := 0, len(profile.StatStatusPairs)
	for left < right {
		if profile.StatStatusPairs[left].Stat.FrontendQuestionId == strings.TrimLeft(id, "0") {
			return profile.StatStatusPairs[left]
		}
		left++
	}
	panic(id + " not exists")
}

func buildReadme() {
	class, err := ioutil.ReadDir("./")
	check(err)
	for _, c := range class {
		if c.IsDir() && c.Name()[0] != '.' && valueOf(solutionOrder, c.Name()) {
			currentClassName = c.Name()
			problems[currentClassName] = []Problem{}
			getSolutions(c, currentClassName)
		}
	}

	stubReadme, err := ioutil.ReadFile(StubPrefix + ReadmeMdStub)
	difficulty := []string{
		Easy, Medium, Hard,
	}
	check(err)

	var readmeBuf []byte
	var directoryIndex []byte

	for _, index := range solutionOrder {
		className := normalizeClassTitle(index)
		directoryIndex = append(directoryIndex, fmt.Sprintf("- [%s](#%s)\n\r", className, className)...)
		problems := problems[index]
		stubClass, err := ioutil.ReadFile(StubPrefix + ReadmeClassStub)
		check(err)
		class := bytes.Replace(stubClass, dummyClassBuf, []byte(className), 1)

		for k, problem := range problems {
			var questionId, questionTitle, questionDifficulty, questionAcceptance, questionTitleSlug string
			if k == 0 || problem.Stat.QuestionId != problems[k-1].Stat.QuestionId {
				questionId = fmt.Sprintf("%04s", problem.Stat.FrontendQuestionId)
				questionTitle = problem.Stat.QuestionTitle
				questionTitleSlug = problem.Stat.QuestionTitleSlug
				questionAcceptance = fmt.Sprintf("%.1f%s", float64(problem.Stat.TotalAcs)*100/float64(problem.Stat.TotalSubmitted), "%")
				questionDifficulty = difficulty[problem.Difficulty.Level-1]
			}

			class = append(class, fmt.Sprintf("\n| %s | [%-32s](https://leetcode-cn.com/problems/%s) | %s | %s | [Go](%s)| %s",
				questionId,
				questionTitle,
				questionTitleSlug,
				questionAcceptance,
				questionDifficulty,
				problem.File,
				problem.Algorithm,
			)...)
		}
		readmeBuf = append(readmeBuf, class...)
	}

	readme := bytes.Replace(stubReadme, dummyTableBuf, readmeBuf, 1)
	readme = bytes.Replace(readme, dummyIndexBuf, directoryIndex, 1)
	check(ioutil.WriteFile(ReadmeName, readme, os.ModePerm))
}

type Stub struct {
	Stub string
	Real string
}

type Problem struct {
	Stat       Stat       `json:"stat"`
	Difficulty Difficulty `json:"difficulty"`
	File       string
	Algorithm  string
}

type Difficulty struct {
	Level int `json:"level"`
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
	IsNewQuestion       bool   `json:"is_new_question"`
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

func normalizeClassTitle(title string) string {
	words := strings.Split(title, "_")
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, "")
}
