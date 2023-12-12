# Array

Array[数组]是用于在相邻位置存储均匀元素的数据结构。 必须在存储数据之前提供数组的大小。

[TOC]

## 1、基本操作

数组的性质决定了其基本操作的复杂度；

* 无序数组

  + 查找：O(n)
  + 插入：O(1)
  + 删除：O(n)

* 有序数组

  + 查找：O(Logn)【使用二分查找】
  + 插入：O(n)【最坏的情况下需要移动所有元素】
  + 删除：O(n)【最坏的情况下需要移动所有元素】

## 2、常用技巧

### 2.1 双指针算法

双指针，是指在遍历对象的过程中，不在普通地使用单个指针进行访问，而是两个相同方向（快慢指针）或相反方向（对撞指针）的指针进行遍历，从而达到相应的目的。

**快慢指针**：一个指针负责维护目标索引（因为符合条件的元素较少，所以这个指针移动较慢，也称为慢指针），一个指针负责遍历数组（也称为快指针）。

相关题目：

- [026删除排序数组中的重复项](026.md)
- [027移除元素](027.md)
- [209长度最小的子数组](209.md)



**对撞指针**：两个指针分别从最左侧（left指针）和最右侧（right指针）向数组中间进行遍历。

相关题目：

- [011盛最多水的容器](011.md)
- [015三数之和](015.md)
- [034在排序数组中查找元素的第一个和最后一个位置](034.md)





#### 从尾部到头部

如果知道的数组的长度，可从尾部按照某种规则依次插入元素，例如[88 合并两个有序数组](088.md)

#### 就地替换

in-place操作是指在不申请额外空间的情况下，通过索引找到目标值并进行就地替换的操作，该操作可以极大地减少时间和空间复杂度。



### 2.2 前缀和技巧





---

相关题目

#### 46 全排列

题目要求：给定一个**没有重复**数字的序列，返回其所有可能的全排列。

算法：回溯法

```go
func permute(nums []int) [][]int {
  res := make([][]int, 0, total(len(nums)))
  func bt(nums, temp []int) {
    if len(nums) == 0 {
      res = append(res, temp)
      return
    }
    for i := range nums {
      temp = append(temp, nums[i])
      n := make([]int, 0, len(nums)-1)
      n = append(n, nums[:i]...)
      n = append(n, nums[i+1:]...)
      bt(n, temp)
    }
  }(nums, make([]int, 0))
  return res
}

func total(n int) int {
  if n <= 0 {return 0}
  res := 1
  for n >= 1 {
    res *= n
    n--
  }
  return res
}
```



#### 54 螺旋矩阵

题目要求：给定一个m * n的二维矩阵，顺时针打印。

算法1：for循环，用x,y坐标加以限制

```go
func spiralOrder(matrix [][]int) []int {
  if len(matrix) < 1 { return []int{} }
  startX, endX := 0, len(matrix[0])-1
  startY, endY := 0, len(matrix)-1
  res := make([]int, 0, (endX+1)*(endY+1))
  for {
    for i := startX, i <= endX; i++ {
      res = append(res, matrix[startY][i]) // left -> right
    } 
    startY++
    if stratY > endY { break }
    for i := startY; i <= endY; i++ {
      res = append(res, matrix[i][endX])   // top -> bottom
    }
    endX--
    if startX > endX { break }
    for i := endX; i >= startX; i-- {
      res = append(res, matrix[endY][i])  // right -> left
    }
    endY--
    if startY > endY { break }
    for i := endY; i >= startY; i-- {
      res = append(res, matrix[i][startX])  // bottom -> top
    }
    startX++
    if startX > endX { break }
  }
  return res
}
```

算法2: 每次打印一轮

```go
func spiralOrder(matrix [][]int) []int {
  if len(matrix) < 1 { return []int{}}
  m, n := len(matrix)-1, len(matrix[0]) - 1
  res := make([]int, 0, (m+1)*(n+1))
  k := 0
  for k <= len(matrix)/2 {
    for i := k; i <= n-k; i++ {
      res = append(res, matrix[k][i])  // left -> right
    }
    for i := k+1; i <= m-k; i++ {
      res = append(res, matrix[i][n-k])  // top -> bottom
    } 
    for i := n-k-1; i >= k; i-- {
      res = append(res, matrix[m-k][i])  // right -> left
    }
    for i := m-k-1; i >= k+1; i-- {
      res = append(res, matrix[i][k])    // bottom -> top
    }
    k++
  }
  return res
}
```

#### 118 杨辉三角

题目要求：给定一个非负整数，返回其杨辉三角二维数组。

算法：下一行由上一行生成，生成规则如下：

```go
1）rows[i][0] = rows[i-1][0]
2）rows[i][j] = rows[i-1][j] + rows[i-1][j-1]
3）rows[i][end] = 1
```

```go
// date 2020/02/02
func generate(numRows int) [][]int {
  if numRows <= 0 { return [][]int{} }
  res := make([][]int, numRows)
  for i := 0; i < numRows; i++ {
    t := make([]int, 0)
    if i == 0 {
      t = append(t, 1)
    }
    res[i] = t
  }
  for i := 1; i < numRows; i++ {
    n := len(res[i-1])
    for j := 0; j < n; j++ {
      if j == 0 {
        res[i][j] = res[i-1][j]
      } else {
        res[i][j] = res[i-1][j] + res[i-1][j-1]
      }
    }
    res[i] = append(res[i], 1)
  }
  return res
}
```

#### 447 回旋镖的数量【中等】

思路分析

算法1：以每个点当做第一个点，计算该点与其他点的距离，如果这个距离出现过那么必定有两个答案（j，k顺序交换）

两层循环，外层循环用map记录距离

```go
// date 2020/01/01
func numberOfBoomerangs(points [][]int) int {
  res, d, n := 0, 0, len(points)
  for i := 0; i < n; i++ {
    m := make(map[int]int)
    for j := 0; j < n; j++ {
      if j == i {continue}
      d = (points[i][0] - points[j][0]) * (points[i][0] - points[j][0]) + (points[i][1] - points[j][1]) * (points[i][1] - points[j][1])
      if _, ok := m[d]; ok {
        res += m[d] * 2
        m[d]++
      } else {
        m[d] = 1
      }
    }
  }
  return res
}
```
