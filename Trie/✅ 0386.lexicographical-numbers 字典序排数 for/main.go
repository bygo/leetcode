package main

// https://leetcode-cn.com/problems/lexicographical-numbers

//    1			 2
//  10  11..  20

// ❓ 字典序排数
// ⚠️ 子节点10 大于 父节点1
// ⚠️ 右节点11 大于 左节点10

func lexicalOrder(n int) []int {
	var numsOrder []int
	var num = 1
	for len(numsOrder) < n {
		// 进入到最低层 (乘以10进入下层)
		for num <= n {
			numsOrder = append(numsOrder, num)
			num *= 10
		}

		// 遍历完毕 or 超出范围 (除以10返回上层)
		for num%10 == 9 || n < num {
			num /= 10
		}
		// 同层遍历 or 返回上层后+1
		num++
	}
	return numsOrder
}
