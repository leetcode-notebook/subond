## Linked List[链表]

[TOC]

### [基本操作](linkedlist_basic.md)

### 1.常用技巧

#### 1.1 快慢指针

在数组章节我们学习的双指针技巧，通常是两个指针从前后一起遍历，向中间逼近；或者读写指针维护不同的数据元素；而在链表的运用中，由于单向链表只能单向遍历，通过调整两个指针的遍历速度完成某种特定任务，因此也称为”快慢指针“技巧。

参见题目19, 141, 142, 160。

回顾：运用双指针技巧的几点注意事项，1）在运用Next节点时一定要判断节点是否为空，举例`fast.Next.Next`一定要判断`fast != nil && fast.Next != nil`；2）注意循环结束的条件。



#### 1.2 哑结点

哑结点，也称为哨兵结点，主要应用与某些头节点不明确的场景。相关题目：

1. [002两数相加](002.md)
2. [021合并两个有序链表](021.md)





#### 1.3 经典题目

关于的链表的经典问题，主要包括题目206反转链表，203移除链表元素，234回文链表和328奇偶链表。





### 2. 相关题目

#### 147 Insertion Sort List【对链表进行插入排序】【M】

思路分析

直接按插入排序操作即可，注意已排序的链表的表尾一定要指向nil。

```go
func insertionSortList(head *ListNode) *ListNode {
  if head == nil || head.Next == nil { return head }
  pre, newHead := head.Next, head
  // newHead为已排序的链表，置表尾为nil
  newHead.Next = nil
  for pre != nil {
    // 保存当前节点的下一个节点
    after := pre.Next
    // 如果当前节点比已排序节点的表头还小，则更新已排序的表头
    if newHead.Val > pre.Val {
      pre.Next = newHead
      newHead = pre
      pre = after
      continue
    }
    // 将当前节点插入已排序的链表中
    p1, p1pre := newHead.Next, newHead
    for p1 != nil {
      if p1.Val > pre.Val {
        pre.Next = p1
        p1pre.Next = pre
        break
      }
      p1pre, p1 = p1, p1.Next
    }
    // 如果当前节点大于所有已排序链表，插入已排序链表的表尾
    if p1 == nil {
      pre.Next = nil
      p1pre.Next = pre
    }
    pre = after
  }
  return newHead
}
```

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

#### 160 Intersection of Two Linked List【E】

题目要求：https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

思路分析：两个指针走过相同的节点数，如果能够相遇，则即为交点。

```go
// 算法：双指针 len a + b = b + a
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {return nil}
    pa, pb := headA, headB
    for pa != pb {
        pa, pb = pa.Next, pb.Next
        if pa == nil && pb == nil {
            return nil
        }
        if pa == nil {
            pa = headB
        }
        if pb == nil {
            pb = headA
        }
    }
    return pa
}
// date 2020/03/29
// 第二版
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil { return nil }
    var res *ListNode
    p1, p2 := headA, headB
    for p1 != nil || p2 != nil {
        if p1 == nil {
            p1 = headB
        }
        if p2 == nil {
            p2 = headA
        }
        if p1 == p2 {
            res = p1
            break
        }
        p1 = p1.Next
        p2 = p2.Next
    }
    return res
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

#### 234 回文链表【E】

题目要求：https://leetcode.com/problems/palindrome-linked-list/

思路分析：快慢指针，找到中间结点，反转前半段链表

```go
// 算法：快慢指针，找到中间点；反转前半段的链表
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil { return true }
    pre, slow, fast := head, head, head
    // 反转前半段链表
    var tail *ListNode
    for fast != nil && fast.Next != nil {
        pre = slow
        slow = slow.Next
        fast = fast.Next.Next
        pre.Next = tail
        tail = pre
    }
    // 奇数个结点
    if fast != nil {
        slow = slow.Next
    }
    for pre != nil && slow != nil {
        if pre.Val != slow.Val {return false }
        pre = pre.Next
        slow = slow.Next
    }
    return true
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

#### 707 设计链表【M】

题目链接：https://leetcode-cn.com/problems/design-linked-list/

思路分析：

```go
type MyLinkedList struct {
    Len int
    Head *LinkedNode
}

type LinkedNode struct {
    Val int
    Next *LinkedNode
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
    return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
    if index < 0 || this.Len == 0 || this.Head == nil || index >= this.Len { return -1 }
    if index == 0 { return this.Head.Val }
    pre := this.Head
    for index > 0 {
        if pre == nil { break }
        pre = pre.Next
        index--
    }
    if index > 0 { return -1}
    return pre.Val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
    node := &LinkedNode{Val: val}
    if this.Head == nil {
        this.Head = node
        this.Len = 1
    } else {
        node.Next = this.Head
        this.Head = node
        this.Len += 1
    }
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    node := &LinkedNode{Val: val}
    if this.Head == nil {
        this.Head = node
        this.Len = 1
        return
    }
    pre := this.Head
    for pre.Next != nil {
        pre = pre.Next
    }
    pre.Next = node
    this.Len += 1
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index > this.Len { return }
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    if index == this.Len {
        this.AddAtTail(val)
        return
    }
    node := &LinkedNode{Val: val}
    pre := this.Head
    for index > 1 {
        pre = pre.Next
        index--
    }
    node.Next = pre.Next
    pre.Next = node
    this.Len += 1
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.Len { return }
    if index == 0 {
        this.Head = this.Head.Next
        this.Len -= 1
        if this.Len == 0 {
            this.Head = nil
        }
        return
    }
    pre := this.Head
    for index > 1 {
        pre = pre.Next
        index--
    }
    n := pre.Next.Next
    pre.Next = n
    this.Len -= 1
    if this.Len == 0 {
        this.Head = nil
    }
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
```

#### 876 Middle of the Linked List

题目要求：https://leetcode-cn.com/problems/middle-of-the-linked-list/

思路分析：快慢指针

```go
// 算法1：快慢指针
func middleNode(head *ListNode) *ListNode {
  if head == nil {return head}
  slow, fast := head, head
  for fast != nil && fast.Next != nil {
    slow = slow.Next
    fast = fast.Next.Next
  }
  return slow
}
```

### [关于链表的其他问题](linkedlist_others.md)

