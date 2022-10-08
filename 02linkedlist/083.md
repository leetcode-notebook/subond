## 83 删除排序链表中的重复元素-简单

> 给定一个已排序的链表的头 `head` ， *删除所有重复的元素，使每个元素只出现一次* 。返回 *已排序的链表* 。
>
> https://leetcode.cn/problems/remove-duplicates-from-sorted-list/



算法分析：

如果当前指针和当前指针的Next不为空，直接判断去重。

```go
// date 2022/10/08
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    pre := head
    for pre != nil && pre.Next != nil {
        if pre.Val == pre.Next.Val {
            pre.Next = pre.Next.Next
            continue
        }
        pre = pre.Next
    }
    return head
}
```

