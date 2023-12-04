## 34 在排序数组中查找元素的第一个和最后一个位置-中等

题目：

给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。



分析：

复习二分查找，[详见](../binary_search.md)

```go
// date 2022/09/15
func searchRange(nums []int, target int) []int {
    res := make([]int, 2, 2)

    l := searchFirstEqual(nums, target)
    r := searchLastEqual(nums, target)

    res[0] = l
    res[1] = r

    return res
}

func searchFirstEqual(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] >= target {
            right = mid - 1
        } else if nums[mid] < target {
            left = mid + 1
        }
    }
    if left < len(nums) && nums[left] == target {
        return left
    }
    return -1
}

// 从 right 找边界
func searchLastEqual(nums []int, target int) int {
    n := len(nums)
    left, right := 0, n-1
    for left <= right {
        mid := (left + right) / 2
        if nums[mid] > target {
            right = mid - 1
        } else if nums[mid] <= target {  // 从右侧找边界，left 不断地向 right 靠近
            left = mid + 1
        }
    }
    if right >= 0 && nums[right] == target {
        return right
    }
    return -1
}
```

