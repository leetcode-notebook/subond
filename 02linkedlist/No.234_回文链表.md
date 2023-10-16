## 234 回文链表-简单

题目：

给你一个单链表的头节点 head，请你判断该链表是否为回文链表。如果是，返回 true，否则，返回 false。



分析：

快慢指针，找到中间节点，反转前半段。

```go
// date 2023/10/16
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    pre, slow, fast := head, head, head
    var tail *ListNode
    for fast != nil && fast.Next != nil {
        pre = slow
        slow = slow.Next
        fast = fast.Next.Next

        pre.Next = tail
        tail = pre
    }
		// 计数个节点
    if fast != nil {
        slow = slow.Next
    }
    for pre != nil && slow != nil {
        if pre.Val != slow.Val {
            return false
        }
        pre = pre.Next
        slow = slow.Next
    }
    return true
}
```

