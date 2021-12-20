package main

// https://leetcode-cn.com/problems/throne-inheritance

type ThroneInheritance struct {
	king  string
	child map[string][]string
	dead  map[string]bool
}

func Constructor(kingName string) (t ThroneInheritance) {
	return ThroneInheritance{
		kingName,
		map[string][]string{},
		map[string]bool{}}
}

func (t *ThroneInheritance) Birth(p, c string) {
	t.child[p] = append(t.child[p], c)
}

func (t *ThroneInheritance) Death(name string) {
	t.dead[name] = true
}

func (t *ThroneInheritance) GetInheritanceOrder() []string {
	var res []string
	var preorder func(string)
	preorder = func(name string) {
		if !t.dead[name] {
			res = append(res, name)
		}
		for _, childName := range t.child[name] {
			preorder(childName)
		}
	}
	preorder(t.king)
	return res
}
