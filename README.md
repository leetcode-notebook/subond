# leetcode

数据结构和算法是一个程序员的基石，本仓库用于个人学习基本数据结构和算法。

题目完成度：[263/500]

[TOC]

---

# 一、基本数据结构

##  [数组Array](array/array.md)

数组属于线性表的一种，底层使用连续的空间进行存储，因此在声明数组的时候需要指定数组的大小，即需要申请多大的内存空间。

### 相关题目
   - 27 Remove Element[两种算法]
   - 189 旋转数组
   - 从排序数组中删除重复项
   - 80 从排序数组中删除重复项 II
   - 54 螺旋矩阵
   - 整数反转
   - 118 杨辉三角【E】


## [链表Linked List](linkedlist/linkedlist.md)

链表也属于线性表的一种，底层使用非连续的空间进行存储，因此在声明链表的时候不需要指定链表大小。

### 相关题目
    - 2 两数相加
    - 19 Remove Nth Node From End of List
    - 24 Swap Nodes in Pairs【M】
    - 61 Rotate List【旋转链表】
    - 82 Remove Duplicates From Sorted List II
    - 83 Remove Duplicates From Sorted List
    - 86 Partition List
    - 92 反转链表II【M】
    - 160 Intersection of Two Linked List【两个链表的交点】
    - 143 Reorder List【中等】143 Reorder List
    - 142 Linked List Cycle II【中等】
    - 141  Linked List Cycle
    - 147 对链表进行插入排序【M】
    - 148 排序链表【M】
    - 328 奇偶链表
    - 876 Middle of the Linked List
    - 234 回文链表
    - 206 反转链表
    - 203 移除链表元素
    - 合并两个有序链表

## [队列Queue](queue/queue.md)

### 基本定义
### 相关题目
    - 621 Task Scheduler【任务调度器】
    - 622 Design Circular Queue 
    - 641 Design Circular Deque【M】
    - 933 Number of Recent Calls

## [栈Stack](stack/stack.md)




## [字符串String](string/string.md)

golang语言中byte是uint8的别名，rune是int32的别名，字符串中的单个字符有byte表示。

### 相关题目
    - 14 Longest Common Prefix
    - 有效的字母异位词
    - 反转字符串
    - 246 中心对称数【E】
    - 58 最后一个单词的长度
    - 521  最长特殊序列I【E】

## [哈希表HashTable](hashTable/hashTable.md)

### 哈希表的原理
### 哈希集合的设计
### 哈希映射的设计
### 相关题目
    - 1 两数之和
    - 202 快乐数
    - 36 有效的数独
    - 170 两数之和III - 数据结构设计
    - 349 两个数组的交集
    - 652 寻找重复的树【中等】
    - 49 字母异位词分组【M】
    - 四数相加II
    - 常数时间插入，删除和获取
    - 146 LRU缓存机制【M】

## [数学题目](math/math.md)

### 相关题目
    - 8 String to Integer
    - 9 Palindrome Number
    - 15 三数之和

---

# 二、抽象数据结构

优先队列是一种抽象的数据结构，可以使用堆实现。

## [堆Heap](heap/heap.md)

[堆](heap/heap.md)

### 堆的定义
### 堆的算法
### 堆的排序
### 相关题目
    - 912 排序数组【M】
    - 1046 最后一块石头的重量
    - 215 数组中第K个最大元素
    - 面试题40 最小的K个数
    - 面试题17.09 第K个数

##  [二叉树Binary Tree](binaryTree/bianryTree.md)

二叉树属于非线性表的一种，与链表不同的是，二叉树有两个后继节点。

[二叉树](binaryTree/binaryTree.md)

### 什么是二叉树
### 二叉树的遍历
### 递归解决二叉树问题
### 相关题目
    - 从中序和后序遍历序列构造二叉树[前序和中序]
    - 98 验证二叉搜索树【M】
    - 104 二叉树的最大深度【E】
    - 111 二叉树的最小深度【E】
    - 101 对称二叉树【E】
    - 156 上下翻转二叉树
    - 257 二叉树的所有路径
    - 270 最接近的二叉搜素树值【E】
    - 543 二叉树的直径【E】
    - 545 二叉树的边界【M】
    - 563 二叉树的坡度【E】
    - 617 合并二叉树【E】
    - 654 最大二叉树【M】
    - 655 输出二叉树【M】
    - 662 二叉树的最大宽度【M】
    - 814 二叉树剪枝【M】
    - 993 二叉树的堂兄弟节点
    - 1104 二叉树寻路【M】
    - 面试题27 二叉树的镜像
    - 路径总和
    - 不同的二叉搜索树


## [二叉搜索树Binary Search Tree](binaryTree/bst.md)

[二叉搜索树](binaryTree/bst.md)

### 什么是二叉搜索树
### 二叉搜索树的基本操作
    - 查询
    - 最小关键字元素
    - 最大关键字元素
    - 前继节点
    - 后继节点
    - 插入
    - 删除
### 相关题目
    - 98 验证二叉搜索树【M】
    - 108 将有序数组转换为二叉搜索树【M】
    - 235 二叉搜素树的最近公共祖先
    - 450 删除二叉搜索树中的节点
    - 1038 从二叉搜索树到更大和树【M】
    - 1214 查找两棵二叉搜索树之和【M】
    - 面试题 17.12 BiNode【E】
    - 面试题54 二叉搜索树的第K大节点

# 三、基本算法理论

算法不仅仅是一种解决问题的思路，更是一种思想。

## 排序算法

## [二分查找](binarySearch/binarySearch.md)

### 相关题目
    - 704 二分查找
    - 162 寻找峰值
    - 658 找到K个最接近的元素 

## [贪心算法](greedy/greed_algorithm.md)

### 算法介绍
### 相关题目
    - 55 跳跃游戏
    - 45 跳远游戏II

## 递归思想



## [广度优先搜索](bfs/bfs.md)



## [深度优先搜索](dfs/dfs.md)



## [动态规划](dynamicProgramming/dp.md)

### 算法介绍
### 解题思路
### 相关题目
    - 64 最小路径和【M】
    - 70 爬楼梯
    - 72 编辑距离【H】
    - 121 买卖股票的最佳时机
    - 122 买卖股票的最佳时机II
    - 123 买卖股票的最佳时机III
    - 188 买卖股票的最佳时机IV
    - 746 使用最少花费爬楼梯
    - 264 丑数【M】
    - 300 最长上升子序列【M】

## 剪枝思想

剪枝是搜索算法的优化手段，目的是为了去掉不必要的搜索。

# [其他题目](others.md)

- 位1的个数
- 汉明距离
- 292 Nim游戏
- 89 格雷编码
