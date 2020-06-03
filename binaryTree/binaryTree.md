## 二叉树

[TOC]

### 什么是二叉树

树是一种数据结构，树中的每个节点都包含一个键值和所有子节点的列表，对于二叉树来说，每个节点最多有两个子树结构，分别称为左子树和右子树。

#### 二叉树的深度

二叉树的深度是指二叉树的根结点到叶子结点的距离，最大深度即根结点到叶子结点的最大距离；最小深度即根结点到叶子结点的最小距离；常见的问题包括求取二叉树的最大/最大深度，一般可以使用递归算法或者DFS。

### 二叉树的遍历

根据访问root节点的先后顺序，可分为前序遍历，中序遍历和后序遍历。递归版本的遍历，只要理解其思想就很好写；而对于非递归版本的遍历，需要深入理解其结点的遍历顺序，并记录下来之前经过的结点，所以一定会用到栈。

从广度搜索和深度搜索的角度来讲，层序遍历属于广度优先搜索，而前序，中序，后序遍历均为深度优先搜索。

#### 前序遍历

前序遍历是指先访问root节点，然后左子树，最后是右子树。

```go
// 递归版
func preOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root == nil { return res }
  // 先访问根结点
  res = append(res, root.Val)
  // 后访问左子树
  if root.Left != nil {
    res = append(res, preOrder(root.Left)...)
  }
  // 最后访问右子树
  if root.Right != nil {
    res = append(res, preOrder(root.Right)...)
  }
  return res
}
// 迭代版
// 前序遍历：root->left->right
func preorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    if root == nil { return res }
    stack := make([]*TreeNode, 0)
    for root != nil || len(stack) != 0 {
        // find root and left
        for root != nil {
            res = append(res, root.Val)  // 先将根节点加入结果集
            stack = append(stack, root)  // 将遍历过的结点加入栈，后序出栈依次访问结点的右子树
            root = root.Left             // 遍历当前结点的左子树
        }
        // find the right
        if len(queue) != 0 {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            root = root.Right
        }
    }
    return res
}
```

#### 中序遍历

前序遍历是指访问左子树，然后root节点，最后是右子树。

```go
// 递归版
func inOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root == nil { return res }
  // 先访问左子树
  if root.Left != nil {
    res = append(res, preOrder(root.Left)...)
  }
  // 后访问根结点
  res = append(res, root.Val)
  // 最后访问右子树
  if root.Right != nil {
    res = append(res, preOrder(root.Right)...)
  }
  return res
}
// 迭代版
func inorderTraversal(root *TreeNode) []int {
    if root == nil {return nil}
    stack := make([]*TreeNode, 0)
    res := make([]int, 0)
    for root != nil || len(stack) != 0 {
        // find the left
        for root != nil {
            stack = append(stack, root)  // 把当前结点添加进去
            root = root.Left
        }
        // 取出最后一个结点
        if len(stack) != 0 {
            root = queue[len(stack)-1]
            res = append(res, root.Val)
            stack = stack[:len(stack)-1]
            root = root.Right
        }
    }
    return res
}
```

#### 后序遍历

前序遍历是指先访问左子树，然后右子树，最后是root节点。

```go
// 前序遍历的伪代码
// 递归版
func postOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root == nil { return res }
  // 先访问左子树
  if root.Left != nil {
    res = append(res, preOrder(root.Left)...)
  }
  // 后访问右子树
  if root.Right != nil {
    res = append(res, preOrder(root.Right)...)
  }
  // 最后访问根结点
  res = append(res, root.Val)
  return res
}
// 迭代版
func postorderTraversal(root *TreeNode) []int {
    if root == nil {return nil}
    res := make([]int, 0)
    stack := make([]*TreeNode, 0)
    stack = append(stack, root)
    var pre, cur *TreeNode   // 记录前驱节点和当前节点
    for len(stack) != 0 {
        // 出栈 当前结点
        cur = stack[len(stack)-1]
        // 如果当前结点为叶子结点，则直接加入结果集
        // 如果当前结点不是叶子结点，但是前驱结点为当前结点的左右子树时(说明当前结点的左右子树已经处理完毕)，也加入结果集
        if cur.Left == nil && cur.Right ==  nil || pre != nil && (pre == cur.Left || pre == cur.Right) {
            res = append(res, cur.Val)
            // 出栈，继续检查
            stack = stack[:len(stack)-1]
            pre = cur
        } else {
            // 因为在出栈的时候检查结点，并追加到结果中
            // 所以，先入栈右子树，后入栈左子树
            if cur.Right != nil {
                stack = append(stack, cur.Right)
            }
            if cur.Left != nil {
                stack = append(stack, cur.Left)
            }
        }
    }
    
    return res
}
```

#### 层序遍历

层序遍历是指逐层遍历树的结构，也称为广度优先搜索，算法从一个根节点开始，先访问根节点，然后遍历其相邻的节点，其次遍历它的二级、三级节点。

借助队列数据结构，先入先出的顺序，实现层序遍历。

```go
// date 2020/03/21
// 层序遍历
// bfs广度优先搜索
// 算法一：使用队列，逐层遍历
func levelOrder(root *TreeNode) [][]int {
  res := make([][]int, 0)
  queue := make([]*TreeNode, 0)
  // 将根结点入队
  queue = append(queue, root)
  var n int
  for len(queue) != 0 {
    lres := make([]int, 0)
    n = len(queue)
    // 出队 计算每层结点，并检查其左右结点
    for i := 0; i < n; i++ {
      lres = append(lres, queue[i].Val)
      if queue[i].Left != nil {
        queue = append(queue, stack[i].Left)
      }
      if queue[i].Right != nil {
        queue = append(queue, stack[i].Right)
      }
    }
    queue = queue[n:]
    res = append(res, lres)
  }
  return res
}
// dfs深度优先搜索
func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    if root == nil {return nil}
    var dfs_ func(root *TreeNode, level int)
    dfs_ = func(root *TreeNode, level int) {
        if root == nil { return }
        // 如果没有当前层的结果集，则增加一个
        if level == len(res) {
            res = append(res, make([]int, 0))
        }
        // 向当前层的结果集合追加元素
        res[level] = append(res[level], root.Val)
        dfs_(root.Left, level+1)
        dfs_(root.Right, level+1)
    }
    dfs_(root, 0)
    return res
}
```

**注意**，从上面两种实现的方式来看，层序遍历既可以使用广度优先搜索，也可以使用深度优先搜索。

### 递归解决树的问题

递归通常是解决树的相关问题最有效和最常用的方法之一，分为**自顶向下**和**自底向上**两种。

**自顶向下**的解决方案

自顶向下意味着在每个递归层级，需要先计算一些值，然后递归调用并将这些值传递给子节点，视为**前序遍历**。参见题目【二叉树的最大深度】

```go
// 通用的框架，自顶向下
1. return specific value for null node
2. update the answer if needed                      // answer <-- params
3. left_ans = top_down(root.left, left_params)      // left_params <-- root.val, params
4. right_ans = top_down(root.right, right_params)   // right_params <-- root.val, params 
5. return the answer if needed                      // answer <-- left_ans, right_ans
```



**自底向上**的解决方案

自底向上意味着在每个递归层级，需要先对子节点递归调用函数，然后根据返回值和根节点本身得到答案，视为**后序遍历**。

```go
// 通用的框架，自底向上
1. return specific value for null node
2. left_ans = bottom_up(root.left)          // call function recursively for left child
3. right_ans = bottom_up(root.right)        // call function recursively for right child
4. return answers                           // answer <-- left_ans, right_ans, root.val
```



**总结**：

当遇到树的问题时，先思考一下两个问题：

> 1.你能确定一些参数，从该节点自身解决出发寻找答案吗？
>
> 2.你可以使用这些参数和节点本身的值来决定什么应该是传递给它子节点的参数吗？

如果答案是肯定的，那么可以尝试使用自顶向下的递归来解决问题。

当然也可以这样思考：

> 对于树中的任意一个节点，如果你知道它子节点的结果，你能计算出该节点的答案吗？

如果答案是肯定的，可以尝试使用自底向上的递归来解决问题。



### 平衡二叉树

一棵高度平衡二叉树的定义为:

> 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

### 对称二叉树

如果一棵二叉树关于根结点左右对称，则称为对称二叉树，也称为镜像二叉树。

### [相关题目](leetcode_bt.md)