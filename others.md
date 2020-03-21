#### 位1的个数

```go
// date 2020/03/18
// 算法：移位比较
func hammingWeight(num uint32) int {
    var c int
    for num > 0 {
        c += int(num & 0x1)
        num >>= 1
    }
    return c
}
```

#### 汉明距离

```go
// date 2020/03/18
func hammingDistance(x int, y int) int {
    return hammingWeight(x^y)
}

func hammingWeight(num int) int {
    var c int
    for num > 0 {
        c += num & 0x1
        num >>= 1
    }
    return c
}
```

#### 292 Nim游戏

算法思路：只要是4的倍数，你一定输。

```go
// date 2020/01/09
func canWinNim(n int) bool {
  if n % 4 == 0 {
    return false
  }
  return true
}
```

#### 89 格雷编码

题目要求：给定一个格雷编码的总位数n，输出器格雷编码序列（格雷编码序列：相邻的两个数字只有一位不同）。

算法：镜像反射法

```go
// date 2020/01/11
/* 算法：G(n) = G(n-1) + R(n-1)  // R(n-1)为G(n-1)反序+head
0 0 00
  1 01
    11
    10
*/
func grayCode(n int) []int {
  res, head := make([]int, 0), 1
  res = append(res, 0)
  for i := 0; i < n; i++ {
    n := len(res) - 1
    for j = n; j >= 0; j-- {
      res = append(res, head + res[j])
    }
    head <<= 1
  }
  return res
}
```

