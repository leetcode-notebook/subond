### 二叉树

#### 什么是二叉树

树是一种数据结构，树中的每个节点都包含一个键值和所有子节点的列表，对于二叉树来说，每个节点最多有两个子树结构，分别称为左子树和右子树。



#### 二叉树的遍历

根据访问root节点的先后顺序，可分为前序遍历，中序遍历和后序遍历。

**前序遍历**

前序遍历是指先访问root节点，然后左子树，最后是右子树。

```go
// 前序遍历的伪代码
// 递归版
func preOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root != nil {
    res = append(res, root.Val)
    res = append(res, preOrder(root.Left)...)
    res = append(res, preOrder(root.Right)...)
  }
  return res
}
```

**中序遍历**

前序遍历是指访问左子树，然后root节点，最后是右子树。

```go
// 前序遍历的伪代码
// 递归版
func inOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root != nil {
    res = append(res, inOrder(root.Left)...)
    res = append(res, root.Val)
    res = append(res, inOrder(root.Right)...)
  }
  return res
}
```

**后序遍历**

前序遍历是指先访问左子树，然后右子树，最后是root节点。

```go
// 前序遍历的伪代码
// 递归版
func postOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root != nil {
    res = append(res, postOrder(root.Left)...)
    res = append(res, postOrder(root.Right)...)
    res = append(res, root.Val)
  }
  return res
}
```

**层序遍历**

层序遍历是指逐层遍历树的结构，也称为广度优先搜索，算法从一个根节点开始，先访问根节点，然后遍历其相邻的节点，其次遍历它的二级、三级节点。

借助栈数据结构，可以帮助我们实现层序遍历。

```go
// 层序遍历
func levelOrder(root *TreeNode) [][]int {
  res := make([][]int, 0)
  stack := make([]*TreeNode, 0)
  stack = append(stack, root)
  for len(stack) != 0 {
    lres, n := make([]int, 0), len(stack)
    for i := 0; i < n; i++ {
      if stack[i] != nil {
        stack = append(stack, stack[i].Left)
        stack = append(stack, stack[i].Right)
        lres = append(lres, stack[i].Val)
      }
    }
    stack = stack[n:]
    if len(lres) > 0 {
      res = append(res, lres)
    }
  }
  return res
}
```



#### 递归解决树的问题

递归通常是解决树的相关问题最有效和最常用的方法之一，分为**自顶向下**和**自底向上**两种。

**自顶向下**的解决方案

自顶向下意味着在每个递归层级，需要先计算一些值，然后递归调用并将这些值传递给子节点，视为前序遍历。参见题目【二叉树的最大深度】

**自底向上**的解决方案

自底向上意味着在每个递归层级，需要先对子节点递归调用函数，然后根据返回值和根节点本身得到答案。

**总结**：

当遇到树的问题时，先思考一下两个问题：

> 1.你能确定一些参数，从该节点自身解决出发寻找答案吗？
>
> 2.你可以使用这些参数和节点本身的值来决定什么应该是传递给它子节点的参数吗？

如果答案是肯定的，那么可以尝试使用自顶向下的递归来解决问题。

当然也可以这样思考：对于树中的任意一个节点，如果你知道它子节点的结果，你能计算出该节点的答案吗？如果答案是肯定的，可以尝试使用自底向上的递归来解决问题。



#### 相关题目

- 从中序和后序遍历序列构造二叉树[前序和中序]
- 104 二叉树的最大深度【E】
- 101 对称二叉树【E】
- 270 最接近的二叉搜素树值【E】
- 543 二叉树的直径【E】
- 545 二叉树的边界【M】
- 563 二叉树的坡度【E】
- 617 合并二叉树【E】
- 面试题27 二叉树的镜像
- 路径总和
- 不同的二叉搜索树

####  从中序和后序遍历序列构造二叉树[前序和中序]

思路分析

利用根节点在中序遍历的位置

```go
// 递归, 中序和后序
func buildTree(inorder, postorder []int) *TreeNode {
  if len(postorder) == 0 {return nil}
  n := len(postorder)
  root = &TreeNode{
    Val: postorder[n-1],
  }
  index := -1
  for i := 0; i < len(inorder); i++ {
    if inorder[i] == root.Val {
      index = i
      break
    }
  }
  if index >= 0 && index < n {
    root.Left = buildTree(inorder[:index], postorder[:index])
    root.Right = buildTree(inorder[index+1:], postorder[index:n-1]))
  }
  return root
}
// 递归，前序和中序
func buildTree(preorder, inorder []int) *TreeNode {
  if 0 == len(preorder) {return nil}
  n := len(preorder)
  root := &TreeNode{
    Val: preorder[0],
  }
  index := -1
  for i := 0; i < len(inorder); i++ {
    if inorder[i] == root.Val{
      index = i
      break
    }
  }
  if index >= 0 && index < n {
    root.Left = buildTree(preorder[1:index+1], inorder[:index])
    root.Right = buildTree(preorder[index+1:], inorder[index+1:])
  }
  return root
}
```



#### 二叉树的最大深度

题目要求：给定一棵二叉树，找出其最大深度。

思路分析

算法：递归，递归调用分别得到左子树和右子树的深度，然后比较取最大值，并+1返回结果

```go
// 自底向上的思想
func maxDepth(root *TreeNode) int {
  if root == nil {return 0}
  res, left, right := 1, 0, 0
  if root.Left != nil { left = maxDepth(root.Left) }
  if root.Right != nil { right = maxDepth(root.Right) }
  if left > right {
    res = left + 1
  } else {
    res = right + 1
  }
  return res
}
```

算法：迭代

```go
func maxDepth(root *TreeNode) int {
  if root == nil {return 0}
  queue := make([]*TreeNode, 0)
  queue = append(queue, root)
  depth, n := 0, 0
  for len(queue) != 0 {
    n = len(queue)
    for i := 0; i < n; i++ {
      if queue[i].Left != nil {
        queue = append(queue, queue[i].Left)
      }
      if queue[i].Right != nil {
        queue = append(queue, queue[i].Right)
      }
    }
    queue = queue[n:]
    depth++
  }
  return depth
}
```

#### 对称二叉树

思路分析

算法1：

如果一个树的左子树和右子树对称，那么这个树就是对称的。那么两个树在什么情况下互为镜像呢？

> 如果两个树同时满足两个条件，则互为镜像。
>
> 1.它们的根节点具有相同的值；
>
> 2.每个树的右子树和另一个树的左子树对称。

```go
// 算法一
func isSymmetrics(root *TreeNode) bool {
  return isMirror(root, root)
}
func isMirror(t1, t2 *TreeNode) bool {
  if t1 == nil && t2 == nil {return true}
  if t1 == nil || t2 == nil {return false}
  return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}
```

算法2：类似层序遍历，将子树的左右结点依次放入队列，通过判断队列中的连续的两个值是否相同来确认是否对称。

```go
// 算法二
func isSymmetric(root *TreeNode) bool {
    if root == nil {return true}
    queue := make([]*TreeNode, 0)
    queue = append(queue, root, root)
    for 0 != len(queue) {
        n := len(queue)
        for i := 0; i < n; i += 2 {
            t1, t2 := queue[i], queue[i+1]
            if t1 == nil && t2 == nil {continue}
            if t1 == nil || t2 == nil {return false}
            if t1.Val != t2.Val {return false}
            queue = append(queue, t1.Left, t2.Right, t1.Right, t2.Left)
        }
        queue = queue[n:]
    }
    return true
}
```

#### 257 二叉树的所有路径

题目要求：给定一个二叉树，返回所有从根节点到叶子节点的路径。

算法分析：

```go
// date 2020/02/23
func binaryTreePaths(root *TreeNode) []string {
  res := make([]string, 0)
  if root == nil { return res }
  if root.Left == nil && root.Right == nil {
    res = append(res, fmt.Sprintf("%d", root.Val))
    return res
  } else if root.Left != nil {
    for _, v := range binaryTreePaths(root.Left) {
      res = append(res, fmt.Sprintf("%d->%s", root.Val, v))
    }
  } else if root.Right != nil {
    for _, v := range binaryTreePaths(root.Right) {
      res = append(res, fmt.Sprintf("%d->%s", root.Val, v))
    }
  }
  return res
}
```



#### 270 最接近的二叉搜索树值

题目要求：给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的数值。

算法分析：使用层序遍历，广度优先搜索。

```go
// date 2020/02/23
func closestValue(root *TreeNode, target float64) int {
    if root == nil { return 0 }
    var res int
    cur := math.Inf(1)
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    for len(stack) != 0 {
        n := len(stack)
        for i := 0; i < n; i++ {
            if math.Abs(float64(stack[i].Val) - target) < cur {
                cur = math.Abs(float64(stack[i].Val) - target)
                res = stack[i].Val
            }
            if stack[i].Left != nil { stack = append(stack, stack[i].Left) }
            if stack[i].Right != nil { stack = append(stack, stack[i].Right) }
        }
        stack = stack[n:]
    }
    return res
}
```



#### 543 二叉树的直径

题目要求：给定一棵二叉树，返回其直径。

算法分析：左右子树深度和的最大值。

```go
// date 2020/02/23
func diameterOfBinaryTree(root *TreeNode) int {
    v1, _ := findDepth(root)
    return v1
}

func findDepth(root *TreeNode) (int, int) {
  if root == nil { return 0, 0 }
  var v int
  v1, l := findDepth(root.Left)
  v2, r := findDepth(root.Right)
  if v1 > v2 {
    v = v1
  } else {
    v = v2
  }
  if l+r > v { v = l+r }
  if l > r { return v, l+1}
  return v, r+1
}
```



#### 545 二叉树的边界

题目要求：给定一个二叉树，按逆时针返回其边界。

思路分析：分别求其左边界，右边界，和所有的叶子节点，用visited辅助去重，算法如下：

```go
// date 2020/02/23
func boundaryOfBinaryTree(root *TreeNode) []int {
  if root == nil { return []int{} }
  res := make([]int, 0)
  left, right := make([]*TreeNode, 0), make([]*TreeNode, 0)
  visited := make(map[*TreeNode]int, 0)
  res = append(res, root.Val)
  visited[root] = 1
  // find left node
  pre := root.Left
  for pre != nil {
    left = append(left, pre)
    visited[pre] = 1
    if pre.Left != nil {
      pre = pre.Left
    } else {
      pre = pre.Right
    }
  }
  // find right node
  pre = root.Right
  for pre != nil {
    right = append(right, pre)
    visited[pre] = 1
    if pre.Right != nil {
      pre = pre.Right
    } else {
      pre = pre.Left
    }
  }
  leafs := findLeafs(root)
  // make res
  for i := 0; i < len(left); i++ {
    res = append(res, left[i].Val)
  }
  for i := 0; i < len(leafs); i++ {
    if _, ok := visited[leafs[i]]; ok { continue }
    res = append(res, ,leafs[i].Val)
  }
  for i := len(right) - 1; i >= 0; i-- {
    res = append(res, right[i].Val)
  }
  return res
}

func findLeafs(root *TreeNode) []*TreeNode {
  res := make([]*TreeNode, 0)
  if root == nil { return res }
  if root.Left == nil && root.Right == nil {
    res = append(res, root)
    return res
  }
  res = append(res, findLeafs(root.Left)...)
  res = append(res, findLeafs(root.Right)...)
  return res
}
```



#### 563 二叉树的坡度

题目要求：给定一棵二叉树，返回其坡度。

算法一：根据定义，递归计算。

```go
func findTilt(root *TreeNode) int {
    if root == nil || root.Left == nil && root.Right == nil { return 0 }
    left := sum(root.Left)
    right := sum(root.Right)
    var res int
    if left > right {
        res = left - right
    } else {
        res = right - left
    }
    return res + findTilt(root.Left) + findTilt(root.Right)
}

func sum(root *TreeNode) int {
    if root == nil { return 0 }
    return sum(root.Left) + sum(root.Right) + root.Val
}
```

#### 617 合并二叉树

给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。

你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。

```go
// date 2020/02/23
// 递归
func mergeTrees(t1, t2 *TreeNode) *TreeNode {
  if t1 == nil { return t2 }
  if t2 == nil { return t1 }
  t1.Val += t2.Val
  t1.Left = mergeTrees(t1.Left, t2.Left)
  t1.Right = mergeTrees(t1.Right, t2.Right)
  return t1
}
```

#### 面试题27 二叉树的镜像

题目要求：给定一棵二叉树，返回其镜像二叉树。

```go
// date 2020/02/23
func mirrorTree(root *TreeNode) *TreeNode {
  if root == nil { return root }
  left := root.Left
  root.Left = root.Right
  root.Right = left
  root.Left = mirrorTree(root.Left)
  root.Right = mirrorTree(root.Right)
  return root
}
```



#### 路径总和

思路分析

算法1：自顶向下的递归思想

```go
// 自顶向下递归
func hasPathSum(root *TreeNode, sum int) bool {
  if root == nil {return false}
  sum -= root.Val
  if root.Left == nil && root.Right == nil {return sum == 0}
  return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
```

算法2：迭代，利用queue保存当前结果，类型层序遍历。

```go
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {return false}
    queue, csum := make([]*TreeNode, 0), make([]int, 0)
    queue = append(queue, root)
    csum = append(csum, sum - root.Val)
    for 0 != len(queue) {
        n := len(queue)
        for i := 0; i < n; i++ {
            if queue[i].Left == nil && queue[i].Right == nil && csum[i] == 0 {return true}
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
                csum = append(csum, csum[i] - queue[i].Left.Val)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
                csum = append(csum, csum[i] - queue[i].Right.Val)
            }
        }
        queue = queue[n:]
        csum = csum[n:]
    }
    return false
}
```

#### 不同的二叉搜索树 II

题目要求：给定一个整数n，生成所有由1...n为节点所组成的二叉搜索树。

思路分析

题目看着很难，实际利用递归思想没有那么难。

从序列`1..n`取出数字`1` 作为当前树的根节点，那么就有i-1个元素用来构造左子树，二另外的i+1..n个元素用来构造右子树。最后我们将得到G(i-1)棵不同的左子树，G(n-i)棵不同的右子树，其中G为卡特兰数。

这样就可以以i为根节点，和两个可能的左右子树序列，遍历一遍左右子树序列，将左右子树和根节点链接起来。

```go
func generateTrees(n int) []*TreeNode {
  if n == 0 { return nil }
  return generateTreesWithNode(1, n)
}

func generateTreesWithNode(start end int) []*TreeNode {
    res := make([]*TreeNode, 0)
    if start > end {
      // 表示空树
      res = append(res, nil)
      return res
    }
    for i := start; i <= end; i++ {
      left := generateTreesWithNode(start, i-1)
      right := generateTreesWithNode(i+1, end)
      for _, l := range left {
        for _, r := range right {
          res = append(res, &TreeNode{
            Val: i,
            Left: l,
            Right: r,
          })
        }
      }
    }
    return res
  }
```



#### 101 平衡二叉树

题目要求：给定一个二叉树，判断其是否为高度平衡二叉树。【一个二叉树的每个节点的左右两个子树的高度差的绝对值不超过1，视为高度平衡二叉树】

算法一：递归

```go
// date 2020/02/18
func isBalanced(root *TreeNode) bool {
  if root == nil { return true }
  if !isBalanced(root.Left) || !isBalanced(root.Right) { return false }
  if Math.Abs(float64(getHeight(root.Left)-getHeight(root.Right))) > float64(1) { return false }
  return true
}
func getHeight(root *TreeNode) int {
  if root == nil { return 0 }
  if root.Left == nil && root.Right == nil { return 1 }
  l, r := getHeight(root.Left), getHeight(root.Right)
  if l > r { return l+1 }
  return r+1
}
```

#### 703 数据流中第K大元素

解题思路：

0。使用数据构建一个容量为k的小顶堆。

1。如果小顶堆堆满后，向其增加元素，若大于堆顶元素，则重新建堆，否则掉弃。

2。堆顶元素即为第k大元素。

```go
// date 2020/02/19
type KthLargest struct {
  minHeap []int
  size
}

func Constructor(k int, nums []int) KthLargest {
  res := KthLargest{
    minHeap: make([]int, k),
  }
  for _, v := range nums {
    res.Add(v)
  }
}

func (this *KthLargest) Add(val int) int {
  if this.size < len(this.minHeap) {
    this.size++
    this.minHeap[this.size-1] = val
    if this.size == len(this.minHeap) {
      this.makeMinHeap()
    }
  } else if this.minHeap[0] < val {
    this.minHeap[0] = val
    this.minHeapify(0)
  }
  return this.minHeap[0]
}

func (this *KthLargest) makeMinHeap() {
  // 为什么最后一个根节点是n >> 1 - 1
  // 长度为n，最后一个叶子节点的索引是n-1; left = i >> 1 + 1; right = i >> 1 + 2
  // 父节点是 n >> 1 - 1
  for i := len(this.minHeap) >> 1 - 1; i >= 0; i-- {
    this.minHeapify(i)
  }
}

func (this *KthLargest) minHeapify(i int) {
  if i > len(this.minHeap) {return}
  temp, n := this.minHeap[i], len(this.minHeap)
  left := i << 1 + 1
  right := i << 1 + 2
  for left < n {
    right = left + 1
    // 找到left, right中较小的那个
    if right < n && this.minHeap[left] > this.minHeap[right] { left++ }
    // left, right 均小于root否和小顶堆，跳出
    if this.minHeap[left] >= this.minHeap[i] { break }
    this.minHeap[i] = this.minHeap[left]
    this.minHeap[left] = temp
		// left发生变化，继续调整
    i = left
    left = i << 1 + 1
  }
}
```

