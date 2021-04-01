# LeetCode

![](https://img.shields.io/badge/Language-Go-%2300ADD8)
[![](https://img.shields.io/badge/Template-语雀-%2325b864)](https://www.yuque.com/daizhuansheng/lc/cfg57t)
---

# 刷题步骤

## `1.新建题目`
第 `1` 题 two-sum

`go run main.go -a=1`

第 `100` 题 same-tree

`go run main.go -a=100`

## `2.模板流`
[![](https://img.shields.io/badge/Template-语雀-%2325b864)](https://www.yuque.com/daizhuansheng/lc/cfg57t)

## `3.目录约定`  `生成 readme.md`

`分类` : `1 级目录`

`题名` : `2 级目录`

`算法名` : `3-N 级目录`

执行 `go run main.go` 自动生成 `readme.md`

# 模板
- [Tree](https://github.com/bygo/leetcode/tree/master/templates/tree)
    - [前序遍历](https://github.com/bygo/leetcode/blob/master/templates/tree/preorder.go)
    - [前序遍历-迭代](https://github.com/bygo/leetcode/blob/master/templates/tree/preorder_stack.go)
    - [前序遍历-Morris-链表](https://github.com/bygo/leetcode/blob/master/templates/tree/preorder_morris_break.go)
    - [前序遍历-Morris-保持树结构](https://github.com/bygo/leetcode/blob/master/templates/tree/preorder_morris_keep.go)
    
    - [中序遍历](https://github.com/bygo/leetcode/blob/master/templates/tree/inorder.go)
    - [中序遍历-迭代](https://github.com/bygo/leetcode/blob/master/templates/tree/inorder_stack.go)
    - [中序遍历-Morris-链表](https://github.com/bygo/leetcode/blob/master/templates/tree/inorder_morris_break.go)
    - [中序遍历-Morris-保持树结构](https://github.com/bygo/leetcode/blob/master/templates/tree/inorder_morris_keep.go)

    - [后序遍历](https://github.com/bygo/leetcode/blob/master/templates/tree/postorder.go)
    - [后序遍历-迭代](https://github.com/bygo/leetcode/blob/master/templates/tree/postorder_stack.go)
    - [后序遍历-Morris-破坏树结构](https://github.com/bygo/leetcode/blob/master/templates/tree/postorder_morris_break.go)
    - [后序遍历-Morris-保持树结构](https://github.com/bygo/leetcode/blob/master/templates/tree/postorder_morris_keep.go)

    - [层次遍历](https://github.com/bygo/leetcode/blob/master/templates/tree/bfs.go)

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
| 0001 | [Two Sum                         ](https://leetcode-cn.com/problems/two-sum) | 50.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/0001.two-sum/hash_table/main.go)| hash table
| 0004 | [Median of Two Sorted Arrays     ](https://leetcode-cn.com/problems/median-of-two-sorted-arrays) | 39.8% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/array/0004.median-of-two-sorted-arrays/binary_search/main.go)| binary search
| 0011 | [Container With Most Water       ](https://leetcode-cn.com/problems/container-with-most-water) | 64.5% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0011.container-with-most-water/two_pointer/main.go)| two pointer
| 0015 | [3Sum                            ](https://leetcode-cn.com/problems/3sum) | 31.6% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0015.3sum/two_pointer/main.go)| two pointer
| 0016 | [3Sum Closest                    ](https://leetcode-cn.com/problems/3sum-closest) | 45.9% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0016.3sum-closest/two_pointer/main.go)| two pointer
| 0018 | [4Sum                            ](https://leetcode-cn.com/problems/4sum) | 40.1% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0018.4sum/two_pointer/main.go)| two pointer
| 0026 | [Remove Duplicates from Sorted Array](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array) | 53.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/0026.remove-duplicates-from-sorted-array/two_pointer/main.go)| two pointer
| 0027 | [Remove Element                  ](https://leetcode-cn.com/problems/remove-element) | 59.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/0027.remove-element/two_pointer/main.go)| two pointer
| 0030 | [Substring with Concatenation of All Words](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words) | 34.4% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/array/0030.substring-with-concatenation-of-all-words/two_pointer/main.go)| two pointer
| 0033 | [Search in Rotated Sorted Array  ](https://leetcode-cn.com/problems/search-in-rotated-sorted-array) | 41.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0033.search-in-rotated-sorted-array/binary_search/main.go)| binary search
| 0034 | [Find First and Last Position of Element in Sorted Array](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array) | 42.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0034.find-first-and-last-position-of-element-in-sorted-array/binary_search/main.go)| binary search
| 0035 | [Search Insert Position          ](https://leetcode-cn.com/problems/search-insert-position) | 47.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/0035.search-insert-position/binary_search/main.go)| binary search
| 0036 | [Valid Sudoku                    ](https://leetcode-cn.com/problems/valid-sudoku) | 62.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0036.valid-sudoku/dfs/main.go)| dfs
| 0037 | [Sudoku Solver                   ](https://leetcode-cn.com/problems/sudoku-solver) | 67.1% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/array/0037.sudoku-solver/dfs/main.go)| dfs
| 0041 | [First Missing Positive          ](https://leetcode-cn.com/problems/first-missing-positive) | 40.8% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/array/0041.first-missing-positive/bit_map/main.go)| bit map
| 0042 | [Trapping Rain Water             ](https://leetcode-cn.com/problems/trapping-rain-water) | 55.0% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/array/0042.trapping-rain-water/two_pointer/main.go)| two pointer
| 0075 | [Sort Colors                     ](https://leetcode-cn.com/problems/sort-colors) | 58.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/array/0075.sort-colors/main.go)| 
| 0121 | [Best Time to Buy and Sell Stock ](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock) | 56.5% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/0121.best-time-to-buy-and-sell-stock/main.go)| 
| 0122 | [Best Time to Buy and Sell Stock II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii) | 67.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/0122.best-time-to-buy-and-sell-stock-ii/main.go)| 
| 1295 | [Find Numbers with Even Number of Digits](https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits) | 81.2% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/1295.find-numbers-with-even-number-of-digits/main.go)| 
| 1365 | [How Many Numbers Are Smaller Than the Current Number](https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number) | 82.6% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/1365.how-many-numbers-are-smaller-than-the-current-number/main.go)| 
| 1389 | [Create Target Array in the Given Order](https://leetcode-cn.com/problems/create-target-array-in-the-given-order) | 82.9% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/array/1389.create-target-array-in-the-given-order/main.go)| 

## LinkedList
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0002 | [Add Two Numbers                 ](https://leetcode-cn.com/problems/add-two-numbers) | 39.8% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0002.add-two-numbers/main.go)| 
| 0019 | [Remove Nth Node From End of List](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list) | 41.5% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0019.remove-nth-node-from-end-of-list/two_pointer/main.go)| two pointer
| 0021 | [Merge Two Sorted Lists          ](https://leetcode-cn.com/problems/merge-two-sorted-lists) | 65.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0021.merge-two-sorted-lists/main.go)| 
| 0023 | [Merge k Sorted Lists            ](https://leetcode-cn.com/problems/merge-k-sorted-lists) | 54.7% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0023.merge-k-sorted-lists/main.go)| 
| 0024 | [Swap Nodes in Pairs             ](https://leetcode-cn.com/problems/swap-nodes-in-pairs) | 69.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0024.swap-nodes-in-pairs/main.go)| 
| 0025 | [Reverse Nodes in k-Group        ](https://leetcode-cn.com/problems/reverse-nodes-in-k-group) | 64.5% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0025.reverse-nodes-in-k-group/main.go)| 
| 0206 | [Reverse Linked List             ](https://leetcode-cn.com/problems/reverse-linked-list) | 71.5% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/linked_list/0206.reverse-linked-list/main.go)| 

## Math
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0007 | [Reverse Integer                 ](https://leetcode-cn.com/problems/reverse-integer) | 35.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/0007.reverse-integer/main.go)| 
| 0008 | [String to Integer (atoi)        ](https://leetcode-cn.com/problems/string-to-integer-atoi) | 21.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/math/0008.string-to-integer-atoi/main.go)| 
| 0009 | [Palindrome Number               ](https://leetcode-cn.com/problems/palindrome-number) | 58.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/0009.palindrome-number/main.go)| 
| 0012 | [Integer to Roman                ](https://leetcode-cn.com/problems/integer-to-roman) | 64.9% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/math/0012.integer-to-roman/main.go)| 
| 0013 | [Roman to Integer                ](https://leetcode-cn.com/problems/roman-to-integer) | 62.4% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/0013.roman-to-integer/main.go)| 
| 0029 | [Divide Two Integers             ](https://leetcode-cn.com/problems/divide-two-integers) | 20.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/math/0029.divide-two-integers/main.go)| 
| 0031 | [Next Permutation                ](https://leetcode-cn.com/problems/next-permutation) | 36.6% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/math/0031.next-permutation/main.go)| 
| 0066 | [Plus One                        ](https://leetcode-cn.com/problems/plus-one) | 45.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/0066.plus-one/main.go)| 
| 0191 | [Number of 1 Bits                ](https://leetcode-cn.com/problems/number-of-1-bits) | 73.8% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/0191.number-of-1-bits/main.go)| 
| 0292 | [Nim Game                        ](https://leetcode-cn.com/problems/nim-game) | 69.6% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/0292.nim-game/main.go)| 
| 0319 | [Bulb Switcher                   ](https://leetcode-cn.com/problems/bulb-switcher) | 49.8% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/math/0319.bulb-switcher/main.go)| 
| 1266 | [Minimum Time Visiting All Points](https://leetcode-cn.com/problems/minimum-time-visiting-all-points) | 82.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/1266.minimum-time-visiting-all-points/main.go)| 
| 1281 | [Subtract the Product and Sum of Digits of an Integer](https://leetcode-cn.com/problems/subtract-the-product-and-sum-of-digits-of-an-integer) | 83.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/1281.subtract-the-product-and-sum-of-digits-of-an-integer/main.go)| 
| 1313 | [Decompress Run-Length Encoded List](https://leetcode-cn.com/problems/decompress-run-length-encoded-list) | 82.9% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/1313.decompress-run-length-encoded-list/main.go)| 
| 1403 | [Minimum Subsequence in Non-Increasing Order](https://leetcode-cn.com/problems/minimum-subsequence-in-non-increasing-order) | 68.9% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/math/1403.minimum-subsequence-in-non-increasing-order/main.go)| 

## Stack
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0020 | [Valid Parentheses               ](https://leetcode-cn.com/problems/valid-parentheses) | 43.9% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/stack/0020.valid-parentheses/main.go)| 
| 0032 | [Longest Valid Parentheses       ](https://leetcode-cn.com/problems/longest-valid-parentheses) | 34.7% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/stack/0032.longest-valid-parentheses/main.go)| 

## String
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0003 | [Longest Substring Without Repeating Characters](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters) | 36.8% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/string/0003.longest-substring-without-repeating-characters/main.go)| 
| 0005 | [Longest Palindromic Substring   ](https://leetcode-cn.com/problems/longest-palindromic-substring) | 33.7% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/string/0005.longest-palindromic-substring/main.go)| 
| 0006 | [ZigZag Conversion               ](https://leetcode-cn.com/problems/zigzag-conversion) | 49.8% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/string/0006.zigzag-conversion/main.go)| 
| 0010 | [Regular Expression Matching     ](https://leetcode-cn.com/problems/regular-expression-matching) | 31.0% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/string/0010.regular-expression-matching/main.go)| 
| 0014 | [Longest Common Prefix           ](https://leetcode-cn.com/problems/longest-common-prefix) | 39.3% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/string/0014.longest-common-prefix/main.go)| 
| 0017 | [Letter Combinations of a Phone Number](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number) | 56.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/string/0017.letter-combinations-of-a-phone-number/main.go)| 
| 0028 | [Implement strStr()              ](https://leetcode-cn.com/problems/implement-strstr) | 39.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/string/0028.implement-strstr/main.go)| 
| 0038 | [Count and Say                   ](https://leetcode-cn.com/problems/count-and-say) | 57.4% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/string/0038.count-and-say/main.go)| 
| 0043 | [Multiply Strings                ](https://leetcode-cn.com/problems/multiply-strings) | 44.7% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/string/0043.multiply-strings/main.go)| 
| 0415 | [Add Strings                     ](https://leetcode-cn.com/problems/add-strings) | 52.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/string/0415.add-strings/main.go)| 
| 0771 | [Jewels and Stones               ](https://leetcode-cn.com/problems/jewels-and-stones) | 85.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/string/0771.jewels-and-stones/hash_table/main.go)| hash table

## Tree
|  #   | Title                                  | Acceptance  | Difficulty  | Solution         | Algorithm
| ---- | -------------------------------------- |  ---------- | ----------- |----------------- | ----------
| 0022 | [Generate Parentheses            ](https://leetcode-cn.com/problems/generate-parentheses) | 77.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0022.generate-parentheses/dfs/catalan/main.go)| dfs.catalan
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0022.generate-parentheses/dfs/recursive/main.go)| dfs.recursive
| 0039 | [Combination Sum                 ](https://leetcode-cn.com/problems/combination-sum) | 72.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0039.combination-sum/dfs/recursive/main.go)| dfs.recursive
| 0040 | [Combination Sum II              ](https://leetcode-cn.com/problems/combination-sum-ii) | 63.9% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0040.combination-sum-ii/dfs/recursive/main.go)| dfs.recursive
| 0094 | [Binary Tree Inorder Traversal   ](https://leetcode-cn.com/problems/binary-tree-inorder-traversal) | 75.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0094.binary-tree-inorder-traversal/dfs/inorder/recursive/main.go)| dfs.inorder.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0094.binary-tree-inorder-traversal/dfs/inorder/stack/main.go)| dfs.inorder.stack
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0094.binary-tree-inorder-traversal/dfs/morris/break/main.go)| dfs.morris.break
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0094.binary-tree-inorder-traversal/dfs/morris/keep/main.go)| dfs.morris.keep
| 0095 | [Unique Binary Search Trees II   ](https://leetcode-cn.com/problems/unique-binary-search-trees-ii) | 67.8% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0095.unique-binary-search-trees-ii/dfs/catalan/main.go)| dfs.catalan
| 0096 | [Unique Binary Search Trees      ](https://leetcode-cn.com/problems/unique-binary-search-trees) | 69.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0096.unique-binary-search-trees/catalan/main.go)| catalan
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0096.unique-binary-search-trees/dp/main.go)| dp
| 0098 | [Validate Binary Search Tree     ](https://leetcode-cn.com/problems/validate-binary-search-tree) | 33.9% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0098.validate-binary-search-tree/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0098.validate-binary-search-tree/dfs/inorder/recursive/main.go)| dfs.inorder.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0098.validate-binary-search-tree/dfs/inorder/stack/main.go)| dfs.inorder.stack
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0098.validate-binary-search-tree/dfs/preorder/recursive/main.go)| dfs.preorder.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0098.validate-binary-search-tree/dfs/preorder/stack/main.go)| dfs.preorder.stack
| 0099 | [Recover Binary Search Tree      ](https://leetcode-cn.com/problems/recover-binary-search-tree) | 62.0% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0099.recover-binary-search-tree/dfs/inorder/recursive/main.go)| dfs.inorder.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0099.recover-binary-search-tree/dfs/inorder/stack/main.go)| dfs.inorder.stack
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0099.recover-binary-search-tree/dfs/morris/keep.go)| dfs.morris
| 0100 | [Same Tree                       ](https://leetcode-cn.com/problems/same-tree) | 60.4% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0100.same-tree/dfs/recursive/main.go)| dfs.recursive
| 0101 | [Symmetric Tree                  ](https://leetcode-cn.com/problems/symmetric-tree) | 54.3% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0101.symmetric-tree/dfs/recursive/main.go)| dfs.recursive
| 0102 | [Binary Tree Level Order Traversal](https://leetcode-cn.com/problems/binary-tree-level-order-traversal) | 64.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0102.binary-tree-level-order-traversal/bfs/queue/main.go)| bfs.queue
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0102.binary-tree-level-order-traversal/dfs/recursive/main.go)| dfs.recursive
| 0103 | [Binary Tree Zigzag Level Order Traversal](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal) | 57.1% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0103.binary-tree-zigzag-level-order-traversal/bfs/queue/main.go)| bfs.queue
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0103.binary-tree-zigzag-level-order-traversal/dfs/recursive/main.go)| dfs.recursive
| 0104 | [Maximum Depth of Binary Tree    ](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree) | 76.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0104.maximum-depth-of-binary-tree/dfs/main.go)| dfs
| 0105 | [Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal) | 69.5% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0105.construct-binary-tree-from-preorder-and-inorder-traversal/main.go)| 
| 0106 | [Construct Binary Tree from Inorder and Postorder Traversal](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal) | 71.5% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0106.construct-binary-tree-from-inorder-and-postorder-traversal/main.go)| 
| 0107 | [Binary Tree Level Order Traversal II](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii) | 68.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0107.binary-tree-level-order-traversal-ii/bfs/queue/main.go)| bfs.queue
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0107.binary-tree-level-order-traversal-ii/dfs/recursive/main.go)| dfs.recursive
| 0108 | [Convert Sorted Array to Binary Search Tree](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree) | 75.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0108.convert-sorted-array-to-binary-search-tree/dfs/recursive/main.go)| dfs.recursive
| 0109 | [Convert Sorted List to Binary Search Tree](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree) | 76.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0109.convert-sorted-list-to-binary-search-tree/array/main.go)| array
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0109.convert-sorted-list-to-binary-search-tree/inorder/main.go)| inorder
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0109.convert-sorted-list-to-binary-search-tree/recursive/main.go)| recursive
| 0110 | [Balanced Binary Tree            ](https://leetcode-cn.com/problems/balanced-binary-tree) | 55.5% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0110.balanced-binary-tree/postorder/main.go)| postorder
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0110.balanced-binary-tree/top/main.go)| top
| 0111 | [Minimum Depth of Binary Tree    ](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree) | 47.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0111.minimum-depth-of-binary-tree/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0111.minimum-depth-of-binary-tree/dfs/main.go)| dfs
| 0112 | [Path Sum                        ](https://leetcode-cn.com/problems/path-sum) | 52.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0112.path-sum/dfs/main.go)| dfs
| 0113 | [Path Sum II                     ](https://leetcode-cn.com/problems/path-sum-ii) | 61.9% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0113.path-sum-ii/dfs/main.go)| dfs
| 0114 | [Flatten Binary Tree to Linked List](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list) | 72.1% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0114.flatten-binary-tree-to-linked-list/preorder.morris/main.go)| preorder.morris
| 0116 | [Populating Next Right Pointers in Each Node](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node) | 69.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0116.populating-next-right-pointers-in-each-node/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0116.populating-next-right-pointers-in-each-node/dfs/main.go)| dfs
| 0117 | [Populating Next Right Pointers in Each Node II](https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii) | 59.6% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0117.populating-next-right-pointers-in-each-node-ii/bfs/main.go)| bfs
| 0124 | [Binary Tree Maximum Path Sum    ](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum) | 43.5% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0124.binary-tree-maximum-path-sum/dfs/main.go)| dfs
| 0129 | [Sum Root to Leaf Numbers        ](https://leetcode-cn.com/problems/sum-root-to-leaf-numbers) | 66.9% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0129.sum-root-to-leaf-numbers/dfs/main.go)| dfs
| 0144 | [Binary Tree Preorder Traversal  ](https://leetcode-cn.com/problems/binary-tree-preorder-traversal) | 69.5% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0144.binary-tree-preorder-traversal/dfs/recursive/main.go)| dfs.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0144.binary-tree-preorder-traversal/dfs/stack/main.go)| dfs.stack
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0144.binary-tree-preorder-traversal/morris/break/main.go)| morris.break
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0144.binary-tree-preorder-traversal/morris/keep/main.go)| morris.keep
| 0145 | [Binary Tree Postorder Traversal ](https://leetcode-cn.com/problems/binary-tree-postorder-traversal) | 74.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0145.binary-tree-postorder-traversal/dfs/recursive/main.go)| dfs.recursive
| 0156 | [Binary Tree Upside Down         ](https://leetcode-cn.com/problems/binary-tree-upside-down) | 72.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0156.binary-tree-upside-down/link/main.go)| link
| 0173 | [Binary Search Tree Iterator     ](https://leetcode-cn.com/problems/binary-search-tree-iterator) | 80.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0173.binary-search-tree-iterator/main.go)| 
| 0199 | [Binary Tree Right Side View     ](https://leetcode-cn.com/problems/binary-tree-right-side-view) | 65.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0199.binary-tree-right-side-view/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0199.binary-tree-right-side-view/dfs/main.go)| dfs
| 0222 | [Count Complete Tree Nodes       ](https://leetcode-cn.com/problems/count-complete-tree-nodes) | 77.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0222.count-complete-tree-nodes/dfs/main.go)| dfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0222.count-complete-tree-nodes/two_pointer/main.go)| two pointer
| 0226 | [Invert Binary Tree              ](https://leetcode-cn.com/problems/invert-binary-tree) | 78.2% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0226.invert-binary-tree/dfs/main.go)| dfs
| 0230 | [Kth Smallest Element in a BST   ](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst) | 73.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0230.kth-smallest-element-in-a-bst/dfs/main.go)| dfs
| 0235 | [Lowest Common Ancestor of a Binary Search Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree) | 66.2% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0235.lowest-common-ancestor-of-a-binary-search-tree/dfs/main.go)| dfs
| 0236 | [Lowest Common Ancestor of a Binary Tree](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree) | 66.8% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0236.lowest-common-ancestor-of-a-binary-tree/dfs/main.go)| dfs
| 0250 | [Count Univalue Subtrees         ](https://leetcode-cn.com/problems/count-univalue-subtrees) | 64.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0250.count-univalue-subtrees/dfs/main.go)| dfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0250.count-univalue-subtrees/main.go)| 
| 0255 | [Verify Preorder Sequence in Binary Search Tree](https://leetcode-cn.com/problems/verify-preorder-sequence-in-binary-search-tree) | 46.1% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0255.verify-preorder-sequence-in-binary-search-tree/main.go)| 
| 0257 | [Binary Tree Paths               ](https://leetcode-cn.com/problems/binary-tree-paths) | 66.8% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0257.binary-tree-paths/dfs/main.go)| dfs
| 0270 | [Closest Binary Search Tree Value](https://leetcode-cn.com/problems/closest-binary-search-tree-value) | 53.2% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0270.closest-binary-search-tree-value/main.go)| 
| 0272 | [Closest Binary Search Tree Value II](https://leetcode-cn.com/problems/closest-binary-search-tree-value-ii) | 63.6% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0272.closest-binary-search-tree-value-ii/main.go)| 
| 0285 | [Inorder Successor in BST        ](https://leetcode-cn.com/problems/inorder-successor-in-bst) | 63.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0285.inorder-successor-in-bst/main.go)| 
| 0297 | [Serialize and Deserialize Binary Tree](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree) | 54.3% | Hard | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0297.serialize-and-deserialize-binary-tree/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0297.serialize-and-deserialize-binary-tree/dfs/main.go)| dfs
| 0298 | [Binary Tree Longest Consecutive Sequence](https://leetcode-cn.com/problems/binary-tree-longest-consecutive-sequence) | 56.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0298.binary-tree-longest-consecutive-sequence/main.go)| 
| 0366 | [Find Leaves of Binary Tree      ](https://leetcode-cn.com/problems/find-leaves-of-binary-tree) | 75.7% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0366.find-leaves-of-binary-tree/main.go)| 
| 0404 | [Sum of Left Leaves              ](https://leetcode-cn.com/problems/sum-of-left-leaves) | 56.8% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0404.sum-of-left-leaves/dfs/main.go)| dfs
| 0429 | [N-ary Tree Level Order Traversal](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal) | 68.5% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0429.n-ary-tree-level-order-traversal/bfs/main.go)| bfs
| 0437 | [Path Sum III                    ](https://leetcode-cn.com/problems/path-sum-iii) | 56.6% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0437.path-sum-iii/dfs/main.go)| dfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0437.path-sum-iii/hash_table/main.go)| hash table
| 0449 | [Serialize and Deserialize BST   ](https://leetcode-cn.com/problems/serialize-and-deserialize-bst) | 55.1% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0449.serialize-and-deserialize-bst/dfs/preorder/main.go)| dfs.preorder
| 0450 | [Delete Node in a BST            ](https://leetcode-cn.com/problems/delete-node-in-a-bst) | 46.7% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0450.delete-node-in-a-bst/dfs/main.go)| dfs
| 0501 | [Find Mode in Binary Search Tree ](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree) | 50.2% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0501.find-mode-in-binary-search-tree/dfs/recursive/main.go)| dfs.recursive
| 0508 | [Most Frequent Subtree Sum       ](https://leetcode-cn.com/problems/most-frequent-subtree-sum) | 66.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0508.most-frequent-subtree-sum/dfs/main.go)| dfs
| 0513 | [Find Bottom Left Tree Value     ](https://leetcode-cn.com/problems/find-bottom-left-tree-value) | 72.5% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0513.find-bottom-left-tree-value/bfs/main.go)| bfs
| 0515 | [Find Largest Value in Each Tree Row](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row) | 63.7% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0515.find-largest-value-in-each-tree-row/bfs/main.go)| bfs
| 0530 | [Minimum Absolute Difference in BST](https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst) | 61.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0530.minimum-absolute-difference-in-bst/dfs/main.go)| dfs
| 0538 | [Convert BST to Greater Tree     ](https://leetcode-cn.com/problems/convert-bst-to-greater-tree) | 66.7% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0538.convert-bst-to-greater-tree/dfs/main.go)| dfs
| 0543 | [Diameter of Binary Tree         ](https://leetcode-cn.com/problems/diameter-of-binary-tree) | 52.8% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0543.diameter-of-binary-tree/dfs/main.go)| dfs
| 0559 | [Maximum Depth of N-ary Tree     ](https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree) | 71.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0559.maximum-depth-of-n-ary-tree/bfs/main.go)| bfs
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0559.maximum-depth-of-n-ary-tree/dfs/main.go)| dfs
| 0563 | [Binary Tree Tilt                ](https://leetcode-cn.com/problems/binary-tree-tilt) | 58.9% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0563.binary-tree-tilt/dfs/main.go)| dfs
| 0572 | [Subtree of Another Tree         ](https://leetcode-cn.com/problems/subtree-of-another-tree) | 47.3% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0572.subtree-of-another-tree/dfs/main.go)| dfs
| 0589 | [N-ary Tree Preorder Traversal   ](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal) | 74.1% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0589.n-ary-tree-preorder-traversal/dfs/recursive/main.go)| dfs.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0589.n-ary-tree-preorder-traversal/dfs/stack/main.go)| dfs.stack
| 0590 | [N-ary Tree Postorder Traversal  ](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal) | 75.5% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0590.n-ary-tree-postorder-traversal/dfs/recursive/main.go)| dfs.recursive
|  | [                                ](https://leetcode-cn.com/problems/) |  |  | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0590.n-ary-tree-postorder-traversal/dfs/stack/main.go)| dfs.stack
| 0606 | [Construct String from Binary Tree](https://leetcode-cn.com/problems/construct-string-from-binary-tree) | 56.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0606.construct-string-from-binary-tree/dfs/main.go)| dfs
| 0617 | [Merge Two Binary Trees          ](https://leetcode-cn.com/problems/merge-two-binary-trees) | 78.8% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0617.merge-two-binary-trees/dfs/recursive/main.go)| dfs.recursive
| 0623 | [Add One Row to Tree             ](https://leetcode-cn.com/problems/add-one-row-to-tree) | 54.1% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0623.add-one-row-to-tree/bfs/main.go)| bfs
| 0637 | [Average of Levels in Binary Tree](https://leetcode-cn.com/problems/average-of-levels-in-binary-tree) | 68.8% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0637.average-of-levels-in-binary-tree/bfs/main.go)| bfs
| 0653 | [Two Sum IV - Input is a BST     ](https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst) | 58.2% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0653.two-sum-iv-input-is-a-bst/dfs/main.go)| dfs
| 0654 | [Maximum Binary Tree             ](https://leetcode-cn.com/problems/maximum-binary-tree) | 81.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0654.maximum-binary-tree/dfs/main.go)| dfs
| 0669 | [Trim a Binary Search Tree       ](https://leetcode-cn.com/problems/trim-a-binary-search-tree) | 67.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0669.trim-a-binary-search-tree/dfs/main.go)| dfs
| 0671 | [Second Minimum Node In a Binary Tree](https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree) | 46.2% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0671.second-minimum-node-in-a-binary-tree/dfs/main.go)| dfs
| 0687 | [Longest Univalue Path           ](https://leetcode-cn.com/problems/longest-univalue-path) | 42.9% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0687.longest-univalue-path/dfs/main.go)| dfs
| 0700 | [Search in a Binary Search Tree  ](https://leetcode-cn.com/problems/search-in-a-binary-search-tree) | 75.5% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0700.search-in-a-binary-search-tree/dfs/main.go)| dfs
| 0701 | [Insert into a Binary Search Tree](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree) | 72.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0701.insert-into-a-binary-search-tree/dfs/main.go)| dfs
| 0783 | [Minimum Distance Between BST Nodes](https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes) | 56.3% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0783.minimum-distance-between-bst-nodes/dfs/main.go)| dfs
| 0814 | [Binary Tree Pruning             ](https://leetcode-cn.com/problems/binary-tree-pruning) | 70.2% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0814.binary-tree-pruning/dfs/main.go)| dfs
| 0872 | [Leaf-Similar Trees              ](https://leetcode-cn.com/problems/leaf-similar-trees) | 62.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0872.leaf-similar-trees/dfs/main.go)| dfs
| 0938 | [Range Sum of BST                ](https://leetcode-cn.com/problems/range-sum-of-bst) | 78.0% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0938.range-sum-of-bst/main.go)| 
| 0951 | [Flip Equivalent Binary Trees    ](https://leetcode-cn.com/problems/flip-equivalent-binary-trees) | 65.9% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0951.flip-equivalent-binary-trees/dfs/main.go)| dfs
| 0965 | [Univalued Binary Tree           ](https://leetcode-cn.com/problems/univalued-binary-tree) | 68.7% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0965.univalued-binary-tree/dfs/main.go)| dfs
| 0979 | [Distribute Coins in Binary Tree ](https://leetcode-cn.com/problems/distribute-coins-in-binary-tree) | 70.3% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/0979.distribute-coins-in-binary-tree/dfs/main.go)| dfs
| 1008 | [Construct Binary Search Tree from Preorder Traversal](https://leetcode-cn.com/problems/construct-binary-search-tree-from-preorder-traversal) | 72.9% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/1008.construct-binary-search-tree-from-preorder-traversal/dfs/main.go)| dfs
| 1022 | [Sum of Root To Leaf Binary Numbers](https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers) | 69.2% | Easy | [Go](https://github.com/temporaries/leetcode/tree/master/tree/1022.sum-of-root-to-leaf-binary-numbers/dfs/main.go)| dfs
| 1302 | [Deepest Leaves Sum              ](https://leetcode-cn.com/problems/deepest-leaves-sum) | 81.4% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/1302.deepest-leaves-sum/dfs/main.go)| dfs
| 1315 | [Sum of Nodes with Even-Valued Grandparent](https://leetcode-cn.com/problems/sum-of-nodes-with-even-valued-grandparent) | 81.0% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/1315.sum-of-nodes-with-even-valued-grandparent/dfs/main.go)| dfs
| 1448 | [Count Good Nodes in Binary Tree ](https://leetcode-cn.com/problems/count-good-nodes-in-binary-tree) | 70.8% | Medium | [Go](https://github.com/temporaries/leetcode/tree/master/tree/1448.count-good-nodes-in-binary-tree/dfs/main.go)| dfs