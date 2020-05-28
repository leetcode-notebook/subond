### 排序算法



相关题目

- 252 会议室
- 280 摆动排序【M】
- 56 合并区间【M】
- 912 排序数组【M】

- 1356 根据数字二进制下1的数目排序【E】

#### 252 会议室

思路分析

算法：直接对区间进行排序，判断是否有交叉

```go
func canAttendMeetings(intervals [][]int) bool {
  sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0], intervals[j][0]
  })
  for i := 0; i < len(intervals) - 1; i++ {
    if intervals[i][1] > intervals[i+1][0] {
      return false
    }
  }
  return true
}
```

#### 280 摆动排序

题目要求：给定数组，要求使其奇数位上的元素都大于其两边的元素值。

思路分析：遍历一遍，逐个纠正。

```go
// date 2020/02/28
func wiggleSort(nums []int) {
  var t int
  for i := 0; i < len(nums)-1; i++ {
    if (i & 0x1 == 0) == (nums[i] > nums[i+1]) {
      t = nums[i+1]
      nums[i+1] = nums[i]
      nums[i] = t
    }
  }
}
```



#### 56 合并区间

思路分析

算法：先对区间进行排序；维护start, end值，注意处理最后一个

```go
func merge(intervals [][]int) [][]int {
  if len(intervals) <= 1 {return intervals}
  sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
  })
  res, i, n := make([][]int, 0), 0, len(intervals) // i保存当前start索引
  // init start, end
  start, end := intervals[0][0], intervals[0][1]
  for i < n {
    i++
    // 查找以start开始的最大区间，并更新end值
    for i < n {
      if intervals[i][1] <= end || intervals[i][0] <= end {
        if intervals[i][1] > end { end = intervals[i][1] }
        i++
      } else { break }
    }
    t := make([]int, 2)
    t[0], t[1] = start, end
    res = append(res, t)
    // i >= n 说明处理完毕，直接退出
    if i >= n {break}
    // update start, end
    start, end = intervals[i][0], intervals[i][1]
  }
}
```

#### 912 排序数组

##### 快速排序

快速排序的思想是分治算法，partition将数组分成两个子数组，如果两个子数组有序，则这个数组有序。

```go
// date 2020/03/01
func sortArray(nums []int) []int {
    if len(nums) == 1 { return nums }
    QuickSort(nums, 0, len(nums)-1)
    return nums
}

// 快速排序
// 返回值为index
func partition(nums []int, l, r int) int {
  p := nums[l]
  index := l
  for i := l; i <= r; i++ {
    if nums[i] <= p {
      swap(nums, index, i)
      index++
    }
  }
  swap(nums, l, index-1)
  return index-1
}

func QuickSort(nums []int, l, r int) {
  if l >= r { return }
  p := partition(nums, l, r)
  QuickSort(nums, l, p-1)
  QuickSort(nums, p+1, r)
}
func swap(nums []int, i, j int) {
    t := nums[i]
    nums[i] = nums[j]
    nums[j] = t
}
```



##### 选择排序

思路分析：选择未排序数组中最小的一个，然后交换值应该所在的位置。时间复杂度O(n^2) 空间复杂度O(1)

```go
func selectSort(nums []int) {
  n := len(nums)
  for i := 0; i < n; i++ {
    for j := i+1; j < n; j++ {
      if nums[i] > nums[j] {
        swap(nums, i, j)
      }
    }
  }
}
func swap(nums []int, i, j int) {
  t := nums[i]
  nums[i] = nums[j]
  nums[j] = t
}
```

##### 插入排序

算法逻辑：一个元素肯定是有序的，然后依次遍历剩下的元素，将其插入到前面已排序的数组的中。

时间复杂度O(n^2) 空间复杂度O(1)

```go
func insertSort(nums []int) {
  temp, k := 0, 0
  for i := 1; i < len(nums); i++ {
    temp, k = nums[i], i-1
    for k > 0 && nums[k] > temp {
      nums[k+1] = nums[k]
      k--
    }
    nums[k+1] = temp
  }
}
```

##### 希尔排序

算法思想

插入排序的思想是使数组相邻的两个元素有序，而希尔排序是使数组中任意间隔为h的元素有序，然后h不断减小，又称为插入排序的快速排序。

```go
func sortArray(nums []int) []int {
    n, h := len(nums), 1
    for h < n/3 { h = 3*h + 1}
    for h >= 1 {
        for i := h; i < n; i++ {
            for j := i; j >= h && nums[j] < nums[j-h]; j -= h {
                nums[j], nums[j-h] = nums[j-h], nums[j]
            }
        }
        h /= 3
    }
}
```

##### 归并排序

算法思想

如果两个数组有序，可以通过归并的方式将其合并成一个有序的数组。

```go
// date 2020/01/05
// 自顶向下的归并排序，递归
// 时间复杂度O(nlgn)
func sortArray(nums []int) []int {
    if len(nums) < 2 { return nums }
    mid := len(nums) >> 1
    l := sortArray(nums[:mid])
    r := sortArray(nums[mid:])
    return merge(l, r)
}

// 归并排序
// 归并排序的merge算法
func merge(a1, a2 []int) []int {
    if len(a1) == 0 { return a2 }
    if len(a2) == 0 { return a1 }
    n1, n2 := len(a1), len(a2)
    n := n1 + n2
    res := make([]int, n)
    i, j := 0, 0
    for k := 0; k < n; k++ {
        if i >= n1 {
            res[k] = a2[j]
            j++
        } else if j >= n2 {
            res[k] = a1[i]
            i++
        } else if a1[i] < a2[j] {
            res[k] = a1[i]
            i++
        } else {
            res[k] = a2[j]
            j++
        }
    }
    return res
}
```

##### 堆排序

算法思想：使用大顶堆

```go
func sortArray(nums []int) []int {
    makeMaxHeap(nums)
    n := len(nums)
    for n > 0 {
        swap(nums, 0, n-1)
        n--
        maxHeapify(nums, 0, n)
    }
    return nums
}

// 堆排序
// 构建大顶堆
func makeMaxHeap(nums []int) {
    for i := len(nums) >> 1 - 1; i >= 0; i-- {
        maxHeapify(nums, i, len(nums))
    }
}
// 自上而下的堆有序化
func maxHeapify(nums []int, start, end int) {
    if start > end { return }
    temp, l, r := nums[start], start << 1 + 1, start << 1 + 2
    for l < end {
        r = l + 1
        if r < end && nums[r] > nums[l] { l++ }
        if nums[start] > nums[l] { break }
        nums[start] = nums[l]
        nums[l] = temp
        start = l
        l = start << 1 + 1
    }
} 
// swap
func swap(nums []int, i, j int) {
    t := nums[i]
    nums[i] = nums[j]
    nums[j] = t
}
```

#### 922 按奇偶排序数组II【E】

思路分析：记录每个不"在位"的索引，然后统一调整。注意不能两两依次调整，因为前后两个”不在位“的原因可以一样。

```go
// date 2020/02/28
func sortArrayByParityII(A []int) []int {
  p1, p2 := make([]int, 0), make([]int, 0)
  for i := 0; i < len(A); i++ {
    if i & 0x1 == 0 && A[i] & 0x1 != 0 { p1 = append(p1, i) }
    if i & 0x1 == 1 && A[i] & 0x1 != 1 { p2 = append(p2, i) }
  }
  i, n := 0, len(p1)
  var t int
  for i < n {
    t = A[p1[i]]
    A[p1[i]] = A[p2[i]]
    A[p2[i]] = t
    i++
  }
  return A
}
```



#### 1086 前五科的均分

思路分析：先对id排序，再对成绩排序。

```go
// date 2020/02/28
func highFive(items [][]int) [][]int {
  sort.Slice(items, func(i, j int) bool {
    if items[i][0] == items[j][0] {
      return items[i][1] > items[j][1] // 对成绩降序排序
    }
    return items[i][0] < items[j][0] // 对id升序排序
  })
  res, id, temp, c := make([][]int, 0), items[0][0], 0, 0
  for _, it := range items {
    if it[0] == id && c < 5 {
      temp += it[1]
      k++
      if k == 5 {
        r := make([]int, 2)
        r[0], r[1] = id, temp/5
        res = append(res, r)
      }
    } else if it[0] != id {
      id = it[0]
      temp = it[1]
      k = 1
    }
  }
  return res
}
```



#### 1356 根据数字二进制下1的数目排序

算法分析：求一个数字的二进制下2的个数。

```go
// date 2020/02/28
// 暴力解法
func sortByBits(arr []int) []int {
  sort.Slice(arr, func(i, j int) bool {
    if bits(arr[i]) == bits(arr[j]) {
      return arr[i] < arr[j]
    }
    return bits(arr[i]) < bits(arr[j])
  })
  return arr
}
func bits(x int) int {
  res := 0
  for x > 0 {
    res += x & 0x1
    x >>= 1
  }
  return res
}
```



