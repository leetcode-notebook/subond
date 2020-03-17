#### 动态规划





### 相关题目

- 64 最小路径和【M】
- 70 爬楼梯
- 72 编辑距离【H】
- 121 买卖股票的最佳时机
- 122 买卖股票的最佳时机II
- 123 买卖股票的最佳时机III
- 188 买卖股票的最佳时机IV
- 746 使用最少花费爬楼梯
- 264 丑数【M】
- 300 最长上升子序列【M】



#### 64 最小路径和

题目要求：https://leetcode-cn.com/problems/minimum-path-sum/

思路分析：自底向上递推

递推公式：`opt[i, j] += min(opt[i+1][j], opt[i][j+1])`，注意边界情况

```go
// date 2020/03/08
// 时间复杂度O(M*N),空间复杂度O(1)
func minPathSum(grid [][]int) int {
    m := len(grid)
    if m == 0 { return 0 }
    n := len(grid[0])
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            // 如果存在右节点和下节点，选择其中最小的那个
            if i + 1 < m && j + 1 < n {
                if grid[i+1][j] < grid[i][j+1] {
                    grid[i][j] += grid[i+1][j]
                } else {
                    grid[i][j] += grid[i][j+1]
                }
            // 如果只有右节点，则选择右节点
            } else if j + 1 < n {
                grid[i][j] += grid[i][j+1]
            // 如果只有下节点，则选择下节点
            } else if i + 1 < m {
                grid[i][j] += grid[i+1][j]
            }
        }
    }
    return grid[0][0]
}
```



#### 70 爬楼梯

题目要求：https://leetcode-cn.com/problems/climbing-stairs/

思路分析：这个问题和斐波那契数列很像，可以有三种解法，一个比一个更优。

递推公式等于c[n] = c[n-1] + c[n-2]，基本情况n=1, c[n]=1;n=2,c[n]=2

算法一：直接递推

```go
// date 2020/03/08
// 算法一：自顶向下直接递推，超时。
// 时间复杂度O(2^N),空间复杂度O(1)
func climbStairs(n int) int {
  if n < 3 { return n}
  return climbStairs(n-1) + climbStairs(n-2)
}
// 算法二：自底向上递推，并且记录中间结果。
// 时间复杂度O(N),空间复杂度O(N)
func climbStairs(n int) int {
  if n < 3 { return n }
  opt := make([]int, n+1)
  opt[1], opt[2] = 1, 2
  for i := 3; i <= n; i++ {
    opt[i] = opt[i-1] + opt[i-2]
  }
  return opt[n]
}
// 算法三：自底向上的递推，只保留有用的结果
// 时间复杂度O(N),空间复杂度O(1)
func climbStairs(n int) int {
  if n < 3 { return n }
  first, second := 1, 2
  res := 0
  for i := 3; i <= n; i++ {
    res = first + second
    first = second
    second = res
  }
  return res
}
```

#### 72 编辑距离

题目要求：https://leetcode-cn.com/problems/edit-distance/

思路分析：

```go
// date 2020/03/16
func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        // base case
        dp[i] = make([]int, n+1)
        dp[i][0] = i
    }
    // base case
    for j := 0; j <= n; j++ {
        dp[0][j] = j
    }
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
            }
        }
    }
    return dp[m][n]
}
func min(x, y int) int {
    if x < y { return x }
    return y
}

// 动态规划
// dp[i][j] 将word1[0...i]变成word2[0...j]所需要的最小步数
// if word1[i] == word2[j]; dp[i][j] = dp[i-1][j-1]
// else dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
// dp[i-1][j] insert word1
// dp[i][j-1] delete word1
// dp[i-1][j-1] replace
```



#### 121 买卖股票的最佳时机

题目要求：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

思路分析：只能交易一次。

遍历一遍，记录之前的最小值。如果当前值大于当前最小值，则计算利润；如果当前值小于当前最小值，则更新当前最小值，具体算法如下

```go
// date 2020/03/14
// 时间复杂度O(N),空间复杂度O(1)
func maxProfit(prices []int) int {
  if 0 == len(prices) { return 0 }
  res, cur_min := 0, prices[0]
  for _, v := range prices {
    if v > cur_min && res < v - cur_min { res = v - cur_min }
    if v < cur_min { cur_min = v }
  }
  return res
}
```

#### 122 买卖股票的最佳时机II

题目要求：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

思路分析：可以买卖无数次

遍历一遍，只要后面的值大于前面的值，即算作利润。

```go
// date 2020/03/14
// 时间复杂度O(N),空间复杂度O(1)
func maxProfit(prices []int) int {
  profit := 0
  for i := 0; i < len(prices)-1; i++ {
    if prices[i] < prices[i+1] {
      profit =+ prices[i+1] - prices[i]
    }
  }
  return profit
}
```

#### 123 买卖股票的最佳时机III

题目要求：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

思路分析：最多可以交易两笔

遍历一遍，分别从数组[0...i]和[i+1...N]中获取最大利润

```go
// date 2020/03/14
// 时间复杂度O(N^2),空间复杂度O(1)
func maxProfix(prices []int) int {
  res, v1, v2 := 0, 0, 0
  for i := 0; i < len(prices); i++ {
    v1 = maxProfitOnce(prices[:i])
    v2 = maxProfitOnce(prices[i+1:])
    if v1 + v2 > res { res = v1 + v2 }
  }
  return res
}

func maxProfitOnce(prices []int) int {
  if 0 == len(prices) { return 0 }
  res, cur_min := 0, prices[0]
  for _, v := range prices {
    if v > cur_min && res < v - cur_min { res = v - cur_min }
    if v < cur_min { cur_min = v }
  }
  return res
}
```

#### 188 买卖股票的最佳时机IV

题目要求：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/

思路分析：

#### 746 使用最小花费爬楼梯

题目要求：

考察知识：动态规划

算法分析

```go
// date 2020/03/08
// 算法一：自顶向下递推
func minCostClimbingStairs(cost []int) int {
  n := len(cost)
  if n == 0 { return 0 }
  if n == 1 { return cost[0] }
  if n == 2 { return min(cost[0], cost[1]) }
  return min(minCostClimbingStairs(cost[:n-1]) + cost[n-1], minCostClimbingStairs(cost[:n-2]) + cost[n-2])
}
func min(x, y int) int {
  if x < y { return x }
  return y
}
// 算法二：自底向上递推
func minCostClimbingStairs(cost []int) int {
  n := len(cost)
  if n == 0 { return 0 }
  if n == 1 { return cost[0] }
  if n == 2 { return min(cost[0], cost[1]) }
  opt := make([]int, n)
  opt[0], opt[1] = cost[0], cost[1]
  for i := 2; i < n; i++ {
    opt[i] = min(opt[i-2], opt[i-1]) + cost[i]
  }
  return min(opt[n-1], opt[n-2])
}
// 算法三：算法二的优化版
func minCostClimbingStairs(cost []int) int {
  n := len(cost)
  if n == 0 { return 0 }
  if n == 1 { return cost[0] }
  first, second := cost[0], cost[1]
  res := min(first, second)
  for i := 2; i < n; i++ {
    res = min(first, second) + cost[i]
    first = second
    second = res
  }
  return min(first, second)
}
```





#### 264 丑树



题目要求：https://leetcode-cn.com/problems/ugly-number-ii/

算法分析：

```go
// date 2020/03/01
// 动态规划
func nthUglyNumber(n int) int {
    if n < 1 { return 1 }
    dp := make([]int, n)
    dp[0] = 1
    p1, p2, p3 := 0, 0, 0
    for i := 1; i < n; i++ {
        dp[i] = min(2 * dp[p1], min(3 * dp[p2], 5 * dp[p3]))
        if dp[i] == 2 * dp[p1] { p1++ }
        if dp[i] == 3 * dp[p2] { p2++ }
        if dp[i] == 5 * dp[p3] { p3++ }
    }
    return dp[n-1]
}

func min(x, y int) int {
    if x < y { return x }
    return y
}
```

#### 300 最长上升子序列

题目要求：https://leetcode-cn.com/problems/longest-increasing-subsequence/

思路分析：

```
定义dp方程，dp[i]表示包含nums[i]元素的最长上升子序列
dp[i] = maxvalue + 1
maxvalue = 0 <= j < i && nums[i] > nums[j]; max(maxvalue, dp[j])
具体算法如下：
```

```go
// date 2020/03/14
// 时间复杂度O(N^2),空间复杂度O(N)
func lengthOfLIS(nums []int) int {
  if len(nums) == 0 { return 0 }
  n, res, maxvalue := len(nums), 1, 1
  dp := make([]int, n)
  dp[0] = 1
  for i := 1; i < n; i++ {
    // 找到nums[i]所满足的最长上升子序列
    maxvalue = 0 // 初值为0
    for j := 0; j < i; j++ {
      if nums[i] > nums[j] { maxvalue = max(maxvalue, dp[j]) }
    }
    // 更新dp[i]
    dp[i] = maxvalue + 1
    res = max(res, dp[i])
  }
  return res
}

func max(x, y int) int {
  if x > y { return x }
  return y
}
```

思路分析：维护一个最长上升子序列数组lis，然后每次更新它，最后数组lis的长度就是结果。

更新有序数组可以用二分查找，时间复杂度为O(logN)。

```go
// date 2020/03/16
func lengthOfLIS(nums []int) int {
  if len(nums) == 0 { return 0}
  lis := make([]int, 0)
  lis = append(lis, nums[0])
  for i := 1; i < len(nums); i++ {
    // 在lis数组找到第一个比nums[i]大的元素，并替换它，否则进行追加
    if nums[i] > lis[len(lis)-1] {
      lis = append(lis, nums[i])
      continue
    }
    for j := 0; j < len(lis); j++ {
      if nums[i] <= lis[j] {
        lis[j] = nums[i]
        break
      }
    }
  }
  return len(lis)
}
```

