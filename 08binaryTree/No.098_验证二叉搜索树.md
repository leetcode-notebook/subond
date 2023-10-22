## 98 验证二叉搜索树-中等

题目：

给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
		节点的右子树只包含 大于 当前节点的数。
		所有左子树和右子树自身必须也是二叉搜索树。



分析：

迭代版本的中序遍历。【推荐该算法】

```go
// date 2023/10/22
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    stack := make([]*TreeNode, 0, 16)
    preVal := math.Inf(-1)
    for root != nil || len(stack) != 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        cur := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if float64(cur.Val) <= preVal {
            return false
        }
        preVal = float64(cur.Val)
        root = cur.Right
    }
    return true
}
```



算法1：遍历其中序遍历序列，并判断其序列是否为递增

```go
// date 2020/03/21
// 算法1：遍历其中序遍历序列，并判断其序列是否为递增
func isValidBST(root *TreeNode) bool {
    nums := inOrder(root)
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] >= nums[i+1] { return false }
    }
    return true
}

func inOrder(root *TreeNode) []int {
    if root == nil { return []int{} }
    res := make([]int, 0)
    if root.Left != nil { res = append(res, inOrder(root.Left)...) }
    res = append(res, root.Val)
    if root.Right != nil { res = append(res, inOrder(root.Right)...) }
    return res
}
```



算法2：

```go
// 算法2：递归
// 递归判断当前结点的值是否位于上下边界之中
// 时间复杂度O(N)
func isValidBST(root *TreeNode) bool {
  if root == nil { return true }
  var isValid func(root *TreeNode, min, max float64) bool
  isValid = func(root *TreeNode, min, max float64) bool {
    if root == nil { return true }
    if float64(root.Val) <= min || float64(root.Val) >= max { return false }
    return isValid(root.Left, min, float64(root.Val)) && isValid(root.Right, float64(root.Val), max)
  }
  return isValid(root, math.Inf(-1), math.Inf(0))
}
```



算法3：

```go
// 算法3
// 其思想是迭代版的中序遍历，因为栈中已经保留所有的结点，只需要出栈判断即可
// 时间复杂度O(N)，空间复杂度O(1),最优解
func isValidBST(root *TreeNode) bool {
    if root == nil { return true }
    stack := make([]*TreeNode, 0)
    preValue := math.Inf(-1)

    for len(stack) != 0 || root != nil {
        // 将左子树全部放入栈
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // 出栈，即取出左子树的最后一个节点
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        // 判断其是否大于前继结点
        if float64(root.Val) <= preValue { return false }
        // 更新前继节点和前继节点的值
        preValue = float64(root.Val)
        root = root.Right
    }
    return true
}
```

