## Linked List[链表]

[TOC]

### [基本操作](linkedlist_basic.md)

### 1.常用技巧

#### 1.1 快慢指针

在数组章节我们学习的双指针技巧，通常是两个指针从前后一起遍历，向中间逼近；或者读写指针维护不同的数据元素；而在链表的运用中，由于单向链表只能单向遍历，通过调整两个指针的遍历速度完成某种特定任务，因此也称为”快慢指针“技巧。

参见题目

- 19—删除链表中的倒数第N个节点
- 141—判断链表是否有环
- 142—判断链表是否有环，如果有，返回环入口
- 160—相交链表，求两个链表的交点



回顾：运用双指针技巧的几点注意事项，1）在运用Next节点时一定要判断节点是否为空，举例`fast.Next.Next`一定要判断`fast != nil && fast.Next != nil`；2）注意循环结束的条件。



#### 1.2 哑结点

哑结点，也称为哨兵结点，主要应用与某些头节点不明确的场景。相关题目：

1. [002两数相加](002.md)
2. [021合并两个有序链表](021.md)





#### 1.3 经典题目

关于的链表的经典问题，主要包括题目206反转链表，203移除链表元素，234回文链表和328奇偶链表。





### 2. 相关题目

#### 148 排序链表

题目要求在O(nlogn)时间复杂度和常数级空间复杂度完成对链表的排序。

思路分析：归并排序

1.找出链表的终点，切成两个链表

2.递归链表进行排序

3.合并链表，利用哑结点。

```go
func sortList(head *ListNode) *ListNode {
  if head == nil || head.Next == nil {return head}
  // find the mid node
  slow, fast := head, head.Next
  for fast != nil || fast.Next != nil {
    slow, fast = slow.Next, fast.Next.Next
  }
  mid := slow.Next
  slow.Next = nil
  left, right := sortList(head), sortList(mid)
}

//算法二：归并排序，非递归版
func sortList(head *ListNode) *ListNode {
  dummy := &ListNode{Val: -1}
  dummy.Next = head
  pre := head
  length := 0
  for pre != nil {
    length, pre = length+1, pre.Next
  }
  for i := 1; i < length; i <<= 1 {
    cur := dummy.Next
    tail := dummy
    for cur != nil {
      left := cur
      right := cut(left, size)
      cur = cut(right, size)
      tail.Next = merge(left, right)
      for tail != nil {
        tail = tail.Next
      }
    }
  }
  return dummy.Next
}

func cut(head *ListNode, int n) *ListNode {
  pre := head
  for n > 0 && pre != nil {
    n, pre = n-1, pre.Next
  }
  if pre == nil {return nil}
  after := pre.Next
  pre.Next = nil
  return after
}

func merge(h1, h2 *ListNode) *ListNode {
  dummy := &ListNode{Val: -1}
  pre := dummy
  for h1 != nil && h2 != nil {
    if h1.Val < h2.Val {
      pre.Next, h1 = h1, h1.Next
    } else {
      pre.Next, h2 = h2, h2.Next
    }
    pre = pre.Next
  }
  if h1 != nil {
    pre.Next = h1
  }
  if h2 != nil {
    pre.Next = h2
  }
  return dummy.Next
}
```

#### 203 移除链表元素【E】

题目链接：https://leetcode.com/problems/remove-linked-list-elements/

思路分析：

算法1：利用哑结点，直接链接所有目标节点。

```go
// 算法1：哑结点【处理起来容易】
func removeElements(head *ListNode, val int) *ListNode {
  if head == nil { return head }
  dummy := &ListNode{Val: -1}
  pre := dummy
  for head != nil {
    if head.Val != val {
      pre.Next = head
      pre = pre.Next
    }
    head = head.Next
  }
  pre.Next = head
  return dummy.Next
}
```

算法2：先找到新的头节点，然后在依次链接所有目标节点。

```go
// 算法2
func removeElements(head *ListNode, val int) *ListNode {
    var newHead *ListNode
    // 找到新的头结点
    for head != nil {
        if head.Val != val {
            newHead = head
            head = head.Next
            break
        }
        head = head.Next
    }
    // 如果头结点为空，表示没有找到符合条件的头结点，直接返回空
    if newHead == nil { return newHead }
    pre := newHead
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

#### 328 奇偶链表【M】

题目链接：https://leetcode.com/problems/odd-even-linked-list/

算法分析

算法1：利用哑结点，时间复杂度O(n)

```go
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    // 两个哑结点，分别记录奇数和偶数的结点
    d1 := &ListNode{Val: -1}
    d2 := &ListNode{Val: -1}
    p1, p2, p := d1, d2, head
    for p != nil {
        // 奇数结点
        p1.Next = p
        p1 = p1.Next
        if p.Next == nil {
            break
        }
        // 偶数结点
        p = p.Next
        p2.Next = p
        p2 = p2.Next
        p = p.Next
    }
    p2.Next = nil
    p1.Next = d2.Next
    return d1.Next
}
```



