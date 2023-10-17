## 24 两两交换链表中的节点-中等

题目：

给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。





分析：

算法1：递归实现，更加巧妙，【推荐该算法】

```go
// date 2023/10/17
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    // basic case
    if head == nil || head.Next == nil {
        return head
    }
    var newHead *ListNode
    unSwap := head.Next.Next // save un-swap node
    newHead = head.Next      // find the new head
    newHead.Next = head      // update next pointer of new head
    head.Next = swapPairs(unSwap)  // swap others
    return newHead
}
```



算法2：

先找出新的头节点，并保存两两交换后的第二个节点，用于更新其Next指针。

```go
// date 2022/10/09
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var nHead *ListNode
    // 1. 找到新的头节点
    nHead = head.Next
    // save unswap list
    after := nHead.Next
    p1, p2 := head, head.Next
    p1.Next = after
    p2.Next = p1
    // 2. 继续交换
    pre := p1
    p1 = after
    for p1 != nil && p1.Next != nil {
        p2 = p1.Next
        after = p2.Next
        p1.Next = after
        p2.Next = p1
        pre.Next = p2
        pre = p1
        p1 = after
    }
    return nHead
}
```

