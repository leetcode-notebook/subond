## 33 搜索旋转排序数组-中等

题目：

整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。

你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。



分析：

通过第一个元素来判断从前半段查找，还是从后半段查找；并且在查找过程中注意翻转的点，一旦到达直接终止查找。

```go
// date 2023/11/02
func search(nums []int, target int) int {
    res := -1
    s := len(nums)
    if s == 0 {
        return res
    }
    if nums[0] <= target {
        i := 0
        for i < s {
            if nums[i] == target {
                res = i
                break
            }
            if i + 1 < s && nums[i+1] < nums[i] {
                break
            }
            i++
        }
    } else {
        i := s-1
        for i >= 0 {
            if nums[i] == target {
                res = i
                break
            }
            if i > 0 && nums[i-1] > nums[i] {
                break
            }
            i--
        }
    }

    return res
}
```
