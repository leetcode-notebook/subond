## 70 爬楼梯-简单

题目：

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？


分析：

直接递推。c[i] = c[i-1] + c[i-2]

```go
// date 2023/11/07
func climbStairs(n int) int {
    if n < 3 {
        return n
    }
    p1, p2 := 1, 2
    var res int
    for i := 3; i <= n; i++ {
        res = p1 + p2
        p1, p2 = p2, res
    }

    return res
}
```
