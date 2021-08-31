package main

// https://leetcode-cn.com/problems/employee-importance

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	m := map[int]*Employee{}
	for i := range employees {
		m[employees[i].Id] = employees[i]
	}

	cur := m[id]
	cnt := cur.Importance

	var queue = []*Employee{m[id]}
	for 0 < len(queue) {
		cur = queue[0]
		for i := range cur.Subordinates {
			cnt += m[cur.Subordinates[i]].Importance
			queue = append(queue, m[cur.Subordinates[i]])
		}
		queue = queue[1:]
	}

	return cnt
}
