package main

// https://leetcode.cn/problems/employee-importance

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	users := map[int]*Employee{}
	for i := range employees {
		users[employees[i].Id] = employees[i]
	}

	cur := users[id]
	cnt := cur.Importance

	var queue = []*Employee{users[id]}
	for 0 < len(queue) {
		cur = queue[0]
		for i := range cur.Subordinates {
			cnt += users[cur.Subordinates[i]].Importance
			queue = append(queue, users[cur.Subordinates[i]])
		}
		queue = queue[1:]
	}

	return cnt
}
