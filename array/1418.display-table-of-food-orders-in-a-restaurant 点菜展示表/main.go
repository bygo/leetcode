package main

import (
	"sort"
	"strconv"
)

// Link: https://leetcode-cn.com/problems/display-table-of-food-orders-in-a-restaurant

func displayTable(orders [][]string) [][]string {
	nameSet := map[string]struct{}{}
	foodsCnt := map[int]map[string]int{}
	for _, order := range orders {
		id, _ := strconv.Atoi(order[1])
		name := order[2]
		nameSet[name] = struct{}{}

		if foodsCnt[id] == nil {
			foodsCnt[id] = map[string]int{}
		}
		foodsCnt[id][name] += 1
	}

	var names []string
	for name := range nameSet {
		names = append(names, name)
	}
	sort.Strings(names)

	var ids []int
	for id := range foodsCnt {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	l1 := len(names)
	l2 := len(ids)
	table := make([][]string, l2+1)
	table[0] = []string{"Table"}
	table[0] = append(table[0], names...)
	for i, id := range ids {
		table[i+1] = make([]string, l1+1)
		table[i+1][0] = strconv.Itoa(id)
		for j, name := range names {
			table[i+1][j+1] = strconv.Itoa(foodsCnt[id][name])
		}
	}
	return table
}
