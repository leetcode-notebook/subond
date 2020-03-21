### 贪心算法



### 相关题目

- 122 买卖股票的最佳时机II



#### 122 买卖股票的最佳时机

题目要求：https://leetcode-cn.com/problems/strobogrammatic-number/

算法分析：

因为题目要求只能持有一种股票，同时可以操作无数次，所以可以采用贪心算法，计算每天的利润。

```go
// date 2020/03/21
func maxProfit(prices []int) int {
    profit := 0
    for i := 0; i < len(prices) - 1; i++ {
        if prices[i] < prices[i+1] {
            profit += prices[i+1] - prices[i]
        }
    }
    return profit
}
```

