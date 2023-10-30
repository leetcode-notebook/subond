## 148 排序链表-中等

题目：

给你链表的头节点 head，请你将其按升序排列并返回排序后的链表。



分析：

算法1：

自顶向下的归并排序。

1. 找到链表的中点，以中点为界分成两个子链。
2. 对子链分别进行排序。
3. 对排序的子链，进行合并。

这里依赖基本 case，即子链中只有一个元素时，就是递归的返回原点。

```go
// date 2023/10/17
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    spre, slow, fast := head, head, head
    for fast != nil && fast.Next != nil {
        spre = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    if fast != nil {
        spre = slow
        slow = slow.Next
    }
    spre.Next = nil
    left, right := sortList(head), sortList(slow)
    dumy := &ListNode{}
    pre := dumy
    for left != nil && right != nil {
        if left.Val < right.Val {
            pre.Next = left
            left = left.Next
        } else {
            pre.Next = right
            right = right.Next
        }
        pre = pre.Next
    }
    if left != nil {
        pre.Next = left
    }
    if right != nil {
        pre.Next = right
    }
    return dumy.Next
}
```



