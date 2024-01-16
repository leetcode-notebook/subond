## 238 除自身以外数组的乘积-中等

题目：

给你一个整数数组 `nums`，返回 *数组 `answer` ，其中 `answer[i]` 等于 `nums` 中除 `nums[i]` 之外其余各元素的乘积* 。

题目数据 **保证** 数组 `nums`之中任意元素的全部前缀元素和后缀的乘积都在 **32 位** 整数范围内。

请 **不要使用除法，**且在 `O(*n*)` 时间复杂度内完成此题。



> **示例 1:**
>
> ```
> 输入: nums = [1,2,3,4]
> 输出: [24,12,8,6]
> ```
>
> **示例 2:**
>
> ```
> 输入: nums = [-1,1,0,-3,3]
> 输出: [0,0,9,0,0]
> ```



**解题思路**

这道题可用前缀乘积和后缀乘积解决，详见解法1。具体为，初始化两个数组 left 和 right，分别表示前缀乘积和后缀乘积。

`left[i]` 表示前`i-1`个元素的乘积，`right[i]`表示后`[i+1, n-1]`个元素的乘积，那么对应每个元素除自身的乘积就是`ans[i] = left[i] * right[i]`。

时间复杂度`O(N)`，空间复杂度`O(N)`。



仔细观察，你会发现，上面的解法1可以优化，即计算前缀乘积的时候，直接保存在结果数组中；然后使用变量后序遍历数组计算后缀乘积 right，直接更新结果数组也可以。详见解法2。



```go
// date 2024/01/16
// 解法1
// 前缀乘积 * 后缀乘积
func productExceptSelf(nums []int) []int {
    n := len(nums)
    left := make([]int, n)
    right := make([]int, n)
    for i := 0; i < n; i++ {
        if i == 0 {
            left[i] = 1
        } else {
            left[i] = left[i-1] * nums[i-1]
        }
    }
    for i := n-1; i >= 0; i-- {
        if i == n-1 {
            right[i] = 1
        } else {
            right[i] = nums[i+1] * right[i+1]
        }
    }
    ans := make([]int, n)
    for i := 0; i < n; i++ {
        ans[i] = left[i] * right[i]
    }
    return ans
}

// 解法2
// 对解法1的优化
func productExceptSelf(nums []int) []int {
    n := len(nums)
    if n < 2 {
        return nums
    }

    ans := make([]int, n)
    ans[0] = 1
    // fill left
    for i := 1; i < n; i++ {
        ans[i] = ans[i-1] * nums[i-1]
    }

    // fill right
    right := 1
    for i := n-1; i >= 0; i-- {
        ans[i] = ans[i] * right
        right *= nums[i]
    }
    return ans
}
```

