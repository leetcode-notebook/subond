### 插入元素



[TOC]

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

