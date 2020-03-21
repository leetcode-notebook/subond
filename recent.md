## Recent Doing

日期 2020年1月6日 每日三题

#### 20 有效的字符串

思路分析：stack数据结构，压栈及出栈

算法如下：时间复杂度O(n) 空间复杂度O(1)

```go
// date 2020/01/06
/* 算法：
 0. 如果字符串长度为单数，直接返回false.
 1. 遇到左括号，入栈
 2. 遇到右括号，判断栈是否为空，如果为空，返回false
 3. 如果不为空，出栈
 4. 遍历结束后，判断栈是否为空。
*/
func isValid(s string) bool {
  if len(s) & 0x1 == 1 {return false}
  stack := make([]rune, 0)
  var c rune
  for _, v := range s {
    if v == '(' || v == '{' || v == '[' {
      stack = append(stack, v)
    } else {
      if len(stack) == 0 {return false}
      c = stack[len(stack)-1]
      if v == ')' && c == '(' || v == '}' && c == '{' || v == ']' && c == '[' {
        stack = stack[:len(stack)-1]
      }
    }
  }
  return len(stack) == 0
}
```

### 23 合并K个排序链表

题目要求：给定K个有序链表，返回合并后的排序链表。

算法1：递归算法，自底向上的两两合并。每个链表遍历一遍，所示时间复杂度O(n*k)

```go
// date 2020/01/06
// 算法：哑结点+递归，两两合并
func mergeKLists(lists []*ListNode) *ListNode {
  // basic case
  if len(lists) == 0 {return nil}
  if len(lists) == 1 {return lists[0]}
  p1 := lists[0]
  p2 := mergeKLists(lists[1:])
  dummy := &ListNode{Val: -1}
  pre := dummy
  for p1 != nil && p2 != nil {
    if p1.Val < p2.Val {
      pre.Next = p1
      p1 = p1.Next
    } else {
      pre.Next = p2
      p2 = p2.Next
    }
    pre = pre.Next
  }
  if p1 != nil {
    pre.Next = p1
  }
  if p2 != nil {
    pre.Next = p2
  }
  return dummy.Next
}
```

#### 155 最小栈

题目要求：实现栈的push, pop, top以及常数时间内的getMin()操作。

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

日期 2020年1月11日

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

#### 235 二叉搜索树的最近公共祖先

题目要求：给定一个二叉搜索树，及其两个节点，返回这两个节点的最近公共祖先。

算法：递归，充分利用搜索树的条件。

```go
// date 2020/01/11
/* 算法
1. 如果两个节点中的其中一个节点等于根节点，那么根节点就是结果
2. 如果两个节点的值，一个大于根节点，一个小于根节点，那么根节点即是最近的公共祖先
3. 如果两个节点的值均小于根节点的值，则从左子树寻找结果，否则从右子树寻找结果。
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  if root == q || root == p {return root}
  if p.Val < root.Val && root.Val < q.Val || p.Val > root.Val && root.Val > q.Val {return root}
  if p.Val < root.Val && q.Val < root.Val {
    return lowestCommonAncestor(root.Left, p, q)
  }
  return lowestCommonAncestor(root.Right, p, q)
}
```

#### 230 二叉搜索树中第K小的元素

题目要求：给定一个二叉搜索树，返回其第K小的元素，你可以假设K总是有效的，即1<=k<=n。

算法：因为是二叉搜索树，先得到其中序遍历结果，然后直接返回k-1所在的元素。

```go
// date 2020/01/11
func kthSmallest(root *TreeNode) int {
  data := inOrder(root)
  if len(data) < k {return 0}
  return data[k-1]
}

func inOrder(root *TreeNode) []int {
  if root == nil {return nil}
  res := make([]int, 0)
  if root.Left != nil { res = append(res, inOrder(root.Left)...) }
  res = append(res, root.Val)
  if root.Right != nil { res = append(res, inOrder(root.Right)...) }
}
```

#### 146 LRU缓冲机制

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

#### 寻找数组的中心索引

题目要求：给定一个数组，返回其中心索引。

备注：中心索引是指其左边个元素之和等于其右边个元素之和。

算法：先求和，再依次遍历。时间复杂度O(2n)

```go
// date 2020/02/01
func pivotIndex(nums []int) int {
  if len(nums) < 1 { return -1 }
  total, sum := 0, 0
  for _, v := range nums { total+= v}
  for i := -1; i < len(nums) - 1; i++ {
    if sum == total - nums[i+1] {return i+1}
    sum += nums[i+1]
    total -= nums[i+1]
  }
  return -1
}
```

#### 至少是其他数字两倍的最大数

题目要求：给定一个数组，如果数组最大的元素至少是其他元素的两倍，则返回其索引，否则返回-1。

算法：最大元素只要大于次大元素的两倍即可。

```go
func dominantIndex(nums []int) int {
  if len(nums) <= 1 { return len(nums)-1 }
  res := -1
  max1, max2 := 0, 0
  for i, v := range nums {
    if v > max1 {
      max2 = max1
      max1 = v
      res = i
    } else if v > max2 {
      max2 = v
    }
  }
  if max1 >= 2 * max2 { return res }
  return -1
}
```

#### 加1

题目要求：给定一个非负整数的数组表现形式，返回其+1的值。

细节题：注意进位。

```go 
func plusOne(digits []int) []int {
  carry, temp := 1, 0
  for i := 0; i < len(digits); i++ {
    temp = digits[i]+carry
    digits[i] = temp % 10
    carry = temp / 10
  }
  if carry == 1 {
    return append([]int{1}, digits...)
  }
  return digits...
}
```



#### 46 全排列

题目要求：给定一个**没有重复**数字的序列，返回其所有可能的全排列。

算法：回溯法

```go
func permute(nums []int) [][]int {
  res := make([][]int, 0, total(len(nums)))
  func bt(nums, temp []int) {
    if len(nums) == 0 {
      res = append(res, temp)
      return
    }
    for i := range nums {
      temp = append(temp, nums[i])
      n := make([]int, 0, len(nums)-1)
      n = append(n, nums[:i]...)
      n = append(n, nums[i+1:]...)
      bt(n, temp)
    }
  }(nums, make([]int, 0))
  return res
}

func total(n int) int {
  if n <= 0 {return 0}
  res := 1
  for n >= 1 {
    res *= n
    n--
  }
  return res
}
```



#### 78 子集

题目要求：给定一个不含重复元素的数组，返回其所有可能的子集。

算法：组合+非递归实现

```go
// date 2020/01/11
/* 算法
1. 外层遍历nums，内层遍历结果集，并取出中间结果集
2. 用中间结果集+num组成新的结果集，并追加到结果中。
*/
func subsets(nums int) [][]int {
  res := make([][]int, 0)
  res = append(res, make([]int, 0))
  for _, v := range nums {
    n := len(res)
    for i := 0; i < n; i++ {
      t := make([]int, 0)
      t = append(t, res[i]...)
      t = append(t, v)
      res = append(res, t)
    }
  }
  return res
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

#### 62 不同路径

题目要求：给定一个二维数组m x n，返回机器人从二维数组的左上角->右下角的所有不同路径。

算法1：排列组合 res = f(m+n-2)/f(m-1)/f(n-1)  f为阶乘

算法2：递归

```go
// date 2020/01/11
/* 算法2：递归方程为dp[i][j] = dp[i-1][j] + dp[i][j-1]
   当i == 0 || j == 0; dp[i][j] = 1
*/
func uniquePaths(m, n int) int {
  dp := make([][]int, m)
  for i := 0; i < m; i++ {
    dp[i] = make([]int, n)
  }
  for i := 0; i < n; i++ {dp[0][i] = 1}
  for i := 0; i < m; i++ {dp[i][0] = 1}
  for i := 1; i < m; i++ {
    for j := 1; j < n; j++ {
      dp[i][j] = dp[i-1][j] + dp[i][j-1]
    }
  }
  return dp[m-1][n-1]
}
```





---

#### 字符串中第一个唯一的字符

```go
func firstUniqChar(s string) int {
  m := make(map[uint8]int, len(s))
  for i := 0; i < len(s); i++ {
    if _, ok := m[s[i]]; ok {
      m[s[i]] = -1
    } else {
      m[s[i]] = i
    }
  }
  for i := 0; i < len(s); i++ {
    if v, ok := m[s[i]]; ok && v != -1 {
      return v
    }
  }
  return -1
}
```



移除链表元素

```go
// Date 11.25
func removeElements(head *ListNode, val int) *ListNode {
  var newHead *ListNode
  for head != nil {
    if head.Val == val {
      head = head.Next
    } else {
      newHead = head
      break
    }
  }
  // if the newHead == nil,return
  if newHead == nil {return newHead}
  pre := newHead
  head = head.Next
  for head != nil {
    if head.Val != val {
      pre.Next = head
      pre = pre.Next
    }
    head = head.Next
  }
  pre.Next = nil
  return newHead
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

### 日期2020/02/02

#### 翻转字符串里的单词

题目要求：给定一个字符串，返回其按单词翻转，并去除多余空格的结果。

算法：从字符串的尾部遍历，找到单个单词，并翻转放入其结果，注意空格。

```go
// date 2020/02/02
func reverseWords(s string) string {
  res, temp := make([]byte, 0), make([]byte, 0) // 存放结果和单个单词
  end := len(s)-1
  for end >= 0 && s[end] == ' ' { end-- }
  for i := end; i >= 0; i-- {
    if s[i] != ' ' {
      temp = append(temp, s[i])
    } else {
      // find the word
      res = append(res, reverseWord(temp)...)
      res = append(res, ' ')
      for i >= 0 && s[i] == ' ' { i-- }
      i++
      temp = make([]byte, 0)
    }
  }
  if 0 != len(temp) {
    res = append(res, reverseWord(temp)...)
  }
  if len(res) > 0 && res[len(res)-1] == ' ' {
    res = res[:len(res)-1]
  }
  return string(res[:len(res)])
}
func reverseWord(word []byte) []byte {
  s, e := 0, len(word)-1
  var t byte
  for s < e {
    t = word[s]
    word[s] = word[e]
    word[e] = t
    s++
    e--
  }
  return word
}
```

