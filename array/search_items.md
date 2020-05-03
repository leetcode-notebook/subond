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

