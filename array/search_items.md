## 查找元素



#### 有效的"山型"数组

题目链接：https://leetcode.com/problems/valid-mountain-array/

题目要求：判断一个数组是否为“山型”数组。山型数组的定义如下：

```sh
1. A.length >= 3
2. 0 < i < A.length-1 && A[0] < A[1] < ... < A[i] && A[i] > A[i+1] > ... > A[A.length-1]
```

思路分析：

```go
// date 2020/05/20
func validMountainArray(A []int) bool {
    if len(A) < 3 { return false }
    top, index := A[0], 0
    n := len(A)
    for i := 1; i < n; i++ {
        if A[i] == A[i-1] { return false }
        if A[i] > top {
            top = A[i]
            index = i
        }
    }
    // 斜坡
    if index == 0 || index == n-1 { return false }
    
    for i := 1; i < index; i++ {
        if A[i] <= A[i-1] { return false }
    }
    for i := index+1; i < n; i++ {
        if A[i-1] <= A[i] { return false }
    }
    
    return true
}
```

#### 485 最大连续1的个数【简单】

题目要求：给定一个只包含1和0的数组，返回其中最大连续1的个数。

算法：双指针

```go
// date 2020/02/02
func findMaxConsecutiveOnes(nums []int) int {
  res, n := 0, len(nums)
  start, end := 0, 0
  for start < n && nums[start] == 0 { start++ }
  end = start
  for end < n {
    // find the first 0
    if nums[end] != 1 {
      if res < (end-start) {
        res = end - start
      }
      for end < n && nums[end] == 0 {end++}
      start = end
    }
    end++
  }
  if res < (end - start) {
    res = end - start
  }
  return res
}
```

#### 盛最多水的容器

思路分析:双指针

```go
// date 2019/12/28
// 双指针
func maxArea(height []int) int {
  var res, temp int
  i, j := 0, len(height) - 1
  for i < j {
    if height[i] < height[j] {
      temp = height[i] * (j-i)
      i++
    } else {
      temp = height[j] * (j-i)
      j--
    }
    if temp > res {
      res = temp
    }
  }
  return res
}
```

