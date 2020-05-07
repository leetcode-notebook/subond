## 删除元素



[TOC]

### 相关题目

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

#### 27 移除元素

题目要求：给定一个数组和一个值val，移除数组中所有值为val的元素，并返回数组的新长度。

算法1：维护index表示每个非val的下标。时间复杂度O(n) 空间复杂度O(n)

```go
// date 2019/12/28
// date 2020/05/02
func removeElement(nums []int, val int) int {
    index, n := 0, len(nums) - 1
    for i := 0; i <= n; i++ {
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

