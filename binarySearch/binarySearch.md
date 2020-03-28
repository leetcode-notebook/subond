## 二分查找

二分查找描述了在有序集合中查找特定值的过程。



### 相关题目
- 704 二分查找
- 162 寻找峰值
- 658 找到K个最接近的元素 

#### 704 二分查找

题目要求：给定一个非降序数组和一个目标值，如果数组中存在目标值则返回其索引，否则返回-1。

算法：二分查找

```go
func search(nums []int, target int) int {
  left, right := 0, len(nums)-1
  var mid int
  for left <= right {
    mid = left + (right-left) / 2
    if nums[mid] == target { return mid }
    if nums[mid] > target {
      right = mid-1
    } else {
      left = mid+1
    }
  }
  return -1
}
```

#### 162 寻找峰值

题目要求：给定一个数组，返回其峰值的索引。

算法1：遍历法，分为升序，降序和峰值在某一点。

```go
// date 2020/02/02
func findPeakElement(nums []int) int {
  for i := 0; i < len(nums)-1; i++ {
    if nums[i] > nums[i+1] { return i }
  }
  return len(nums)-1
}
```

算法2：递归二分查找

时间复杂度O(logN)，空间复杂度O(logN)

```go
func findPeakElement(nums []int) int {
  return search(nums, 0, len(nums)-1)
}
func search(nums []int, l, r int) int {
  if l == r { return l }
  mid := (l+r)/2
  if nums[mid] > nums[mid+1] {
    return search(nums, l, mid)
  }
  return search(nums, mid+1, r)
}
```

算法3：迭代二分查找

时间复杂度O(logN)，空间复杂度O(1)

```go
func findPeakElement(nums []int) int {
  l, r := 0, len(nums)-1
  var mid int
  for l < r {
    mid = (l+r) / 2
    if nums[mid] > nums[mid+1] {
      r = mid
    } else {
      l = mid+1
    }
  }
  return l
}
```

#### 658 找到K个最接近的元素