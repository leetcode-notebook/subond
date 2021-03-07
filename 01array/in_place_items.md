## 就地替换

[TOC]

### 介绍

in-place操作是指在不申请额外空间的情况下，通过索引找到目标值并进行就地替换的操作，该操作可以极大地减少时间和空间复杂度。

### 相关题目

#### Replace Elements with Greatest Element on Right Side

题目链接：
题目要求：https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/
思路分析：

```go
func replaceElements(arr []int) []int {
    if len(arr) < 1 { return arr }
    right_max := -1
    var temp int
    for i := len(arr)-1; i >= 0 ; i-- {
        temp = arr[i]
        arr[i] = right_max
        if temp > right_max { 
            right_max = temp
        }
    }
    return arr
}
```

#### 移动零

题目链接：

题目要求：给定一个数组nums，将其所有的零移动至数组的末尾，其他非零元素保持相对位置不变。

思路分析：

维护好非零元素应该使用的索引，遍历数组，将非零元素放在其索引。

```go
func moveZeros(nums []int) {
  index_nzero := 0
  for i := 0; i < len(nums); i++ {
    if nums[i] != 0 {
      if index_nzero < i {
        nums[index_nzero] = nums[i]
        nums[i] = 0
      }
      index_nzero++
    }
  }
}
```

#### Sort Array By Parity

题目链接：

题目要求：给定一个数组，通过就地修改使所有偶数位于所有基数的前面

思路分析：

```go
// date 2020/05/3
func sortArrayByParity(A []int) []int {
    i, j := 0, len(A)-1
    for i < j {
        if A[i] & 0x1 == 0 { i++; continue }
        if A[i] & 0x1 == 1 {
            for j > i && A[j] & 0x1 == 1 { j-- }
            if j == i { break }
            swap(A, i, j)
            i++
            j--
        }
    }
    return A
}

func swap(A []int, i, j int) {
    temp := A[i]
    A[i] = A[j]
    A[j] = temp
}
```

#### Squares of a Sorted Array

题目链接：

题目要求：给定一个非递增的有序数组，返回其非递增的平方数组

思路要求：

```go
// date 2020/05/03
func sortedSquares(A []int) []int {
    left, right := 0, len(A)-1
    res := make([]int, len(A))
    index, x, y := len(A)-1, 0, 0
    for left <= right {
        x, y = square(A[left]), square(A[right])
        if x > y {
            res[index] = x
            left++
        } else {
            res[index] = y
            right--
        }
        index--
    }
    return res
}

func square(x int) int {
    return x * x
}
```

