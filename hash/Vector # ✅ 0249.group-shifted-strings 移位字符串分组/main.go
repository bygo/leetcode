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
			b = append(b, (char[j]-char[0]+26)%26)
			j++
		}
		id := getId(string(b))
		res[id] = append(res[id], char)
	}
	return res
}
