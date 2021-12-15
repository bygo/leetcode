# 0098.validate-binary-search-tree 验证二叉搜索树 
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var queue = []*TreeNode{root}
	var minQ = []int{-1 << 63}
	var maxQ = []int{1<<63 - 1}
	for 0 < len(queue) {
		var cnt = len(queue)
		for i, root := range queue[:cnt] {
			min, max := minQ[i], maxQ[i]
			if root.Val <= min || max <= root.Val {
				return false
			}
			if root.Left != nil {
				minQ = append(minQ, min)
				maxQ = append(maxQ, root.Val)
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				minQ = append(minQ, root.Val)
				maxQ = append(maxQ, max)
				queue = append(queue, root.Right)
			}
		}
		queue, minQ, maxQ = queue[cnt:], minQ[cnt:], maxQ[cnt:]
	}

	return true
}

```

# 0102.binary-tree-level-order-traversal 二叉树的层序遍历 
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var depth int
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		res = append(res, []int{})
		for _, q := range queue[:cnt] {
			res[depth] = append(res[depth], q.Val)
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[cnt:]
		depth++
	}

	return res
}

```

# 0103.binary-tree-zigzag-level-order-traversal 二叉树的锯齿形层序遍历 
```go
package bfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var queue = []*TreeNode{root}
	var depth, cnt int
	for 0 < len(queue) {
		cnt = len(queue)
		res = append(res, []int{})
		for _, q := range queue[:cnt] {
			res[depth] = append(res[depth], q.Val)
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		if depth%2 == 1 {
			var i = 0
			var j = len(res[depth]) - 1
			for i < j {
				res[depth][i], res[depth][j] = res[depth][j], res[depth][i]
				i++
				j--
			}
		}
		depth++
		queue = queue[cnt:]
	}

	return res
}

```

# 0107.binary-tree-level-order-traversal-ii 二叉树的层序遍历 II 
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue = []*TreeNode{root}
	var depth int
	for 0 < len(queue) {
		var cnt = len(queue)
		res = append(res, []int{})
		for _, q := range queue[:cnt] {
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, q.Right)
			}
			res[depth] = append(res[depth], q.Val)
		}
		queue = queue[cnt:]
		depth++
	}

	var i = 0
	var j = len(res) - 1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}

```

# 0111.minimum-depth-of-binary-tree 二叉树的最小深度 
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

func minDepth(root *TreeNode) int {
	var depth int
	if root == nil {
		return depth
	}
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		depth++
		var cnt = len(queue)
		for _, q := range queue[:cnt] {
			if q.Left == nil && q.Right == nil {
				return depth
			}
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[cnt:]
	}
	return depth
}

```

# 0116.populating-next-right-pointers-in-each-node 填充每个节点的下一个右侧节点指针 
```go
package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var queue = []*Node{root}
	var head = root
	var pre *Node
	for 0 < len(queue) {
		var cnt = len(queue)
		pre = nil
		for _, v := range queue[:cnt] {
			if pre != nil {
				pre.Next = v
			}
			pre = v
			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[cnt:]
	}

	return head
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var parent = root
	for parent.Left != nil {
		var p = parent
		for p != nil {
			p.Left.Next = p.Right
			if p.Next != nil {
				p.Right.Next = p.Next.Left
			}
			p = p.Next
		}
		parent = parent.Left
	}

	return root
}

```

# 0117.populating-next-right-pointers-in-each-node-ii 填充每个节点的下一个右侧节点指针 II 
```go
package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// populating next right pointers in each node ii
// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var queue = []*Node{root}
	var head = root
	var last *Node
	for 0 < len(queue) {
		var cnt = len(queue)
		last = nil
		for _, v := range queue[:cnt] {
			if last != nil {
				last.Next = v
			}
			last = v
			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[cnt:]
	}
	return head
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var parent *Node
	parent = root

	for parent != nil {
		var pre, down *Node
		down = nil
		for parent != nil {
			if parent.Left != nil {
				if pre == nil {
					down = parent.Left
				} else {
					pre.Next = parent.Left
				}
				pre = parent.Left
			}

			if parent.Right != nil {
				if pre == nil {
					down = parent.Right
				} else {
					pre.Next = parent.Right
				}
				pre = parent.Right
			}
			parent = parent.Next
		}
		parent = down
	}
	return root
}

```

# 0199.binary-tree-right-side-view 二叉树的右视图 
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// binary tree right side view
// https://leetcode-cn.com/problems/binary-tree-right-side-view/

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		for i, q := range queue[:cnt] {
			if i == cnt-1 {
				res = append(res, q.Val)
			}
			if q.Left != nil {
				queue = append(queue, q.Left)
			}

			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[cnt:]
	}
	return res
}

```

# 0297.serialize-and-deserialize-binary-tree 二叉树的序列化与反序列化 
```go
package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// serialize and deserialize binary tree
// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var res []string
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		for _, q := range queue[:cnt] {
			if q != nil {
				res = append(res, strconv.Itoa(q.Val))
				queue = append(queue, q.Left, q.Right)
			} else {
				res = append(res, "#")
			}
		}
		queue = queue[cnt:]
	}

	return strings.Join(res, ",")
}

func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	var res = strings.Split(data, ",")
	var root = &TreeNode{}
	root.Val, _ = strconv.Atoi(res[0])
	var queue = []*TreeNode{root}
	res = res[1:]
	for 0 < len(queue) {
		var cnt = len(queue)
		for _, q := range queue[:cnt] {
			if res[0] != "#" {
				l, _ := strconv.Atoi(res[0])
				q.Left = &TreeNode{Val: l}
				queue = append(queue, q.Left)
			}

			if res[1] != "#" {
				r, _ := strconv.Atoi(res[1])
				q.Right = &TreeNode{Val: r}
				queue = append(queue, q.Right)
			}
			res = res[2:]
		}
		queue = queue[cnt:]
	}
	return root
}

```

# 0428.serialize-and-deserialize-n-ary-tree 序列化和反序列化 N 叉树 
```go
package main

import (
	"strconv"
	"strings"
)

type Node struct {
	Val      int
	Children []*Node
}

// serialize and Deserialize N-ary Tree
// https://leetcode-cn.com/problems/serialize-and-deserialize-n-ary-tree/

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (c *Codec) serialize(root *Node) string {
	if root == nil {
		return ""
	}

	var queue = []*Node{root}
	var res = []string{strconv.Itoa(root.Val)}
	for 0 < len(queue) {
		var cnt = len(queue)
		res = append(res, "d")
		for i, p := range queue[:cnt] {
			if i != 0 {
				res = append(res, "#")
			}
			for j, c := range p.Children {
				if j != 0 {
					res = append(res, ",")
				}
				queue = append(queue, c)
				res = append(res, strconv.Itoa(c.Val))
			}
		}
		queue = queue[cnt:]
	}

	return strings.Join(res, "")
}

func (c *Codec) deserialize(data string) *Node {
	if data == "" {
		return nil
	}
	var pre, parent = &Node{}, &Node{}
	var queue = []*Node{pre}
	for _, d := range strings.Split(data, "d") {
		var cnt = len(queue)
		for k, n := range strings.Split(d, "#") {
			parent = queue[k]
			for _, q := range strings.Split(n, ",") {
				if q != "" {
					v, _ := strconv.Atoi(q)
					c := &Node{Val: v}
					parent.Children = append(parent.Children, c)
					queue = append(queue, c)
				}
			}
		}
		queue = queue[cnt:]
	}

	return pre.Children[0]
}

```

# 0429.n-ary-tree-level-order-traversal N 叉树的层序遍历 
```go
package main

type Node struct {
	Val      int
	Children []*Node
}

// n-ary-tree level order traversal
// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var depth int
	var queue = []*Node{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		res = append(res, []int{})
		for _, v := range queue[:cnt] {
			res[depth] = append(res[depth], v.Val)
			for _, c := range v.Children {
				queue = append(queue, c)
			}
		}
		depth++
		queue = queue[cnt:]
	}
	return res
}

```

# 0433.minimum-genetic-mutation 最小基因变化 
```go
package main

// https://leetcode-cn.com/problems/minimum-genetic-mutation

func minMutation(start string, end string, bank []string) int {
	m := make(map[string]struct{})
	for _, b := range bank {
		m[b] = struct{}{}
	}
	if start == end {
		return 0
	}

	if _, ok := m[end]; !ok {
		return -1
	}
	queue := []string{start}
	res := 1
	base := []byte{'A', 'C', 'G', 'T'}
	for {
		cnt := len(queue)
		if cnt == 0 {
			break
		}
		for i := 0; i < cnt; i++ {
			b := []byte(queue[i])
			for j := 0; j < 8; j++ {
				for k := 0; k < 4; k++ {
					if b[j] != base[k] {
						b[j], base[k] = base[k], b[j]
						s := string(b)
						if s == end {
							return res
						}
						if _, ok := m[s]; ok {
							queue = append(queue, s)
						}
						b[j], base[k] = base[k], b[j]
					}
				}
			}
		}
		res++
		queue = queue[cnt:]
	}
	return -1
}

```

# 0513.find-bottom-left-tree-value 找树左下角的值 
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// find bottom left tree value
// https://leetcode-cn.com/problems/find-bottom-left-tree-value/

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int

	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		res = queue[0].Val
		for _, v := range queue[:cnt] {
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		queue = queue[cnt:]
	}
	return res
}

```

# 0515.find-largest-value-in-each-tree-row 在每个树行中找最大值 
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// find largest value in each tree row
// https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var queue = []*TreeNode{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		var max = -1 << 63
		for _, v := range queue[:cnt] {
			if v.Left != nil {
				queue = append(queue, v.Left)
			}

			if v.Right != nil {
				queue = append(queue, v.Right)
			}
			if max < v.Val {
				max = v.Val
			}
		}
		res = append(res, max)
		queue = queue[cnt:]
	}

	return res
}

```

# 0559.maximum-depth-of-n-ary-tree N 叉树的最大深度 
```go
package main

type Node struct {
	Val      int
	Children []*Node
}

// maximum depth of n-ary-tree
// https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/submissions/

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	var depth int
	var queue = []*Node{root}
	for 0 < len(queue) {
		var cnt = len(queue)
		for _, q := range queue[:cnt] {
			for _, c := range q.Children {
				queue = append(queue, c)
			}
		}
		queue = queue[cnt:]
		depth++
	}
	return depth
}

```

# 0623.add-one-row-to-tree 在二叉树中增加一行 
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// add one row to tree
// https://leetcode-cn.com/problems/add-one-row-to-tree/

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}

	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}

	var queue = []*TreeNode{root}
	var depth = 1

	for 0 < len(queue) {
		var cnt = len(queue)
		depth++
		if depth == d {
			for _, q := range queue {
				nl := &TreeNode{Val: v, Left: q.Left}
				q.Left = nl
				nr := &TreeNode{Val: v, Right: q.Right}
				q.Right = nr
			}
			break
		}

		for _, q := range queue[:cnt] {
			if q.Left != nil {
				queue = append(queue, q.Left)
			}

			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		queue = queue[cnt:]
	}

	return root
}

```

# 0637.average-of-levels-in-binary-tree 二叉树的层平均值 
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// average of levels in binary tree
// https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var res []float64
	for 0 < len(queue) {
		var sum int
		var cnt = len(queue)
		for _, q := range queue[:cnt] {
			sum += q.Val
			if q.Left != nil {
				queue = append(queue, q.Left)
			}
			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		res = append(res, float64(sum)/float64(cnt))
		queue = queue[cnt:]
	}
	return res
}

```

# 1161.maximum-level-sum-of-a-binary-tree 最大层内元素和 
```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maximum level sum of a binary tree
// https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue = []*TreeNode{root}
	var res, depth, sum int
	var max = -1 << 63
	for 0 < len(queue) {
		var cnt = len(queue)
		depth++
		for _, q := range queue[:cnt] {
			sum += q.Val
			if q.Left != nil {
				queue = append(queue, q.Left)
			}

			if q.Right != nil {
				queue = append(queue, q.Right)
			}
		}
		if max < sum {
			max = sum
			res = depth
		}
		sum = 0
		queue = queue[cnt:]
	}

	return res
}

```

# 1306.jump-game-iii 跳跃游戏 III 
```go
package main

// can i jump to point which the value is 0 ?
// https://leetcode-cn.com/problems/jump-game-iii

func canReach(arr []int, start int) bool {
	var l1 = len(arr) - 1
	var queue = []int{start}
	var visited = map[int]bool{}
	for 0 < len(queue) {
		var cnt = len(queue)
		for _, idx := range queue[:cnt] {
			if arr[idx] == 0 {
				return true
			}
			var l = idx - arr[idx]
			var r = idx + arr[idx]
			if 0 <= l && !visited[l] {
				queue = append(queue, l)
				visited[l] = true
			}

			if r <= l1 && !visited[r] {
				queue = append(queue, r)
				visited[r] = true
			}
		}
		queue = queue[cnt:]
	}

	return false
}

```

