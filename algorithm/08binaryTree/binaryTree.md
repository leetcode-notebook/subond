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
    res = append(res, inOrder(root.Left)...)
  }
  // 后访问根结点
  res = append(res, root.Val)
  // 最后访问右子树
  if root.Right != nil {
    res = append(res, inOrder(root.Right)...)
  }
  return res
}
// 迭代版
// 需要stack结构
// left->root->right
func inorderTraversal(root *TreeNode) []int {
    if root == nil { return nil }
    // 初始化队列
    stack := make([]*TreeNode, 0, 16)
    res := make([]int, 0, 16)
    for root != nil || 0 != len(stack) {
        // 先把左子树全部放入队列
        if root != nil {
            stack = append(stack, root)
            root = root.Left
            continue
        }
        // 去除当前结点，查看其右子树
        if 0 != len(stack) {
            root = stack[len(stack)-1]
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
// left->right-root
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



---

### Binary Indexed Tree \[二叉索引树\]

  * 1.[结构介绍](#结构介绍)
  * 2.[问题定义](#问题定义)
  * 3.[表面形式](#表现形式)
    * 3.1[更新操作](#更新操作)
    * 3.2[求和操作](#求和操作)
  * 4.[工作原理](#BIT工作原理)
  * 5.[代码实现](#代码实现)

#### 结构介绍

二叉索引树，Binary Indexed Tree，又名树状数组，或者Fenwick Tree，因为本算法由Fenwick创造。

#### 问题定义

给定一个数组array[0....n-1]，实现两个函数：1）求取前i项和；2）更新第i项值。

分析：

其中一个比较简单的解法就是通过循环累计(0-i-1)求和(需要O(n)时间)，根据下标更新值(需要O(1)时间)。另一种解法是创建一个前n项和数组(即，第i项存储前i项的和)，这样求和需要O(1)时间，更新需要O(n)时间。

而使用BIT数据结构，可以在O(Logn)时间内进行求和操作和更新操作。下面具体介绍一下。

#### 表现形式

BIT这种数据结构使用一般数组进行存储，每个节点存储 **部分** 输入数组的元素之和。BIT数组的元素个数与输入数组的元素个数一致，均为n。BIT数组的初始值均为0，通过更新操作获取输入数组部分元素之和，并存储在相应的位置上。

#### 更新操作

更新操作是指根据输入数组的元素来更新BIT数组的元素，进而根据BIT数组求前i项和。

```
update(index, val): Updates BIT for operation arr[index] += val
// Note that arr[] is not changed here.  It changes
// only BI Tree for the already made change in arr[].
1) Initialize index as index+1.
2) Do following while index is smaller than or equal to n.
...a) Add value to BITree[index]
...b) Go to parent of BITree[index].  Parent can be obtained by removing
     the last set bit from index, i.e., index = index + (index & (-index))
```

令BIT数组为：BIT[i] = Input[i-lowbit(i)+1] + Input[i-lowbit(i)+2] + ... + Input[i];

即从最左边的孩子，到自身的和，如BIT[12]= A9+A10+A11+A12

#### 求和操作

求和操作是指利用BIT数组求原来输入数组的前i项和。

```
getSum(index): Returns sum of arr[0..index]
// Returns sum of arr[0..index] using BITree[0..n].  It assumes that
// BITree[] is constructed for given array arr[0..n-1]
1) Initialize sum as 0 and index as index+1.
2) Do following while index is greater than 0.
...a) Add BITree[index] to sum
...b) Go to parent of BITree[index].  Parent can be obtained by removing
     the last set bit from index, i.e., index = index - (index & (-index))
3) Return sum.
```


说明：

1. 如果节点y是节点x的父节点，那么节点y可以由节点x通过移除Lowbit(x)获得。**Lowbit(natural)为自然数(即1,2,3…n)的二进制形式中最右边出现1的值**。即通过以下公式获得 $parent(i) = i - i & (-i)$。

2. 节点y的子节点x(对BIT数组而言)存储从y(不包括y)到x(包括x)(对输入数组而言)的的元素之和。即BIT数组的值由输入数组两者之间下标的节点对应关系获得。


计算前缀和Sum(i)的计算：顺着节点i往左走，边走边“往上爬”，把经过的BIT[i]累加起来即可。

![bit演示图](http://on64c9tla.bkt.clouddn.com/Algorithm/bit.jpg)

#### BIT工作原理

BIT工作原理得益于所有的正整数，均可以表示为多个2的幂次方之和。例如，19 = 16 + 2 + 1。BIT的每个节点都存储n个元素的总和，其中n是2的幂。例如，在上面的getSum（）的第一个图中，前12个元素的和可以通过最后4个元素（从9到12） ）加上8个元素的总和（从1到8）。 数字n的二进制表示中的设置位数为O（Logn）。 因此，我们遍历getSum（）和update（）操作中最多的O（Logn）节点。 构造的时间复杂度为O（nLogn），因为它为所有n个元素调用update（）。

#### 代码实现

```
void updateBIT(int BITree[], int n, int index, int val) {
  index = index + 1;
  while(index <= n) {
    BITree[index] += val;
    index += index & (-index);
  }
}

int getSum(int BITree[], int index) {
  int sum = 0;
  index = index + 1;
  while(index > 0) {
    sum += BITree[index];
    index -= index & (-index);
  }
  return sum;
}

int *conBIT(int arr[], int n) {
  int *BITree = new int[n + 1];
  for(int i = 0; i <= n; i++)
    BITree[i] = 0;
  for(int i = 0; i < n; i++)
    updateBIT(BITree, n, i, arr[i]);
  return BITree;
}
```





---

### 二叉搜索树

顾名思义，二叉搜索树是以一棵二叉树来组织的，其满足一下条件：x为二叉树的一个节点，如果y是x左子树的一个节点，那么`y.Val <= x.Val`；如果y是x右子树的一个节点，则`y.Val >= x.Val`

根据二叉树的性质可得知，二叉搜索树的中序遍历为升序序列，相反如果交换左右子树的遍历顺序亦可得到降序序列。

```go
// inorder伪代码
func inOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root != nil {
    res = append(res, inOrder(root.Left)...)
    res = append(res, root.Val)
    res = append(res, inOrder(root.Right)...)
  }
  return res
}

func decOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root != nil {
    res = append(res, decOrder(root.Right)...)
    res = append(res, root.Val)
    res = append(res, decOrder(root.Left)...)
  }
  return res
}
```

#### 二叉搜索树的基本操作

二叉搜索树的基本操作包括查询，最小关键字元素，最大关键字元素，前继节点，后继节点，插入和删除，这些基本操作所花费的时间与二叉搜索树的高度成正比。

**查询**

如果二叉搜索树的高度为h，则查询的时间复杂度为O(h)。

```go
// 查询伪代码
// 递归版
func searchBST(root *TreeNode, key int) *TreeNode {
  if root == nil || root.Val == key { return root }
  if root.Val < key {
    return searchBST(root.Right, key)
  }
  return searchBST(root.Left, key)
}
// 迭代版
func searchBST(root *TreeNode, key int) *TreeNode {
  for root != nil && root.Val != key {
    if root.Val < key {
      root = root.Right
    } else {
      root = root.Left
    }
  }
  return root
}
```

**最大关键字元素**

如果二叉搜素树的高度为h，则最大关键字元素的操作时间复杂度为O(h)。

```go
// 最大关键字元素的伪代码
func maximumBST(root *TreeNode) *TreeNode {
  for root.Right != nil { root = root.Right }
  return root
}
// 递归版
func maximumBST(root *TeeNode) *TreeNode {
  if root.Right != nil {
    return maximumBST(root.Right)
  }
  return root
}
```

**最小关键字元素**

如果二叉搜素树的高度为h，则最大关键字元素的操作时间复杂度为O(h)。

```go
// 最小关键字元素的伪代码
func minimumBST(root *TreeNode) *TreeNode {
  for root.Left != nil { root = root.Left }
  return root
}
// 递归版
func minimumBST(root *TreeNode) *TreeNode {
  if root.Left != nil {
    return minimum(root.Left)
  }
  return root
}
```

**前继节点**

前继节点是指给定一个节点x，如果存在x的前继节点则返回前继节点，否则返回nil。

如果二叉搜素树的高度为h，则查找前继节点的操作时间复杂度为O(h)。

```go
func predecessorBST(x *TreeNode) *TreeNode {
  // x节点存在左子树，则返回左子树中最大的节点
  if x.Left != nil {
    return maximum(root.Left)
  }
  // 如果x节点没有左子树，则找到其父节点，且父节点的右子树包含x
  y := x.Present
  for y != nil && x == y.Left {
    x = y
    y = y.Present
  }
  return y
}
```

**后继节点**

后继节点是指给定一个节点x，如果存在x的后继节点则返回后继节点，否则返回nil。

如果二叉搜素树的高度为h，则查找后继节点的操作时间复杂度为O(h)。

```go
// 后继节点的伪代码
func successorBST(x *TreeNode) *TreeNode {
  // x节点存在右子树，则返回右子树中最小的节点
  if x.Right != nil {
    return minimum(root.Right)
  }
  // 如果x节点没有右子树，则找到其父节点，且父节点的左子树包含x
  y := x.Present
  for y != nil && x == y.Right {
    x = y
    y = y.Present
  }
  return y
}
```

**插入**

插入和删除会引起二叉搜索树所表示的动态集合的变化。一定要修改数据结构来反映这种变化，单修改要保持二叉搜索树的性质不变。

```go
// x.Val = val, x.Right = x.Left = x.Present = nil
// 插入的伪代码
func insertBST(root, x *TreeNode) {
  var y *TreeNode
  for root != nil {
    y = root
    if x.Val < root.Val {
      root = root.Left
    } else {
      root = root.Right
    }
  }
  x.Present = y
  if y == nil { // empty tree
    root = x
  } else if x.Val < y.Val {
    y.Left = x
  } else {
    y.Right = x
  }
}
// 没有父指针
// 递归版
func insertBST(root, x *TreeNode) *TreeNode {
  if root == nil {
    root = x
    return root
  }
  if root.Val < x.Val {
    return insertBST(root.Right)
  }
  return insertBST(root.Left)
}
// 迭代版
func insertBST(root, x *TreeNode) *TreeNode {
  if root == nil {
    root == x
    return root
  }
  pre, head := root, root
  for head != nil {
    pre = head
    if root.Val < x.Val {
      head = head.Right
    } else {
      head = head.Left
    }
  }
  if y.Val < x.Val {
    y.Right = x
  } else {
    y.Left = x
  }
  return root
}
```

**删除**

从一棵二叉搜索树中删除一个节点x的整个策略分为三个情况，分别是：

1）如果x没有孩子节点，那么直接删除即可；

2）如果x只有一个孩子节点，那么将该孩子节点提升到x即可；

3）如果x有两个孩子，那么需要找到x的后继节点y（一定存在x的右子树中），让y占据x的位置，那么x原来的右子树部分称为y的右子树，x的左子树称为y的新的左子树。

```go
func deleteBST(root *TreeNode, key int) *TreeNode{
  if root == nil { return root }
  if root.Val < key {
    return deleteBST(root.Right, key)
  } else if root.Val > key {
    return deleteBST(root.Left, key)
  } else {
    // no child node
    if root.Left == nil && root.Right == nil {
      root = nil
    } else if root.Right != nil {
      root.Val = postNodeVal(root.Right)
      root.Right = deleteBST(root.Right, root.Val)
    } else {
      root.Val = preNodeVal(root.Left)
      root.Left = deleteBST(root.Left, root.Val)
    }
  }
  return root
}

// find post node val in right
func postNodeVal(root *TreeNode) int {
  for root.Left != nil { root = root.Left }
  return root.Val
}
// find pre node val in left
func preNodeVal(root *TreeNode) int {
  for root.Right != nil { root = root.Right }
  return root.Val
}
```

### 二叉搜索树相关题目

- 108 将有序数组转换为二叉搜索树【M】
- 235 二叉搜素树的最近公共祖先
- 230 二叉搜索树中第K小的元素
- 450 删除二叉搜索树中的节点
- 1038 从二叉搜索树到更大和树【M】
- 1214 查找两棵二叉搜索树之和【M】
- 面试题 17.12 BiNode【E】
- 面试题54 二叉搜索树的第K大节点

#### 108 将有序数组转换为二叉搜索树

题目要求：给定一个升序数据，返回一棵高度平衡二叉树。

思路分析：

二叉搜索树的定义：

1）若任意节点的左子树不为空，则左子树上所有的节点值均小于它的根节点值；

2）若任意节点的右子树不为空，则右子树上所有的节点值均大于它的根节点值；

3）任意节点的左，右子树均为二叉搜索树；

4）没有键值相等的点。

如何构造一棵树，可拆分成无数个这样的子问题：构造树的每个节点，以及节点之间的关系。对于每个节点来说，都需要：

1）选取节点；

2）构造该节点的左子树；

3）构造该节点的右子树。

算法一：对于升序序列，可以选择中间的节点作为根节点，然后递归调用。因为每个节点只遍历一遍，所以时间复杂度为O(n)；因为递归调用，空间复杂度为O(log(n))。

```go
//date 2020/02/20
// 算法一：递归版
func sortedArrayToBST(nums []int) *TreeNode {
  if len(nums) == 0 { return nil }
  // 选择根节点
  mid := len(nums) >> 1
  // 构造左子树和右子树
  left := nums[:mid]
  right := nums[mid+1:]
  
  return &TreeNode{
    Val: nums[mid],
    Left: sortedArrayToBST(left),
    Right: sortedArrayToBST(right),
  }
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


#### 235 二叉搜索树的最近公共祖先

题目要求：给定一个二叉搜素树和两个节点，返回其最近公共祖先。

算法：递归，并充分利用二叉搜索树的性质`left.Val < root.Val < right.Val`。

```go
// date 2020/02/18
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  if root == nil || root == p || root == q {return root }
  if p == nil { return q }
  if q == nil { return p }
  if root.Val > q.Val && root.Val > p.Val { return lowestCommonAncestor(root.Left) }
  if root.Val < q.Val && root.Val < p.Val { return lowestCommonAncestor(root.Right) }
  return root
}
```



#### 450 删除二叉搜索树中的节点

题目要求：给定一个二叉搜索树和key，删除key节点，并返回树的根节点

算法一：二叉搜索树的中序遍历序列是个升序序列，则欲要删除的节点需要用其前驱节点或后驱节点来代替。

1）如果为叶子节点，则直接删除；

2）如果不是叶子节点且有右子树，则用后驱节点代替。后驱节点为右子树中的较低位置；

3）如果不是叶子节点且有左子树，则用前驱节点代替。前驱节点为左子树中的较高位置；  

```go
func deleteBST(root *TreeNode, key int) *TreeNode{
  if root == nil { return root }
  if root.Val < key {
    return deleteBST(root.Right, key)
  } else if root.Val > key {
    return deleteBST(root.Left, key)
  } else {
    // no child node
    if root.Left == nil && root.Right == nil {
      root = nil
    } else if root.Right != nil {
      root.Val = postNodeVal(root.Right)
      root.Right = deleteBST(root.Right, root.Val)
    } else {
      root.Val = preNodeVal(root.Left)
      root.Left = deleteBST(root.Left, root.Val)
    }
  }
  return root
}

// find post node val in right
func postNodeVal(root *TreeNode) int {
  for root.Left != nil { root = root.Left }
  return root.Val
}
// find pre node val in left
func preNodeVal(root *TreeNode) int {
  for root.Right != nil { root = root.Right }
  return root.Val
}
```

#### 1038 从二叉搜索树到更大和树

题目要求：给出二叉**搜索**树的根节点，该二叉树的节点值各不相同，修改二叉树，使每个节点 `node` 的新值等于原树中大于或等于 `node.val` 的值之和

思路分析：

`root.Val = root.Val + root.Right.Val`

`root.Left.Val = root.Left.Val + root.Val`

首先得到二叉搜索树的中序遍历序列的逆序列，然后从序列的第一个开始，将累加值记录到当前树的节点上。

```go
// date 2020/02/23
func bsttoGST(root *TreeNode) *TreeNode{
  array := decOrder(root)
  val := 0
  for i := 0; i < len(array); i++ {
    val += array[i]
    setNewVal(root, array[i], val)
  }
}
func decOrder(root *TreeNode) []int {
  res := make([]int, 0)
  if root != nil {
    res = append(res, decOrder(root.Right)...)
    res = append(res, root.Val)
    res = append(res, decOrder(root.Left)...)
  }
  return res
}

func setNewVal(root *TreeNode, key, val int) {
  if root != nil {
    if root.Val == key {
      root.Val = val
    } else if root.Val > key {
      setNewVal(root.Left, key, val)
    } else {
      setNewVal(root.Right, key, val)
    }
  }
}
```

#### 面试题 BiNode

题目要求：将一颗二叉搜索树转换成单链表，Right指针指向下一个节点。

思路分析：

要求返回头节点，所以先转换left。

```go
// date 2020/02/19
func convertBiNode(root *TreeNode) *TreeNode {
  if root == nil { return nil }
  // 左子树不为空，需要转换
  if root.Left != nil {
    // 得到左子树序列，并找到最后一个节点
    left :=  convertBiNode(root.Left)
    pre := left
    for pre.Right != nil { pre = pre.Right }
    // 将最后一个节点right指向root
    pre.Right = root
    root.Left = nil
    // 得到右子树序列
    root.Right = convertBiNode(root.Right)
    return left
  }
  root.Right = convertBiNode(root.Right)
  return root
}
```

#### 面试题54 二叉搜索树的第K大节点

题目要求：给定一个二叉搜索树，找出其中第k大节点。

算法一：二叉搜索树的中序遍历为升序序列，按照Right->root->Left的顺序遍历可得到降序序列，然后返回第k节点。

```go
// date 2020/02/18
func kthLargest(root *TreeNode, k int) int {
  if root == nil || k < 0 { return -1 }
  nums := decOrder(root)
  if len(nums) < k {return -1}
  return nums[k-1]
}

// right->root->left
func decOrder(root *TreeNode) []int {
  if root == nil { return []int{} }
  res := make([]int, 0)
  if root.Right != nil {
    res = append(res, decOrder(root.Right)...)
  }
  res = append(res, root.Val)
  if root.Left != nil {
    res = append(res, decOrder(root.Left)...)
  }
  return res
}
```



### [相关题目](leetcode_bt.md)

#### 95 不同的二叉搜索树 II【中等】

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

func generateTreesWithNode(start, end int) []*TreeNode {
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

#### 98 验证二叉搜索树【中等】

题目要求：https://leetcode-cn.com/problems/validate-binary-search-tree/

算法分析：

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

#### 101 对称二叉树【简单】

题目链接：https://leetcode.com/problems/symmetric-tree/

思路分析

算法1：

如果一个树的左子树和右子树对称，那么这个树就是对称的。那么两个树在什么情况下互为镜像呢？

> 如果两个树同时满足两个条件，则互为镜像。
>
> 1.它们的根节点具有相同的值；
>
> 2.每个树的右子树和另一个树的左子树对称。

```go
// 算法一:递归
func isSymmetrics(root *TreeNode) bool {
  var isMirror func(r1, r2 *TreeNode) bool
  isMirror = func(r1, r2 *TreeNode) bool {
    if r1 == nil && r2 == nil { return true }
    if r1 == nil || r2 == nil { return false }
    return r1.Val == r2.Val && isMirror(r1.Left, r2.Right) && isMirror(r1.Right, r2.Left)
  }
  return isMirror(root, root)
}
```

算法2：类似层序遍历，将子树的左右结点依次放入队列，通过判断队列中的连续的两个值是否相同来确认是否对称。

```go
// 算法二:迭代
func isSymmetric(root *TreeNode) bool {
    if root == nil { return true }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root, root)
    var n int
    var t1, t2 *TreeNode
    for len(queue) != 0 {
        n = len(queue)
        for i := 0; i < n; i += 2 {
            t1, t2 = queue[i], queue[i+1]
            if t1 == nil && t2 == nil { continue }
            if t1 == nil || t2 == nil { return false }
            if t1.Val != t2.Val { return false }
            queue = append(queue, t1.Left, t2.Right, t1.Right, t2.Left)
        }
        queue = queue[n:]
    }
    return true
}
```

#### 104 二叉树的最大深度【简单】

题目链接：https://leetcode.com/problems/maximum-depth-of-binary-tree/

题目要求：给定一棵二叉树，找出其最大深度。

思路分析

算法：递归，递归调用分别得到左子树和右子树的深度，然后比较取最大值，并+1返回结果

```go
// 算法一: 递归，采用自底向上的递归思想
// 时间复杂度O(N)，空间复杂度O(NlogN)
// 自底向上的递归
func maxDepth(root *TreeNode) int {
    if root == nil { return 0 }
    l, r := maxDepth(root.Left), maxDepth(root.Right)
    if l > r { return l+1 }
    return r + 1
}

// 算法二
// bfs广度优先搜索
// 时间复杂度O(N)，空间复杂度O(N)
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

// 算法三 dfs深度优先搜索
// 时间复杂度O(N)，空间复杂度O(NlogN)
// 自顶向下的递归
func maxDepth(root *TreeNode) int {
    // dfs
    var dfs func(root *TreeNode, level int) int
    dfs = func(root *TreeNode, level int) int {
        if root == nil { return level }
        if root.Left == nil && root.Right == nil { return level+1 }
        if root.Left == nil {
            return dfs(root.Right, level+1)
        }
        if root.Right == nil {
            return dfs(root.Left, level+1)
        }
        l, r := dfs(root.Left, level+1), dfs(root.Right, level+1)
        if l > r { return l }
        return r
    }
    return dfs(root, 0)
}
```

#### 105 从前序与中序遍历序列构造二叉树【中等】

算法分析：1）从中序序列中找到根节点，递归左右子树。

```go
// date 2020/02/26
func buildTree(preorder []int, inorder []int) *TreeNode {
  if len(preorder) == 0 { return nil }
  root := &TreeNode{Val: preorder[0]}
  index := 0
  for i, v := range inorder {
    if v == preorder[0] {
      index = i
      break
    }
  }
  root.Left = buildTree(preorder[1:index+1], inorder[:index])
  root.Right = buildTree(preorder[index+1:], inorder[index+1:])
  return root
}
```

####  106 从中序和后序遍历序列构造二叉树【中等】

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

#### 110 平衡二叉树【简单】

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

#### 111 二叉树的最小深度【简单】

题}要求：https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

算法分析

```go
// date 2020/03/21
// 递归算法
func minDepth(root *TreeNode) int {
    if root == nil { return 0 }
    if root.Left == nil { return 1 + minDepth(root.Right) }
    if root.Right == nil { return 1 + minDepth(root.Left) }
    l, r := minDepth(root.Left), minDepth(root.Right)
    if l > r { return r+1 }
    return l+1
}
// bfs广度优先搜索
func minDepth(root *TreeNode) int {
    if root == nil { return 0 }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    var n, depth int
    for len(queue) != 0 {
        n = len(queue)
        for i := 0; i < n; i++ {
            if queue[i].Left == nil && queue[i].Right == nil { return depth+1 }
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
        }
        depth++
        queue = queue[n:]
    }
    return depth
}
// dfs深度优先搜索
func minDepth(root *TreeNode) int {
    var dfs func(root *TreeNode, depth int) int
    dfs = func(root *TreeNode, depth int) int {
        // 空结点直接返回
        if root == nil { return depth }
        // 叶子结点
        if root.Left == nil && root.Right == nil { return depth+1 }
        if root.Left == nil {
            return dfs(root.Right, depth+1)
        }
        if root.Right == nil {
            return dfs(root.Left, depth+1)
        }
        l, r := dfs(root.Left, depth+1), dfs(root.Right, depth+1)
        if l < r { return l }
        return r
    }
    return dfs(root, 0)
}
```

#### 112 路径总和Path Sum【简单】

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

#### 156 上下翻转二叉树【中等】

题目要求：给定一个二叉树，其中所有的右节点要么是具有兄弟节点（拥有相同父节点的左节点）的叶节点，要么为空，将此二叉树上下翻转并将它变成一棵树， 原来的右节点将转换成左叶节点。返回新的根。

算法分析：

```go
// date 2020/02/25
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	var parent, parent_right *TreeNode
  for root != nil {
    // 0.保存当前left, right
    root_left := root.Left
    root.Left = parent_right  // 重新构建root,其left为上一论的right
    parent_right := root.Right // 保存right
    root.Right = parent // 重新构建root,其right为上一论的root
    parent = root
    root = root_left
  }
  return parent
}
```

#### 199 二叉树的右视图【中等】

算法分析：层序遍历的最后一个节点值。

```go
// date 2020/02/26
func rightSideView(root *TreeNode) []int {
  res := make([]int, 0)
  if root == nil { return res }
  stack := make([]*TreeNode, 0)
  stack = append(stack, root)
  for len(stack) != 0 {
    n := len(stack)
    res = append(res, stack[0].Val)
    for i := 0; i < n; i++ {
      if stack[i].Right != nil { stack = append(stack, stack[i].Right) }
      if stack[i].Left != nil { stack = append(stack, stack[i].Left) }
    }
    stack = stack[n:]
  }
  return res
}
```

#### 236 二叉树的最近公共祖先【中等】

题目要求：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

思路分析：

```go
// date 2020/03/29
// 递归
// 在左右子树中分别查找p,q结点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
     if root == nil || root == p || root == q { return root }
     if p == q { return p }
     left := lowestCommonAncestor(root.Left, p, q)
     right := lowestCommonAncestor(root.Right, p, q)
     // 左右均不为空，表示p,q分别位于左右子树中
     if left != nil && right != nil { return root }
     // 左为空，表示p,q位于右子树,否则位于左子树
     if left == nil { return right }
     return left
}
```

#### 257 二叉树的所有路径【简单】

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

```go
class MinStack {
public:
    /** initialize your data structure here. */
    stack<int> _stack;
    int _min = INT_MAX;
    MinStack() {
        
    }
    
    void push(int x) {
        if(_min >= x){
            if(!_stack.empty()){
                _stack.push(_min);
            }
            _min = x;
        }
        _stack.push(x);
    }
    
    void pop() {
        if(_stack.empty())
            return;
        if(_stack.size() == 1)
            _min = INT_MAX;
        else if(_min == _stack.top()){//下一个元素是下一个最小值
            _stack.pop();
            _min = _stack.top();
        }
        _stack.pop();
    }
    
    int top() {
        return _stack.top();
    }
    
    int getMin() {
        return _min;
    }
};
```

#### 270 最接近的二叉搜索树值【简单】

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

#### 538 把二叉树转换成累加树【简单】

算法：逆序的中序遍历，查找。

```go
// date 2020/02/26
func convertBST(root *TreeNode) *TreeNode {
  var sum int
  decOrder(root, &sum)
  return root
}

func decOrder(root *TreeNode, sum *int) {
  if root == nil { return }
  decOrder(root.Right, sum)
  *sum += root.Val
  root.Val = *sum
  decOrder(root.Left, sum)
}
```

#### 543 二叉树的直径【简单】

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

#### 545 二叉树的边界【中等】

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

#### 563 二叉树的坡度【简单】

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

#### 617 合并二叉树【简单】

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

#### 654 最大二叉树【中等】

题目要求：

给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下：

二叉树的根是数组中的最大元素。
			左子树是通过数组中最大值左边部分构造出的最大二叉树。
			右子树是通过数组中最大值右边部分构造出的最大二叉树。
			通过给定的数组构建最大二叉树，并且输出这个树的根节点。

算法分析：

找到数组中的最大值，构建根节点，然后递归调用。

```go
// date 2020/02/25
// 递归版
func constructMaximumBinaryTree(nums []int) *TreeNode {
  if len(nums) == 0 { return nil }
  if len(nums) == 1 { return &TreeNode{Val: nums[0]} }
  p := 0
  for i, v := range nums {
    if v > nums[p] { p = i }
  }
  root := &TreeNode{Val: nums[p]}
  root.Left = constructMaximumBinaryTree(nums[:p])
  root.Right = constructMaximumBinaryTree(nums[p+1:])
  return root
}
```

#### 655 输出二叉树【中等】

题目要求：将二叉树输出到m*n的二维字符串数组中。

算法思路：先构建好res，然后逐层填充。

```go
// date 2020/02/25
func printTree(root *TreeNode) [][]string {
  depth := findDepth(root)
  length := 1 << depth - 1
  res := make([][]string, depth)
  for i := 0; i < depth; i++ {
    res[i] = make([]string, length)
  }
  fill(res, root, 0, 0, length)
}

func fill(res [][]string, t *TreeNode, i, l, r int) {
  if t == nil { return }
  res[i][(l+r)/2] = fmt.Sprintf("%d", root.Val)
  fill(res, t.Left, i+1, l, (l+r)/2)
  fill(res, t.Right, i+1, (l+r+1)/2, r)
}

func findDepth(root *TreeNode) int {
  if root == nil { return 0 }
  l, r := findDepth(root.Left), findDepth(root.Right)
  if l > r { return l+1 }
  return r+1
}
```

#### 662 二叉树的最大宽度【中等】

题目要求：给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。

每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。
		链接：https://leetcode-cn.com/problems/maximum-width-of-binary-tree

算法分析：使用栈stack逐层遍历，计算每一层的宽度，注意全空节点的时候要返回。时间复杂度O(n)，空间复杂度O(logn)。

```go
// date 2020/02/26
func widthOfBinaryTree(root *TreeNode) int {
  res := 0
  stack := make([]*TreeNode, 0)
  stack = append(stack, root)
  var n, i, j int
  var isAllNil bool
  for 0 != len(stack) {
    i, j, n = 0, len(stack)-1, len(stack)
    // 去掉每一层两头的空节点
    for i < j && stack[i] == nil { i++ }
    for i < j && stack[j] == nil { j-- }
    // 计算结果
    if j-i+1 > res { res = j-i+1 }
    // 添加下一层
    isAllNil = true
    for ; i <= j; i++ {
      if stack[i] != nil {
        stack = append(stack, stack[i].Left, stack[i].Right)
        isAllNil = false
      } else {
        stack = append(stack, nil, nil)
      }
    }
    if isAllNil { break }
    stack = stack[n:]
  }
  return res
}
```

#### 669 修剪二叉搜索树【简单】

题目要求：给定一个二叉搜索树，同时给定最小边界L 和最大边界 R。通过修剪二叉搜索树，使得所有节点的值在[L, R]中 (R>=L) 。你可能需要改变树的根节点，所以结果应当返回修剪好的二叉搜索树的新的根节点。

算法分析：

```go
// date 2020/02/25
func trimBST(root *TreeNode, L, R int) *TreeNode {
  if root == nil { return nil }
  if root.Val < L { return trimBST(root.Right, L, R) }
  if root.Val > R { return trimBST(root.Left, L, R) }
  root.Left = trimBST(root.Left, L, R)
  root.Right = trimBST(root.Right, L, R)
  return root
}
```

#### 703 数据流中第K大元素【简单】

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

#### 814 二叉树剪枝【中等】

题目要求：给定一个二叉树，每个节点的值要么是0，要么是1。返回移除了所有不包括1的子树的原二叉树。

算法分析：

使用containOne()函数判断节点以及节点的左右子树是否包含1。

```go
// date 2020/02/25
func pruneTree(root *TreeNode) *TreeNode {
  if containsOne(root) { return root }
  return nil
}

func containsOne(root *TreeNode) bool {
  if root == nil { return false }
  l, r := containsOne(root.Left), containsOne(root.Right)
  if !l { root.Left = nil }
  if !r { root.Right = nil }
  return root.Val == 1 || l || r
}
```

#### 993 二叉树的堂兄弟节点【简单】

题目链接：https://leetcode-cn.com/problems/cousins-in-binary-tree/

题目要求：给定一个二叉树和两个节点元素值，判断两个节点值是否是堂兄弟节点。

堂兄弟节点的定义：两个节点的深度一样，但父节点不一样。

算法分析：

使用findFatherAndDepth函数递归找到每个节点的父节点和深度，并进行比较。

```go
// date 2020/02/25
func isCousins(root *TreeNode, x, y int) bool {
  xf, xd := findFatherAndDepth(root, x)
  yf, yd := findFatherAndDepth(root, y)
  return xd == yd && xf != yf
}
func findFatherAndDepth(root *TreeNode, x int) (*TreeNode, int) {
  if root == nil || root.Val == x { return nil, 0 }
  if root.Left != nil && root.Left.Val == x { return root, 1 }
  if root.Right != nil && root.Right.Val == y { return root, 1 }
  l, lv := findFatherAndDepth(root.Left, x)
  r, rv := findFatherAndDepth(root.Right, x)
  if l != nil { return l, lv+1 }
  return r, rv+1
}
```

#### 1104 二叉树寻路【中等】

题目要求：https://leetcode-cn.com/problems/path-in-zigzag-labelled-binary-tree/

算法分析：

```go
// date 2020/02/25
func pathInZigZagTree(label int) []int {
  res := make([]int, 0)
  for label != 1 {
    res = append(res, label)
    label >>= 1
    y := highBit(label)
    lable = ^(label-y)+y
  }
  res = append(res, 1)
  i, j := 0, len(res)-1
  for i < j {
    t := res[i]
    res[i] = res[j]
    res[j] = t
    i++
    j--
  }
  return res
}label = label ^(1 << (label.bit_length() - 1)) - 1

func bitLen(x int) int {
  res := 0
  for x > 0 {
    x >>= 1
    res++
  }
  return res
}
```

#### 面试题27 二叉树的镜像【简单】

题目链接：https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/

题目要求：给定一棵二叉树，返回其镜像二叉树。

```go
// date 2020/02/23
// 递归算法
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

