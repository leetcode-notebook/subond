## Array【线性表】

题目列表

* 对角线遍历【M】

- 27 Remove Element[两种算法]
- 189 旋转数组
- 从排序数组中删除重复项
- 80 从排序数组中删除重复项 II
- 54 螺旋矩阵
- 整数反转
- 118 杨辉三角【E】

#### 对角线遍历二维数组

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



#### 27 移除元素

题目要求：给定一个数组和一个值val，移除数组中所有值为val的元素，并返回数组的新长度。

算法1：维护index表示每个非val的下标。时间复杂度O(n) 空间复杂度O(n)

```go
// date 2019/12/28
func removeElement(nums []int, val int) int {
  index := -1
  for i := 0; i < len(nums); i++ {
    if nums[i] != val {
      index++
      nums[index] = nums[i]
    }
  }
  return index+1
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
      nums[n] = val
    } else {
      i++
    }
  }
  return tail
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



#### 整数反转

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



#### 27 Remove Element[两个算法]

题目要求：

给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

思路分析：

将目标元素交换值数组的后面

```go
// date: 2019/12/01
// 算法思路：遍历数组将所有等于val的元素移动到数组的尾部，记录尾部元素的索引
func removeElement(nums []int, val int) int {
  i, last_index := 0, len(nums) - 1
  for ; i <= last_index; i++ {
    if nums[i] == val {
      // 
      nums[i] = nums[last_index]
      nums[last_index] = val
      last_index--
      i--
    }
  }
  return last_index+1
}
// date: 2019/12/28
// 算法思路：nums[1..i]只维护不等于val的元素
func removeElement(nums []int, val int) int {
  index_new := -1
  for i := 0; i < len(nums); i++ {
    if nums[i] != val {
      index_new++
      nums[index_new] = nums[i]
    }
  }
  return index_new+1
}
```

#### 合并两个有序数组

思路分析

算法1：从后往前按个排序每个元素。

```go
func merge(nums1 []int, m int, nums2 []int, n int) {
  ia, ib, icurr := m-1, n-1, m+n-1
  for ia >= 0 && ib >= 0 {
    if nums1[ia] >= nums2[ib] {
      nums1[icurr] = nums1[ia]
      ia--
    } else {
      nums1[icurr] = nums2[ib]
      ib--
    }
    icurr--
  }
  if ib >= 0 {
    nums1[icurr] = nums2[ib]
    ib--
    icurr--
  }
}
```

#### 移动零

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

#### 颜色分类

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

#### 数组中的第K个最大元素

```go
func findKthLargest(nums []int, k int) int {
  if k > len(nums) {return}
  return nums[QuickSort(nums, 0, len(nums)-1, k)]
}

func QuickSort(nums []int, l, r, k) int {
  if l >= r {return l}
  p := partition(nums, l, r)
  if p+1 == k {
    return p
  }
  if p+1 > k {
    return QuickSort(nums, l, p-1, k)
  } else {
    return QuickSort(nums, p+1, r, k)
  }
}

func partition(nums []int, l, r int) int {
  v, i, j := nums[l], l, l+1
  var temp int
  for j <= r {
    if nums[j] >= v {
	    temp = nums[i+1]
      nums[i+1] = nums[j]
      nums[j] = temp
      i++
    }
    j++
  }
  temp = nums[i]
  nums[i] = v
  nums[l] = temp
  return i
}
```

#### 两数之和II

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

#### 长度最小的子数组

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

#### 存在重复元素 II

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

#### 存在重复元素 III

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

#### 回旋镖的数量

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

