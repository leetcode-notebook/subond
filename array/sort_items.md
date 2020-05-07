### 排序数组



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
