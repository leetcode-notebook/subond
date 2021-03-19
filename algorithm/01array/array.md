## Array【线性表】

Array[数组]是用于在相邻位置存储均匀元素的数据结构。 必须在存储数据之前提供数组的大小。

[TOC]

### 基本操作

数组的性质决定了其基本操作的复杂度；

* 无序数组

  + 查找：O(n)
  + 插入：O(1)
  + 删除：O(n)

* 有序数组

  + 查找：O(Logn)【使用二分查找】
  + 插入：O(n)【最坏的情况下需要移动所有元素】
  + 删除：O(n)【最坏的情况下需要移动所有元素】

### 常用技巧

#### 双指针技巧

**双指针技巧**：一个指针负责维护目标索引，一个指针负责遍历数组。相关题目：移除元素，从排序数组中删除重复项。

#### 从尾部到头部

如果知道的数组的长度，可从尾部按照某种规则依次插入元素，例如[88 合并两个有序数组](# 88 合并两个有序数组)

#### 就地替换

in-place操作是指在不申请额外空间的情况下，通过索引找到目标值并进行就地替换的操作，该操作可以极大地减少时间和空间复杂度。

### 相关题目

#### 1两数之和

题目要求：给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出 **和为目标值** 的那 **两个** 整数，并返回它们的数组下标。

```golang
// date 2021/02/21
// 解法一：直接两层循环
func twoSum(nums []int, target int) []int {
    res := make([]int, 2, 2)
    n := len(nums)
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            if nums[i] + nums[j] == target {
                res[0], res[1] = i, j
                return res
            }
        }
    }
    return res
}
// 解法二：因为题目中有提到每种输入只会对应一种输出，所以可以用map存储已经遍历过的元素和脚标，从而达到O(1)的查找
func twoSum(nums []int, target int) []int {
    res := make([]int, 2, 2)
    set := make(map[int]int, len(nums))
    n := len(nums)
    for i := 0; i < n; i++ {
        if j, ok := set[target-nums[i]]; ok {
            res[0], res[1] = j, i
            return res
        } else {
            set[nums[i]] = i
        }
    }
    return res
}
```

#### 扩展题目：两数之和

**问题延伸：给定一个无序且元素可重复的数组a[],以及一个数sum，求a[]中是否存在两个元素的和等于sum，并输出这两个元素的下标。答案可能不止一种，请输出所有可能的答案结果集。**

这是VMware面试中的一道题，以下是个人的解法。这个算法的时间复杂度为O(n)，但是需要辅助空间。

```golang
// data 2021/02/21
// 解法一：两层循环
func twoSum(nums []int, target int) [][]int {
  res := make([][]int, 0, 16)
  n := len(nums)
  for i := 0; i < n; i++ {
    for j := i+1; j < n; j++ {
      if nums[i] + nums[j] == target {
        res = append(res, []int{i,j})
      }
    }
  }
  return res
}
```



#### 7 整数反转

```go
func reverse(x int) int {
  mMAX, mMIN := 1 << 31 - 1, -1 << 31
  res, flag := 0, 1
  if x < 0 {
    flag, x = -1, -x
  }
  for x > 0 {
    res = res * 10 + x % 10
    x /= 10
  }
  res *= flag
  if res > 0 && res > mMAX || res < 0 && res < mMIN {return 0}
  return res
}
```



#### 11 盛最多水的容器【中等】

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



#### 26 从排序数组中删除重复项

题目要求：给定一个排序数组，**原地**删除重复出现的元素，使得每个元素只出现一次，并返回新的数组长度。

思路分析

算法1：双指针，一个负责遍历元素，一个负责更新元素。

```go
func removeDuplicates(nums []int) int {
  if len(nums) <= 1 {return len(nums)}
  index, key := 1, nums[0]
  for i := 1; i < len(nums); i++ {
    if nums[i] != key {
      key = nums[i]
      nums[index] = key
      index++
    }
  }
  return index
}
// date 2019/12/28
// 算法 nums[0...i]只维护不重复的元素
func removeDuplicates(nums []int) int {
  if len(nums) <= 1 {return len(nums)}
  i, j := 0, 1  // 第一个元素肯定不重复
  for j < len(nums) {
    if nums[j] != nums[i] { // 出现不重复的元素
      // 指向同一个元素，不需要更新
      if i + 1 != j {
        nums[i+1] = nums[j]
      }
      i++
    }
    j++
  }
  return i+1
}
// date 2019/12/28
// 第二个算法的简洁版
func removeDuplicates(nums []int) int {
  if len(nums) <= 1 {return len(nums)}
  index := 0
  for i := 1; i < len(nums); i++ {
    if nums[i] != nums[index] {
      index++
      nums[index] = nums[i]  // 即是指向同一个元素，更新也没问题
    }
  }
  return index+1
}
```

#### 27 移除元素【简单】

题目链接：https://leetcode-cn.com/problems/remove-element/

题目要求：给你一个数组 `nums` 和一个值 `val`，你需要 **[原地](https://baike.baidu.com/item/原地算法)** 移除所有数值等于 `val` 的元素，并返回移除后数组的新长度。

算法1：

```golang
// date 2021/03/13
// index作为新的索引，只记录符合条件的(不等于val)的元素
func removeElement(nums []int, val int) int {
    var index int
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[index] = nums[i]
            index++
        }
    }
    return index
}
```

算法2：交换的思想，将每个等于val的元素，交换至数组的尾部，维护尾部索引，返回新的尾部索引。时间复杂度O(n) 空间复杂度O(n)

```go
func removeElements(nums []int, val int) int {
  i, tail := 0, len(nums)
  for i < tail; {
    if nums[i] == val {
      tail--
      nums[i] = nums[tail]
      nums[tail] = val
    } else {
      i++
    }
  }
  return tail
}
```



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

#### 53 最大子序和

题目链接：https://leetcode-cn.com/problems/maximum-subarray/

题目要求：给定一个整数数组，找到一个具有最大和的连续子数组，并返回其值。

思路分析：

```go
// date 2020/05/03
func maxSubArray(nums []int) int {
    if len(nums) == 0 { return 0 }
    res := nums[0]
    var temp, cur_max int
    for _, v := range nums {
        cur_max + v
        // 如果当前结果小于零，则重新计算;否则将当前元素加入结果集
        if 0 > cur_max {
            cur_max = v
        } else {
            cur_max += v
        }
        // 当前结果集大于最终结果集时，更新最终结果集
        if cur_max > res { res = cur_max }
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



#### 66 加一【简单】

题目要求：给定一个非负整数的数组表现形式，返回其+1的值。

细节题：注意进位。

```go 
func plusOne(digits []int) []int {
  carry, temp := 1, 0
  for i := 0; i < len(digits); i++ {
    temp = digits[i]+carry
    digits[i] = temp % 10
    carry = temp / 10
  }
  if carry == 1 {
    return append([]int{1}, digits...)
  }
  return digits...
}
```



#### 75 颜色分类【中等】

思路分析：两个指针

i维护元素为0的脚标，j维护元素为2的脚标

```go
func sortColors(nums []int) {
  k, i, j := 0, 0, len(nums) - 1
  for k <= j {
    if nums[k] == 2 {
      nums[k] = nums[j]
      nums[j] = 2
      j--
    } else if nums[k] == 0 {
      nums[k] = nums[i]
      nums[i] = 0
      i++
      k++
    } else {
      k++
    }
  }
}
```



#### 78 子集

题目要求：给定一个不含重复元素的数组，返回其所有可能的子集。

算法：组合+非递归实现

```go
// date 2020/01/11
/* 算法
1. 外层遍历nums，内层遍历结果集，并取出中间结果集
2. 用中间结果集+num组成新的结果集，并追加到结果中。
*/
func subsets(nums int) [][]int {
  res := make([][]int, 0)
  res = append(res, make([]int, 0))
  for _, v := range nums {
    n := len(res)
    for i := 0; i < n; i++ {
      t := make([]int, 0)
      t = append(t, res[i]...)
      t = append(t, v)
      res = append(res, t)
    }
  }
  return res
}
```



#### 80 从排序数组中删除重复项 II

题目要求：原地删除排序数组中的重复项，使每个元素最多出现两次。

思路分析：

```go
// date 2019/12/28
// 算法 nums[0...index]只保存符合要求的元素
func removeDuplicates2(nums []int) int {
  if len(nums) <= 2 {return len(nums)}
  var index, i int
  // basic case, update index and i
  if nums[0] == nums[1] {
    index, i = 1, 2
  } else {
    index, i = 0, 1
  }
  value, n := 0, len(nums)
  for ; i < n; i++ {
    // find new value
    if nums[index] != nums[i] {
      value = nums[i]
      index++
      nums[index] = value
      if i + 1 < n && nums[i+1] == value {
        index++
        nums[index] = value
      }
    }
  }
  return index+1
}
// 简洁版
func removeDuplicates2(nums []int) int {
  if len(nums) <= 2 {return len(nums)}
  index := 1
  before, after := 0, 2
  for after < len(nums) {
    if nums[index] != nums[after] || (nums[index] == nums[after] && nums[after] != nums[before]) {
      before = index
      index++
      nums[index] = nums[after]
    }
    after++
  }
  return index+1
}
```



#### 88 合并两个有序数组

题目要求：https://leetcode.com/problems/merge-sorted-array/

思路分析：

```go
// date 2020/04/19
func merge(nums1 []int, m int, nums2 []int, n int)  {
    ia, ib, icurr := m-1, n-1, m+n-1
    for ia >= 0 || ib >= 0 {
        if ia < 0 {
            nums1[icurr] = nums2[ib]
            ib--
            icurr--
            continue
        }
        if ib < 0 {
            nums1[icurr] = nums1[ia]
            ia--
            icurr--
            continue
        }
        if nums1[ia] > nums2[ib] {
            nums1[icurr] = nums1[ia]
            ia--
        } else {
            nums1[icurr] = nums2[ib]
            ib--
        }
        icurr--
    }
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



#### 167 两数之和II【简单】

题目要求：输入数组排序的，且为升序

思路分析：利用双指针逼近结果

```go
// date 2019/12/28
// 双指针
func twoSum(nums []int, target int) []int {
  res := make([]int, 0, 2)
  i, j := 0, len(nums) - 1
  var temp int
  for i < j {
    temp = nums[i] + nums[j]
    if temp == target {
      res = append(res, i+1, j+1)
      break
    } else if temp > target {
      j--
    } else {
      i++
    }
  }
  return res
}
```



#### 189 旋转数组

题目要求：给定一个数组，将数组中的元素向右移动K个位置，K是非负整数。

思路分析

算法1：暴力法，先移动一步，然后移动K步。时间复杂度O(n*k)，空间复杂度O(1)

```go
// date 2019/12/21
func rotate(nums []int, k int) {
  k = k % len(nums)
  for k > 0 {
    moveOneStep(nums)
    k--
  }
}
func moveOneStep(nums []int) {
  if len(nums) <= 1 {return}
  t := nums[len(nums)-1]
  for i := len(nums) - 1; i >= 1; i-- {
    nums[i] = nums[i-1]
  }
  nums[0] = t
}
```

算法2：新建额外数组，将原来数组中的每一个元素按照右移K的位置的方式，放入新数组，然后在遍历新数组重新赋值原始数组。时间复杂度O(n) 空间复杂度O(n)

```go
// date 2020/01/04
func rotate(nums []int, k int) {
  k = k % len(nums)
  a := make([]int, len(nums))
  for i := 0; i < len(nums); i++ {
    a[(i+k)%n] = nums[i]
  }
  for i := 0; i < len(nums); i++ {
    nums[i] = a[i]
  }
}
```

算法3：反转数组

当我们旋转数组k次，k%n个尾部元素会移动到数组的头部，剩下的元素被向后移动。那么，先将数组整体反转一次，然后反转前k个元素，在反转后面的n-k个元素，即是想要的结果。

```go
// date 2020/01/04
func rotate(nums []int, k int) {
  k %= len(nums)
  reverse(nums, 0, len(nums)-1)
  reverse(nums, 0, k - 1)
  reverse(nums, k, len(nums)-1)
  
}
func reverse(nums[]int, start, end int) {
  for start < end {
    temp := nums[end]
    nums[end] = nums[start]
    nums[start] = temp
    end--
    start++
  }
}
```



#### 209 长度最小的子数组【中等】

思路分析：

1统计所有的值，如果其和小于s，则return 0

2.从当前值开始递增，直到大于等于s并更新结果

```go
// date 2019/12/28
// 步长法
func minSubArrayLen(s int, nums []int) int {
  res := len(nums)
  var sum, total int
  for i := 0; i < len(nums); i++ {
    if nums[i] >= s {return 1}
    total += nums[i]
    sum = nums[i]
    for j := i+1; j < len(nums); j++ {
      sum += nums[j]
      if sum >= s {
        if j - i + 1 < res {
          res = j - i + 1
        }
        break
      }
    }
  }
  if total < s {return 0}
  return res
}
```

算法二 双指针

双指针 维护滑动窗口[i, j]，使其在不满足total >= s时，j++；或者满足条件total -= nums[i]

```go
// date 2019/12/28
// 双指针
func minSubArrayLen(s int, nums []int) int {
  i, j := 0, -1
  total, res := 0, len(nums) + 1
  for i < len(nums) {
    if j + 1 < len(nums) && total < s {
      j++
      total += nums[j]
    } else {
      total -= nums[i]
      i++
    }
    if total >= s && j - i + 1 < res {
      res = j-i+1
    }
  }
  if res = len(nums) + 1 {return 0}
  return res
}
```



#### 215 数组中的第K个最大元素

题目要求：https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

思路分析：降序的快排。把任务切分成小任务swap。

```go
// date 2020/03/29
func findKthLargest(nums []int, k int) int {
    if k > len(nums) {return -1}
    // 快排排序
    QuickSort(nums, 0, len(nums)-1, k-1)
    return nums[k-1]
}

func QuickSort(nums []int, l, r, k int) int {
    p := partition(nums, l, r)
    // 如果已经部分有序，且找到k则返回，减少排序次数
    if p == k {
        return p
    } else if p > k {
        // 左边继续排序
        return QuickSort(nums, l, p-1, k)
    }
    // 右边继续排序
    return QuickSort(nums, p+1, r, k)
}

func partition(nums []int, l, r int) int {
    // v基值,i维护大于等于基值的索引，j遍历数组
    v, i, j := nums[l], l, l+1
    for j <= r {
        // 如果元素大于基值，则进行交换
        if nums[j] >= v {
            swap(nums, i+1, j)
            i++
        }
        j++
    }
    swap(nums, i, l)
    return i
}

func swap(nums []int, x, y int) {
    t := nums[x]
    nums[x] = nums[y]
    nums[y] = t
}
// 进化版
func findKthLargest(nums []int, k int) int {
    if k > len(nums) {return -1}
    // 快排排序
    var swap func(i, j int)
    var partition func(left, right int) int
    var quickSort func(left, right, k int)
    swap = func(i, j int) {
        t := nums[i]
        nums[i] = nums[j]
        nums[j] = t
    }
    partition = func(left, right int) int {
        v, i, j := nums[left], left, left+1
        for j <= right {
            if nums[j] >= v {
                swap(i+1, j)
                i++
            }
            j++
        }
        swap(i, left)
        return i
    }
    
    quickSort = func(left, right, k int) {
        p := partition(left, right)
        if p == k {
            return
        } else if p > k {
            quickSort(left, p-1, k)
        } else {
            quickSort(p+1, right, k)
        }
    }
    
    quickSort(0, len(nums)-1, k-1)
    return nums[k-1]
}
```



#### 219 存在重复元素 II【简单】

思路分析

算法1：利用map

```go
// data 2020/01/01
// map
func containsNearbyDuplicate(nums []int, k int) bool {
  m := make(map[int]int)
  for index, v := range nums {
    if j, ok := m[v]; ok && index - j <= k {
      return true
    }
    m[v] = index
  }
  return false
}
```

算法2：利用双指针

```go
// data 2020/02/02
// two points
func containNearbyDuplicate(nums []int, k int) bool {
  i, j, n := 0, 0, len(nums)
  for i < n {
    j = i + 1
    for j - i <= k && j < n{
      if nums[i] == nums[j] {return true}
      j++
    }
    i++
  }
  return false
}
```

算法3：滑动窗口

```go
// date 2020/01/01
func containsNearbyDuplicates(nums []int, k int) bool {
  queue := make([]int, 0)
  for _, v := range nums {
    if inQueue(queue, v) {return true}
    queue = append(queue, v)
    if len(queue) > k {
      queue = queue[1:]
    }
  }
  return false
}

func inQueue(nums []int, n int) bool {
  for _, v := range nums {
    if v == n {return true}
  }
  return false
}
```

#### 220 存在重复元素 III【中等】

思路分析

算法1 双指针，用双指针控制窗口大小，时间复杂度O(n)

```go
// date 2020/01/01
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
  i, j, n := 0, 0, len(nums)
  for i < n {
    j = i + 1
    for j - i <= k && j < n {
      if nums[i] < nums[j] && nums[j] - nums[i] <= t {return true}
      if nums[i] >= nums[j] && nums[i] - nums[j] <= t {return true}
      j++
    }
    i++
  }
  return false
}
```



#### 283 [移动零](https://leetcode-cn.com/problems/move-zeroes/)

题目要求：给定一个数组nums，将其所有的零移动至数组的末尾，其他非零元素保持相对位置不变。

思路分析：

维护好非零元素应该使用的索引，遍历数组，将非零元素放在其索引。

```go
// date 2021/03/13
func moveZeros(nums []int) {
  index_nzero := 0
  for i := 0; i < len(nums); i++ {
    if nums[i] != 0 {
      // 只要当前面有零的时候才操作
      if index_nzero < i {
        nums[index_nzero] = nums[i]
        nums[i] = 0
      }
      index_nzero++
    }
  }
}
```



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



#### 498 对角线遍历

题目要求：给定一个二维数组，返回其对角线遍历序列。

```go
// date 2020/02/21
func findDiagonalOrder(matrix [][]int) []int {
  res := make([]int, 0)
  m := len(matrix)
  if m == 0 { return res }
  n := len(matrix[0])
  if n == 0 { return res }
  bXFlag := true
  pm, pn := 0, 0
  x, y := 0, 0
  for i := 0; i < m+n-1; i++ {
    if bXFlag {
      pm = m
      pn = n
    } else {
      pm = n
      pn = m
    }
    // 初始值
    // 每一趟x尽量取最大，要么是上限，要么是i
    if i < pm {
      x = i
    } else {
      x = pm - 1
    }
    y = i - x
    // 结束值
    // x减到0，或者y达到上限
    for x >= 0 && y < pn {
      if bXFlag {
        res = append(res, matrix[x][y])
      } else {
        res = append(res, matrix[y][x])
      }
      x--
      y++
    }
    bXFlag = !bXFlag
  }
  return res
}
```



#### 674 最长连续递增序列

题目要求：https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/

思路分析：

```go
// date 2020/03/29
// 算法1
func findLengthOfLCIS(nums []int) int {
    if len(nums) < 1 { return 0 }
    // res记录结果，left记录连续递增序列中最左边的元素
    res, left := 1, 0
    for i := 1; i < len(nums); i++ {
        if nums[i] <= nums[i-1] {
            if res < i - left { res = i - left }
            left = i
        }
    }
    if res < len(nums) - left { res = len(nums) - left }
    return res
}
// 算法2:记录当前最大
func findLengthOfLCIS(nums []int) int {
  if len(nums) < 1 { return 0 }
  res, cur_max := 1, 1
  for i := 0; i < len(nums)-1; i++ {
    if nums[i] < nums[i+1] {
      cur_max++
    } else {
      cur_max = 1
    }
    if res < cur_max { res = cur_max }
  }
  return res
}
```



#### 724 寻找数组的中心索引【简单】

题目要求：给定一个数组，返回其中心索引。

备注：中心索引是指其左边个元素之和等于其右边个元素之和。

算法：先求和，再依次遍历。时间复杂度O(2n)

```go
// date 2020/02/01
func pivotIndex(nums []int) int {
  if len(nums) < 1 { return -1 }
  total, sum := 0, 0
  for _, v := range nums { total+= v}
  for i := -1; i < len(nums) - 1; i++ {
    if sum == total - nums[i+1] {return i+1}
    sum += nums[i+1]
    total -= nums[i+1]
  }
  return -1
}
```



#### 747 至少是其他数字两倍的最大数【简单】

题目要求：给定一个数组，如果数组最大的元素至少是其他元素的两倍，则返回其索引，否则返回-1。

算法：最大元素只要大于次大元素的两倍即可。

```go
func dominantIndex(nums []int) int {
  if len(nums) <= 1 { return len(nums)-1 }
  res := -1
  max1, max2 := 0, 0
  for i, v := range nums {
    if v > max1 {
      max2 = max1
      max1 = v
      res = i
    } else if v > max2 {
      max2 = v
    }
  }
  if max1 >= 2 * max2 { return res }
  return -1
}
```



#### 905 [按奇偶排序数组](https://leetcode-cn.com/problems/sort-array-by-parity/) 【简单】

题目要求：给定一个数组，通过就地修改使所有偶数位于所有基数的前面。你可以返回满足此条件的任何数组作为答案。

思路分析：

```go
// date 2021/03/13
func sortArrayByParity(A []int) []int {
    i, j := 0, len(A)-1
    for i < j {
        if A[i] & 0x1 == 0 { i++; continue }
        if A[j] & 0x1 == 1 { j--; continue }
        A[i], A[j] = A[j], A[i]
        i++
        j--
    }
    return A
}
```

#### 977 [有序数组的平方](https://leetcode-cn.com/problems/squares-of-a-sorted-array/)【简单】

题目要求：给定一个非递增的有序数组，返回其非递增的平方数组

思路要求：

```golang
// date 2021/03/13
func sortedSquares(nums []int) []int {
    total := len(nums)
    index, res := total-1, make([]int, total)
    i, j := 0, len(nums)-1
    for i <= j {
        if i == j {
            res[index] = nums[i] * nums[i]
            break
        }
        si, sj := nums[i] * nums[i], nums[j] * nums[j]
        if si > sj {
            res[index] = si
            i++
        } else {
            res[index] = sj
            j--
        }
        index--
    }
    return res
}
```

#### 1089 重复零[简单]

题目分析：https://leetcode.com/problems/duplicate-zeros/

思路分析：

```go
// date 2020/04/19
func duplicateZeros(arr []int)  {
    n := len(arr)
    j := 0
    for i := 0; i < n; i++ {
        if arr[i] == 0 {
            j = n - 2
            for j > i {
                arr[j+1] = arr[j]
                j--
            }
            arr[j+1] = 0
            i++
        }
    }
}
```

#### 1295 找出具有偶数位数字的个数[简单]

题目要求：https://leetcode.com/problems/find-numbers-with-even-number-of-digits/

思路分析：

```go
// date 2020/04/19
func findNumbers(nums []int) int {
    res := 0
    for _, v := range nums {
        if isEvenDigitsOfNumber(v) {
            res++
        }
    }
    return res
}

func isEvenDigitsOfNumber(num int) bool {
    digits := 0 
    for num > 0 {
        digits++
        num /= 10
    }
    return digits & 0x1 == 0
}
```

#### 1299 [将每个元素替换为右侧最大元素](https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/)

题目要求：

给你一个数组 `arr` ，请你将每个元素用它右边最大的元素替换，如果是最后一个元素，用 `-1` 替换。

完成所有替换操作后，请你返回这个数组。

思路分析：

```go
// date 2021/03/13
func replaceElements(arr []int) []int {
    cur, right_max := -1, -1
    for i := len(arr)-1; i >= 0; i-- {
        // 保存当前值
        cur = arr[i]
        // 将当前值赋值为右侧最大值
        arr[i] = right_max
        // 更新右侧最大值
        if cur > right_max {
            right_max = cur
        }
    }
    return arr
}
```

