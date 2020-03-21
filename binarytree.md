### 二叉树



#### 相关题目

- 102 二叉树的层次遍历【M】



#### 102 二叉树的层次遍历

题目要求：https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

算法分析：bfs广度优先搜索和dfs深度优先搜索

```go
// date 2020/03/21
// bfs广度优先搜索
func levelOrder(root *TreeNode) [][]int {
  res := make([][]int, 0)
  if root == nil { return res }
  queue := make([]*TreeNode, 0)
  queue = append(queue, root)
  for len(queue) != 0 {
    n := len(queue)
    temp_res := make([]int, 0)
    for i := 0; i < n; i++ {
      temp_res = append(temp_res, queue[i].Val)
      if queue[i].Left != nil { queue = append(queue, queue[i].Left) }
      if queue[i].Right != nil { queue = append(queue, queue[i].Right) }
    }
    res = append(res, temp_res)
    queue = queue[n:]
  }
  return res
}
// dfs深度优先搜索
func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    var dfs_ func(root *TreeNode, level int)
    dfs_ = func(root *TreeNode, level int) {
        if root == nil { return }
        if level == len(res) {
            res = append(res, make([]int, 0))
        }
        res[level] = append(res[level], root.Val)
        dfs_(root.Left, level+1)
        dfs_(root.Right,level+1)
    }
    dfs_(root, 0)
    return res
}
```

