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
