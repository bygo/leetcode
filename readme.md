# `Go版本的一题多解（所有解）`

# 刷题步骤

## `1.新建刷题目录`
比如第 `1` 题 two-sum

`go run main.go -a=1`

比如第 `100` 题 same-tree

`go run main.go -a=100`

## `2.模板流`
## `3.目录约定`  `生成 readme.md`

`分类` : `1 级目录`

`题名` : `2 级目录`

`算法` : `3-N 级目录`

执行 `go run main.go` 自动生成 `readme.md`

# 模板
- [Tree](https://github.com/temporaries/leetcode/tree/master/templates/tree)
    - [前序遍历](https://github.com/temporaries/leetcode/blob/master/templates/tree/preorder.go)
    - [前序遍历-迭代](https://github.com/temporaries/leetcode/blob/master/templates/tree/preorder_stack.go)
    - [前序遍历-Morris-链表](https://github.com/temporaries/leetcode/blob/master/templates/tree/preorder_morris_break.go)
    - [前序遍历-Morris-保持树结构](https://github.com/temporaries/leetcode/blob/master/templates/tree/preorder_morris_keep.go)
    
    - [中序遍历](https://github.com/temporaries/leetcode/blob/master/templates/tree/inorder.go)
    - [中序遍历-迭代](https://github.com/temporaries/leetcode/blob/master/templates/tree/inorder_stack.go)
    - [中序遍历-Morris-链表](https://github.com/temporaries/leetcode/blob/master/templates/tree/inorder_morris_break.go)
    - [中序遍历-Morris-保持树结构](https://github.com/temporaries/leetcode/blob/master/templates/tree/inorder_morris_keep.go)

    - [后序遍历](https://github.com/temporaries/leetcode/blob/master/templates/tree/postorder.go)
    - [层次遍历](https://github.com/temporaries/leetcode/blob/master/templates/tree/bfs.go)

# 数据结构

---
- [Array](#Array)
- [LinkedList](#LinkedList)
- [Math](#Math)
- [Stack](#Stack)
- [String](#String)
- [Tree](#Tree)

---



## Array
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0001 | [Two Sum                         ](https://leetcode-cn.com/problems/two-sum) | 48.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/0001.two-sum/hash_table/main.go)| hash table
| 0004 | [Median of Two Sorted Arrays     ](https://leetcode-cn.com/problems/median-of-two-sorted-arrays) | 37.2% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/array/0004.median-of-two-sorted-arrays/binary_search/main.go)| binary search
| 0011 | [Container With Most Water       ](https://leetcode-cn.com/problems/container-with-most-water) | 62.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0011.container-with-most-water/two_pointer/main.go)| two pointer
| 0015 | [3Sum                            ](https://leetcode-cn.com/problems/3sum) | 26.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0015.3sum/two_pointer/main.go)| two pointer
| 0016 | [3Sum Closest                    ](https://leetcode-cn.com/problems/3sum-closest) | 43.6% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0016.3sum-closest/two_pointer/main.go)| two pointer
| 0018 | [4Sum                            ](https://leetcode-cn.com/problems/4sum) | 37.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0018.4sum/two_pointer/main.go)| two pointer
| 0026 | [Remove Duplicates from Sorted Array](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array) | 49.4% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/0026.remove-duplicates-from-sorted-array/two_pointer/main.go)| two pointer
| 0027 | [Remove Element                  ](https://leetcode-cn.com/problems/remove-element) | 57.9% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/0027.remove-element/two_pointer/main.go)| two pointer
| 0030 | [Substring with Concatenation of All Words](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words) | 30.1% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/array/0030.substring-with-concatenation-of-all-words/two_pointer/main.go)| two pointer
| 0033 | [Search in Rotated Sorted Array  ](https://leetcode-cn.com/problems/search-in-rotated-sorted-array) | 36.6% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0033.search-in-rotated-sorted-array/binary_search/main.go)| binary search
| 0034 | [Find First and Last Position of Element in Sorted Array](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array) | 39.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0034.find-first-and-last-position-of-element-in-sorted-array/binary_search/main.go)| binary search
| 0035 | [Search Insert Position          ](https://leetcode-cn.com/problems/search-insert-position) | 45.4% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/0035.search-insert-position/binary_search/main.go)| binary search
| 0036 | [Valid Sudoku                    ](https://leetcode-cn.com/problems/valid-sudoku) | 59.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0036.valid-sudoku/dfs/main.go)| dfs
| 0037 | [Sudoku Solver                   ](https://leetcode-cn.com/problems/sudoku-solver) | 60.7% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/array/0037.sudoku-solver/dfs/main.go)| dfs
| 0041 | [First Missing Positive          ](https://leetcode-cn.com/problems/first-missing-positive) | 37.9% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/array/0041.first-missing-positive/bit_map/main.go)| bit map
| 0042 | [Trapping Rain Water             ](https://leetcode-cn.com/problems/trapping-rain-water) | 50.4% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/array/0042.trapping-rain-water/two_pointer/main.go)| two pointer
| 1295 | [Find Numbers with Even Number of Digits](https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits) | 82.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/1295.find-numbers-with-even-number-of-digits/main.go)| 
| 1365 | [How Many Numbers Are Smaller Than the Current Number](https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number) | 82.5% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/1365.how-many-numbers-are-smaller-than-the-current-number/main.go)| 
| 1389 | [Create Target Array in the Given Order](https://leetcode-cn.com/problems/create-target-array-in-the-given-order) | 81.9% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/1389.create-target-array-in-the-given-order/main.go)| 

## LinkedList
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0002 | [Add Two Numbers                 ](https://leetcode-cn.com/problems/add-two-numbers) | 37.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0002.add-two-numbers/main.go)| 
| 0019 | [Remove Nth Node From End of List](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list) | 38.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0019.remove-nth-node-from-end-of-list/two_pointer/main.go)| two pointer
| 0021 | [Merge Two Sorted Lists          ](https://leetcode-cn.com/problems/merge-two-sorted-lists) | 60.8% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0021.merge-two-sorted-lists/main.go)| 
| 0023 | [Merge k Sorted Lists            ](https://leetcode-cn.com/problems/merge-k-sorted-lists) | 49.7% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0023.merge-k-sorted-lists/main.go)| 
| 0024 | [Swap Nodes in Pairs             ](https://leetcode-cn.com/problems/swap-nodes-in-pairs) | 65.1% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0024.swap-nodes-in-pairs/main.go)| 
| 0025 | [Reverse Nodes in k-Group        ](https://leetcode-cn.com/problems/reverse-nodes-in-k-group) | 57.6% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0025.reverse-nodes-in-k-group/main.go)| 

## Math
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0007 | [Reverse Integer                 ](https://leetcode-cn.com/problems/reverse-integer) | 34.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/0007.reverse-integer/main.go)| 
| 0008 | [String to Integer (atoi)        ](https://leetcode-cn.com/problems/string-to-integer-atoi) | 20.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/math/0008.string-to-integer-atoi/main.go)| 
| 0009 | [Palindrome Number               ](https://leetcode-cn.com/problems/palindrome-number) | 57.3% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/0009.palindrome-number/main.go)| 
| 0012 | [Integer to Roman                ](https://leetcode-cn.com/problems/integer-to-roman) | 63.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/math/0012.integer-to-roman/main.go)| 
| 0013 | [Roman to Integer                ](https://leetcode-cn.com/problems/roman-to-integer) | 61.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/0013.roman-to-integer/main.go)| 
| 0029 | [Divide Two Integers             ](https://leetcode-cn.com/problems/divide-two-integers) | 19.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/math/0029.divide-two-integers/main.go)| 
| 0031 | [Next Permutation                ](https://leetcode-cn.com/problems/next-permutation) | 33.1% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/math/0031.next-permutation/main.go)| 
| 0191 | [Number of 1 Bits                ](https://leetcode-cn.com/problems/number-of-1-bits) | 66.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/0191.number-of-1-bits/main.go)| 
| 0292 | [Nim Game                        ](https://leetcode-cn.com/problems/nim-game) | 69.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/0292.nim-game/main.go)| 
| 0319 | [Bulb Switcher                   ](https://leetcode-cn.com/problems/bulb-switcher) | 45.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/math/0319.bulb-switcher/main.go)| 
| 1403 | [Minimum Subsequence in Non-Increasing Order](https://leetcode-cn.com/problems/minimum-subsequence-in-non-increasing-order) | 71.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/1403.minimum-subsequence-in-non-increasing-order/main.go)| 

## Stack
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0020 | [Valid Parentheses               ](https://leetcode-cn.com/problems/valid-parentheses) | 41.3% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/stack/0020.valid-parentheses/main.go)| 
| 0032 | [Longest Valid Parentheses       ](https://leetcode-cn.com/problems/longest-valid-parentheses) | 30.1% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/stack/0032.longest-valid-parentheses/main.go)| 

## String
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0003 | [Longest Substring Without Repeating Characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters) | 33.6% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/string/0003.longest-substring-without-repeating-characters/main.go)| 
| 0005 | [Longest Palindromic Substring   ](https://leetcode-cn.com/problems/longest-palindromic-substring) | 29.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/string/0005.longest-palindromic-substring/main.go)| 
| 0006 | [ZigZag Conversion               ](https://leetcode-cn.com/problems/zigzag-conversion) | 47.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/string/0006.zigzag-conversion/main.go)| 
| 0010 | [Regular Expression Matching     ](https://leetcode-cn.com/problems/regular-expression-matching) | 27.1% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/string/0010.regular-expression-matching/main.go)| 
| 0014 | [Longest Common Prefix           ](https://leetcode-cn.com/problems/longest-common-prefix) | 36.8% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/string/0014.longest-common-prefix/main.go)| 
| 0017 | [Letter Combinations of a Phone Number](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number) | 53.1% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/string/0017.letter-combinations-of-a-phone-number/main.go)| 
| 0028 | [Implement strStr()              ](https://leetcode-cn.com/problems/implement-strstr) | 39.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/string/0028.implement-strstr/main.go)| 
| 0038 | [Count and Say                   ](https://leetcode-cn.com/problems/count-and-say) | 54.8% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/string/0038.count-and-say/main.go)| 
| 0043 | [Multiply Strings                ](https://leetcode-cn.com/problems/multiply-strings) | 42.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/string/0043.multiply-strings/main.go)| 
| 0415 | [Add Strings                     ](https://leetcode-cn.com/problems/add-strings) | 49.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/string/0415.add-strings/main.go)| 
| 0771 | [Jewels and Stones               ](https://leetcode-cn.com/problems/jewels-and-stones) | 82.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/string/0771.jewels-and-stones/hash_table/main.go)| hash table

## Tree
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0022 | [Generate Parentheses            ](https://leetcode-cn.com/problems/generate-parentheses) | 73.8% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0022.generate-parentheses/dfs/catalan/main.go)| dfs.catalan
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0022.generate-parentheses/dfs/recursive/main.go)| dfs.recursive
| 0039 | [Combination Sum                 ](https://leetcode-cn.com/problems/combination-sum) | 68.7% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0039.combination-sum/dfs/recursive/main.go)| dfs.recursive
| 0040 | [Combination Sum II              ](https://leetcode-cn.com/problems/combination-sum-ii) | 61.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0040.combination-sum-ii/dfs/recursive/main.go)| dfs.recursive
| 0094 | [Binary Tree Inorder Traversal   ](https://leetcode-cn.com/problems/binary-tree-inorder-traversal) | 70.9% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0094.binary-tree-inorder-traversal/dfs/inorder/recursive/main.go)| dfs.inorder.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0094.binary-tree-inorder-traversal/dfs/inorder/stack/main.go)| dfs.inorder.stack
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0094.binary-tree-inorder-traversal/dfs/morris/break/main.go)| dfs.morris.break
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0094.binary-tree-inorder-traversal/dfs/morris/keep/main.go)| dfs.morris.keep
| 0095 | [Unique Binary Search Trees II   ](https://leetcode-cn.com/problems/unique-binary-search-trees-ii) | 62.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0095.unique-binary-search-trees-ii/dfs/catalan/main.go)| dfs.catalan
| 0096 | [Unique Binary Search Trees      ](https://leetcode-cn.com/problems/unique-binary-search-trees) | 65.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0096.unique-binary-search-trees/catalan/main.go)| catalan
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0096.unique-binary-search-trees/dp/main.go)| dp
| 0098 | [Validate Binary Search Tree     ](https://leetcode-cn.com/problems/validate-binary-search-tree) | 29.7% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0098.validate-binary-search-tree/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0098.validate-binary-search-tree/dfs/inorder/recursive/main.go)| dfs.inorder.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0098.validate-binary-search-tree/dfs/inorder/stack/main.go)| dfs.inorder.stack
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0098.validate-binary-search-tree/dfs/preorder/recursive/main.go)| dfs.preorder.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0098.validate-binary-search-tree/dfs/preorder/stack/main.go)| dfs.preorder.stack
| 0099 | [Recover Binary Search Tree      ](https://leetcode-cn.com/problems/recover-binary-search-tree) | 56.1% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0099.recover-binary-search-tree/dfs/inorder/recursive/main.go)| dfs.inorder.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0099.recover-binary-search-tree/dfs/inorder/stack/main.go)| dfs.inorder.stack
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0099.recover-binary-search-tree/dfs/morris/keep.go)| dfs.morris
| 0100 | [Same Tree                       ](https://leetcode-cn.com/problems/same-tree) | 57.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0100.same-tree/dfs/recursive/main.go)| dfs.recursive
| 0101 | [Symmetric Tree                  ](https://leetcode-cn.com/problems/symmetric-tree) | 50.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0101.symmetric-tree/dfs/recursive/main.go)| dfs.recursive
| 0102 | [Binary Tree Level Order Traversal](https://leetcode-cn.com/problems/binary-tree-level-order-traversal) | 61.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0102.binary-tree-level-order-traversal/bfs/queue/main.go)| bfs.queue
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0102.binary-tree-level-order-traversal/dfs/recursive/main.go)| dfs.recursive
| 0103 | [Binary Tree Zigzag Level Order Traversal](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal) | 54.1% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0103.binary-tree-zigzag-level-order-traversal/bfs/queue/main.go)| bfs.queue
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0103.binary-tree-zigzag-level-order-traversal/dfs/recursive/main.go)| dfs.recursive
| 0104 | [Maximum Depth of Binary Tree    ](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree) | 72.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0104.maximum-depth-of-binary-tree/dfs/main.go)| dfs
| 0105 | [Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal) | 64.8% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0105.construct-binary-tree-from-preorder-and-inorder-traversal/main.go)| 
| 0106 | [Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal) | 67.7% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0106.construct-binary-tree-from-inorder-and-postorder-traversal/main.go)| 
| 0107 | [Binary Tree Level Order Traversal II](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii) | 64.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0107.binary-tree-level-order-traversal-ii/bfs/queue/main.go)| bfs.queue
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0107.binary-tree-level-order-traversal-ii/dfs/recursive/main.go)| dfs.recursive
| 0108 | [Convert Sorted Array to Binary Search Tree](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree) | 70.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0108.convert-sorted-array-to-binary-search-tree/dfs/recursive/main.go)| dfs.recursive
| 0109 | [Convert Sorted List to Binary Search Tree](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree) | 71.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0109.convert-sorted-list-to-binary-search-tree/array/main.go)| array
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0109.convert-sorted-list-to-binary-search-tree/inorder/main.go)| inorder
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0109.convert-sorted-list-to-binary-search-tree/recursive/main.go)| recursive
| 0110 | [Balanced Binary Tree            ](https://leetcode-cn.com/problems/balanced-binary-tree) | 51.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0110.balanced-binary-tree/postorder/main.go)| postorder
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0110.balanced-binary-tree/top/main.go)| top
| 0111 | [Minimum Depth of Binary Tree    ](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree) | 41.9% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0111.minimum-depth-of-binary-tree/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0111.minimum-depth-of-binary-tree/dfs/main.go)| dfs
| 0112 | [Path Sum                        ](https://leetcode-cn.com/problems/path-sum) | 49.4% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0112.path-sum/dfs/main.go)| dfs
| 0113 | [Path Sum II                     ](https://leetcode-cn.com/problems/path-sum-ii) | 59.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0113.path-sum-ii/dfs/main.go)| dfs
| 0114 | [Flatten Binary Tree to Linked List](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list) | 68.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0114.flatten-binary-tree-to-linked-list/preorder.morris/main.go)| preorder.morris
| 0116 | [Populating Next Right Pointers in Each Node](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node) | 57.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0116.populating-next-right-pointers-in-each-node/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0116.populating-next-right-pointers-in-each-node/dfs/main.go)| dfs
| 0117 | [Populating Next Right Pointers in Each Node II](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii) | 46.8% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0117.populating-next-right-pointers-in-each-node-ii/bfs/main.go)| bfs
| 0124 | [Binary Tree Maximum Path Sum    ](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum) | 39.8% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0124.binary-tree-maximum-path-sum/dfs/main.go)| dfs
| 0129 | [Sum Root to Leaf Numbers        ](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers) | 62.6% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0129.sum-root-to-leaf-numbers/dfs/main.go)| dfs
| 0144 | [Binary Tree Preorder Traversal  ](https://leetcode-cn.com/problems/binary-tree-preorder-traversal) | 65.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0144.binary-tree-preorder-traversal/dfs/recursive/main.go)| dfs.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0144.binary-tree-preorder-traversal/dfs/stack/main.go)| dfs.stack
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0144.binary-tree-preorder-traversal/morris/break/main.go)| morris.break
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0144.binary-tree-preorder-traversal/morris/keep/main.go)| morris.keep
| 0145 | [Binary Tree Postorder Traversal ](https://leetcode-cn.com/problems/binary-tree-postorder-traversal) | 70.9% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0145.binary-tree-postorder-traversal/dfs/recursive/main.go)| dfs.recursive
| 0173 | [Binary Search Tree Iterator     ](https://leetcode-cn.com/problems/binary-search-tree-iterator) | 72.1% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0173.binary-search-tree-iterator/main.go)| 
| 0199 | [Binary Tree Right Side View     ](https://leetcode-cn.com/problems/binary-tree-right-side-view) | 63.9% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0199.binary-tree-right-side-view/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0199.binary-tree-right-side-view/dfs/main.go)| dfs
| 0222 | [Count Complete Tree Nodes       ](https://leetcode-cn.com/problems/count-complete-tree-nodes) | 68.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0222.count-complete-tree-nodes/dfs/main.go)| dfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0222.count-complete-tree-nodes/two_pointer/main.go)| two pointer
| 0226 | [Invert Binary Tree              ](https://leetcode-cn.com/problems/invert-binary-tree) | 74.6% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0226.invert-binary-tree/dfs/main.go)| dfs
| 0230 | [Kth Smallest Element in a BST   ](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst) | 69.6% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0230.kth-smallest-element-in-a-bst/dfs/main.go)| dfs
| 0235 | [Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree) | 63.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0235.lowest-common-ancestor-of-a-binary-search-tree/dfs/main.go)| dfs
| 0236 | [Lowest Common Ancestor of a Binary Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree) | 61.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0236.lowest-common-ancestor-of-a-binary-tree/dfs/main.go)| dfs
| 0257 | [Binary Tree Paths               ](https://leetcode-cn.com/problems/binary-tree-paths) | 63.2% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0257.binary-tree-paths/dfs/main.go)| dfs
| 0297 | [Serialize and Deserialize Binary Tree](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree) | 46.0% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0297.serialize-and-deserialize-binary-tree/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0297.serialize-and-deserialize-binary-tree/dfs/main.go)| dfs
| 0404 | [Sum of Left Leaves              ](https://leetcode-cn.com/problems/sum-of-left-leaves) | 54.4% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0404.sum-of-left-leaves/dfs/main.go)| dfs
| 0429 | [N-ary Tree Level Order Traversal](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal) | 65.1% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0429.n-ary-tree-level-order-traversal/bfs/main.go)| bfs
| 0437 | [Path Sum III                    ](https://leetcode-cn.com/problems/path-sum-iii) | 54.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0437.path-sum-iii/dfs/main.go)| dfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0437.path-sum-iii/hash_table/main.go)| hash table
| 0449 | [Serialize and Deserialize BST   ](https://leetcode-cn.com/problems/serialize-and-deserialize-bst) | 50.9% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0449.serialize-and-deserialize-bst/dfs/preorder/main.go)| dfs.preorder
| 0501 | [Find Mode in Binary Search Tree ](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree) | 44.2% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0501.find-mode-in-binary-search-tree/dfs/recursive/main.go)| dfs.recursive
| 0559 | [Maximum Depth of N-ary Tree     ](https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree) | 68.9% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0559.maximum-depth-of-n-ary-tree/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0559.maximum-depth-of-n-ary-tree/dfs/main.go)| dfs
| 0589 | [N-ary Tree Preorder Traversal   ](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal) | 72.8% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0589.n-ary-tree-preorder-traversal/dfs/recursive/main.go)| dfs.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0589.n-ary-tree-preorder-traversal/dfs/stack/main.go)| dfs.stack
| 0590 | [N-ary Tree Postorder Traversal  ](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal) | 72.8% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0590.n-ary-tree-postorder-traversal/dfs/recursive/main.go)| dfs.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0590.n-ary-tree-postorder-traversal/dfs/stack/main.go)| dfs.stack
| 0617 | [Merge Two Binary Trees          ](https://leetcode-cn.com/problems/merge-two-binary-trees) | 75.8% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0617.merge-two-binary-trees/dfs/recursive/main.go)| dfs.recursive