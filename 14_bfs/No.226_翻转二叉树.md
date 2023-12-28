## 226 翻转二叉树-中等

题目：

给你一棵二叉树的根节点 `root` ，翻转这棵二叉树，并返回其根节点。



**分析：**

算法1：直接递归

```go
// date 2023/12/28
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        or := root.Right
        root.Right = root.Left
        root.Left = or

        dfs(root.Right)
        dfs(root.Left)
    }

    dfs(root)

    return root
}
```



算法2：

广度优先搜索，依次交换队列中的每个节点。

```go
// date 2023/12/28
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    if root.Left == nil && root.Right == nil {
        return root
    }

    queue := make([]*TreeNode, 0, 16)
    queue = append(queue, root)

    for len(queue) != 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            cur := queue[i]
            or := cur.Right
            cur.Right = cur.Left
            cur.Left = or

            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }

        queue = queue[n:]
    }

    return root
}
```

