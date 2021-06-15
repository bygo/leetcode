# 二分法
- `mid := l + (r-l)/2` 防止溢出 下降取整
- `l<=r` 最后判断

### [二分查找](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree)

```go
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		cur := nums[mid]
		if cur < target {
			l = mid + 1
		} else if target < cur {
			r = mid - 1
		} else if target == cur {
			return mid
		}
	}
	return -1
}
```

