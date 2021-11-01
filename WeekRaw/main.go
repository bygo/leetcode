package WeekRaw

import "sort"

type Node struct {
	Pndex  int
	Index  int
	Parent *Node
	C      []*Node
	UserId int
	IsLock bool
}
type LockingTree struct {
	m map[int]*Node
}

func Constructor(parent []int) LockingTree {
	m := map[int]*Node{}
	un := []*Node{}
	for i := range parent {
		n := &Node{}
		if i == 0 {
			m[i] = n
		} else {
			p := m[parent[i]] // 父节点
			if p == nil {     // 不存在
				n.Pndex = parent[i]
				n.Index = i
				un = append(un, n)
			} else { //存在
				n.Parent = p
				p.C = append(p.C, n)
				m[i] = n
			}
		}
	}

	for 0 < len(un) {
		cur := un[0]
		un = un[1:]
		v, ok := m[cur.Pndex]
		if ok {
			v.C = append(v.C, cur)
			cur.Parent = v
			m[cur.Index] = cur
		} else {
			un = append(un, cur)
		}
	}
	return LockingTree{m: m}
}

func (this *LockingTree) Lock(num int, user int) bool {
	n := this.m[num]
	if !n.IsLock {
		n.IsLock = true
		n.UserId = user
		return true
	}
	return false
}

func (this *LockingTree) Unlock(num int, user int) bool {
	n := this.m[num]
	if user == n.UserId && n.IsLock {
		n.IsLock = false
		n.UserId = 0
		return true
	}
	return false
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	head := this.m[num]
	n := head

	for n != nil {
		if n.IsLock {
			return false
		}
		n = n.Parent
	}

	var queue = head.C
	var res = false
	for 0 < len(queue) {
		cur := queue[0]
		queue = queue[1:]
		if cur.IsLock {
			res = true
			cur.IsLock = false
			cur.UserId = 0
		}
		if cur.C != nil {
			queue = append(queue, cur.C...)
		}
	}
	if res {
		head.IsLock = true
		head.UserId = user
	}
	return res
}

/**
 * Your LockingTree object will be instantiated and called as such:
 * obj := Constructor(parent);
 * param_1 := obj.Lock(num,user);
 * param_2 := obj.Unlock(num,user);
 * param_3 := obj.Upgrade(num,user);
 */

func findFarmland(land [][]int) [][]int {
	m := map[int]int{}
	l1 := len(land)
	l2 := len(land[0])
	var res [][]int
	for i := range land {
		for j := range land[i] {
			if m[i*300+j] != 1 {
				if land[i][j] == 1 {
					k := i
					l := j
					m[i*300+j] = 1
					for k+1 < l1 && land[k+1][l] == 1 {
						k++
						m[k*300+l] = 1
					}

					for l+1 < l2 && land[k][l+1] == 1 {
						l++
						for ii := i; ii <= k; ii++ {
							m[ii*300+l] = 1
						}
					}

					res = append(res, []int{i, j, k, l})
				}
			}

		}
	}
	return res
}

func countQuadruplets(nums []int) int {
	m := map[int]int{}

	var res int

	for i := range nums {
		res += m[nums[i]]
		for j := range nums[:i] {
			for k := range nums[:j] {
				m[nums[i]+nums[j]+nums[k]]++
			}
		}
	}

	return res
}

type r struct {
	a int
	d int
}

func numberOfWeakCharacters(properties [][]int) int {
	var arr []r
	for i := range properties {
		arr = append(arr, r{
			a: properties[i][0],
			d: properties[i][1],
		})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[j].a < arr[i].a
	})

	var max = -1 << 63
	var res int
	var stack []int
	top := len(arr) - 1
	for i := range arr {
		if arr[i].d < max {
			res++
		}
		stack = append(stack, arr[i].d)

		if i < top && arr[i].a != arr[i+1].a {
			for 0 < len(stack) {
				if max < stack[0] {
					max = stack[0]
				}
				stack = stack[1:]
			}
		}

	}
	return res
}
