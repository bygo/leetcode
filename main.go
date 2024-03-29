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

type Stub struct {
	Stub string
	Real string
}

type Problem struct {
	Stat      Stat `json:"stat"`
	File      string
	Algorithm string

	Difficulty `json:"difficulty"`
}

type Difficulty struct {
	Level `json:"level"`
}

type Level int

const (
	LevelEasy Level = iota
	LevelMedium
	LevelHard
	LevelUnknown
)

var difficulty = map[Level]string{
	LevelEasy:    "Easy",
	LevelMedium:  "Medium",
	LevelHard:    "Hard",
	LevelUnknown: "Unknown",
}

type Stat struct {
	QuestionId          int    `json:"question_id"`
	QuestionTitle       string `json:"question__title"`
	QuestionTitleSlug   string `json:"question__title_slug"`
	QuestionHide        bool   `json:"question__hide"`
	QuestionCN          string
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

const (
	Repository = "https://github.com/bygo/leetcode"

	ApiProblemsAll     = "https://leetcode.cn/api/problems/all/"
	AllJsonFile        = "all.json"
	AllJsonUpdatedText = "update `all.json`"

	StubPrefix = "Stub/"
	RealPrefix = ""

	ReadmeClassStub = "class.stub"
	ReadmeMdStub    = "readmeStub.md"
	ReadmeName      = "readme.md"

	DefaultProblemStub = "default"
)

var profile Profile
var currentClassName string

var problems = map[string][]Problem{}
var solutionOrder = []string{
	"two_pointer",
	"bit",
	"bfs",
	"stack",
	"enum",
	"tree",
	"divide&conquer",
	"back",
	"sort",
	"search",
	"linked_list",
	"binary_search",
	"catalan",
	"stack_monotonic",
	"sort_counter",
	"hash",
	"greedy",
	"dp",
	"classic",
	"math",
	"sql",
}

var problemStubs = map[string]*Stub{
	"default": {"problem.stub", "main.go"},
	"tree":    {"tree.stub", "main.go"},
	"link":    {"link.stub", "main.go"},
	"sql":     {"sql.stub", "1.sql"},
}

var dummyBufClass = []byte("@DummyClass")
var dummyBufTable = []byte("@DummyTable")
var dummyBufIndex = []byte("@DummyIndex")
var dummyBufLink = []byte("@DummyLink")
var dummyBufLinkTitle = []byte("@DummyTitle")
var dummyBufHeadline = []byte("@DummyHeadline")
var dummyBufAC = []byte("1059")

var ignorePrefix = []string{
	//"LCP", "Offer"
}

type NodeReadme struct {
	Dir      string
	Language string
}

var NodeReadmeArr = []NodeReadme{
	{"dp", "go"},
	{"sql", "sql"},
	{"bfs", "go"},
	{"hash", "go"},
}

func main() {
	problemID := flag.String("p", "", "problem id")
	problemTyp := flag.String("t", "", "problem type")
	cmd := flag.String("c", "", "other cmd")

	flag.Parse()

	if *cmd == "n" {
		for _, n := range NodeReadmeArr {
			genNodeReadme(n.Language, n.Dir)
		}
	}

	body, err := ioutil.ReadFile(AllJsonFile)
	if err != nil || *problemID == "" {
		resp, err := http.Get(ApiProblemsAll)
		Check(err)
		defer func() {
			err = resp.Body.Close()
			Check(err)
			println(AllJsonUpdatedText)
		}()

		body, err = ioutil.ReadAll(resp.Body)
		Check(err)
	}

	Check(ioutil.WriteFile(AllJsonFile, body, os.ModePerm))

	Check(json.Unmarshal(body, &profile))

	if *problemID == "" {
		buildReadme()
		return
	}
	p := getProblem(*problemID)
	v, ok := problemStubs[*problemTyp]
	if !ok {
		v = problemStubs[DefaultProblemStub]
	}

	problemStub, err := ioutil.ReadFile(StubPrefix + v.Stub)
	Check(err)

	i, err := strconv.Atoi(p.Stat.FrontendQuestionId)
	Check(err)

	path := fmt.Sprintf("%04d.%s", i, p.Stat.QuestionTitleSlug)

	problemStub = bytes.Replace(problemStub, dummyBufLinkTitle, []byte(p.Stat.QuestionTitle), 1)
	problemStub = bytes.Replace(problemStub, dummyBufLink, []byte(fmt.Sprintf("https://leetcode.cn/problems/%s", p.Stat.QuestionTitleSlug)), 1)

	//Check(os.Mkdir(path, os.ModePerm))

	Check(ioutil.WriteFile(fmt.Sprintf("%s", path+v.Real), problemStub, os.ModePerm))
}

func getSolutions(dir os.FileInfo, path string) {
	if dir.IsDir() {
		dirs, err := ioutil.ReadDir(path)
		Check(err)
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
	//algorithm := strings.Replace(strings.Join(desc[2:l-1], "."), "_", " ", -1)
	//title := strings.Split(desc[1], ".")
	title := desc[1]
	for _, prefix := range ignorePrefix {
		if strings.HasPrefix(title, prefix) {
			return
		}
	}
	problem := getProblem(title)
	if problem.Stat.QuestionId == 0 {
		problem.Stat.QuestionId = 9999
		problem.Stat.TotalAcs = 0
		problem.Stat.TotalSubmitted = 0
		problem.Difficulty.Level = LevelUnknown
		problem.Stat.QuestionTitle = title
		problem.Stat.QuestionTitleSlug = " "
	}

	problem.File = Repository + "/blob/master/" + strings.Replace(path, " ", "%20", -1)
	//problem.Algorithm = algorithm
	problems[currentClassName] = append(problems[currentClassName], problem)
}

func getProblem(id string) Problem {
	left, right := 0, len(profile.StatStatusPairs)
	var algorithm, code, title, cn string
	index := 0
	if id[index] == '(' {
		for id[index] != ')' {
			index++
		}
		algorithm = id[1:index]
		index += 2
	}
	code = parseInt(id[index:])

	titleLeft := index
	for index < len(id) && id[index] <= 127 {
		index++
	}
	title = id[titleLeft:index]
	cn = id[index:]

	for left < right {
		if profile.StatStatusPairs[left].Stat.FrontendQuestionId == code {
			profile.StatStatusPairs[left].Stat.QuestionTitle = title
			profile.StatStatusPairs[left].Algorithm = algorithm
			profile.StatStatusPairs[left].Stat.QuestionCN = cn

			return profile.StatStatusPairs[left]
		}
		left++
	}
	return Problem{}
}

func parseInt(id string) string {
	var sum []byte
	for i := range id {
		if '0' <= id[i] && id[i] <= '9' {
			if 0 != len(sum) || id[i] != '0' {
				sum = append(sum, id[i])
			}
		} else if 0 < len(sum) {
			return string(sum)
		}
	}
	return string(sum)
}

func buildReadme() {
	class, err := ioutil.ReadDir("./")
	Check(err)
	for _, c := range class {
		if c.IsDir() && c.Name()[0] != '.' && ValueOf(solutionOrder, c.Name()) {
			currentClassName = c.Name()
			problems[currentClassName] = []Problem{}
			getSolutions(c, currentClassName)
		}
	}

	stubReadme, err := ioutil.ReadFile(StubPrefix + ReadmeMdStub)
	Check(err)

	var readmeBuf []byte
	var directoryIndex []byte

	_ = func() {
		for _, index := range solutionOrder {
			className := NormalizeClassTitle(index)
			directoryIndex = append(directoryIndex, fmt.Sprintf("- [%s](#%s)\n\r", className, className)...)
			problems := problems[index]
			stubClass, err := ioutil.ReadFile(StubPrefix + ReadmeClassStub)
			Check(err)
			class := bytes.Replace(stubClass, dummyBufClass, []byte(className), 1)

			for k, problem := range problems {
				var questionAlgorithm, questionTitle, questionDifficulty, questionTitleSlug, questionCN string
				questionAlgorithm = " "
				if k == 0 || problem.Stat.QuestionId != problems[k-1].Stat.QuestionId {
					//questionId = fmt.Sprintf("%04s", problem.Stat.FrontendQuestionId)
					questionAlgorithm = problem.Algorithm
					questionTitle = problem.Stat.QuestionTitle
					questionCN = problem.Stat.QuestionCN
					questionTitleSlug = problem.Stat.QuestionTitleSlug
					//questionAcceptance = fmt.Sprintf("%.1f%s", float64(problem.Stat.TotalAcs)*100/float64(problem.Stat.TotalSubmitted), "%")
					questionDifficulty = difficulty[problem.Difficulty.Level-1]
				}

				class = append(class, fmt.Sprintf("\n [%-32s](https://leetcode.cn/problems/%s) | %s | %s | [Go](%s) | %s",
					questionTitle,
					questionTitleSlug,
					questionCN,
					questionDifficulty,
					problem.File,
					questionAlgorithm,
				)...)
			}
			readmeBuf = append(readmeBuf, class...)
		}
	}

	const prefix = "https://github.com/bygo/leetcode/tree/master/"
	for _, index := range solutionOrder {
		className := NormalizeClassTitle(index)
		directoryIndex = append(directoryIndex, fmt.Sprintf("- [%s](%s%s)\n\r", className, prefix, className)...)
		Check(err)
	}

	readme := bytes.Replace(stubReadme, dummyBufHeadline, dummyBufAC, 1)
	//readme = bytes.Replace(readme, dummyBufTable, readmeBuf, 1)
	readme = bytes.Replace(readme, dummyBufIndex, directoryIndex, 1)
	Check(ioutil.WriteFile(ReadmeName, readme, os.ModePerm))
}

func genNodeReadme(language string, dir string) {
	var md bytes.Buffer

	all, err := ioutil.ReadDir(dir)
	Check(err)

	for _, item := range all {
		sqlDir, _ := ioutil.ReadDir(fmt.Sprintf("./%s/%s", dir, item.Name()))
		var b = bytes.NewBufferString(fmt.Sprintf("# %s \n", item.Name()))
		for _, sqlFile := range sqlDir {
			b.WriteString("```" + language + "\n")
			s, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", dir, item.Name(), sqlFile.Name()))
			b.Write(s)
			b.WriteString("\n```\n\n")
			md.WriteString(b.String())
		}
	}

	err = ioutil.WriteFile(dir+"/readme.md", md.Bytes(), 0755)
	Check(err)
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ValueOf(str []string, value string) bool {
	for _, v := range str {
		if v == value {
			return true
		}
	}
	return false
}

func NormalizeClassTitle(title string) string {
	words := strings.Split(title, "_")
	for i, w := range words {
		words[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(words, "")
}
