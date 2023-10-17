## 023 合并K个升序链表-困难

题目：

给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



分析：

复习递归算法，掌握自底向上和自上向下的递归方法。

算法1：递归算法，自底向上的两两合并。每个链表遍历一遍，所示时间复杂度O(n*k)，空间复杂度O(1)。

```go
// date 2023/10/17
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }
    k := len(lists)
    l1, l2 := mergeKLists(lists[:k-1]), lists[k-1]
    dumy := &ListNode{}
    pre := dumy
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            pre.Next = l1
            l1 = l1.Next
        } else {
            pre.Next = l2
            l2 = l2.Next
        }
        pre = pre.Next
    }
    if l1 != nil {
        pre.Next = l1
    }
    if l2 != nil {
        pre.Next = l2
    }
    return dumy.Next
}
```

