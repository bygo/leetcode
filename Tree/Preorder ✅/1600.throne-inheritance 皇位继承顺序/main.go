package main

// https://leetcode-cn.com/problems/throne-inheritance

// ❓ 皇位继承顺序
// ⚠️ 多叉树前序遍历

type ThroneInheritance struct {
	king          string
	ParentMpChild map[string][]string
	dead          map[string]struct{}
}

func Constructor(kingName string) (t ThroneInheritance) {
	return ThroneInheritance{
		kingName,
		map[string][]string{},
		map[string]struct{}{},
	}
}

func (t *ThroneInheritance) Birth(parent, child string) {
	t.ParentMpChild[parent] = append(t.ParentMpChild[parent], child)
}

func (t *ThroneInheritance) Death(name string) {
	t.dead[name] = struct{}{}
}

func (t *ThroneInheritance) GetInheritanceOrder() []string {
	var strsName []string
	var dfs func(string)
	dfs = func(parent string) {
		_, ok := t.dead[parent]
		if !ok {
			strsName = append(strsName, parent)
		}
		for _, childName := range t.ParentMpChild[parent] {
			dfs(childName)
		}
	}
	dfs(t.king)
	return strsName
}
