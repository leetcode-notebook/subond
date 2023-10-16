## 147 对链表进行插入排序-中等

题目：

给定单链表的头 head，使用 **插入排序** 对链表进行排序，并返回 排序后的链表的头。



分析：

为了代码简洁我们写一个已排序的链表中插入单个节点的辅助函数。

```go
// date 2023/10/16
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    pre := head.Next

    nHead := head
    nHead.Next = nil

    for pre != nil {
        after := pre.Next
        pre.Next = nil
        nHead = insertOneNode(nHead, pre)
        pre = after
    }
    return nHead
}

func insertOneNode(head, node *ListNode) *ListNode {
    if head == nil || node == nil {
        return head
    }
    if head.Val > node.Val {
        node.Next = head
        head = node
        return head
    }
    pre, pp := head, head
    for pre != nil {
        if pre.Val <= node.Val {
            pp = pre
            pre = pre.Next
            continue
        }
        node.Next = pre
        pp.Next = node
        return head
    }
    if pre == nil {
        node.Next = nil
        pp.Next = node
    }
    return head
}
```

