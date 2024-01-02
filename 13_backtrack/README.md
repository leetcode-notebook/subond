## 回溯算法

回溯算法和二叉树的 DFS 类似，本质上是一种暴力穷举算法。

抽象地说，解决一个回溯问题，实际上就是遍历一颗决策树的过程，树的每个叶子节点存放着一个合法答案。

我们把整棵树遍历一遍，把叶子节点上的答案全部收集起来，就能得到所有的合法答案。



站在回溯决策树的一个节点上，你只需要思考三个问题：

1. 路径：那些已经做出的选择
2. 选择列表：就是你当前可以做的选择
3. 结束条件：到达决策树的底层，无法再做选择的条件



代码方面，回溯算法的基本框架是：

```go
result = make([][]int, 0, 16)

func backtrace(选择列表, 路径) {
  if 满足结束条件 {
    result = append(result, path)
    return
  }
  for 选择 := range 选择列表 {
    做出选择
    backtrace(路径, 选择列表) // 或者叫 待选择列表
  }
}
```



## 相关题目

第一类：很经典的回溯算法

- 第 39 题，组合总和，【重点看，很经典】
- 第 40 题，组合总和，不可重复使用【重点看，很经典】
- 第 46 题，全排列，【重点看，很经典】
- 第 77 题，组合
- 第 78 题，子集
- 第 90 题，子集2，结果去重
- 第 216 题，组合总和3，不可重复，和为目标值【很经典】



第二类：重在理解 DFS

- 第 17 题，电话号码的的字母组合
- 第 22 题，括号生产
- 第 79 题，单词搜索
- 第 93 题，复原IP地址
- 第 494 题，目标和，顺序递归回溯
- 第 784 字母大小写全排列，顺序递归回溯
- 第 797 题，所有可能的路径，图算法【很基本，看看】