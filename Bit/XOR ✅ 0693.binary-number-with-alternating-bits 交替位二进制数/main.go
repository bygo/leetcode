package main

import (
	"sort"
)

// https://leetcode-cn.com/problems/binary-number-with-alternating-bits

func hasAlternatingBits(n int) bool {
	// 全部变1
	n = n ^ (n >> 1) + 1
	// 全部进位成功 只剩1个1
	return n&(n-1) == 0
}

func minimumCost(cost []int) int {
	mp := [101]int{}
	for i := range cost {
		mp[cost[i]] ++
	}

	var cnt int
	var tot int
	for i := 100; 1 <= i; i-- {
		for j := 0; j < mp[i]; j++ {
			cnt++
			if cnt%3 == 0 {
				continue
			}
			tot += i
		}
	}
	return tot
}

func numberOfArrays(differences []int, lower int, upper int) int {
	var max = -1 << 63
	var min = 1<<63 - 1
	var sum int
	for _, difference := range differences {
		sum += difference
		if sum < min {
			min = sum
		}
		if max < sum { // -46
			max = sum
		}
	}

	if min < 0 && 0 < max {
		max = max - min
	} else if min < 0 {
		max = 0 - min
	} else if 0 < max {

	}
	var diff = upper - lower + 1
	res := diff - max
	if res < 0 {
		return 0
	}
	return res
}

func main() {
	numberOfWays("SSSSPPPPPSSPSPSSPPPPPPSPSPSSPPPSPSPSPSSPPPSSSSPSSPSPSPSSSSPPSSPSSSSSSPPPSPSSPPSSPSPSPSSPSSSSSSPSSSPSSSPSSPSSPPPSSPSSPPSSSPSSPSPSSSSPSPPSSPSSPPSSPPSPSPPPSPSPPSPPSPSPPPPPPPPSPSSSPPSPSSPSSSSPPSPSPPPPSPPSSPSPSSPSSPSPPPPPPSSPSPPPPSSPSPPPPPSSSSPSSSSSPPSPPPSPSSPPSPSPSSSSPPPSPPSSSPSPPSSSSSSSSPSSSSPSPSSSPSPPPPPSPSPSPPPSSSPSSPSPSSSPPPPSSPSPSSSSPSSSSSPPSPSSSSPSSSSSSPPPSSPPPPPPSPSSSPPSPPSSSPPPSSPPPSPPSSPPSPSSSSSSPPSSSPPSPSPSPPSPSSPPSPSSPSPPSPSPPSSPSSPPSPPSPSPPSPPSSSSPSPPPSSPPSSPPPSPPSSPSPPSPSPSSPSPPSSSPPSSSPSSPSPPSPSPSPSPSPSSPPPSPSSSSSPPSPSSSPSSSPSSSSSSPSPPPPPPPSSPPSPPSPPSPSSSSPSSSPPPPSPSSPSSPSSPPSSSSPPSPSPSPPPSSSPSPSSPPSSSPPSSPSPPPPPSSSPPSSPSPPPSSSPPPPPSPPPPPSPSPSSSPPSPSSPPSPPSSSPSSSPPSSPPSPSSSPSPPSSSSPSPPSSSPSSPSPPPSPPPSPSSPSPPSSSPPSPSSPSPPPPSSSPSSPPSPSPSPSPSSSPSPSPPPPSSSSSSPPPPPPSPPPSPSPPSPSPPSPPPPPSSSPSPPPSPPSPSPPSSSPPSPSSPPPPPSSPPPPPSSPSSPPSPSPSSSPPPSSSSPPPSPPPPSSSPSPPPSPPPSPPSPSSSPPPPPSSSSPPPPSSSSPSPSPSPPSPSPPSPSPSPSSPPPSSPPPPPSSPSPPPPPPSSPSSPSPSSPPSPPSSSPPSPSSSSSSSSSSSPSSPSPSPPPSPSPPPPSPSSSPPPSSPSPSSSSPSPPPSPSPPSPSPPPSPSPPSSSPSPPSSSPPPPPPSPPSSPPPPSPPPPPSPPPPPSPSPSSPPPSPSSSSSPSSSSPPPSSSPPSPPPPPSSPSSSPSPSPSSSPPSSPSPSPPSSPSPSPSPSPPSSSSPPPSSPPPSSSPPSPSPPPPSPSPSPPPSSPSSPSSPPPPSSSPPPPSSSPPPPPPSPSSSPSSPSPSSSSPSPSPSPPSPSSSPPSSSSSPSSPPPPSSSSPPPSPSSPPSPSSPPPPPPPSSPSPPPSSPPPPSSSSSSPSPPPSPPSSSPSPSSSPSSPSSSSSPSSPSPSSPSSPPSSSSPPSPSSPPSSSPSPPSSSPSSPSSSPSPSPPSSPPPPSSPSPSPPSPPPPPSPSSSPPSPPPPSSPPSPSPSSPSPSPSSPPPPPPSSSPPPPSPSPPSSPSPSSSPSSPPPPSPSSSSPPSPPSSSPPSPPPSSPSPSPPPSPSSSSSPPSPSPPPSSPPPSSSPPSSSPPSPSSPPSSSPPSPPPPPPSSSPSSPSSPSPPPSPSPPPSPSSPSSSSSSSSPPPSSSSPPSSSPSSSPSSPPPSSPSPPPPPPPPSSPPSPPPSPPPSSPSSPPPPPSSSSPPSSPPPPSS")
}

type pair struct {
	price int
	row   int
	col   int
}

var dirs = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func highestRankedKItems(grid [][]int, pricing []int, start []int, k int) [][]int {
	var p pair
	p.row = start[0]
	p.col = start[1]
	p.price = grid[p.row][p.col]
	grid[p.row][p.col] = 0

	rowL := len(grid)
	colL := len(grid[0])
	var que = []pair{p}
	var order [][]int
	for {
		cnt := len(que)
		if cnt == 0 {
			break
		}

		sort.Slice(que, func(i, j int) bool {
			return que[i].price < que[j].price ||
				que[i].price == que[j].price && que[i].row < que[j].row ||
				que[i].price == que[j].price && que[i].row == que[j].row && que[i].col < que[j].col
		})

		if k <= len(order) {
			break
		}

		for i := 0; i < cnt; i++ {
			q := que[i]
			if pricing[0] <= q.price && q.price <= pricing[1] {
				order = append(order, []int{q.row, q.col})
			}
			for _, dir := range dirs {
				row := q.row + dir[0]
				col := q.col + dir[1]
				if 0 <= row && row < rowL && 0 <= col && col < colL && grid[row][col] != 0 {
					que = append(que, pair{
						price: grid[row][col],
						row:   row,
						col:   col,
					})
					grid[row][col] = 0
				}
			}
		}
		que = que[cnt:]
	}
	orderL := len(order)
	if orderL < k {
		return order
	}
	return order[:k]
}

func numberOfWays(s string) int {
	var sL = len(s)
	var cnt = 0
	var cntF int
	var comb []int
	for i := 0; i < sL; i++ {
		ch := s[i]
		if ch == 'S' {
			if 0 < cnt && cnt%2 == 0 {
				comb = append(comb, cntF)
			}
			cntF = 0
			cnt++
		} else {
			cntF++
		}
	}

	println(cnt)
	if 0 < cnt && cnt%2 == 0 {
		comb = append(comb, 0)
	}

	combL := len(comb)
	if combL == 0 || cnt%2 == 1 {
		return 0
	}

	tot := 1
	for _, c := range comb {
		tot = tot * (c + 1) % 1000000007
	}
	return tot % 1000000007
}
