## 530 二叉搜索树的最小绝对差-中等

题目：

给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。



分析：

递归中序遍历，求相邻元素的最小值。

```go
// date 2023/10/30
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    res := 1 << 32 - 1
    preVal := -1

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }
        dfs(root.Left)

        if preVal != -1 && root.Val - preVal < res {
            res = root.Val - preVal
        }
        preVal = root.Val
        dfs(root.Right)
    }

    dfs(root)

    return res
}
```
