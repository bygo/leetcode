# LeetCode

[![](https://img.shields.io/badge/Language-Go-%2300ADD8)](https://golang.org/)
[![](https://img.shields.io/badge/AC-842-%23F781BE)](https://leetcode-cn.com/u/bygo/)

---

# Template
- [Tree](https://github.com/bygo/leetcode/tree/master/templates/tree)
- [LinkedList](https://github.com/bygo/leetcode/tree/master/templates/linked_list)
- [Stack](https://github.com/bygo/leetcode/tree/master/templates/stack)
- [Array](https://github.com/bygo/leetcode/tree/master/templates/array)

---
- [TwoPointer](#TwoPointer)
- [Bit](#Bit)
- [Bfs](#Bfs)
- [Stack](#Stack)
- [Enum](#Enum)
- [Tree](#Tree)
- [Divide&conquer](#Divide&conquer)
- [Back](#Back)
- [Sort](#Sort)
- [Search](#Search)
- [LinkedList](#LinkedList)
- [BinarySearch](#BinarySearch)
- [Catalan](#Catalan)
- [StackMonotonic](#StackMonotonic)
- [SortCounter](#SortCounter)
- [Hash](#Hash)
- [Greedy](#Greedy)
- [Dp](#Dp)
- [Classic](#Classic)
- [Math](#Math)
- [Sql](#Sql)

---



## TwoPointer
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0003.longest-substring-without-repeating-characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters) |  无重复字符的最长子串 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0003.longest-substring-without-repeating-characters%20无重复字符的最长子串/main.go) | 
 [0019.remove-nth-node-from-end-of-list](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list) |  删除链表的倒数第 N 个结点 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0019.remove-nth-node-from-end-of-list%20删除链表的倒数第%20N%20个结点/main.go) | 
 [0027.remove-element             ](https://leetcode-cn.com/problems/remove-element) |  移除元素 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0027.remove-element%20移除元素/main.go) | 
 [0030.substring-with-concatenation-of-all-words](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words) |  串联所有单词的子串 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0030.substring-with-concatenation-of-all-words%20串联所有单词的子串/main.go) | 
 [0042.trapping-rain-water        ](https://leetcode-cn.com/problems/trapping-rain-water) |  接雨水 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0042.trapping-rain-water%20接雨水/main.go) | 
 [0075.sort-colors                ](https://leetcode-cn.com/problems/sort-colors) |  颜色分类 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0075.sort-colors%20颜色分类/main.go) | 
 [0076.minimum-window-substring   ](https://leetcode-cn.com/problems/minimum-window-substring) |  最小覆盖子串 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0076.minimum-window-substring%20最小覆盖子串/main.go) | 
 [0080.remove-duplicates-from-sorted-array-ii](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii) |  删除有序数组中的重复项 II | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0080.remove-duplicates-from-sorted-array-ii%20删除有序数组中的重复项%20II/main.go) | 
 [0424.longest-repeating-character-replacement](https://leetcode-cn.com/problems/longest-repeating-character-replacement) |  替换后的最长重复字符 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0424.longest-repeating-character-replacement%20替换后的最长重复字符/main.go) | 
 [0443.string-compression         ](https://leetcode-cn.com/problems/string-compression) |  压缩字符串 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0443.string-compression%20压缩字符串/main.go) | 
 [0524.longest-word-in-dictionary-through-deleting](https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting) |  通过删除字母匹配到字典里最长单词 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0524.longest-word-in-dictionary-through-deleting%20通过删除字母匹配到字典里最长单词/main.go) | 
 [0718.maximum-length-of-repeated-subarray](https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray) |  最长重复子数组 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0718.maximum-length-of-repeated-subarray%20最长重复子数组/main.go) | 
 [0930.binary-subarrays-with-sum  ](https://leetcode-cn.com/problems/binary-subarrays-with-sum) |  和相同的二元子数组 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/0930.binary-subarrays-with-sum%20和相同的二元子数组/main.go) | 
 [1099.two-sum-less-than-k        ](https://leetcode-cn.com/problems/two-sum-less-than-k) |  小于 K 的两数之和 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/1099.two-sum-less-than-k%20小于%20K%20的两数之和/main.go) | 
 [1790.check-if-one-string-swap-can-make-strings-equal](https://leetcode-cn.com/problems/check-if-one-string-swap-can-make-strings-equal) |  仅执行一次字符串交换能否使两个字符串相等 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/two_pointer/1790.check-if-one-string-swap-can-make-strings-equal%20仅执行一次字符串交换能否使两个字符串相等/main.go) | 

## Bit
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0067.add-binary                 ](https://leetcode-cn.com/problems/add-binary) |  二进制求和 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bit/0067.add-binary%20二进制求和/main.go) | 
 [0168.excel-sheet-column-title   ](https://leetcode-cn.com/problems/excel-sheet-column-title) |  Excel表列名称 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bit/0168.excel-sheet-column-title%20Excel表列名称/main.go) | 
 [0371.sum-of-two-integers        ](https://leetcode-cn.com/problems/sum-of-two-integers) |  两整数之和 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bit/0371.sum-of-two-integers%20两整数之和/main.go) | 
 [0401.binary-watch               ](https://leetcode-cn.com/problems/binary-watch) |  二进制手表 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bit/0401.binary-watch%20二进制手表/main.go) | 
 [0421.maximum-xor-of-two-numbers-in-an-array](https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array) |  数组中两个数的最大异或值 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bit/0421.maximum-xor-of-two-numbers-in-an-array%20数组中两个数的最大异或值/main.go) | 
 [0461.hamming-distance           ](https://leetcode-cn.com/problems/hamming-distance) |  汉明距离 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bit/0461.hamming-distance%20汉明距离/main.go) | 
 [0476.number-complement          ](https://leetcode-cn.com/problems/number-complement) |  数字的补数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bit/0476.number-complement%20数字的补数/main.go) | 
 [0477.total-hamming-distance     ](https://leetcode-cn.com/problems/total-hamming-distance) |  汉明距离总和 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bit/0477.total-hamming-distance%20汉明距离总和/main.go) | 
 [1689.partitioning-into-minimum-number-of-deci-binary-numbers](https://leetcode-cn.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers) |  十-二进制数的最少数目 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bit/1689.partitioning-into-minimum-number-of-deci-binary-numbers%20十-二进制数的最少数目/main.go) | 
 [Offer                           ](https://leetcode-cn.com/problems/3sum) |  0015.er-jin-zhi-zhong-1de-ge-shu-lcof 二进制中1的个数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bit/Offer%200015.er-jin-zhi-zhong-1de-ge-shu-lcof%20二进制中1的个数/main.go) | 

## Bfs
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0098.validate-binary-search-tree](https://leetcode-cn.com/problems/validate-binary-search-tree) |  验证二叉搜索树 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0098.validate-binary-search-tree%20验证二叉搜索树/main.go) | 
 [0102.binary-tree-level-order-traversal](https://leetcode-cn.com/problems/binary-tree-level-order-traversal) |  二叉树的层序遍历 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0102.binary-tree-level-order-traversal%20二叉树的层序遍历/main.go) | 
 [0103.binary-tree-zigzag-level-order-traversal](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal) |  二叉树的锯齿形层序遍历 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0103.binary-tree-zigzag-level-order-traversal%20二叉树的锯齿形层序遍历/main.go) | 
 [0107.binary-tree-level-order-traversal-ii](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii) |  二叉树的层序遍历 II | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0107.binary-tree-level-order-traversal-ii%20二叉树的层序遍历%20II/main.go) | 
 [0111.minimum-depth-of-binary-tree](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree) |  二叉树的最小深度 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0111.minimum-depth-of-binary-tree%20二叉树的最小深度/main.go) | 
 [0116.populating-next-right-pointers-in-each-node](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node) |  填充每个节点的下一个右侧节点指针 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0116.populating-next-right-pointers-in-each-node%20填充每个节点的下一个右侧节点指针/main.go) | 
 [0117.populating-next-right-pointers-in-each-node-ii](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii) |  填充每个节点的下一个右侧节点指针 II | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0117.populating-next-right-pointers-in-each-node-ii%20填充每个节点的下一个右侧节点指针%20II/main.go) | 
 [0199.binary-tree-right-side-view](https://leetcode-cn.com/problems/binary-tree-right-side-view) |  二叉树的右视图 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0199.binary-tree-right-side-view%20二叉树的右视图/main.go) | 
 [0297.serialize-and-deserialize-binary-tree](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree) |  二叉树的序列化与反序列化 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0297.serialize-and-deserialize-binary-tree%20二叉树的序列化与反序列化/main.go) | 
 [0428.serialize-and-deserialize-n-ary-tree](https://leetcode-cn.com/problems/serialize-and-deserialize-n-ary-tree) |  序列化和反序列化 N 叉树 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0428.serialize-and-deserialize-n-ary-tree%20序列化和反序列化%20N%20叉树/main.go) | 
 [0429.n-ary-tree-level-order-traversal](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal) |  N 叉树的层序遍历 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0429.n-ary-tree-level-order-traversal%20N%20叉树的层序遍历/main.go) | 
 [0433.minimum-genetic-mutation   ](https://leetcode-cn.com/problems/minimum-genetic-mutation) |  最小基因变化 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0433.minimum-genetic-mutation%20最小基因变化/main.go) | 
 [0513.find-bottom-left-tree-value](https://leetcode-cn.com/problems/find-bottom-left-tree-value) |  找树左下角的值 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0513.find-bottom-left-tree-value%20找树左下角的值/main.go) | 
 [0515.find-largest-value-in-each-tree-row](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row) |  在每个树行中找最大值 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0515.find-largest-value-in-each-tree-row%20在每个树行中找最大值/main.go) | 
 [0559.maximum-depth-of-n-ary-tree](https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree) |  N 叉树的最大深度 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0559.maximum-depth-of-n-ary-tree%20N%20叉树的最大深度/main.go) | 
 [0623.add-one-row-to-tree        ](https://leetcode-cn.com/problems/add-one-row-to-tree) |  在二叉树中增加一行 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0623.add-one-row-to-tree%20在二叉树中增加一行/main.go) | 
 [0637.average-of-levels-in-binary-tree](https://leetcode-cn.com/problems/average-of-levels-in-binary-tree) |  二叉树的层平均值 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/bfs/0637.average-of-levels-in-binary-tree%20二叉树的层平均值/main.go) | 
 [1161.maximum-level-sum-of-a-binary-tree](https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree) |  最大层内元素和 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/1161.maximum-level-sum-of-a-binary-tree%20最大层内元素和/main.go) | 
 [1306.jump-game-iii              ](https://leetcode-cn.com/problems/jump-game-iii) |  跳跃游戏 III | Medium | [Go](https://github.com/bygo/leetcode/blob/master/bfs/1306.jump-game-iii%20跳跃游戏%20III/main.go) | 

## Stack
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0020.valid-parentheses          ](https://leetcode-cn.com/problems/valid-parentheses) |  有效(){}[]括号 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0020.valid-parentheses%20有效(){}[]括号/main.go) | 
 [0032.longest-valid-parentheses  ](https://leetcode-cn.com/problems/longest-valid-parentheses) |  最长有效()括号 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0032.longest-valid-parentheses%20最长有效()括号/main.go) | 
 [0071.simplify-path              ](https://leetcode-cn.com/problems/simplify-path) |  简化路径 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0071.simplify-path%20简化路径/main.go) | 
 [0084.largest-rectangle-in-histogram](https://leetcode-cn.com/problems/largest-rectangle-in-histogram) |  统计矩形 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0084.largest-rectangle-in-histogram%20统计矩形/main.go) | 
 [0085.maximal-rectangle          ](https://leetcode-cn.com/problems/maximal-rectangle) |  最大矩形 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0085.maximal-rectangle%20最大矩形/main.go) | 
 [0094.binary-tree-inorder-traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal) |  二叉树的中序遍历 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0094.binary-tree-inorder-traversal%20二叉树的中序遍历/main.go) | 
 [0103.binary-tree-zigzag-level-order-traversal](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal) |  二叉树的锯齿形层序遍历 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0103.binary-tree-zigzag-level-order-traversal%20二叉树的锯齿形层序遍历/main.go) | 
 [0144.binary-tree-preorder-traversal](https://leetcode-cn.com/problems/binary-tree-preorder-traversal) |  二叉树的前序遍历 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0144.binary-tree-preorder-traversal%20二叉树的前序遍历/main.go) | 
 [0145.binary-tree-postorder-traversal](https://leetcode-cn.com/problems/binary-tree-postorder-traversal) |  二叉树的后序遍历 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0145.binary-tree-postorder-traversal%20二叉树的后序遍历/main.go) | 
 [0150.evaluate-reverse-polish-notation](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation) |  逆波兰表达式求值 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0150.evaluate-reverse-polish-notation%20逆波兰表达式求值/main.go) | 
 [0155.min-stack                  ](https://leetcode-cn.com/problems/min-stack) |  最小栈 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0155.min-stack%20最小栈/main.go) | 
 [0224.basic-calculator           ](https://leetcode-cn.com/problems/basic-calculator) |  加减小括号 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0224.basic-calculator%20加减小括号/main.go) | 
 [0225.implement-stack-using-queues](https://leetcode-cn.com/problems/implement-stack-using-queues) |  用队列实现栈 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0225.implement-stack-using-queues%20用队列实现栈/main.go) | 
 [0227.basic-calculator-ii        ](https://leetcode-cn.com/problems/basic-calculator-ii) |  加减乘除 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0227.basic-calculator-ii%20加减乘除/main.go) | 
 [0232.implement-queue-using-stacks](https://leetcode-cn.com/problems/implement-queue-using-stacks) |  用栈实现队列 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0232.implement-queue-using-stacks%20用栈实现队列/main.go) | 
 [0255.verify-preorder-sequence-in-binary-search-tree](https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree) |  验证前序遍历序列二叉搜索树 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0255.verify-preorder-sequence-in-binary-search-tree%20验证前序遍历序列二叉搜索树/main.go) | 
 [0272.closest-binary-search-tree-value-ii](https://leetcode-cn.com/problems/closest-binary-search-tree-value-ii) |  最接近的二叉搜索树值 II | Hard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0272.closest-binary-search-tree-value-ii%20最接近的二叉搜索树值%20II/main.go) | 
 [0316.remove-duplicate-letters   ](https://leetcode-cn.com/problems/remove-duplicate-letters) |  去除重复字母 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0316.remove-duplicate-letters%20去除重复字母/main.go) | 
 [0331.verify-preorder-serialization-of-a-binary-tree](https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree) |  验证二叉树的前序序列化 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0331.verify-preorder-serialization-of-a-binary-tree%20验证二叉树的前序序列化/main.go) | 
 [0341.flatten-nested-list-iterator](https://leetcode-cn.com/problems/flatten-nested-list-iterator) |  扁平化嵌套列表迭代器 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0341.flatten-nested-list-iterator%20扁平化嵌套列表迭代器/main.go) | 
 [0385.mini-parser                ](https://leetcode-cn.com/problems/mini-parser) |  迷你语法分析器 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0385.mini-parser%20迷你语法分析器/main.go) | 
 [0394.decode-string              ](https://leetcode-cn.com/problems/decode-string) |  字符串解码 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0394.decode-string%20字符串解码/main.go) | 
 [0402.remove-k-digits            ](https://leetcode-cn.com/problems/remove-k-digits) |  移掉K位数字后最小 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0402.remove-k-digits%20移掉K位数字后最小/main.go) | 
 [0430.flatten-a-multilevel-doubly-linked-list](https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list) |  扁平化多级双向链表 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0430.flatten-a-multilevel-doubly-linked-list%20扁平化多级双向链表/main.go) | 
 [0439.ternary-expression-parser  ](https://leetcode-cn.com/problems/ternary-expression-parser) |  三元表达式解析器 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0439.ternary-expression-parser%20三元表达式解析器/main.go) | 
 [0445.add-two-numbers-ii         ](https://leetcode-cn.com/problems/add-two-numbers-ii) |  两数相加 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0445.add-two-numbers-ii%20两数相加/main.go) | 
 [0456.132-pattern                ](https://leetcode-cn.com/problems/132-pattern) |  132模式 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0456.132-pattern%20132模式/main.go) | 
 [0496.next-greater-element-i     ](https://leetcode-cn.com/problems/next-greater-element-i) |  下一个更大元素 I | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0496.next-greater-element-i%20下一个更大元素%20I/main.go) | 
 [0503.next-greater-element-ii    ](https://leetcode-cn.com/problems/next-greater-element-ii) |  下一个更大的值 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0503.next-greater-element-ii%20下一个更大的值/main.go) | 
 [0591.tag-validator              ](https://leetcode-cn.com/problems/tag-validator) |  标签验证器 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0591.tag-validator%20标签验证器/main.go) | 
 [0636.exclusive-time-of-functions](https://leetcode-cn.com/problems/exclusive-time-of-functions) |  函数的独占时间 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0636.exclusive-time-of-functions%20函数的独占时间/main.go) | 
 [0678.valid-parenthesis-string   ](https://leetcode-cn.com/problems/valid-parenthesis-string) |  有效的括号字符串 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0678.valid-parenthesis-string%20有效的括号字符串/main.go) | 
 [0726.number-of-atoms            ](https://leetcode-cn.com/problems/number-of-atoms) |  原子的数量 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0726.number-of-atoms%20原子的数量/main.go) | 
 [0735.asteroid-collision         ](https://leetcode-cn.com/problems/asteroid-collision) |  星球碰撞 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0735.asteroid-collision%20星球碰撞/main.go) | 
 [0739.daily-temperatures         ](https://leetcode-cn.com/problems/daily-temperatures) |  能观测到更高温度的天数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0739.daily-temperatures%20能观测到更高温度的天数/main.go) | 
 [0772.basic-calculator-iii       ](https://leetcode-cn.com/problems/basic-calculator-iii) |  加减乘除小括号 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0772.basic-calculator-iii%20加减乘除小括号/main.go) | 
 [0844.backspace-string-compare   ](https://leetcode-cn.com/problems/backspace-string-compare) |  比较含退格的字符串 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/0844.backspace-string-compare%20比较含退格的字符串/main.go) | 
 [0856.score-of-parentheses       ](https://leetcode-cn.com/problems/score-of-parentheses) |  括号的分数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0856.score-of-parentheses%20括号的分数/main.go) | 
 [0880.decoded-string-at-index    ](https://leetcode-cn.com/problems/decoded-string-at-index) |  索引处的解码字符串 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0880.decoded-string-at-index%20索引处的解码字符串/main.go) | 
 [0895.maximum-frequency-stack    ](https://leetcode-cn.com/problems/maximum-frequency-stack) |  最大频率栈 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/stack/0895.maximum-frequency-stack%20最大频率栈/main.go) | 
 [0901.online-stock-span          ](https://leetcode-cn.com/problems/online-stock-span) |  股票价格跨度 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0901.online-stock-span%20股票价格跨度/main.go) | 
 [0921.minimum-add-to-make-parentheses-valid](https://leetcode-cn.com/problems/minimum-add-to-make-parentheses-valid) |  使括号有效的最少添加 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0921.minimum-add-to-make-parentheses-valid%20使括号有效的最少添加/main.go) | 
 [0946.validate-stack-sequences   ](https://leetcode-cn.com/problems/validate-stack-sequences) |  验证栈序列 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/0946.validate-stack-sequences%20验证栈序列/main.go) | 
 [1003.check-if-word-is-valid-after-substitutions](https://leetcode-cn.com/problems/check-if-word-is-valid-after-substitutions) |  检查替换后的词是否有效 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/1003.check-if-word-is-valid-after-substitutions%20检查替换后的词是否有效/main.go) | 
 [1019.next-greater-node-in-linked-list](https://leetcode-cn.com/problems/next-greater-node-in-linked-list) |  下一个更大的值 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/1019.next-greater-node-in-linked-list%20下一个更大的值/main.go) | 
 [1081.smallest-subsequence-of-distinct-characters](https://leetcode-cn.com/problems/smallest-subsequence-of-distinct-characters) |  去除重复字母 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/1081.smallest-subsequence-of-distinct-characters%20去除重复字母/main.go) | 
 [1190.reverse-substrings-between-each-pair-of-parentheses](https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses) |  反转每对括号间的子串 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/stack/1190.reverse-substrings-between-each-pair-of-parentheses%20反转每对括号间的子串/main.go) | 
 [1221.split-a-string-in-balanced-strings](https://leetcode-cn.com/problems/split-a-string-in-balanced-strings) |  分割平衡字符串 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/1221.split-a-string-in-balanced-strings%20分割平衡字符串/main.go) | 
 [1544.make-the-string-great      ](https://leetcode-cn.com/problems/make-the-string-great) |  整理字符串 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/1544.make-the-string-great%20整理字符串/main.go) | 
 [1598.crawler-log-folder         ](https://leetcode-cn.com/problems/crawler-log-folder) |  文件夹操作日志搜集器 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/stack/1598.crawler-log-folder%20文件夹操作日志搜集器/main.go) | 

## Enum
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0002.add-two-numbers            ](https://leetcode-cn.com/problems/add-two-numbers) |  两数相加 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/0002.add-two-numbers%20两数相加/main.go) | 
 [0006.zigzag-conversion          ](https://leetcode-cn.com/problems/zigzag-conversion) |  Z 字形变换 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/0006.zigzag-conversion%20Z%20字形变换/main.go) | 
 [0014.longest-common-prefix      ](https://leetcode-cn.com/problems/longest-common-prefix) |  最长公共前缀 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/0014.longest-common-prefix%20最长公共前缀/main.go) | 
 [0036.valid-sudoku               ](https://leetcode-cn.com/problems/valid-sudoku) |  有效的数独 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/0036.valid-sudoku%20有效的数独/main.go) | 
 [0038.count-and-say              ](https://leetcode-cn.com/problems/count-and-say) |  外观数列 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/0038.count-and-say%20外观数列/main.go) | 
 [0058.length-of-last-word        ](https://leetcode-cn.com/problems/length-of-last-word) |  最后一个单词的长度 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/0058.length-of-last-word%20最后一个单词的长度/main.go) | 
 [0121.best-time-to-buy-and-sell-stock](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock) |  买卖股票的最佳时机 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/0121.best-time-to-buy-and-sell-stock%20买卖股票的最佳时机/main.go) | 
 [0122.best-time-to-buy-and-sell-stock-ii](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii) |  买卖股票的最佳时机 II | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/0122.best-time-to-buy-and-sell-stock-ii%20买卖股票的最佳时机%20II/main.go) | 
 [0171.excel-sheet-column-number  ](https://leetcode-cn.com/problems/excel-sheet-column-number) |  Excel 表列序号 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/0171.excel-sheet-column-number%20Excel%20表列序号/main.go) | 
 [0423.reconstruct-original-digits-from-english](https://leetcode-cn.com/problems/reconstruct-original-digits-from-english) |  从英文中重建数字 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/0423.reconstruct-original-digits-from-english%20从英文中重建数字/main.go) | 
 [1213.intersection-of-three-sorted-arrays](https://leetcode-cn.com/problems/intersection-of-three-sorted-arrays) |  三个有序数组的交集 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/1213.intersection-of-three-sorted-arrays%20三个有序数组的交集/main.go) | 
 [1295.find-numbers-with-even-number-of-digits](https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits) |  统计位数为偶数的数字 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/1295.find-numbers-with-even-number-of-digits%20统计位数为偶数的数字/main.go) | 
 [1389.create-target-array-in-the-given-order](https://leetcode-cn.com/problems/create-target-array-in-the-given-order) |  按既定顺序创建目标数组 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/1389.create-target-array-in-the-given-order%20按既定顺序创建目标数组/main.go) | 
 [1503.last-moment-before-all-ants-fall-out-of-a-plank](https://leetcode-cn.com/problems/last-moment-before-all-ants-fall-out-of-a-plank) |  所有蚂蚁掉下来前的最后一刻 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/enum/1503.last-moment-before-all-ants-fall-out-of-a-plank%20所有蚂蚁掉下来前的最后一刻/main.go) | 
 [1646.get-maximum-in-generated-array](https://leetcode-cn.com/problems/get-maximum-in-generated-array) |  获取生成数组中的最大值 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/1646.get-maximum-in-generated-array%20获取生成数组中的最大值/main.go) | 
 [1672.richest-customer-wealth    ](https://leetcode-cn.com/problems/richest-customer-wealth) |  最富有客户的资产总量 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/1672.richest-customer-wealth%20最富有客户的资产总量/main.go) | 
 [2042.check-if-numbers-are-ascending-in-a-sentence](https://leetcode-cn.com/problems/check-if-numbers-are-ascending-in-a-sentence) |  检查句子中的数字是否递增 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/enum/2042.check-if-numbers-are-ascending-in-a-sentence%20检查句子中的数字是否递增/main.go) | 
 [5875.final-value-of-variable-after-performing-operations 执行操作后的变量值](https://leetcode-cn.com/problems/5875.final-value-of-variable-after-performing-operations 执行操作后的变量值) |  | Hard | [Go](https://github.com/bygo/leetcode/blob/master/enum/5875.final-value-of-variable-after-performing-operations%20执行操作后的变量值/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/enum/5883.check-if-word-can-be-placed-in-crossword%20判断单词是否能放入填字游戏内/main.go) |  

## Tree
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0022.generate-parentheses       ](https://leetcode-cn.com/problems/generate-parentheses) |  有效的 括号组合 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0022.generate-parentheses%20有效的%20括号组合/main.go) | 
 [0094.binary-tree-inorder-traversal](https://leetcode-cn.com/problems/binary-tree-inorder-traversal) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0094.binary-tree-inorder-traversal/dfs/inorder/recursive/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0094.binary-tree-inorder-traversal/dfs/inorder/stack/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0094.binary-tree-inorder-traversal/dfs/morris/break/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0094.binary-tree-inorder-traversal/dfs/morris/keep/main.go) |  
 [0095.unique-binary-search-trees-ii](https://leetcode-cn.com/problems/unique-binary-search-trees-ii) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0095.unique-binary-search-trees-ii/dfs/catalan/main.go) | 
 [0096.unique-binary-search-trees ](https://leetcode-cn.com/problems/unique-binary-search-trees) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0096.unique-binary-search-trees/catalan/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0096.unique-binary-search-trees/dp/main.go) |  
 [0098.validate-binary-search-tree](https://leetcode-cn.com/problems/validate-binary-search-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0098.validate-binary-search-tree/dfs/inorder/recursive/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0098.validate-binary-search-tree/dfs/inorder/stack/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0098.validate-binary-search-tree/dfs/preorder/recursive/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0098.validate-binary-search-tree/dfs/preorder/stack/main.go) |  
 [0099.recover-binary-search-tree ](https://leetcode-cn.com/problems/recover-binary-search-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0099.recover-binary-search-tree/dfs/inorder/recursive/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0099.recover-binary-search-tree/dfs/inorder/stack/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0099.recover-binary-search-tree/dfs/morris/keep.go) |  
 [0100.same-tree                  ](https://leetcode-cn.com/problems/same-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0100.same-tree/dfs/recursive/main.go) | 
 [0101.symmetric-tree             ](https://leetcode-cn.com/problems/symmetric-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0101.symmetric-tree/dfs/recursive/main.go) | 
 [0102.binary-tree-level-order-traversal](https://leetcode-cn.com/problems/binary-tree-level-order-traversal) |  二叉树的层序遍历 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0102.binary-tree-level-order-traversal%20二叉树的层序遍历/dfs/recursive/main.go) | 
 [0103.binary-tree-zigzag-level-order-traversal](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal) |  二叉树的锯齿形层序遍历 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0103.binary-tree-zigzag-level-order-traversal%20二叉树的锯齿形层序遍历/dfs/recursive/main.go) | 
 [0104.maximum-depth-of-binary-tree](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0104.maximum-depth-of-binary-tree/dfs/main.go) | 
 [0105.construct-binary-tree-from-preorder-and-inorder-traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0105.construct-binary-tree-from-preorder-and-inorder-traversal/main.go) | 
 [0106.construct-binary-tree-from-inorder-and-postorder-traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0106.construct-binary-tree-from-inorder-and-postorder-traversal/main.go) | 
 [0107.binary-tree-level-order-traversal-ii](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0107.binary-tree-level-order-traversal-ii/dfs/recursive/main.go) | 
 [0108.convert-sorted-array-to-binary-search-tree](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0108.convert-sorted-array-to-binary-search-tree/dfs/recursive/main.go) | 
 [0109.convert-sorted-list-to-binary-search-tree](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0109.convert-sorted-list-to-binary-search-tree/array/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0109.convert-sorted-list-to-binary-search-tree/inorder/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0109.convert-sorted-list-to-binary-search-tree/recursive/main.go) |  
 [0110.balanced-binary-tree       ](https://leetcode-cn.com/problems/balanced-binary-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0110.balanced-binary-tree/postorder/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0110.balanced-binary-tree/top/main.go) |  
 [0111.minimum-depth-of-binary-tree](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0111.minimum-depth-of-binary-tree/dfs/main.go) | 
 [0112.path-sum                   ](https://leetcode-cn.com/problems/path-sum) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0112.path-sum/dfs/main.go) | 
 [0113.path-sum-ii                ](https://leetcode-cn.com/problems/path-sum-ii) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0113.path-sum-ii/dfs/main.go) | 
 [0114.flatten-binary-tree-to-linked-list](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0114.flatten-binary-tree-to-linked-list/preorder.morris/main.go) | 
 [0116.populating-next-right-pointers-in-each-node](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0116.populating-next-right-pointers-in-each-node/dfs/main.go) | 
 [0124.binary-tree-maximum-path-sum](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum) |  | Hard | [Go](https://github.com/bygo/leetcode/blob/master/tree/0124.binary-tree-maximum-path-sum/dfs/main.go) | 
 [0129.sum-root-to-leaf-numbers   ](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0129.sum-root-to-leaf-numbers/dfs/main.go) | 
 [0144.binary-tree-preorder-traversal](https://leetcode-cn.com/problems/binary-tree-preorder-traversal) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0144.binary-tree-preorder-traversal/dfs/recursive/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0144.binary-tree-preorder-traversal/dfs/stack/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0144.binary-tree-preorder-traversal/morris/break/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0144.binary-tree-preorder-traversal/morris/keep/main.go) |  
 [0145.binary-tree-postorder-traversal](https://leetcode-cn.com/problems/binary-tree-postorder-traversal) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0145.binary-tree-postorder-traversal/dfs/recursive/main.go) | 
 [0156.binary-tree-upside-down    ](https://leetcode-cn.com/problems/binary-tree-upside-down) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0156.binary-tree-upside-down/link/main.go) | 
 [0173.binary-search-tree-iterator](https://leetcode-cn.com/problems/binary-search-tree-iterator) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0173.binary-search-tree-iterator/main.go) | 
 [0199.binary-tree-right-side-view](https://leetcode-cn.com/problems/binary-tree-right-side-view) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0199.binary-tree-right-side-view/dfs/main.go) | 
 [0222.count-complete-tree-nodes  ](https://leetcode-cn.com/problems/count-complete-tree-nodes) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0222.count-complete-tree-nodes/dfs/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0222.count-complete-tree-nodes/two_pointer/main.go) |  
 [0226.invert-binary-tree         ](https://leetcode-cn.com/problems/invert-binary-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0226.invert-binary-tree/dfs/main.go) | 
 [0230.kth-smallest-element-in-a-bst](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0230.kth-smallest-element-in-a-bst/main.go) | 
 [0235.lowest-common-ancestor-of-a-binary-search-tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0235.lowest-common-ancestor-of-a-binary-search-tree/dfs/main.go) | 
 [0236.lowest-common-ancestor-of-a-binary-tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0236.lowest-common-ancestor-of-a-binary-tree/dfs/main.go) | 
 [0250.count-univalue-subtrees    ](https://leetcode-cn.com/problems/count-univalue-subtrees) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0250.count-univalue-subtrees/dfs/main.go) | 
 [0255.verify-preorder-sequence-in-binary-search-tree](https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0255.verify-preorder-sequence-in-binary-search-tree/main.go) | 
 [0257.binary-tree-paths          ](https://leetcode-cn.com/problems/binary-tree-paths) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0257.binary-tree-paths/dfs/main.go) | 
 [0270.closest-binary-search-tree-value](https://leetcode-cn.com/problems/closest-binary-search-tree-value) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0270.closest-binary-search-tree-value/main.go) | 
 [0272.closest-binary-search-tree-value-ii](https://leetcode-cn.com/problems/closest-binary-search-tree-value-ii) |  | Hard | [Go](https://github.com/bygo/leetcode/blob/master/tree/0272.closest-binary-search-tree-value-ii/main.go) | 
 [0285.inorder-successor-in-bst   ](https://leetcode-cn.com/problems/inorder-successor-in-bst) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0285.inorder-successor-in-bst/main.go) | 
 [0297.serialize-and-deserialize-binary-tree](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree) |  | Hard | [Go](https://github.com/bygo/leetcode/blob/master/tree/0297.serialize-and-deserialize-binary-tree/bfs/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0297.serialize-and-deserialize-binary-tree/dfs/main.go) |  
 [0298.binary-tree-longest-consecutive-sequence](https://leetcode-cn.com/problems/binary-tree-longest-consecutive-sequence) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0298.binary-tree-longest-consecutive-sequence/main.go) | 
 [0314.binary-tree-vertical-order-traversal](https://leetcode-cn.com/problems/binary-tree-vertical-order-traversal) |  二叉树的垂直遍历 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0314.binary-tree-vertical-order-traversal%20二叉树的垂直遍历/main.go) | 
 [0366.find-leaves-of-binary-tree ](https://leetcode-cn.com/problems/find-leaves-of-binary-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0366.find-leaves-of-binary-tree/main.go) | 
 [0404.sum-of-left-leaves         ](https://leetcode-cn.com/problems/sum-of-left-leaves) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0404.sum-of-left-leaves/dfs/main.go) | 
 [0426.convert-binary-search-tree-to-sorted-doubly-linked-list](https://leetcode-cn.com/problems/convert-binary-search-tree-to-sorted-doubly-linked-list) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0426.convert-binary-search-tree-to-sorted-doubly-linked-list/inorder/main.go) | 
 [0428.serialize-and-deserialize-n-ary-tree](https://leetcode-cn.com/problems/serialize-and-deserialize-n-ary-tree) |  | Hard | [Go](https://github.com/bygo/leetcode/blob/master/tree/0428.serialize-and-deserialize-n-ary-tree/dfs/main.go) | 
 [0437.path-sum-iii               ](https://leetcode-cn.com/problems/path-sum-iii) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0437.path-sum-iii/dfs/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0437.path-sum-iii/hash_table/main.go) |  
 [0449.serialize-and-deserialize-bst](https://leetcode-cn.com/problems/serialize-and-deserialize-bst) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0449.serialize-and-deserialize-bst/dfs/preorder/main.go) | 
 [0450.delete-node-in-a-bst       ](https://leetcode-cn.com/problems/delete-node-in-a-bst) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0450.delete-node-in-a-bst/dfs/main.go) | 
 [0501.find-mode-in-binary-search-tree](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0501.find-mode-in-binary-search-tree/dfs/recursive/main.go) | 
 [0508.most-frequent-subtree-sum  ](https://leetcode-cn.com/problems/most-frequent-subtree-sum) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0508.most-frequent-subtree-sum/dfs/main.go) | 
 [0510.inorder-successor-in-bst-ii](https://leetcode-cn.com/problems/inorder-successor-in-bst-ii) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0510.inorder-successor-in-bst-ii/main.go) | 
 [0530.minimum-absolute-difference-in-bst](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0530.minimum-absolute-difference-in-bst/dfs/main.go) | 
 [0538.convert-bst-to-greater-tree](https://leetcode-cn.com/problems/convert-bst-to-greater-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0538.convert-bst-to-greater-tree/dfs/main.go) | 
 [0543.diameter-of-binary-tree    ](https://leetcode-cn.com/problems/diameter-of-binary-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0543.diameter-of-binary-tree/dfs/main.go) | 
 [0559.maximum-depth-of-n-ary-tree](https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0559.maximum-depth-of-n-ary-tree/dfs/main.go) | 
 [0563.binary-tree-tilt           ](https://leetcode-cn.com/problems/binary-tree-tilt) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0563.binary-tree-tilt/dfs/main.go) | 
 [0572.subtree-of-another-tree    ](https://leetcode-cn.com/problems/subtree-of-another-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0572.subtree-of-another-tree/dfs/main.go) | 
 [0589.n-ary-tree-preorder-traversal](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0589.n-ary-tree-preorder-traversal/dfs/recursive/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0589.n-ary-tree-preorder-traversal/dfs/stack/main.go) |  
 [0590.n-ary-tree-postorder-traversal](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0590.n-ary-tree-postorder-traversal/dfs/recursive/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/tree/0590.n-ary-tree-postorder-traversal/dfs/stack/main.go) |  
 [0606.construct-string-from-binary-tree](https://leetcode-cn.com/problems/construct-string-from-binary-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0606.construct-string-from-binary-tree/dfs/main.go) | 
 [0617.merge-two-binary-trees     ](https://leetcode-cn.com/problems/merge-two-binary-trees) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0617.merge-two-binary-trees/dfs/recursive/main.go) | 
 [0653.two-sum-iv-input-is-a-bst  ](https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0653.two-sum-iv-input-is-a-bst/dfs/main.go) | 
 [0654.maximum-binary-tree        ](https://leetcode-cn.com/problems/maximum-binary-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0654.maximum-binary-tree/dfs/main.go) | 
 [0669.trim-a-binary-search-tree  ](https://leetcode-cn.com/problems/trim-a-binary-search-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0669.trim-a-binary-search-tree/dfs/main.go) | 
 [0671.second-minimum-node-in-a-binary-tree](https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0671.second-minimum-node-in-a-binary-tree/dfs/main.go) | 
 [0687.longest-univalue-path      ](https://leetcode-cn.com/problems/longest-univalue-path) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0687.longest-univalue-path/dfs/main.go) | 
 [0700.search-in-a-binary-search-tree](https://leetcode-cn.com/problems/search-in-a-binary-search-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0700.search-in-a-binary-search-tree/dfs/main.go) | 
 [0701.insert-into-a-binary-search-tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0701.insert-into-a-binary-search-tree/dfs/main.go) | 
 [0783.minimum-distance-between-bst-nodes](https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0783.minimum-distance-between-bst-nodes/dfs/main.go) | 
 [0814.binary-tree-pruning        ](https://leetcode-cn.com/problems/binary-tree-pruning) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0814.binary-tree-pruning/dfs/main.go) | 
 [0865.smallest-subtree-with-all-the-deepest-nodes](https://leetcode-cn.com/problems/smallest-subtree-with-all-the-deepest-nodes) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0865.smallest-subtree-with-all-the-deepest-nodes/main.go) | 
 [0872.leaf-similar-trees         ](https://leetcode-cn.com/problems/leaf-similar-trees) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0872.leaf-similar-trees/dfs/main.go) | 
 [0889.construct-binary-tree-from-preorder-and-postorder-traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0889.construct-binary-tree-from-preorder-and-postorder-traversal/main.go) | 
 [0897.increasing-order-search-tree](https://leetcode-cn.com/problems/increasing-order-search-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0897.increasing-order-search-tree/inorder/main.go) | 
 [0938.range-sum-of-bst           ](https://leetcode-cn.com/problems/range-sum-of-bst) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0938.range-sum-of-bst/main.go) | 
 [0951.flip-equivalent-binary-trees](https://leetcode-cn.com/problems/flip-equivalent-binary-trees) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0951.flip-equivalent-binary-trees/dfs/main.go) | 
 [0965.univalued-binary-tree      ](https://leetcode-cn.com/problems/univalued-binary-tree) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/0965.univalued-binary-tree/dfs/main.go) | 
 [0979.distribute-coins-in-binary-tree](https://leetcode-cn.com/problems/distribute-coins-in-binary-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/0979.distribute-coins-in-binary-tree/dfs/main.go) | 
 [1008.construct-binary-search-tree-from-preorder-traversal](https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1008.construct-binary-search-tree-from-preorder-traversal/dfs/main.go) | 
 [1022.sum-of-root-to-leaf-binary-numbers](https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/1022.sum-of-root-to-leaf-binary-numbers/dfs/main.go) | 
 [1026.maximum-difference-between-node-and-ancestor](https://leetcode-cn.com/problems/maximum-difference-between-node-and-ancestor) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1026.maximum-difference-between-node-and-ancestor/dfs/main.go) | 
 [1038.binary-search-tree-to-greater-sum-tree](https://leetcode-cn.com/problems/binary-search-tree-to-greater-sum-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1038.binary-search-tree-to-greater-sum-tree/inorder/main.go) | 
 [1104.path-in-zigzag-labelled-binary-tree](https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1104.path-in-zigzag-labelled-binary-tree/main.go) | 
 [1123.lowest-common-ancestor-of-deepest-leaves](https://leetcode-cn.com/problems/lowest-common-ancestor-of-deepest-leaves) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1123.lowest-common-ancestor-of-deepest-leaves/dfs/main.go) | 
 [1161.maximum-level-sum-of-a-binary-tree](https://leetcode-cn.com/problems/maximum-level-sum-of-a-binary-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1161.maximum-level-sum-of-a-binary-tree/bfs/main.go) | 
 [1302.deepest-leaves-sum         ](https://leetcode-cn.com/problems/deepest-leaves-sum) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1302.deepest-leaves-sum/dfs/main.go) | 
 [1315.sum-of-nodes-with-even-valued-grandparent](https://leetcode-cn.com/problems/sum-of-nodes-with-even-valued-grandparent) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1315.sum-of-nodes-with-even-valued-grandparent/dfs/main.go) | 
 [1325.delete-leaves-with-a-given-value](https://leetcode-cn.com/problems/delete-leaves-with-a-given-value) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1325.delete-leaves-with-a-given-value/postorder/main.go) | 
 [1448.count-good-nodes-in-binary-tree](https://leetcode-cn.com/problems/count-good-nodes-in-binary-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1448.count-good-nodes-in-binary-tree/dfs/main.go) | 
 [1484.clone-binary-tree-with-random-pointer](https://leetcode-cn.com/problems/clone-binary-tree-with-random-pointer) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1484.clone-binary-tree-with-random-pointer/dfs/main.go) | 
 [1490.clone-n-ary-tree           ](https://leetcode-cn.com/problems/clone-n-ary-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1490.clone-n-ary-tree/dfs/main.go) | 
 [1560.most-visited-sector-in-a-circular-track](https://leetcode-cn.com/problems/most-visited-sector-in-a-circular-track) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/tree/1560.most-visited-sector-in-a-circular-track/hash/main.go) | 
 [1600.throne-inheritance         ](https://leetcode-cn.com/problems/throne-inheritance) |  皇位继承顺序 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1600.throne-inheritance%20皇位继承顺序/main.go) | 
 [1602.find-nearest-right-node-in-binary-tree](https://leetcode-cn.com/problems/find-nearest-right-node-in-binary-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1602.find-nearest-right-node-in-binary-tree/bfs/main.go) | 
 [1669.merge-in-between-linked-lists](https://leetcode-cn.com/problems/merge-in-between-linked-lists) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1669.merge-in-between-linked-lists/main.go) | 
 [1740.find-distance-in-a-binary-tree](https://leetcode-cn.com/problems/find-distance-in-a-binary-tree) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/tree/1740.find-distance-in-a-binary-tree/dfs/main.go) | 

## Divide&conquer
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0014.longest-common-prefix      ](https://leetcode-cn.com/problems/longest-common-prefix) |  最长公共前缀 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/divide&conquer/0014.longest-common-prefix%20最长公共前缀/main.go) | 

## Back
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0037.sudoku-solver              ](https://leetcode-cn.com/problems/sudoku-solver) |  解数独 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/back/0037.sudoku-solver%20解数独/main.go) | 
 [0051.n-queens                   ](https://leetcode-cn.com/problems/n-queens) |  N 皇后 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/back/0051.n-queens%20N%20皇后/main.go) | 
 [0052.n-queens-ii                ](https://leetcode-cn.com/problems/n-queens-ii) |  N皇后 II | Hard | [Go](https://github.com/bygo/leetcode/blob/master/back/0052.n-queens-ii%20N皇后%20II/main.go) | 
 [Offer.0038.zi-fu-chuan-de-pai-lie-lcof](https://leetcode-cn.com/problems/count-and-say) |  字符串的排列 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/back/Offer.0038.zi-fu-chuan-de-pai-lie-lcof%20字符串的排列/main.go) | 

## Sort
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0179.largest-number             ](https://leetcode-cn.com/problems/largest-number) |  最大数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sort/0179.largest-number%20最大数/main.go) | 
 [面试题                             ](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number) |  17.14. 最小K个数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sort/面试题%2017.14.%20最小K个数/main.go) | 

## Search
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0011.container-with-most-water  ](https://leetcode-cn.com/problems/container-with-most-water) |  盛最多水的容器 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/search/0011.container-with-most-water%20盛最多水的容器/main.go) | 
 [0026.remove-duplicates-from-sorted-array](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array) |  删除有序数组中的重复项 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/search/0026.remove-duplicates-from-sorted-array%20删除有序数组中的重复项/main.go) | 
 [1838.frequency-of-the-most-frequent-element](https://leetcode-cn.com/problems/frequency-of-the-most-frequent-element) |  最高频元素的频数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/search/1838.frequency-of-the-most-frequent-element%20最高频元素的频数/main.go) | 
 [1894.find-the-student-that-will-replace-the-chalk](https://leetcode-cn.com/problems/find-the-student-that-will-replace-the-chalk) |  找到需要补充粉笔的学生编号 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/search/1894.find-the-student-that-will-replace-the-chalk%20找到需要补充粉笔的学生编号/main.go) | 
 [Offer                           ](https://leetcode-cn.com/problems/maximum-subarray) |  53 - I. 在排序数组中查找数字 I | Easy | [Go](https://github.com/bygo/leetcode/blob/master/search/Offer%2053%20-%20I.%20在排序数组中查找数字%20I/main.go) | 

## LinkedList
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0457.circular-array-loop        ](https://leetcode-cn.com/problems/circular-array-loop) |  环形数组是否存在循环 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/(Array)%200457.circular-array-loop%20环形数组是否存在循环/main.go) | Array
 [0021.merge-two-sorted-lists     ](https://leetcode-cn.com/problems/merge-two-sorted-lists) |  合并两个有序链表 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/(Merge)%200021.merge-two-sorted-lists%20合并两个有序链表/main.go) | Merge
 [0023.merge-k-sorted-lists       ](https://leetcode-cn.com/problems/merge-k-sorted-lists) |  合并K个有序链表 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/(Merge)%200023.merge-k-sorted-lists%20合并K个有序链表/main.go) | Merge
 [0355.design-twitter             ](https://leetcode-cn.com/problems/design-twitter) |  设计推特 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/(Merge)%200355.design-twitter%20设计推特/main.go) | Merge
 [0024.swap-nodes-in-pairs        ](https://leetcode-cn.com/problems/swap-nodes-in-pairs) |  两两交换节点 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/(Swap)%200024.swap-nodes-in-pairs%20两两交换节点/main.go) | Swap
 [0025.reverse-nodes-in-k-group   ](https://leetcode-cn.com/problems/reverse-nodes-in-k-group) |  K 个一组翻转链表 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0025.reverse-nodes-in-k-group%20K%20个一组翻转链表/main.go) | 
 [0061.rotate-list                ](https://leetcode-cn.com/problems/rotate-list) |  拼接链表 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0061.rotate-list%20拼接链表/main.go) | 
 [0082.remove-duplicates-from-sorted-list-ii](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii) |  删除重复全部元素 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0082.remove-duplicates-from-sorted-list-ii%20删除重复全部元素/main.go) | 
 [0083.remove-duplicates-from-sorted-list](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list) |  删除重复多余元素 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0083.remove-duplicates-from-sorted-list%20删除重复多余元素/main.go) | 
 [0086.partition-list             ](https://leetcode-cn.com/problems/partition-list) |  小于X的靠左 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0086.partition-list%20小于X的靠左/main.go) | 
 [0092.reverse-linked-list-ii     ](https://leetcode-cn.com/problems/reverse-linked-list-ii) |  反转链表 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0092.reverse-linked-list-ii%20反转链表/main.go) | 
 [0109.convert-sorted-list-to-binary-search-tree](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree) |  有序链表转二叉搜索树 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0109.convert-sorted-list-to-binary-search-tree%20有序链表转二叉搜索树/array/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0109.convert-sorted-list-to-binary-search-tree%20有序链表转二叉搜索树/inorder/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0109.convert-sorted-list-to-binary-search-tree%20有序链表转二叉搜索树/recursive/main.go) |  
 [0138.copy-list-with-random-pointer](https://leetcode-cn.com/problems/copy-list-with-random-pointer) |  深拷贝随机指针链表 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0138.copy-list-with-random-pointer%20深拷贝随机指针链表/main.go) | 
 [0141.linked-list-cycle          ](https://leetcode-cn.com/problems/linked-list-cycle) |  是否有环 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0141.linked-list-cycle%20是否有环/main.go) | 
 [0142.linked-list-cycle-ii       ](https://leetcode-cn.com/problems/linked-list-cycle-ii) |  环的起点 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0142.linked-list-cycle-ii%20环的起点/main.go) | 
 [0143.reorder-list               ](https://leetcode-cn.com/problems/reorder-list) |  重排链表 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0143.reorder-list%20重排链表/main.go) | 
 [0147.insertion-sort-list        ](https://leetcode-cn.com/problems/insertion-sort-list) |  插入排序 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0147.insertion-sort-list%20插入排序/main.go) | 
 [0160.intersection-of-two-linked-lists](https://leetcode-cn.com/problems/intersection-of-two-linked-lists) |  相交起始点 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0160.intersection-of-two-linked-lists%20相交起始点/main.go) | 
 [0203.remove-linked-list-elements](https://leetcode-cn.com/problems/remove-linked-list-elements) |  删除链表元素 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0203.remove-linked-list-elements%20删除链表元素/main.go) | 
 [0206.reverse-linked-list        ](https://leetcode-cn.com/problems/reverse-linked-list) |  反转链表 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0206.reverse-linked-list%20反转链表/main.go) | 
 [0234.palindrome-linked-list     ](https://leetcode-cn.com/problems/palindrome-linked-list) |  回文链表 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0234.palindrome-linked-list%20回文链表/main.go) | 
 [0237.delete-node-in-a-linked-list](https://leetcode-cn.com/problems/delete-node-in-a-linked-list) |  删除节点 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0237.delete-node-in-a-linked-list%20删除节点/main.go) | 
 [0328.odd-even-linked-list       ](https://leetcode-cn.com/problems/odd-even-linked-list) |  奇偶链表 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0328.odd-even-linked-list%20奇偶链表/main.go) | 
 [0369.plus-one-linked-list       ](https://leetcode-cn.com/problems/plus-one-linked-list) |  加一 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0369.plus-one-linked-list%20加一/main.go) | 
 [0379.design-phone-directory     ](https://leetcode-cn.com/problems/design-phone-directory) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0379.design-phone-directory/main.go) | 
 [0430.flatten-a-multilevel-doubly-linked-list](https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list) |  扁平化多级双向链表 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0430.flatten-a-multilevel-doubly-linked-list%20扁平化多级双向链表/main.go) | 
 [0445.add-two-numbers-ii         ](https://leetcode-cn.com/problems/add-two-numbers-ii) |  两数相加 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0445.add-two-numbers-ii%20两数相加/main.go) | 
 [0707.design-linked-list         ](https://leetcode-cn.com/problems/design-linked-list) |  设计链表 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0707.design-linked-list%20设计链表/main.go) | 
 [0708.insert-into-a-sorted-circular-linked-list](https://leetcode-cn.com/problems/insert-into-a-sorted-circular-linked-list) |  循环有序列表的插入 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0708.insert-into-a-sorted-circular-linked-list%20循环有序列表的插入/main.go) | 
 [0725.split-linked-list-in-parts ](https://leetcode-cn.com/problems/split-linked-list-in-parts) |  分割链表 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0725.split-linked-list-in-parts%20分割链表/main.go) | 
 [0817.linked-list-components     ](https://leetcode-cn.com/problems/linked-list-components) |  链表组件 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0817.linked-list-components%20链表组件/main.go) | 
 [0876.middle-of-the-linked-list  ](https://leetcode-cn.com/problems/middle-of-the-linked-list) |  中间节点 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/0876.middle-of-the-linked-list%20中间节点/main.go) | 
 [1019.next-greater-node-in-linked-list](https://leetcode-cn.com/problems/next-greater-node-in-linked-list) |  下一个更大的值 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1019.next-greater-node-in-linked-list%20下一个更大的值/main.go) | 
 [1171.remove-zero-sum-consecutive-nodes-from-linked-list](https://leetcode-cn.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list) |  从链表中删去总和值为零的连续节点 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1171.remove-zero-sum-consecutive-nodes-from-linked-list%20从链表中删去总和值为零的连续节点/main.go) | 
 [1290.convert-binary-number-in-a-linked-list-to-integer](https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer) |  二进制链表的值 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1290.convert-binary-number-in-a-linked-list-to-integer%20二进制链表的值/main.go) | 
 [1367.linked-list-in-binary-tree ](https://leetcode-cn.com/problems/linked-list-in-binary-tree) |  二叉树的链表 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1367.linked-list-in-binary-tree%20二叉树的链表/main.go) | 
 [1474.delete-n-nodes-after-m-nodes-of-a-linked-list](https://leetcode-cn.com/problems/delete-n-nodes-after-m-nodes-of-a-linked-list) |   删除m节点后移动n节点 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1474.delete-n-nodes-after-m-nodes-of-a-linked-list%20%20删除m节点后移动n节点/main.go) | 
 [1634.add-two-polynomials-represented-as-linked-lists](https://leetcode-cn.com/problems/add-two-polynomials-represented-as-linked-lists) |  多项式合并 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1634.add-two-polynomials-represented-as-linked-lists%20多项式合并/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1634.add-two-polynomials-represented-as-linked-lists%20多项式合并/main.js) |  
 [1669.merge-in-between-linked-lists](https://leetcode-cn.com/problems/merge-in-between-linked-lists) |  删除区间后合并 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1669.merge-in-between-linked-lists%20删除区间后合并/main.go) | 
 [1670.design-front-middle-back-queue](https://leetcode-cn.com/problems/design-front-middle-back-queue) |  队列 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1670.design-front-middle-back-queue%20队列/main.go) | 
 [1721.swapping-nodes-in-a-linked-list](https://leetcode-cn.com/problems/swapping-nodes-in-a-linked-list) |  交换镜像节点 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/1721.swapping-nodes-in-a-linked-list%20交换镜像节点/main.go) | 
 [剑指                              ](https://leetcode-cn.com/problems/generate-parentheses) |  Offer 22. 链表中倒数第k个节点 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/剑指%20Offer%2022.%20链表中倒数第k个节点/main.go) | 
 [剑指                              ](https://leetcode-cn.com/problems/merge-k-sorted-lists) |  Offer II 023. 两个链表的第一个重合节点 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/linked_list/剑指%20Offer%20II%20023.%20两个链表的第一个重合节点/main.go) | 

## BinarySearch
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0004.median-of-two-sorted-arrays](https://leetcode-cn.com/problems/median-of-two-sorted-arrays) |  寻找两个正序数组的中位数 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0004.median-of-two-sorted-arrays%20寻找两个正序数组的中位数/main.go) | 
 [0014.longest-common-prefix      ](https://leetcode-cn.com/problems/longest-common-prefix) |  最长公共前缀 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0014.longest-common-prefix%20最长公共前缀/main.go) | 
 [0033.search-in-rotated-sorted-array](https://leetcode-cn.com/problems/search-in-rotated-sorted-array) |  搜索旋转排序数组 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0033.search-in-rotated-sorted-array%20搜索旋转排序数组/main.go) | 
 [0034.find-first-and-last-position-of-element-in-sorted-array](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array) |  在排序数组中查找元素的第一个和最后一个位置 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0034.find-first-and-last-position-of-element-in-sorted-array%20在排序数组中查找元素的第一个和最后一个位置/main.go) | 
 [0035.search-insert-position     ](https://leetcode-cn.com/problems/search-insert-position) |  搜索插入位置 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0035.search-insert-position%20搜索插入位置/main.go) | 
 [0069.sqrtx                      ](https://leetcode-cn.com/problems/sqrtx) |  x 的平方根 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0069.sqrtx%20x%20的平方根/main.go) | 
 [0081.search-in-rotated-sorted-array-ii](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii) |  搜索旋转排序数组 II | Medium | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0081.search-in-rotated-sorted-array-ii%20搜索旋转排序数组%20II/main.go) | 
 [0153.find-minimum-in-rotated-sorted-array](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array) |  寻找旋转排序数组中的最小值 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0153.find-minimum-in-rotated-sorted-array%20寻找旋转排序数组中的最小值/main.go) | 
 [0154.find-minimum-in-rotated-sorted-array-ii](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii) |  寻找旋转排序数组中的最小值 II | Hard | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0154.find-minimum-in-rotated-sorted-array-ii%20寻找旋转排序数组中的最小值%20II/main.go) | 
 [0162.find-peak-element          ](https://leetcode-cn.com/problems/find-peak-element) |  寻找峰值 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0162.find-peak-element%20寻找峰值/main.go) | 
 [0167.two-sum-ii-input-array-is-sorted](https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted) |  两数之和 II - 输入有序数组 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0167.two-sum-ii-input-array-is-sorted%20两数之和%20II%20-%20输入有序数组/main.go) | 
 [0278.first-bad-version          ](https://leetcode-cn.com/problems/first-bad-version) |  第一个错误版本 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0278.first-bad-version%20第一个错误版本/main.go) | 
 [0367.valid-perfect-square       ](https://leetcode-cn.com/problems/valid-perfect-square) |  有效的完全平方数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0367.valid-perfect-square%20有效的完全平方数/main.go) | 
 [0374.guess-number-higher-or-lower](https://leetcode-cn.com/problems/guess-number-higher-or-lower) |  猜数字大小 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0374.guess-number-higher-or-lower%20猜数字大小/main.go) | 
 [0441.arranging-coins            ](https://leetcode-cn.com/problems/arranging-coins) |  排列硬币 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0441.arranging-coins%20排列硬币/main.go) | 
 [0701.insert-into-a-binary-search-tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree) |  二叉搜索树中的插入操作 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0701.insert-into-a-binary-search-tree%20二叉搜索树中的插入操作/main.go) | 
 [0744.find-smallest-letter-greater-than-target](https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target) |  寻找比目标字母大的最小字母 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0744.find-smallest-letter-greater-than-target%20寻找比目标字母大的最小字母/main.go) | 
 [0852.peak-index-in-a-mountain-array](https://leetcode-cn.com/problems/peak-index-in-a-mountain-array) |  山脉数组的峰顶索引 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/0852.peak-index-in-a-mountain-array%20山脉数组的峰顶索引/main.go) | 
 [1064.fixed-point                ](https://leetcode-cn.com/problems/fixed-point) |  不动点 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/1064.fixed-point%20不动点/main.go) | 
 [1150.check-if-a-number-is-majority-element-in-a-sorted-array](https://leetcode-cn.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array) |  检查一个数是否在数组中占绝大多数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/1150.check-if-a-number-is-majority-element-in-a-sorted-array%20检查一个数是否在数组中占绝大多数/main.go) | 
 [1385.find-the-distance-value-between-two-arrays](https://leetcode-cn.com/problems/find-the-distance-value-between-two-arrays) |  两个数组间的距离值 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/1385.find-the-distance-value-between-two-arrays%20两个数组间的距离值/main.go) | 
 [1539.kth-missing-positive-number](https://leetcode-cn.com/problems/kth-missing-positive-number) |  第 k 个缺失的正整数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/1539.kth-missing-positive-number%20第%20k%20个缺失的正整数/main.go) | 
 [1608.special-array-with-x-elements-greater-than-or-equal-x](https://leetcode-cn.com/problems/special-array-with-x-elements-greater-than-or-equal-x) |  特殊数组的特征值 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/1608.special-array-with-x-elements-greater-than-or-equal-x%20特殊数组的特征值/main.go) | 
 [LCP                             ](https://leetcode-cn.com/problems/4sum) |  18. 早餐组合 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/LCP%2018.%20早餐组合/main.go) | 
 [剑指                              ](https://leetcode-cn.com/problems/container-with-most-water) |  Offer 11. 旋转数组的最小数字 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/剑指%20Offer%2011.%20旋转数组的最小数字/main.go) | 
 [剑指                              ](https://leetcode-cn.com/problems/maximum-subarray) |  Offer 53 - II. 0～n-1中缺失的数字 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/剑指%20Offer%2053%20-%20II.%200～n-1中缺失的数字/main.go) | 
 [面试题                             ](https://leetcode-cn.com/problems/string-to-integer-atoi) |  08.03. 魔术索引 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/面试题%2008.03.%20魔术索引/main.go) | 
 [面试题                             ](https://leetcode-cn.com/problems/regular-expression-matching) |  10.05. 稀疏数组搜索 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/binary_search/面试题%2010.05.%20稀疏数组搜索/main.go) | 

## Catalan
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0022.generate-parentheses       ](https://leetcode-cn.com/problems/generate-parentheses) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/catalan/0022.generate-parentheses/main.go) | 

## StackMonotonic
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0042.trapping-rain-water        ](https://leetcode-cn.com/problems/trapping-rain-water) |  接雨水 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/stack_monotonic/0042.trapping-rain-water%20接雨水/main.go) | 

## SortCounter
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------

## Hash
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0380.insert-delete-getrandom-o1 ](https://leetcode-cn.com/problems/insert-delete-getrandom-o1) |  O(1) 时间插入、删除和获取随机元素 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Array)%200380.insert-delete-getrandom-o1%20O(1)%20时间插入、删除和获取随机元素/main.go) | Array
 [0041.first-missing-positive     ](https://leetcode-cn.com/problems/first-missing-positive) |  缺失的第一个正数 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Bit)%200041.first-missing-positive%20缺失的第一个正数/main.go) | Bit
 [0442.find-all-duplicates-in-an-array](https://leetcode-cn.com/problems/find-all-duplicates-in-an-array) |  数组中重复的数据 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Bit)%200442.find-all-duplicates-in-an-array%20数组中重复的数据/main.go) | Bit
 [0448.find-all-numbers-disappeared-in-an-array](https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array) |  找到所有数组中消失的数字 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Bit)%200448.find-all-numbers-disappeared-in-an-array%20找到所有数组中消失的数字/main.go) | Bit
 [0205.isomorphic-strings         ](https://leetcode-cn.com/problems/isomorphic-strings) |  同构字符串 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Double)%200205.isomorphic-strings%20同构字符串/main.go) | Double
 [0290.word-pattern               ](https://leetcode-cn.com/problems/word-pattern) |  单词规律 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Double)%200290.word-pattern%20单词规律/main.go) | Double
 [0734.sentence-similarity        ](https://leetcode-cn.com/problems/sentence-similarity) |  句子相似性 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Double)%200734.sentence-similarity%20句子相似性/main.go) | Double
 [0387.first-unique-character-in-a-string](https://leetcode-cn.com/problems/first-unique-character-in-a-string) |  字符串中的第一个唯一字符 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark%20Bit%20Order)%200387.first-unique-character-in-a-string%20字符串中的第一个唯一字符/main.go) | Mark Bit Order
 [0389.find-the-difference        ](https://leetcode-cn.com/problems/find-the-difference) |  找不同 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark%20Count%20Bit)%200389.find-the-difference%20找不同/main.go) | Mark Count Bit
 [0217.contains-duplicate         ](https://leetcode-cn.com/problems/contains-duplicate) |  出现重复 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200217.contains-duplicate%20出现重复/main.go) | Mark
 [0219.contains-duplicate-ii      ](https://leetcode-cn.com/problems/contains-duplicate-ii) |  两个重复元素距离小于等于K | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200219.contains-duplicate-ii%20两个重复元素距离小于等于K/main.go) | Mark
 [0242.valid-anagram              ](https://leetcode-cn.com/problems/valid-anagram) |  有效的字母异位词 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200242.valid-anagram%20有效的字母异位词/main.go) | Mark
 [0246.strobogrammatic-number     ](https://leetcode-cn.com/problems/strobogrammatic-number) |  中心对称数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200246.strobogrammatic-number%20中心对称数/main.go) | Mark
 [0299.bulls-and-cows             ](https://leetcode-cn.com/problems/bulls-and-cows) |  猜数字游戏 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200299.bulls-and-cows%20猜数字游戏/main.go) | Mark
 [0348.design-tic-tac-toe         ](https://leetcode-cn.com/problems/design-tic-tac-toe) |  设计井字棋 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200348.design-tic-tac-toe%20设计井字棋/main.go) | Mark
 [0349.intersection-of-two-arrays ](https://leetcode-cn.com/problems/intersection-of-two-arrays) |  两个数组的交集 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200349.intersection-of-two-arrays%20两个数组的交集/main.go) | Mark
 [0350.intersection-of-two-arrays-ii](https://leetcode-cn.com/problems/intersection-of-two-arrays-ii) |  两个数组的交集 II | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200350.intersection-of-two-arrays-ii%20两个数组的交集%20II/main.go) | Mark
 [0359.logger-rate-limiter        ](https://leetcode-cn.com/problems/logger-rate-limiter) |  日志速率限制器 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200359.logger-rate-limiter%20日志速率限制器/main.go) | Mark
 [0383.ransom-note                ](https://leetcode-cn.com/problems/ransom-note) |  赎金信 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200383.ransom-note%20赎金信/main.go) | Mark
 [0409.longest-palindrome         ](https://leetcode-cn.com/problems/longest-palindrome) |  最长回文串 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200409.longest-palindrome%20最长回文串/main.go) | Mark
 [0500.keyboard-row               ](https://leetcode-cn.com/problems/keyboard-row) |  只返回键盘同一行的字母 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200500.keyboard-row%20只返回键盘同一行的字母/main.go) | Mark
 [0575.distribute-candies         ](https://leetcode-cn.com/problems/distribute-candies) |  分糖果 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200575.distribute-candies%20分糖果/main.go) | Mark
 [0594.longest-harmonious-subsequence](https://leetcode-cn.com/problems/longest-harmonious-subsequence) |  最长和谐(相差1)子序列 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200594.longest-harmonious-subsequence%20最长和谐(相差1)子序列/main.go) | Mark
 [0599.minimum-index-sum-of-two-lists](https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists) |  相同值的最小索引总和 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Mark)%200599.minimum-index-sum-of-two-lists%20相同值的最小索引总和/main.go) | Mark
 [0356.line-reflection            ](https://leetcode-cn.com/problems/line-reflection) |  直线镜像 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Math)%200356.line-reflection%20直线镜像/main.go) | Math
 [0169.majority-element           ](https://leetcode-cn.com/problems/majority-element) |  众数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Moore%20投票)%200169.majority-element%20众数/main.go) | Moore 投票
 [0229.majority-element-ii        ](https://leetcode-cn.com/problems/majority-element-ii) |  求众数 II | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Moore%20投票)%200229.majority-element-ii%20求众数%20II/main.go) | Moore 投票
 [0128.longest-consecutive-sequence](https://leetcode-cn.com/problems/longest-consecutive-sequence) |  最长连续序列 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Order)%200128.longest-consecutive-sequence%20最长连续序列/main.go) | Order
 [0362.design-hit-counter         ](https://leetcode-cn.com/problems/design-hit-counter) |  敲击计数器 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Queue)%200362.design-hit-counter%20敲击计数器/main.go) | Queue
 [0720.longest-word-in-dictionary ](https://leetcode-cn.com/problems/longest-word-in-dictionary) |  词典中最长的单词 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Trie)%200720.longest-word-in-dictionary%20词典中最长的单词/main.go) | Trie
 [0049.group-anagrams             ](https://leetcode-cn.com/problems/group-anagrams) |  字母异位词分组 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Vector)%200049.group-anagrams%20字母异位词分组/main.go) | Vector
 [0438.find-all-anagrams-in-a-string](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string) |  找到字符串中所有字母异位词 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Vector)%200438.find-all-anagrams-in-a-string%20找到字符串中所有字母异位词/main.go) | Vector
 [0395.longest-substring-with-at-least-k-repeating-characters](https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters) |  至少包含 K 个重复字符的最长子串 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Window%20Divide)%200395.longest-substring-with-at-least-k-repeating-characters%20至少包含%20K%20个重复字符的最长子串/main.go) | Window Divide
 [0003.longest-substring-without-repeating-characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters) |  无重复的最长子串 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Window)%200003.longest-substring-without-repeating-characters%20无重复的最长子串/main.go) | Window
 [0325.maximum-size-subarray-sum-equals-k](https://leetcode-cn.com/problems/maximum-size-subarray-sum-equals-k) |  和等于 k 的最长子数组长度 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Window)%200325.maximum-size-subarray-sum-equals-k%20和等于%20k%20的最长子数组长度/main.go) | Window
 [0340.longest-substring-with-at-most-k-distinct-characters](https://leetcode-cn.com/problems/longest-substring-with-at-most-k-distinct-characters) |  至多包含 K 个不同字符的最长子串 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Window)%200340.longest-substring-with-at-most-k-distinct-characters%20至多包含%20K%20个不同字符的最长子串/main.go) | Window
 [0697.degree-of-an-array         ](https://leetcode-cn.com/problems/degree-of-an-array) |  度相同的最短连续子数组 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/(Window)%200697.degree-of-an-array%20度相同的最短连续子数组/main.go) | Window
 [0146.lru-cache                  ](https://leetcode-cn.com/problems/lru-cache) |  最近最少使用 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/(lru)%200146.lru-cache%20最近最少使用/main.go) | lru
 [0748.shortest-completing-word   ](https://leetcode-cn.com/problems/shortest-completing-word) |  最短补全词 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0748.shortest-completing-word%20最短补全词/main.go) | 
 [0760.find-anagram-mappings      ](https://leetcode-cn.com/problems/find-anagram-mappings) |  找A在B的位置 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0760.find-anagram-mappings%20找A在B的位置/main.go) | 
 [0771.jewels-and-stones          ](https://leetcode-cn.com/problems/jewels-and-stones) |  找J在S中的数量 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0771.jewels-and-stones%20找J在S中的数量/main.go) | 
 [0811.subdomain-visit-count      ](https://leetcode-cn.com/problems/subdomain-visit-count) |  子域名访问次数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/0811.subdomain-visit-count%20子域名访问次数/main.go) | 
 [0819.most-common-word           ](https://leetcode-cn.com/problems/most-common-word) |  最常见的单词 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0819.most-common-word%20最常见的单词/main.go) | 
 [0859.buddy-strings              ](https://leetcode-cn.com/problems/buddy-strings) |  亲密字符串 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0859.buddy-strings%20亲密字符串/main.go) | 
 [0884.uncommon-words-from-two-sentences](https://leetcode-cn.com/problems/uncommon-words-from-two-sentences) |  两句话中的不常见单词 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0884.uncommon-words-from-two-sentences%20两句话中的不常见单词/main.go) | 
 [0888.fair-candy-swap            ](https://leetcode-cn.com/problems/fair-candy-swap) |  公平的糖果棒交换 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0888.fair-candy-swap%20公平的糖果棒交换/main.go) | 
 [0914.x-of-a-kind-in-a-deck-of-cards](https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards) |  卡牌分组 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0914.x-of-a-kind-in-a-deck-of-cards%20卡牌分组/main.go) | 
 [0929.unique-email-addresses     ](https://leetcode-cn.com/problems/unique-email-addresses) |  独特的电子邮件地址 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0929.unique-email-addresses%20独特的电子邮件地址/main.go) | 
 [0930.binary-subarrays-with-sum  ](https://leetcode-cn.com/problems/binary-subarrays-with-sum) |  和相同的二元子数组 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/0930.binary-subarrays-with-sum%20和相同的二元子数组/main.go) | 
 [0953.verifying-an-alien-dictionary](https://leetcode-cn.com/problems/verifying-an-alien-dictionary) |  验证外星语词典 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0953.verifying-an-alien-dictionary%20验证外星语词典/main.go) | 
 [0961.n-repeated-element-in-size-2n-array](https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array) |  重复 N 次的元素 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0961.n-repeated-element-in-size-2n-array%20重复%20N%20次的元素/main.go) | 
 [0997.find-the-town-judge        ](https://leetcode-cn.com/problems/find-the-town-judge) |  找到小镇的法官 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/0997.find-the-town-judge%20找到小镇的法官/main.go) | 
 [1002.find-common-characters     ](https://leetcode-cn.com/problems/find-common-characters) |  查找常用字符 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1002.find-common-characters%20查找常用字符/main.go) | 
 [1086.high-five                  ](https://leetcode-cn.com/problems/high-five) |  前五科的均分 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1086.high-five%20前五科的均分/main.go) | 
 [1119.remove-vowels-from-a-string](https://leetcode-cn.com/problems/remove-vowels-from-a-string) |  删去字符串中的元音 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1119.remove-vowels-from-a-string%20删去字符串中的元音/main.go) | 
 [1122.relative-sort-array        ](https://leetcode-cn.com/problems/relative-sort-array) |  数组的相对排序 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1122.relative-sort-array%20数组的相对排序/main.go) | 
 [1128.number-of-equivalent-domino-pairs](https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs) |  等价多米诺骨牌对的数量 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1128.number-of-equivalent-domino-pairs%20等价多米诺骨牌对的数量/main.go) | 
 [1133.largest-unique-number      ](https://leetcode-cn.com/problems/largest-unique-number) |  最大唯一数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1133.largest-unique-number%20最大唯一数/main.go) | 
 [1160.find-words-that-can-be-formed-by-characters](https://leetcode-cn.com/problems/find-words-that-can-be-formed-by-characters) |  拼写单词 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1160.find-words-that-can-be-formed-by-characters%20拼写单词/main.go) | 
 [1165.single-row-keyboard        ](https://leetcode-cn.com/problems/single-row-keyboard) |  机械手最少移动次数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1165.single-row-keyboard%20机械手最少移动次数/main.go) | 
 [1189.maximum-number-of-balloons ](https://leetcode-cn.com/problems/maximum-number-of-balloons) |  “气球” 的最大数量 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1189.maximum-number-of-balloons%20“气球”%20的最大数量/main.go) | 
 [1275.find-winner-on-a-tic-tac-toe-game](https://leetcode-cn.com/problems/find-winner-on-a-tic-tac-toe-game) |  找出井字棋的获胜者 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1275.find-winner-on-a-tic-tac-toe-game%20找出井字棋的获胜者/main.go) | 
 [1331.rank-transform-of-an-array ](https://leetcode-cn.com/problems/rank-transform-of-an-array) |  数组序号转换 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1331.rank-transform-of-an-array%20数组序号转换/main.go) | 
 [1346.check-if-n-and-its-double-exist](https://leetcode-cn.com/problems/check-if-n-and-its-double-exist) |  检查整数及其两倍数是否存在 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1346.check-if-n-and-its-double-exist%20检查整数及其两倍数是否存在/main.go) | 
 [1370.increasing-decreasing-string](https://leetcode-cn.com/problems/increasing-decreasing-string) |  上升下降字符串 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1370.increasing-decreasing-string%20上升下降字符串/main.go) | 
 [1394.find-lucky-integer-in-an-array](https://leetcode-cn.com/problems/find-lucky-integer-in-an-array) |  找出数组中的幸运数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1394.find-lucky-integer-in-an-array%20找出数组中的幸运数/main.go) | 
 [1399.count-largest-group        ](https://leetcode-cn.com/problems/count-largest-group) |  统计最大组的数目 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1399.count-largest-group%20统计最大组的数目/main.go) | 
 [1426.counting-elements          ](https://leetcode-cn.com/problems/counting-elements) |  数元素 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1426.counting-elements%20数元素/main.go) | 
 [1436.destination-city           ](https://leetcode-cn.com/problems/destination-city) |  旅行终点站 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1436.destination-city%20旅行终点站/main.go) | 
 [1460.make-two-arrays-equal-by-reversing-sub-arrays](https://leetcode-cn.com/problems/make-two-arrays-equal-by-reversing-sub-arrays) |  通过翻转子数组使两个数组相等 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1460.make-two-arrays-equal-by-reversing-sub-arrays%20通过翻转子数组使两个数组相等/main.go) | 
 [1512.number-of-good-pairs       ](https://leetcode-cn.com/problems/number-of-good-pairs) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1512.number-of-good-pairs/main.go) | 
 [1536.minimum-swaps-to-arrange-a-binary-grid](https://leetcode-cn.com/problems/minimum-swaps-to-arrange-a-binary-grid) |  按照频率将数组升序排序 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/1536.minimum-swaps-to-arrange-a-binary-grid%20按照频率将数组升序排序/main.go) | 
 [1570.dot-product-of-two-sparse-vectors](https://leetcode-cn.com/problems/dot-product-of-two-sparse-vectors) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/1570.dot-product-of-two-sparse-vectors/main.go) | 
 [1603.design-parking-system      ](https://leetcode-cn.com/problems/design-parking-system) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1603.design-parking-system/main.go) | 
 [1624.largest-substring-between-two-equal-characters](https://leetcode-cn.com/problems/largest-substring-between-two-equal-characters) |  两个相同字符之间的最长子字符串 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1624.largest-substring-between-two-equal-characters%20两个相同字符之间的最长子字符串/main.go) | 
 [1640.check-array-formation-through-concatenation](https://leetcode-cn.com/problems/check-array-formation-through-concatenation) |  能否连接形成数组 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1640.check-array-formation-through-concatenation%20能否连接形成数组/main.go) | 
 [1656.design-an-ordered-stream   ](https://leetcode-cn.com/problems/design-an-ordered-stream) |  设计有序流 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1656.design-an-ordered-stream%20设计有序流/main.go) | 
 [1711.count-good-meals           ](https://leetcode-cn.com/problems/count-good-meals) |  大餐计数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/1711.count-good-meals%20大餐计数/main.go) | 
 [1742.maximum-number-of-balls-in-a-box](https://leetcode-cn.com/problems/maximum-number-of-balls-in-a-box) |  盒子中小球的最大数量 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1742.maximum-number-of-balls-in-a-box%20盒子中小球的最大数量/main.go) | 
 [1748.sum-of-unique-elements     ](https://leetcode-cn.com/problems/sum-of-unique-elements) |  唯一元素的和 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1748.sum-of-unique-elements%20唯一元素的和/main.go) | 
 [1763.longest-nice-substring     ](https://leetcode-cn.com/problems/longest-nice-substring) |  最长的美好子字符串 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1763.longest-nice-substring%20最长的美好子字符串/main.go) | 
 [1796.second-largest-digit-in-a-string](https://leetcode-cn.com/problems/second-largest-digit-in-a-string) |  字符串中第二大的数字 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1796.second-largest-digit-in-a-string%20字符串中第二大的数字/main.go) | 
 [1805.number-of-different-integers-in-a-string](https://leetcode-cn.com/problems/number-of-different-integers-in-a-string) |  字符串中不同整数的数目 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1805.number-of-different-integers-in-a-string%20字符串中不同整数的数目/main.go) | 
 [1935.maximum-number-of-words-you-can-type](https://leetcode-cn.com/problems/maximum-number-of-words-you-can-type) |  可以输入的最大单词数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1935.maximum-number-of-words-you-can-type%20可以输入的最大单词数/main.go) | 
 [1941.check-if-all-characters-have-equal-number-of-occurrences](https://leetcode-cn.com/problems/check-if-all-characters-have-equal-number-of-occurrences) |  检查是否所有字符出现次数相同 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/1941.check-if-all-characters-have-equal-number-of-occurrences%20检查是否所有字符出现次数相同/main.go) | 
 [2006.count-number-of-pairs-with-absolute-difference-k](https://leetcode-cn.com/problems/count-number-of-pairs-with-absolute-difference-k) |  差的绝对值为 K 的数对数目 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/2006.count-number-of-pairs-with-absolute-difference-k%20差的绝对值为%20K%20的数对数目/main.go) | 
 [2007.find-original-array-from-doubled-array](https://leetcode-cn.com/problems/find-original-array-from-doubled-array) |  从双倍数组中还原原数组 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/2007.find-original-array-from-doubled-array%20从双倍数组中还原原数组/main.go) | 
 [5859.count-number-of-pairs-with-absolute-difference-k 差的绝对值为 K 的数对数目](https://leetcode-cn.com/problems/5859.count-number-of-pairs-with-absolute-difference-k 差的绝对值为 K 的数对数目) |  | Hard | [Go](https://github.com/bygo/leetcode/blob/master/hash/5859.count-number-of-pairs-with-absolute-difference-k%20差的绝对值为%20K%20的数对数目/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/hash/5860.find-original-array-from-doubled-array%20从双倍数组中还原原数组/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/hash/5877.detect-squares%20检测正方形/main.go) |  
 [LCP                             ](https://leetcode-cn.com/problems/container-with-most-water) |  11. 期望个数统计 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/LCP%2011.%20期望个数统计/main.go) | 
 [LCP.07.                         ](https://leetcode-cn.com/problems/reverse-integer) |  传递信息 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/LCP.07.%20传递信息/main.go) | 
 [LCS                             ](https://leetcode-cn.com/problems/add-two-numbers) |  02. 完成一半题目 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/LCS%2002.%20完成一半题目/main.go) | 
 [vip                             ](https://leetcode-cn.com/problems/vip) |  | Hard | [Go](https://github.com/bygo/leetcode/blob/master/hash/vip/0159.longest-substring-with-at-most-two-distinct-characters%20至多包含两个不同字符的最长子串/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/hash/vip/0244.shortest-word-distance-ii%20最短单词距离%20II/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/hash/vip/0249.group-shifted-strings%20移位字符串分组/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/hash/vip/0266.palindrome-permutation%20回文排列/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/hash/vip/0288.unique-word-abbreviation%20单词的唯一缩写/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/hash/vip/0604.design-compressed-string-iterator%20迭代压缩字符串/main.go) |  
 [剑指                              ](https://leetcode-cn.com/problems/powx-n) |  Offer 50. 第一个只出现一次的字符 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/剑指%20Offer%2050.%20第一个只出现一次的字符/main.go) | 
 [剑指                              ](https://leetcode-cn.com/problems/longest-valid-parentheses) |  Offer II 032. 有效的变位词 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/hash/剑指%20Offer%20II%20032.%20有效的变位词/main.go) | 
 [剑指                              ](https://leetcode-cn.com/problems/sort-colors) |  Offer II 075. 数组相对排序 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/剑指%20Offer%20II%20075.%20数组相对排序/main.go) | 
 [剑指                              ](https://leetcode-cn.com/problems/pascals-triangle-ii) |  Offer II 119. 最长连续序列 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/剑指%20Offer%20II%20119.%20最长连续序列/main.go) | 
 [面试题                             ](https://leetcode-cn.com/problems/two-sum) |  01.01. 判定字符是否唯一 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/hash/面试题%2001.01.%20判定字符是否唯一/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/hash/面试题%2001.02.%20判定是否互为字符重排/main.go) |  
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/hash/面试题%2001.04.%20回文排列/main.go) |  
 [面试题                             ](https://leetcode-cn.com/problems/3sum-closest) |  16.15. 珠玑妙算 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/面试题%2016.15.%20珠玑妙算/main.go) | 
 [面试题                             ](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number) |  17.04. 消失的数字 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/hash/面试题%2017.04.%20消失的数字/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/hash/面试题%2017.10.%20主要元素%20(Moore%20投票)/main.go) |  

## Greedy
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0045.jump-game-ii               ](https://leetcode-cn.com/problems/jump-game-ii) |  跳跃游戏 II | Medium | [Go](https://github.com/bygo/leetcode/blob/master/greedy/0045.jump-game-ii%20跳跃游戏%20II/main.go) | 
 [0055.jump-game                  ](https://leetcode-cn.com/problems/jump-game) |  跳跃游戏 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/greedy/0055.jump-game%20跳跃游戏/main.go) | 
 [0678.valid-parenthesis-string   ](https://leetcode-cn.com/problems/valid-parenthesis-string) |  有效的括号字符串 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/greedy/0678.valid-parenthesis-string%20有效的括号字符串/main.go) | 

## Dp
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0005.longest-palindromic-substring](https://leetcode-cn.com/problems/longest-palindromic-substring) |  最长回文 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0005.longest-palindromic-substring%20最长回文/main.go) | 
 [0010.regular-expression-matching](https://leetcode-cn.com/problems/regular-expression-matching) |  .*正则匹配 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0010.regular-expression-matching%20.*正则匹配/main.go) | 
 [0032.longest-valid-parentheses  ](https://leetcode-cn.com/problems/longest-valid-parentheses) |  最长有效括号 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0032.longest-valid-parentheses%20最长有效括号/main.go) | 
 [0042.trapping-rain-water        ](https://leetcode-cn.com/problems/trapping-rain-water) |  接雨水 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0042.trapping-rain-water%20接雨水/main.go) | 
 [0044.wildcard-matching          ](https://leetcode-cn.com/problems/wildcard-matching) |  *？通配符 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0044.wildcard-matching%20*？通配符/main.go) | 
 [0053.maximum-subarray           ](https://leetcode-cn.com/problems/maximum-subarray) |  最大子序和 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0053.maximum-subarray%20最大子序和/main.go) | 
 [0062.unique-paths               ](https://leetcode-cn.com/problems/unique-paths) |  不同路径 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0062.unique-paths%20不同路径/main.go) | 
 [0063.unique-paths-ii            ](https://leetcode-cn.com/problems/unique-paths-ii) |  不同路径 II | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0063.unique-paths-ii%20不同路径%20II/main.go) | 
 [0064.minimum-path-sum           ](https://leetcode-cn.com/problems/minimum-path-sum) |  最小路径和 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0064.minimum-path-sum%20最小路径和/main.go) | 
 [0070.climbing-stairs            ](https://leetcode-cn.com/problems/climbing-stairs) |  爬楼梯 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0070.climbing-stairs%20爬楼梯/main.go) | 
 [0072.edit-distance              ](https://leetcode-cn.com/problems/edit-distance) |  编辑距离 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0072.edit-distance%20编辑距离/main.go) | 
 [0087.scramble-string            ](https://leetcode-cn.com/problems/scramble-string) |  扰乱字符串 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0087.scramble-string%20扰乱字符串/main.go) | 
 [0091.decode-ways                ](https://leetcode-cn.com/problems/decode-ways) |  解码方法 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0091.decode-ways%20解码方法/main.go) | 
 [0097.interleaving-string        ](https://leetcode-cn.com/problems/interleaving-string) |  交错字符串 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0097.interleaving-string%20交错字符串/main.go) | 
 [0115.distinct-subsequences      ](https://leetcode-cn.com/problems/distinct-subsequences) |  不同的子序列 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0115.distinct-subsequences%20不同的子序列/main.go) | 
 [0118.pascals-triangle           ](https://leetcode-cn.com/problems/pascals-triangle) |  杨辉三角 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0118.pascals-triangle%20杨辉三角/main.go) | 
 [0119.pascals-triangle-ii        ](https://leetcode-cn.com/problems/pascals-triangle-ii) |  杨辉三角 II | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0119.pascals-triangle-ii%20杨辉三角%20II/main.go) | 
 [0120.triangle                   ](https://leetcode-cn.com/problems/triangle) |  三角形最小路径 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0120.triangle%20三角形最小路径/main.go) | 
 [0121.best-time-to-buy-and-sell-stock](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock) |  买卖股票的最佳时机 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0121.best-time-to-buy-and-sell-stock%20买卖股票的最佳时机/main.go) | 
 [0123.best-time-to-buy-and-sell-stock-iii](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii) |  买卖股票的最佳时机 III | Hard | [Go](https://github.com/bygo/leetcode/blob/master/dp/0123.best-time-to-buy-and-sell-stock-iii%20买卖股票的最佳时机%20III/main.go) | 
 [0139.word-break                 ](https://leetcode-cn.com/problems/word-break) |  单词拆分 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0139.word-break%20单词拆分/main.go) | 
 [0279.perfect-squares            ](https://leetcode-cn.com/problems/perfect-squares) |  完全平方数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0279.perfect-squares%20完全平方数/main.go) | 
 [0303.range-sum-query-immutable  ](https://leetcode-cn.com/problems/range-sum-query-immutable) |  区域和检索 - 数组不可变 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0303.range-sum-query-immutable%20区域和检索%20-%20数组不可变/main.go) | 
 [0322.coin-change                ](https://leetcode-cn.com/problems/coin-change) |  零钱兑换 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0322.coin-change%20零钱兑换/main.go) | 
 [0474.ones-and-zeroes            ](https://leetcode-cn.com/problems/ones-and-zeroes) |  一和零 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0474.ones-and-zeroes%20一和零/main.go) | 
 [0509.fibonacci-number           ](https://leetcode-cn.com/problems/fibonacci-number) |  斐波那契数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/0509.fibonacci-number%20斐波那契数/main.go) | 
 [0518.coin-change-2              ](https://leetcode-cn.com/problems/coin-change-2) |  零钱兑换 II | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0518.coin-change-2%20零钱兑换%20II/main.go) | 
 [0523.continuous-subarray-sum    ](https://leetcode-cn.com/problems/continuous-subarray-sum) |  连续子数组和 K的倍数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0523.continuous-subarray-sum%20连续子数组和%20K的倍数/main.go) | 
 [0525.contiguous-array           ](https://leetcode-cn.com/problems/contiguous-array) |  连续数组 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0525.contiguous-array%20连续数组/main.go) | 
 [0650.2-keys-keyboard            ](https://leetcode-cn.com/problems/2-keys-keyboard) |  只有两个键的键盘 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0650.2-keys-keyboard%20只有两个键的键盘/main.go) | 
 [0877.stone-game                 ](https://leetcode-cn.com/problems/stone-game) |  石头游戏 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/0877.stone-game%20石头游戏/main.go) | 
 [1480.running-sum-of-1d-array    ](https://leetcode-cn.com/problems/running-sum-of-1d-array) |  数组前缀和 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/dp/1480.running-sum-of-1d-array%20数组前缀和/main.go) | 
 [1744.can-you-eat-your-favorite-candy-on-your-favorite-day](https://leetcode-cn.com/problems/can-you-eat-your-favorite-candy-on-your-favorite-day) |  你能在你最喜欢的那天吃到你最喜欢的糖果吗？ | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/1744.can-you-eat-your-favorite-candy-on-your-favorite-day%20你能在你最喜欢的那天吃到你最喜欢的糖果吗？/main.go) | 
 [1769.minimum-number-of-operations-to-move-all-balls-to-each-box](https://leetcode-cn.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box) |  移动所有球到每个盒子所需的最小操作数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/dp/1769.minimum-number-of-operations-to-move-all-balls-to-each-box%20移动所有球到每个盒子所需的最小操作数/main.go) | 
 [5876.sum-of-beauty-in-the-array 数组美丽值求和](https://leetcode-cn.com/problems/5876.sum-of-beauty-in-the-array 数组美丽值求和) |  | Hard | [Go](https://github.com/bygo/leetcode/blob/master/dp/5876.sum-of-beauty-in-the-array%20数组美丽值求和/main.go) | 
 [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/bygo/leetcode/blob/master/dp/5881.maximum-difference-between-increasing-elements%20增量元素之间的最大差值/main.go) |  

## Classic
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0005.longest-palindromic-substring](https://leetcode-cn.com/problems/longest-palindromic-substring) |  最长回文子串 (manacher) | Medium | [Go](https://github.com/bygo/leetcode/blob/master/classic/0005.longest-palindromic-substring%20最长回文子串%20(manacher)/main.go) | 
 [0028.implement-strstr           ](https://leetcode-cn.com/problems/implement-strstr) |  实现 strStr() (kmp) | Easy | [Go](https://github.com/bygo/leetcode/blob/master/classic/0028.implement-strstr%20实现%20strStr()%20(kmp)/main.go) | 

## Math
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0007.reverse-integer            ](https://leetcode-cn.com/problems/reverse-integer) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0007.reverse-integer/main.go) | 
 [0008.string-to-integer-atoi     ](https://leetcode-cn.com/problems/string-to-integer-atoi) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0008.string-to-integer-atoi/main.go) | 
 [0009.palindrome-number          ](https://leetcode-cn.com/problems/palindrome-number) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0009.palindrome-number/main.go) | 
 [0012.integer-to-roman           ](https://leetcode-cn.com/problems/integer-to-roman) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0012.integer-to-roman/main.go) | 
 [0013.roman-to-integer           ](https://leetcode-cn.com/problems/roman-to-integer) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0013.roman-to-integer/main.go) | 
 [0029.divide-two-integers        ](https://leetcode-cn.com/problems/divide-two-integers) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0029.divide-two-integers/main.go) | 
 [0031.next-permutation           ](https://leetcode-cn.com/problems/next-permutation) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0031.next-permutation/main.go) | 
 [0043.multiply-strings           ](https://leetcode-cn.com/problems/multiply-strings) |  字符串相乘 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0043.multiply-strings%20字符串相乘/main.go) | 
 [0066.plus-one                   ](https://leetcode-cn.com/problems/plus-one) |  加一 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0066.plus-one%20加一/main.go) | 
 [0165.compare-version-numbers    ](https://leetcode-cn.com/problems/compare-version-numbers) |  比较版本号 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0165.compare-version-numbers%20比较版本号/main.go) | 
 [0166.fraction-to-recurring-decimal](https://leetcode-cn.com/problems/fraction-to-recurring-decimal) |  分数到小数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0166.fraction-to-recurring-decimal%20分数到小数/main.go) | 
 [0191.number-of-1-bits           ](https://leetcode-cn.com/problems/number-of-1-bits) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0191.number-of-1-bits/main.go) | 
 [0263.ugly-number                ](https://leetcode-cn.com/problems/ugly-number) |  丑数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0263.ugly-number%20丑数/main.go) | 
 [0268.missing-number             ](https://leetcode-cn.com/problems/missing-number) |  丢失的数字 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0268.missing-number%20丢失的数字/main.go) | 
 [0292.nim-game                   ](https://leetcode-cn.com/problems/nim-game) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0292.nim-game/main.go) | 
 [0319.bulb-switcher              ](https://leetcode-cn.com/problems/bulb-switcher) |  | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0319.bulb-switcher/main.go) | 
 [0415.add-strings                ](https://leetcode-cn.com/problems/add-strings) |  字符串相加 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0415.add-strings%20字符串相加/main.go) | 
 [0453.minimum-moves-to-equal-array-elements](https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements) |  最小操作次数使数组元素相等 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0453.minimum-moves-to-equal-array-elements%20最小操作次数使数组元素相等/main.go) | 
 [0470.implement-rand10-using-rand7](https://leetcode-cn.com/problems/implement-rand10-using-rand7) |  用 Rand7() 实现 Rand10() | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0470.implement-rand10-using-rand7%20用%20Rand7()%20实现%20Rand10()/main.go) | 
 [0478.generate-random-point-in-a-circle](https://leetcode-cn.com/problems/generate-random-point-in-a-circle) |  在圆内随机生成点 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0478.generate-random-point-in-a-circle%20在圆内随机生成点/main.go) | 
 [0492.construct-the-rectangle    ](https://leetcode-cn.com/problems/construct-the-rectangle) |  构造矩形 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/0492.construct-the-rectangle%20构造矩形/main.go) | 
 [0877.stone-game                 ](https://leetcode-cn.com/problems/stone-game) |  石头游戏 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/math/0877.stone-game%20石头游戏/main.go) | 
 [1025.divisor-game               ](https://leetcode-cn.com/problems/divisor-game) |  除数博弈 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/1025.divisor-game%20除数博弈/main.go) | 
 [1266.minimum-time-visiting-all-points](https://leetcode-cn.com/problems/minimum-time-visiting-all-points) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/1266.minimum-time-visiting-all-points/main.go) | 
 [1281.subtract-the-product-and-sum-of-digits-of-an-integer](https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/1281.subtract-the-product-and-sum-of-digits-of-an-integer/main.go) | 
 [1313.decompress-run-length-encoded-list](https://leetcode-cn.com/problems/decompress-run-length-encoded-list) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/1313.decompress-run-length-encoded-list/main.go) | 
 [1403.minimum-subsequence-in-non-increasing-order](https://leetcode-cn.com/problems/minimum-subsequence-in-non-increasing-order) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/1403.minimum-subsequence-in-non-increasing-order/main.go) | 
 [1791.find-center-of-star-graph  ](https://leetcode-cn.com/problems/find-center-of-star-graph) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/math/1791.find-center-of-star-graph/main.go) | 

## Sql
| Title                                 | cn         | Difficulty  | Solution | Algorithm
| ------------------------------------- | ---------  | ----------- |--------- | ---------
 [0175.combine-two-tables         ](https://leetcode-cn.com/problems/combine-two-tables) |  给所有人加上地址 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0175.combine-two-tables%20给所有人加上地址/1.sql) | 
 [0176.second-highest-salary      ](https://leetcode-cn.com/problems/second-highest-salary) |  找第二名 Null 也返回 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0176.second-highest-salary%20找第二名%20Null%20也返回/1.sql) | 
 [0177.nth-highest-salary         ](https://leetcode-cn.com/problems/nth-highest-salary) |  找第N名 Null 也返回 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0177.nth-highest-salary%20找第N名%20Null%20也返回/1.sql) | 
 [0178.rank-scores                ](https://leetcode-cn.com/problems/rank-scores) |  稠密排行 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0178.rank-scores%20稠密排行/1.sql) | 
 [0180.consecutive-numbers        ](https://leetcode-cn.com/problems/consecutive-numbers) |  连续出现3次的数字 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0180.consecutive-numbers%20连续出现3次的数字/1.sql) | 
 [0181.employees-earning-more-than-their-managers](https://leetcode-cn.com/problems/employees-earning-more-than-their-managers) |  超过经理收入的员工 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0181.employees-earning-more-than-their-managers%20超过经理收入的员工/1.sql) | 
 [0182.duplicate-emails           ](https://leetcode-cn.com/problems/duplicate-emails) |  找重复邮箱 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0182.duplicate-emails%20找重复邮箱/1.sql) | 
 [0183.customers-who-never-order  ](https://leetcode-cn.com/problems/customers-who-never-order) |  从不订购的客户 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0183.customers-who-never-order%20从不订购的客户/1.sql) | 
 [0184.department-highest-salary  ](https://leetcode-cn.com/problems/department-highest-salary) |  部门最高工资的员工 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0184.department-highest-salary%20部门最高工资的员工/1.sql) | 
 [0185.department-top-three-salaries](https://leetcode-cn.com/problems/department-top-three-salaries) |  部门工资前三高的员工 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0185.department-top-three-salaries%20部门工资前三高的员工/1.sql) | 
 [0196.delete-duplicate-emails    ](https://leetcode-cn.com/problems/delete-duplicate-emails) |  删除重复邮箱 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0196.delete-duplicate-emails%20删除重复邮箱/1.sql) | 
 [0197.rising-temperature         ](https://leetcode-cn.com/problems/rising-temperature) |  温度相比昨天是上升的 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0197.rising-temperature%20温度相比昨天是上升的/1.sql) | 
 [0262.trips-and-users            ](https://leetcode-cn.com/problems/trips-and-users) |  非禁止用户取消率 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0262.trips-and-users%20非禁止用户取消率/1.sql) | 
 [0511.game-play-analysis-i       ](https://leetcode-cn.com/problems/game-play-analysis-i) |  首次登陆的时间 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0511.game-play-analysis-i%20首次登陆的时间/1.sql) | 
 [0512.game-play-analysis-ii      ](https://leetcode-cn.com/problems/game-play-analysis-ii) |  首次登陆的设备名称 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0512.game-play-analysis-ii%20首次登陆的设备名称/1.sql) | 
 [0534.game-play-analysis-iii     ](https://leetcode-cn.com/problems/game-play-analysis-iii) |  每人每天累积多少时长 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0534.game-play-analysis-iii%20每人每天累积多少时长/1.sql) | 
 [0550.game-play-analysis-iv      ](https://leetcode-cn.com/problems/game-play-analysis-iv) |  首日后隔天登录玩家的比率 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0550.game-play-analysis-iv%20首日后隔天登录玩家的比率/1.sql) | 
 [0569.median-employee-salary     ](https://leetcode-cn.com/problems/median-employee-salary) |  每个公司薪酬中位数 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0569.median-employee-salary%20每个公司薪酬中位数/1.sql) | 
 [0570.managers-with-at-least-5-direct-reports](https://leetcode-cn.com/problems/managers-with-at-least-5-direct-reports) |  至少5名下属的经理 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0570.managers-with-at-least-5-direct-reports%20至少5名下属的经理/1.sql) | 
 [0571.find-median-given-frequency-of-numbers](https://leetcode-cn.com/problems/find-median-given-frequency-of-numbers) |  频率数字的中位数 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0571.find-median-given-frequency-of-numbers%20频率数字的中位数/1.sql) | 
 [0574.winning-candidate          ](https://leetcode-cn.com/problems/winning-candidate) |  当选者 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0574.winning-candidate%20当选者/1.sql) | 
 [0577.employee-bonus             ](https://leetcode-cn.com/problems/employee-bonus) |  员工奖金 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0577.employee-bonus%20员工奖金/1.sql) | 
 [0578.get-highest-answer-rate-question](https://leetcode-cn.com/problems/get-highest-answer-rate-question) |  回答率最高的问题 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0578.get-highest-answer-rate-question%20回答率最高的问题/1.sql) | 
 [0579.find-cumulative-salary-of-an-employee](https://leetcode-cn.com/problems/find-cumulative-salary-of-an-employee) |  员工累积薪水 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0579.find-cumulative-salary-of-an-employee%20员工累积薪水/1.sql) | 
 [0580.count-student-number-in-departments](https://leetcode-cn.com/problems/count-student-number-in-departments) |  各专业学生人数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0580.count-student-number-in-departments%20各专业学生人数/1.sql) | 
 [0584.find-customer-referee      ](https://leetcode-cn.com/problems/find-customer-referee) |  用户的推荐人 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0584.find-customer-referee%20用户的推荐人/1.sql) | 
 [0585.investments-in-2016        ](https://leetcode-cn.com/problems/investments-in-2016) |  2016年的投资 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0585.investments-in-2016%202016年的投资/1.sql) | 
 [0586.customer-placing-the-largest-number-of-orders](https://leetcode-cn.com/problems/customer-placing-the-largest-number-of-orders) |  订单最多的客户 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0586.customer-placing-the-largest-number-of-orders%20订单最多的客户/1.sql) | 
 [0595.big-countries              ](https://leetcode-cn.com/problems/big-countries) |  大的国家 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0595.big-countries%20大的国家/1.sql) | 
 [0596.classes-more-than-5-students](https://leetcode-cn.com/problems/classes-more-than-5-students) |  超过5名学生的课 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0596.classes-more-than-5-students%20超过5名学生的课/1.sql) | 
 [0597.friend-requests-i-overall-acceptance-rate](https://leetcode-cn.com/problems/friend-requests-i-overall-acceptance-rate) |  好友申请总体通过率 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0597.friend-requests-i-overall-acceptance-rate%20好友申请总体通过率/1.sql) | 
 [0601.human-traffic-of-stadium   ](https://leetcode-cn.com/problems/human-traffic-of-stadium) |  人流量 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0601.human-traffic-of-stadium%20人流量/1.sql) | 
 [0602.friend-requests-ii-who-has-the-most-friends](https://leetcode-cn.com/problems/friend-requests-ii-who-has-the-most-friends) |  谁有最多的好友 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0602.friend-requests-ii-who-has-the-most-friends%20谁有最多的好友/1.sql) | 
 [0603.consecutive-available-seats](https://leetcode-cn.com/problems/consecutive-available-seats) |  连续空余座位 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0603.consecutive-available-seats%20连续空余座位/1.sql) | 
 [0607.sales-person               ](https://leetcode-cn.com/problems/sales-person) |  销售员 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0607.sales-person%20销售员/1.sql) | 
 [0608.tree-node                  ](https://leetcode-cn.com/problems/tree-node) |  树节点 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0608.tree-node%20树节点/1.sql) | 
 [0610.triangle-judgement         ](https://leetcode-cn.com/problems/triangle-judgement) |  判断三角形 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0610.triangle-judgement%20判断三角形/1.sql) | 
 [0612.shortest-distance-in-a-plane](https://leetcode-cn.com/problems/shortest-distance-in-a-plane) |  平面上的最近距离 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0612.shortest-distance-in-a-plane%20平面上的最近距离/1.sql) | 
 [0613.shortest-distance-in-a-line](https://leetcode-cn.com/problems/shortest-distance-in-a-line) |  直线上的最近距离 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0613.shortest-distance-in-a-line%20直线上的最近距离/1.sql) | 
 [0614.second-degree-follower     ](https://leetcode-cn.com/problems/second-degree-follower) |  二级关注者 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0614.second-degree-follower%20二级关注者/1.sql) | 
 [0615.average-salary-departments-vs-company](https://leetcode-cn.com/problems/average-salary-departments-vs-company) |  平均工资：部门与公司比较 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0615.average-salary-departments-vs-company%20平均工资：部门与公司比较/1.sql) | 
 [0618.students-report-by-geography](https://leetcode-cn.com/problems/students-report-by-geography) |  学生地理信息报告 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/0618.students-report-by-geography%20学生地理信息报告/1.sql) | 
 [0619.biggest-single-number      ](https://leetcode-cn.com/problems/biggest-single-number) |  最出现一次的最大数字 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0619.biggest-single-number%20最出现一次的最大数字/1.sql) | 
 [0620.not-boring-movies          ](https://leetcode-cn.com/problems/not-boring-movies) |  有趣的电影 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0620.not-boring-movies%20有趣的电影/1.sql) | 
 [0626.exchange-seats             ](https://leetcode-cn.com/problems/exchange-seats) |  换座位 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/0626.exchange-seats%20换座位/1.sql) | 
 [0627.swap-salary                ](https://leetcode-cn.com/problems/swap-salary) |  变更性别 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/0627.swap-salary%20变更性别/1.sql) | 
 [1045.customers-who-bought-all-products](https://leetcode-cn.com/problems/customers-who-bought-all-products) |  买下所有产品的客户 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1045.customers-who-bought-all-products%20买下所有产品的客户/1.sql) | 
 [1050.actors-and-directors-who-cooperated-at-least-three-times](https://leetcode-cn.com/problems/actors-and-directors-who-cooperated-at-least-three-times) |  合作至少三次的演员和导员 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1050.actors-and-directors-who-cooperated-at-least-three-times%20合作至少三次的演员和导员/1.sql) | 
 [1068.product-sales-analysis-i   ](https://leetcode-cn.com/problems/product-sales-analysis-i) |  产品的年份和价格 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1068.product-sales-analysis-i%20产品的年份和价格/1.sql) | 
 [1069.product-sales-analysis-ii  ](https://leetcode-cn.com/problems/product-sales-analysis-ii) |  产品的销售总额 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1069.product-sales-analysis-ii%20产品的销售总额/1.sql) | 
 [1070.product-sales-analysis-iii ](https://leetcode-cn.com/problems/product-sales-analysis-iii) |  第一年的价格 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1070.product-sales-analysis-iii%20第一年的价格/1.sql) | 
 [1075.project-employees-i        ](https://leetcode-cn.com/problems/project-employees-i) |  项目的平均年限 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1075.project-employees-i%20项目的平均年限/1.sql) | 
 [1076.project-employees-ii       ](https://leetcode-cn.com/problems/project-employees-ii) |  员工最多的项目 销售额最高的销售者 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1076.project-employees-ii%20员工最多的项目%20销售额最高的销售者/1.sql) | 
 [1077.project-employees-iii      ](https://leetcode-cn.com/problems/project-employees-iii) |  项目经济最丰富的员工 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1077.project-employees-iii%20项目经济最丰富的员工/1.sql) | 
 [1082.sales-analysis-i           ](https://leetcode-cn.com/problems/sales-analysis-i) |  销售额最高的销售者 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1082.sales-analysis-i%20销售额最高的销售者/1.sql) | 
 [1083.sales-analysis-ii          ](https://leetcode-cn.com/problems/sales-analysis-ii) |  买 S8 却没有 iPhone 的买家 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1083.sales-analysis-ii%20买%20S8%20却没有%20iPhone%20的买家/1.sql) | 
 [1084.sales-analysis-iii         ](https://leetcode-cn.com/problems/sales-analysis-iii) |  只在春季销售的产品 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1084.sales-analysis-iii%20只在春季销售的产品/1.sql) | 
 [1097.game-play-analysis-v       ](https://leetcode-cn.com/problems/game-play-analysis-v) |  第二天留存率 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1097.game-play-analysis-v%20第二天留存率/1.sql) | 
 [1098.unpopular-books            ](https://leetcode-cn.com/problems/unpopular-books) |  过去一年少于十本的书 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1098.unpopular-books%20过去一年少于十本的书/1.sql) | 
 [1107.new-users-daily-count      ](https://leetcode-cn.com/problems/new-users-daily-count) |  每日新用户 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1107.new-users-daily-count%20每日新用户/1.sql) | 
 [1112.highest-grade-for-each-student](https://leetcode-cn.com/problems/highest-grade-for-each-student) |  最高分数的学科 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1112.highest-grade-for-each-student%20最高分数的学科/1.sql) | 
 [1113.reported-posts             ](https://leetcode-cn.com/problems/reported-posts) |  不同的报告记录 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1113.reported-posts%20不同的报告记录/1.sql) | 
 [1126.active-businesses          ](https://leetcode-cn.com/problems/active-businesses) |  活跃的业务 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1126.active-businesses%20活跃的业务/1.sql) | 
 [1127.user-purchase-platform     ](https://leetcode-cn.com/problems/user-purchase-platform) |  统计单端或双端人数 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1127.user-purchase-platform%20统计单端或双端人数/1.sql) | 
 [1132.reported-posts-ii          ](https://leetcode-cn.com/problems/reported-posts-ii) |  垃圾清除率 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1132.reported-posts-ii%20垃圾清除率/1.sql) | 
 [1141.user-activity-for-the-past-30-days-i](https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-i) |  近30天活跃数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1141.user-activity-for-the-past-30-days-i%20近30天活跃数/1.sql) | 
 [1142.user-activity-for-the-past-30-days-ii](https://leetcode-cn.com/problems/user-activity-for-the-past-30-days-ii) |  平均会话次数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1142.user-activity-for-the-past-30-days-ii%20平均会话次数/1.sql) | 
 [1148.article-views-i            ](https://leetcode-cn.com/problems/article-views-i) |  浏览过自己文章的作者 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1148.article-views-i%20浏览过自己文章的作者/1.sql) | 
 [1149.article-views-ii           ](https://leetcode-cn.com/problems/article-views-ii) |  一天浏览过两篇文章的人 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1149.article-views-ii%20一天浏览过两篇文章的人/1.sql) | 
 [1158.market-analysis-i          ](https://leetcode-cn.com/problems/market-analysis-i) |  统计2019年订单总数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1158.market-analysis-i%20统计2019年订单总数/1.sql) | 
 [1159.market-analysis-ii         ](https://leetcode-cn.com/problems/market-analysis-ii) |  售出第二件不是喜欢的商品 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1159.market-analysis-ii%20售出第二件不是喜欢的商品/1.sql) | 
 [1164.product-price-at-a-given-date](https://leetcode-cn.com/problems/product-price-at-a-given-date) |  变更后的价格 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1164.product-price-at-a-given-date%20变更后的价格/1.sql) | 
 [1173.immediate-food-delivery-i  ](https://leetcode-cn.com/problems/immediate-food-delivery-i) |  当天配送率 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1173.immediate-food-delivery-i%20当天配送率/1.sql) | 
 [1174.immediate-food-delivery-ii ](https://leetcode-cn.com/problems/immediate-food-delivery-ii) |  首日即时配送订单 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1174.immediate-food-delivery-ii%20首日即时配送订单/1.sql) | 
 [1179.reformat-department-table  ](https://leetcode-cn.com/problems/reformat-department-table) |  格式化工资 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1179.reformat-department-table%20格式化工资/1.sql) | 
 [1193.monthly-transactions-i     ](https://leetcode-cn.com/problems/monthly-transactions-i) |  统计 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1193.monthly-transactions-i%20统计/1.sql) | 
 [1194.tournament-winners         ](https://leetcode-cn.com/problems/tournament-winners) |  每组的优胜者 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1194.tournament-winners%20每组的优胜者/1.sql) | 
 [1204.last-person-to-fit-in-the-elevator](https://leetcode-cn.com/problems/last-person-to-fit-in-the-bus) |  最后进入电梯 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1204.last-person-to-fit-in-the-elevator%20最后进入电梯/1.sql) | 
 [1205.monthly-transactions-ii    ](https://leetcode-cn.com/problems/monthly-transactions-ii) |  交易统计 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1205.monthly-transactions-ii%20交易统计/1.sql) | 
 [1211.queries-quality-and-percentage](https://leetcode-cn.com/problems/queries-quality-and-percentage) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1211.queries-quality-and-percentage/1.sql) | 
 [1212.team-scores-in-football-tournament](https://leetcode-cn.com/problems/team-scores-in-football-tournament) |  计算得分 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1212.team-scores-in-football-tournament%20计算得分/1.sql) | 
 [1225.report-contiguous-dates    ](https://leetcode-cn.com/problems/report-contiguous-dates) |  统计连续 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1225.report-contiguous-dates%20统计连续/1.sql) | 
 [1241.number-of-comments-per-post](https://leetcode-cn.com/problems/number-of-comments-per-post) |  查询评论数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1241.number-of-comments-per-post%20查询评论数/1.sql) | 
 [1251.average-selling-price      ](https://leetcode-cn.com/problems/average-selling-price) |  | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1251.average-selling-price/1.sql) | 
 [1264.page-recommendations       ](https://leetcode-cn.com/problems/page-recommendations) |  推荐朋友喜欢的页面 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1264.page-recommendations%20推荐朋友喜欢的页面/1.sql) | 
 [1270.all-people-report-to-the-given-manager](https://leetcode-cn.com/problems/all-people-report-to-the-given-manager) |  递归查询 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1270.all-people-report-to-the-given-manager%20递归查询/1.sql) | 
 [1280.students-and-examinations  ](https://leetcode-cn.com/problems/students-and-examinations) |  学生各科测试次数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1280.students-and-examinations%20学生各科测试次数/1.sql) | 
 [1285.find-the-start-and-end-number-of-continuous-ranges](https://leetcode-cn.com/problems/find-the-start-and-end-number-of-continuous-ranges) |  区间次数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1285.find-the-start-and-end-number-of-continuous-ranges%20区间次数/1.sql) | 
 [1294.weather-type-in-each-country](https://leetcode-cn.com/problems/weather-type-in-each-country) |  11月份的天气 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1294.weather-type-in-each-country%2011月份的天气/1.sql) | 
 [1303.find-the-team-size         ](https://leetcode-cn.com/problems/find-the-team-size) |  团队人数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1303.find-the-team-size%20团队人数/1.sql) | 
 [1308.running-total-for-different-genders](https://leetcode-cn.com/problems/running-total-for-different-genders) |  男女的累积分数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1308.running-total-for-different-genders%20男女的累积分数/1.sql) | 
 [1321.restaurant-growth          ](https://leetcode-cn.com/problems/restaurant-growth) |  最近七天平均值 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1321.restaurant-growth%20最近七天平均值/1.sql) | 
 [1322.ads-performance            ](https://leetcode-cn.com/problems/ads-performance) |  广告效果统计 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1322.ads-performance%20广告效果统计/1.sql) | 
 [1327.list-the-products-ordered-in-a-period](https://leetcode-cn.com/problems/list-the-products-ordered-in-a-period) |  指定日期总数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1327.list-the-products-ordered-in-a-period%20指定日期总数/1.sql) | 
 [1336.number-of-transactions-per-visit](https://leetcode-cn.com/problems/number-of-transactions-per-visit) |  每次访问的交易次数 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1336.number-of-transactions-per-visit%20每次访问的交易次数/1.sql) | 
 [1341.movie-rating               ](https://leetcode-cn.com/problems/movie-rating) |  评论数和评分数统计 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1341.movie-rating%20评论数和评分数统计/1.sql) | 
 [1350.students-with-invalid-departments](https://leetcode-cn.com/problems/students-with-invalid-departments) |  不存在院校的学生 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1350.students-with-invalid-departments%20不存在院校的学生/1.sql) | 
 [1355.activity-participants      ](https://leetcode-cn.com/problems/activity-participants) |  去除第一和倒数第一 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1355.activity-participants%20去除第一和倒数第一/1.sql) | 
 [1364.number-of-trusted-contacts-of-a-customer](https://leetcode-cn.com/problems/number-of-trusted-contacts-of-a-customer) |  互为好友的顾客 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1364.number-of-trusted-contacts-of-a-customer%20互为好友的顾客/1.sql) | 
 [1369.get-the-second-most-recent-activity](https://leetcode-cn.com/problems/get-the-second-most-recent-activity) |  倒数第二次活动 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1369.get-the-second-most-recent-activity%20倒数第二次活动/1.sql) | 
 [1378.replace-employee-id-with-the-unique-identifier](https://leetcode-cn.com/problems/replace-employee-id-with-the-unique-identifier) |  使用uuid替换id | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1378.replace-employee-id-with-the-unique-identifier%20使用uuid替换id/1.sql) | 
 [1384.total-sales-amount-by-year ](https://leetcode-cn.com/problems/total-sales-amount-by-year) |  年度统计 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1384.total-sales-amount-by-year%20年度统计/1.sql) | 
 [1393.capital-gainloss           ](https://leetcode-cn.com/problems/capital-gainloss) |   资本损益 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1393.capital-gainloss%20%20资本损益/1.sql) | 
 [1398.customers-who-bought-products-a-and-b-but-not-c](https://leetcode-cn.com/problems/customers-who-bought-products-a-and-b-but-not-c) |  买AB不买C | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1398.customers-who-bought-products-a-and-b-but-not-c%20买AB不买C/1.sql) | 
 [1407.top-travellers             ](https://leetcode-cn.com/problems/top-travellers) |  旅行距离排名 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1407.top-travellers%20旅行距离排名/1.sql) | 
 [1412.find-the-quiet-students-in-all-exams](https://leetcode-cn.com/problems/find-the-quiet-students-in-all-exams) |  成绩中游学生 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1412.find-the-quiet-students-in-all-exams%20成绩中游学生/1.sql) | 
 [1418.display-table-of-food-orders-in-a-restaurant](https://leetcode-cn.com/problems/display-table-of-food-orders-in-a-restaurant) |  点菜展示表 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1418.display-table-of-food-orders-in-a-restaurant%20点菜展示表/main.go) | 
 [1421.npv-queries                ](https://leetcode-cn.com/problems/npv-queries) |  连表 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1421.npv-queries%20连表/1.sql) | 
 [1435.create-a-session-bar-chart ](https://leetcode-cn.com/problems/create-a-session-bar-chart) |  分段汇总 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1435.create-a-session-bar-chart%20分段汇总/1.sql) | 
 [1440.evaluate-boolean-expression](https://leetcode-cn.com/problems/evaluate-boolean-expression) |  连表比较 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1440.evaluate-boolean-expression%20连表比较/1.sql) | 
 [1445.apples-oranges             ](https://leetcode-cn.com/problems/apples-oranges) |  销量差 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1445.apples-oranges%20销量差/1.sql) | 
 [1454.active-users               ](https://leetcode-cn.com/problems/active-users) |  连续7天在线的用户 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1454.active-users%20连续7天在线的用户/1.sql) | 
 [1459.rectangles-area            ](https://leetcode-cn.com/problems/rectangles-area) |  组合成矩形 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1459.rectangles-area%20组合成矩形/1.sql) | 
 [1468.calculate-salaries         ](https://leetcode-cn.com/problems/calculate-salaries) |  扣税 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1468.calculate-salaries%20扣税/1.sql) | 
 [1479.sales-by-day-of-the-week   ](https://leetcode-cn.com/problems/sales-by-day-of-the-week) |  每周销量统计 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1479.sales-by-day-of-the-week%20每周销量统计/1.sql) | 
 [1485.group-sold-products-by-the-date](https://leetcode-cn.com/problems/group-sold-products-by-the-date) |  group聚合 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1485.group-sold-products-by-the-date%20group聚合/1.sql) | 
 [1495.friendly-movies-streamed-last-month](https://leetcode-cn.com/problems/friendly-movies-streamed-last-month) |  儿童适宜的电影 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1495.friendly-movies-streamed-last-month%20儿童适宜的电影/1.sql) | 
 [1501.countries-you-can-safely-invest-in](https://leetcode-cn.com/problems/countries-you-can-safely-invest-in) |  通话时长高于平均 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1501.countries-you-can-safely-invest-in%20通话时长高于平均/1.sql) | 
 [1511.customer-order-frequency   ](https://leetcode-cn.com/problems/customer-order-frequency) |  六七月消费大于等于100 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1511.customer-order-frequency%20六七月消费大于等于100/1.sql) | 
 [1517.find-users-with-valid-e-mails](https://leetcode-cn.com/problems/find-users-with-valid-e-mails) |  邮箱正则 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1517.find-users-with-valid-e-mails%20邮箱正则/1.sql) | 
 [1527.patients-with-a-condition  ](https://leetcode-cn.com/problems/patients-with-a-condition) |  患病患者 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1527.patients-with-a-condition%20患病患者/1.sql) | 
 [1532.the-most-recent-three-orders](https://leetcode-cn.com/problems/the-most-recent-three-orders) |  最近三笔订单 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1532.the-most-recent-three-orders%20最近三笔订单/1.sql) | 
 [1543.fix-product-name-format    ](https://leetcode-cn.com/problems/fix-product-name-format) |  格式化 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1543.fix-product-name-format%20格式化/1.sql) | 
 [1549.the-most-recent-orders-for-each-product](https://leetcode-cn.com/problems/the-most-recent-orders-for-each-product) |  产品最新订单 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1549.the-most-recent-orders-for-each-product%20产品最新订单/1.sql) | 
 [1555.bank-account-summary       ](https://leetcode-cn.com/problems/bank-account-summary) |  是否透支 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1555.bank-account-summary%20是否透支/1.sql) | 
 [1565.unique-orders-and-customers-per-month](https://leetcode-cn.com/problems/unique-orders-and-customers-per-month) |  月订单数和顾客数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1565.unique-orders-and-customers-per-month%20月订单数和顾客数/1.sql) | 
 [1571.warehouse-manager          ](https://leetcode-cn.com/problems/warehouse-manager) |  计算立方 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1571.warehouse-manager%20计算立方/1.sql) | 
 [1581.customer-who-visited-but-did-not-make-any-transactions](https://leetcode-cn.com/problems/customer-who-visited-but-did-not-make-any-transactions) |  进店未交易 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1581.customer-who-visited-but-did-not-make-any-transactions%20进店未交易/1.sql) | 
 [1587.bank-account-summary-ii    ](https://leetcode-cn.com/problems/bank-account-summary-ii) |  余大于10000的用户 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1587.bank-account-summary-ii%20余大于10000的用户/1.sql) | 
 [1596.the-most-frequently-ordered-products-for-each-customer](https://leetcode-cn.com/problems/the-most-frequently-ordered-products-for-each-customer) |  经常订购的商品 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1596.the-most-frequently-ordered-products-for-each-customer%20经常订购的商品/1.sql) | 
 [1607.sellers-with-no-sales      ](https://leetcode-cn.com/problems/sellers-with-no-sales) |  没有卖出产品的销售 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1607.sellers-with-no-sales%20没有卖出产品的销售/1.sql) | 
 [1613.find-the-missing-ids       ](https://leetcode-cn.com/problems/find-the-missing-ids) |  缺失的id | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1613.find-the-missing-ids%20缺失的id/1.sql) | 
 [1623.all-valid-triplets-that-can-represent-a-country](https://leetcode-cn.com/problems/all-valid-triplets-that-can-represent-a-country) |  三个学校组合不同的代表队 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1623.all-valid-triplets-that-can-represent-a-country%20三个学校组合不同的代表队/1.sql) | 
 [1633.percentage-of-users-attended-a-contest](https://leetcode-cn.com/problems/percentage-of-users-attended-a-contest) |  注册率 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1633.percentage-of-users-attended-a-contest%20注册率/1.sql) | 
 [1635.hopper-company-queries-i   ](https://leetcode-cn.com/problems/hopper-company-queries-i) |  活跃数与行程数 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1635.hopper-company-queries-i%20活跃数与行程数/1.sql) | 
 [1645.hopper-company-queries-ii  ](https://leetcode-cn.com/problems/hopper-company-queries-ii) |  工作率 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1645.hopper-company-queries-ii%20工作率/1.sql) | 
 [1651.hopper-company-queries-iii ](https://leetcode-cn.com/problems/hopper-company-queries-iii) |  平均距离和平均时间 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1651.hopper-company-queries-iii%20平均距离和平均时间/1.sql) | 
 [1661.average-time-of-process-per-machine](https://leetcode-cn.com/problems/average-time-of-process-per-machine) |  平均运行时间 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1661.average-time-of-process-per-machine%20平均运行时间/1.sql) | 
 [1667.fix-names-in-a-table       ](https://leetcode-cn.com/problems/fix-names-in-a-table) |   修复名字 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1667.fix-names-in-a-table%20%20修复名字/1.sql) | 
 [1677.products-worth-over-invoices](https://leetcode-cn.com/problems/products-worth-over-invoices) |  统计 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1677.products-worth-over-invoices%20统计/1.sql) | 
 [1683.invalid-tweets             ](https://leetcode-cn.com/problems/invalid-tweets) |  无效推文 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1683.invalid-tweets%20无效推文/1.sql) | 
 [1693.daily-leads-and-partners   ](https://leetcode-cn.com/problems/daily-leads-and-partners) |  统计 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1693.daily-leads-and-partners%20统计/1.sql) | 
 [1699.number-of-calls-between-two-persons](https://leetcode-cn.com/problems/number-of-calls-between-two-persons) |  通话次数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1699.number-of-calls-between-two-persons%20通话次数/1.sql) | 
 [1709.biggest-window-between-visits](https://leetcode-cn.com/problems/biggest-window-between-visits) |  日期缺失 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1709.biggest-window-between-visits%20日期缺失/1.sql) | 
 [1715.count-apples-and-oranges   ](https://leetcode-cn.com/problems/count-apples-and-oranges) |  统计 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1715.count-apples-and-oranges%20统计/1.sql) | 
 [1729.find-followers-count       ](https://leetcode-cn.com/problems/find-followers-count) |  关注数 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1729.find-followers-count%20关注数/1.sql) | 
 [1731.the-number-of-employees-which-report-to-each-employee](https://leetcode-cn.com/problems/the-number-of-employees-which-report-to-each-employee) |  每位经理有多少下属 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1731.the-number-of-employees-which-report-to-each-employee%20每位经理有多少下属/1.sql) | 
 [1741.find-total-time-spent-by-each-employee](https://leetcode-cn.com/problems/find-total-time-spent-by-each-employee) |  员工花费时间 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1741.find-total-time-spent-by-each-employee%20员工花费时间/1.sql) | 
 [1747.leetflex-banned-accounts   ](https://leetcode-cn.com/problems/leetflex-banned-accounts) |  应被禁止的账户 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1747.leetflex-banned-accounts%20应被禁止的账户/1.sql) | 
 [1757.recyclable-and-low-fat-products](https://leetcode-cn.com/problems/recyclable-and-low-fat-products) |  可回收且低脂的产品 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1757.recyclable-and-low-fat-products%20可回收且低脂的产品/1.sql) | 
 [1767.find-the-subtasks-that-did-not-execute](https://leetcode-cn.com/problems/find-the-subtasks-that-did-not-execute) |  找到未执行的子任务 | Hard | [Go](https://github.com/bygo/leetcode/blob/master/sql/1767.find-the-subtasks-that-did-not-execute%20找到未执行的子任务/1.sql) | 
 [1777.products-price-for-each-store](https://leetcode-cn.com/problems/products-price-for-each-store) |  每家商店价格 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1777.products-price-for-each-store%20每家商店价格/1.sql) | 
 [1783.grand-slam-titles          ](https://leetcode-cn.com/problems/grand-slam-titles) |  大满贯数 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1783.grand-slam-titles%20大满贯数/1.sql) | 
 [1789.primary-department-for-each-employee](https://leetcode-cn.com/problems/primary-department-for-each-employee) |  员工的主要部门 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1789.primary-department-for-each-employee%20员工的主要部门/1.sql) | 
 [1795.rearrange-products-table   ](https://leetcode-cn.com/problems/rearrange-products-table) |  商品价格合并 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1795.rearrange-products-table%20商品价格合并/1.sql) | 
 [1809.ad-free-sessions           ](https://leetcode-cn.com/problems/ad-free-sessions) |  没有展示广告的会话 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1809.ad-free-sessions%20没有展示广告的会话/1.sql) | 
 [1811.find-interview-candidates  ](https://leetcode-cn.com/problems/find-interview-candidates) |  统计 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1811.find-interview-candidates%20统计/1.sql) | 
 [1821.find-customers-with-positive-revenue-this-year](https://leetcode-cn.com/problems/find-customers-with-positive-revenue-this-year) |  正收入 | Easy | [Go](https://github.com/bygo/leetcode/blob/master/sql/1821.find-customers-with-positive-revenue-this-year%20正收入/1.sql) | 
 [1831.maximum-transaction-each-day](https://leetcode-cn.com/problems/maximum-transaction-each-day) |  每天最大交易 | Medium | [Go](https://github.com/bygo/leetcode/blob/master/sql/1831.maximum-transaction-each-day%20每天最大交易/1.sql) | 