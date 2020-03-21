### Math

#### 相关题目

    - 8 String to Integer
    - 9 Palindrome Number
    - 15 三数之和
    - 231 2的幂
    - 149 直线上最多的点数[困难]
    - 斐波那契数列

#### 8 String to Integer

https://leetcode.com/problems/string-to-integer-atoi/

```go
/* 细节题: 需要各种异常场景
1. 忽略开头的空白字符
2. 数字的符号
3. 溢出
4. 无效输入
*/
func myAtoi(str string) int {
    myMAX := 1 << 31 - 1
    myMIN := -1 << 31
    sign := 1
    var res int
    i := 0
    n := len(str)
    for i < n && str[i] == ' ' {i++}
    if i == n {return res}
    if str[i] == '+' {
        i++
    } else if str[i] == '-' {
        i++
        sign = -1
    }
    for ; i < n; i++ {
        if str[i] - '0' < 0 || str[i] - '0' > 9 {break}
        t := int(str[i] - '0')
        if res > myMAX / 10 || res == myMAX / 10 && t > myMAX % 10 {
            if sign == 1 {
                return myMAX
            } else if sign == -1 {
                return myMIN
            }
        }
        res = res * 10 + t
    }
    return res * sign
}
```



#### 9 Palindrome Number

```go
// 算法1：直接比较数字的每一位，更通用的算法
func isPalindrome(x int) bool {
    if x < 0 || (x != 0 && x % 10 == 0) {return false}
    d := 1
    for x / d >= 10 {
        d *= 10
    }
    for x > 0 {
        // 最高位
        q := x / d
        // 最低位
        p := x % 10
        if q != p { return false }
        // 移动至次高位，同时去掉最后一位
        x = x % d / 10
        d /= 100
    }
    return true
}
// 算法2：将数字从中间分开，判断左边和右边的数字是否相等
func isPalindrome(x int) bool {
    if x < 0 || (x != 0 && x % 10 == 0) {return false}
    left := 0
    for x > left {
        left = left * 10  + x % 10
        x /= 10
    }
    return x == left || x == left/10
}
```

思路学习：算法2对于判断链表同样适用。参考题目234https://leetcode.com/problems/palindrome-linked-list/

#### 15 三数之和

思路分析

算法：时间复杂度O(n2)

1-对数组进行升序排序，然后逐一遍历，固定当前元素nums[i]，使用左右指针遍历当前元素后面的所有元素，nums[l]和nums[r]，如果三数之和=0，则添加进结果集

2-当nums[i] > 0 结束循环 // 升序排序，三数之和必然大于零

3-当nums[i] == nums[i-1] 表明用重复元素，跳过

4-sum = 0零时，如果nums[l] == nums[l+1] 重复元素，跳过

5-sum = 0, nums[r] == nums[r-1] 重复元素，跳过

```go
func threeSum(nums []int) [][]int {
  sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
  })
  res, n, sum := make([][]int, 0), len(nums), -1
  l, r := 0, n-1
  for i := 0; i < n; i++ {
    if nums[i] > 0 {break}
    if i > 0 && nums[i] == nums[i-1] {continue}
    l, r = i+1, n-1
    for l < r {
      sum = nums[i] + nums[l] + nums[r]
      if sum == 0 {
        t := make([]int, 0)
        t = append(t, nums[i], nums[l], nums[r])
        res = append(res, t)
        for l < r && nums[l] == nums[l+1] {l++}
        for l < r && nums[r] == nums[r-1] {r--}
        l++
        r--
      } else if sum < 0 {
        l++
      } else {
        r--
      }
    }
  }
  return res
}
```

#### 四数求和

思路分析

算法：原理同三数求和

```go
func fourSum(nums []int, target int) [][]int {
  sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
  })
  res, n, sum := make([][]int, 0), len(nums), 0
  l, r := 0, n-1
  for i := 0; i < n; i++ {
    // nums[i]可以是负数 if nums[i] > target {break}
    if i > 0 && nums[i] == nums[i-1] {continue}
    for j := i+1; j < n; j++ {
      if j > i+1 && nums[j] == nums[j-1] {continue}
      l, r = j+1, n-1
      for l < r {
        sum = nums[i] + nums[j] + nums[l] + nums[r]
        if sum == target {
          t := make([]int, 0)
          t = append(t, nums[i], nums[j], nums[l], nums[r])
          res = append(res, t)
          for l < r && nums[l] == nums[l+1] {l++}
          for l < r && nums[r] == nums[r-1] {r--}
          l++
          r--
        } else if sum < target {
          l++
        } else {
          r--
        }
      }
    }
  }
  return res
}
```

#### 231 2的幂

题目要求：给定一个整数，判断其是否是2的幂次方。

思路分析：判断这个数是否不断的被2整数，直接结果等于1。如果某一次结果为奇数，直接返回false。

```go
// date 2020/01/11
func isPowerOfTwo(n int) bool {
  for n > 1 {
    if n &0x1 == 1 {return false}
    n >>= 1
  }
  return n == 1
}
```

#### 149 直线上最多的点数[困难]

题目要求：给定二维平面上的n个点，求最多有多少个点在同一条直线上。

算法：哈希表，欧几里得最大公约数算法

```go
// data 2020/02/01
func maxPoints(points [][]int) int {
  if len(points) < 3 { return len(points) }
  n, count := len(points), make([]int, len(points))
  size, same := 1, 0
  for i := 0; i < n; i++ {
    m := make(map[[2]int]int, 0)
    count[i] = 1
    size, same = 1, 0
    for j := 0; j < n; j++ {
      if i == j { continue }
      // x坐标一样
      if points[i][0] == points[j][0] {
        if points[i][1] == points[j][1] {
          same++
        }
        size++
      } else {
     // x坐标不一样 求斜率
        dy := points[j][1] - points[i][1]
        dx := points[j][0] - points[i][0]
        g := gcd(dy, dx)
        if g != 0 {
          dy, dx = dy/g, dx/g
        }
        k := [2]int
        k[0], k[1] = dy, dx
        find := false
        for mk, mv := range m {
          if mk[0] == k[0] && mk[1] == k[1] {
            find = true
            m[mk]++
          }
        }
        if !find {
          m[k] = 1
        }
      }
    }
    for _, v := range m {
      if v + 1 > count[i] {
        count[i] = v + 1
      }
    }
    count[i] += same
    if count[i] < size {
      count[i] = size
    }
  }
  maxIndex := 0
  for i := 0; i < n; i++ {
    if count[i] > count[maxIndex] {
      maxIndex = i
    }
  }
  return count[maxIndex]
}

func gcd(x, y int) int {
  if y == 0 { return x }
  return gcd(y, x % y)
}
```

#### 斐波那契数列

思路分析

记忆化递归

```go
func fib(N int) int {
  if N < 2 {return N}
  c, p1, p2 := 1, 1, 0
  for i := 2; i <= N; i++ {
  	c = p1+p2
    p1, p2 = c, p1
  }
  return c
}
```