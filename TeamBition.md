### TeamBition练习题

#### 合并K个排序链表

算法一：两两合并，递归版

```go
// date 2020/02/21
func mergeKLists(lists []*ListNode) *ListNode {
  if len(lists) == 0 { return nil }
  if len(lists) == 1 { return lists[0] }
  nHead := &ListNode{Val: -1}
  pre := nHead
  p1, p2 := lists[0], lists[1]
  for p1 != nil && p2 != nil {
    if p1.Val < p2.Val {
      pre.Next = p1
      p1 = p1.Next
    } else {
      pre.Next = p2
      p2 = p2.Next
    }
    pre = pre.Next
  }
  return mergeKLists(nHead.Next, lists[1:])
}
```

