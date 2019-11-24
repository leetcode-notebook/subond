

## Linked List

### Linked List【链表】

#### 相关题目

- 19 Remove Nth Node From End of List
- 61 Rotate List【旋转链表】
- 82 Remove Duplicates From Sorted List II
- 83 Remove Duplicates From Sorted List
- 86 Partition List
- 160 Intersection of Two Linked List【两个链表的交点】
- 143 Reorder List【中等】143 Reorder List
- 142 Linked List Cycle II【中等】
- 141 
- 876 Middle of the Linked List
- 234 回文链表
- 206 反转链表
- 203 移除链表元素

#### 19 Remove Nth Node From End of List

思路分析：

双指针slow和fast，fast快指针先走n步(n有效，所以最坏的情况是快指针走到表尾，即是删除表头元素)；然后slow指针和fast指针同走，当fast指针走到最后一个元素时，因为slow与fast慢指针差n步，slow刚好为欲删除节点的前驱节点。

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {return nil}
    slow, fast := head, head
    for n > 0 && fast != nil {
        fast = fast.Next
        n--
    }
    // remove the head
    if fast == nil {
        return head.Next
    }
    for fast.Next != nil {
        slow, fast = slow.Next, fast.Next
    }
    slow.Next = slow.Next.Next
    return head
}
```



#### 61 Rotate List【旋转链表】

思路分析：

算法1：直接操作。先求出链表长度L，k = k % L；然后让表头移动step(L-k-1)步到达新的尾节点tail，tail继续走到表尾，并指向表头。

```go
/*
type ListNode {
  Val int
  Next *ListNode
}
  输入: 1->2->3->4->5->NULL, k = 2
  输出: 4->5->1->2->3->NULL
*/
// 算法1：直接操作
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {return head}
    // find the len of linkedlist
    len := 0
    pre := head
    for pre != nil {
        len++
        pre = pre.Next
    }
    // k = k % len
    k = k % len
    step := len - k - 1
    pre = head
    for step > 0 {
        pre = pre.Next
        step--
    }
    tail := pre
    for tail.Next != nil {
        tail = tail.Next
    }
    tail.Next = head
    head = pre.Next
    pre.Next = nil
    return head
}

// 算法2：线形成环，同时计算出len
// 新表尾 n - k - 1;新表头 n - k

```

思路分析：

算法2：计算链表长度的同时将线性链表组织成环形链表(优势：相比算法1省去了重新将表尾元素指向表头的操作)；此时，找到新的表尾L-k-1，新表头L-k。【**推荐使用该算法**】

```go
// 算法2
func rotateRight(head *ListNode, k int) *ListNode {
  if head == nil || head.Next == nil {
    return head
  }
  pre := head
  l := 1
  for pre.Next != nil {
    pre = pre.Next
    l++
  }
  k = k % l
  // 将链表形成环，然后表头走Len-k-1步到达新的尾部，然后断开。
  pre.Next = head
  step := l - k - 1
  pre = head
  for i := 0; i < step; i++ {
    pre = pre.Next
  }
  head = pre.Next
  pre.Next = nil
  return head
}
```



#### 82 Remove Duplicates From Sorted List II

思路分析：

**哑节点+快慢指针**

fast：发现重复元素；slow：存放结果

```go
// 83升级版，所有重复元素都不要，只保留出现过一次的元素
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummy := &ListNode{Val: -1}
    dummy.Next = head
    slow := dummy
    fast := head
    for fast.Next != nil {
        if fast.Val != fast.Next.Val {
            if slow.Next == fast {
                // 说明slow和fast之间没有重复元素，更新slow节点
                slow = fast
            } else {
                // 说明slow和fast之间有重复元素，更新slow的next节点
                slow.Next = fast.Next
            }
        }
        fast = fast.Next
    }
    // coner case
    if slow.Next != nil && slow.Next != fast && slow.Next.Val == fast.Val {
        slow.Next = nil
    }
    return dummy.Next
}
```

#### 83 Remove Duplicates From Sorted List

```go
// 算法1：直观
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

#### 86 Partition List

```go
// 算法：哑结点+两个指针
func partition(head *ListNode, x int) *ListNode {
  before := &ListNode{Val: -1}
  after := &ListNode{Val: -1}
  bhead := before
  ahead := after
  for head != nil {
    if head.Val < x {
      before.Next = head
      before = before.Next
    } else {
      after.Next = head
      after = after.Next
    }
    head = head.Next
  }
  after.Next = nil
  before.Next = ahead.Next
  return bhead.Next
}
```

#### 160 Intersection of Two Linked List【两个链表的交点】

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
```



#### 143 Reorder List 

**注意**：要对每一段代码很熟悉，要精确的知道每段代码的结果是什么样子；然后大问题拆成小问题去解决。

```go
/* 思路分析：
1. 快慢指针找到链表中点
2. 将后半段链表反转
3. 重新排列链表
*/
func reorder(head *ListNode) {
  if head == nil || head.Next == nil || head.Next.Next == nil {return}
  pre, slow, fast := head, head, head
  for fast != nil && fast.Next != nil {
    pre = slow
    slow, fast = slow.Next, fast.Next.Next
  }
  if fast != nil {
    pre, slow = slow, slow.Next
  }
  pre.Next = nil
  
  var tail *ListNode
  for slow.Next != nil {
    slowN := slow.Next
    slow.Next = tail
    tail = slow
    slow = slowN
  }
  slow.Next = tail
  
  pre = head
  for pre != nil && slow != nil {
    preN := pre.Next
    slowN := slow.Next
    pre.Next = slow
    slow.Next = preN
    pre = preN
    slow = slowN
  }
}
```

#### 141 Linked List Cycle

```go
// 算法1：快慢指针
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow, fast = slow.Next, fast.Next.Next
        if slow == fast {return true}
    }
    return false
}
```

#### 142 Linked List Cycle II

```go
// 算法1：利用额外空间，时间复杂度和空间复杂度均为O(n)
func detectCycle(head *ListNode) *ListNode {
  visited := make(map[*ListNode]struct{})
  for head != nil {
    if _, ok := visited[head]; ok {
      return head
    }
    visited[head] = struct{}{}
    head = head.Next
  }
  return nil
}
// 算法2：Floyd算法
/*
假设：环外有F个元素，环上C个元素
快慢指针判断是否有相遇，相遇(h节点)时，快指多走一个环【F+a+b+a】；慢指针走了【F+a】；因为2(F+a) = F+a+b+a,所以f=b
因此，相遇节点和头节点同时走f步，即是环的入口
*/
func detectCycle(head *ListNode) *ListNode {
  slow, fast := head, head
  for fast != nil && fast.Next != nil {
    slow, fast = slow.Next, fast.Next.Next
    if slow == fast {
      break
    }
  }
  if fast == nil || fast.Next == nil {return nil}
  fast = head
  for fast != slow {
    slow, fast = slow.Next, fast.Next
  }
  return fast
}
```



#### 876 Middle of the Linked List

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



#### 234 回文链表

```go
// 算法：快慢指针，找到中间点；反转前半段的链表
func isPalindrome(head *ListNode) bool {
  if head == nil || head.Next == nil { return true }
  pre, slow, fast := head, head, head
  var tail *ListNode
  for fast != nil && fast.Next != nil {
    pre = slow
    slow = slow.Next
    fast = fast.Next.Next
    pre.Next = tail
    tail = pre
  }
  // the last one
  if fast != nil {
    slow = slow.Next
  }
  for pre != nil && slow != nil {
    if pre.Val != slow.Val { return false }
    pre = pre.Next
    slow = slow.Next
  }
  return true
}
```



#### 206 反转链表

```go
// 算法1：迭代
func reverseList(head *ListNode) *ListNode {
  if head == nil || head.Next == nil {return head}
  var tail *ListNode
  pre := head
  for pre != nil {
    // the list end
    if pre.Next == nil {
      pre.Next = tail
      head = pre
      break
    }
    after := pre.Next
    pre.Next = tail
    tail = pre
    pre = after
  }
  return head
}
// 算法2：递归
```



#### 203 移除链表元素

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
// 算法2
func removeElements(head *ListNode, val int) *ListNode {
  if head == nil {return head}
  pre := head
  var newHead *ListNode
  for pre != nil {
    // found new head
    if pre.Val != val {
      newHead = pre
      head = pre
      pre = pre.Next
      break
    }
    pre = pre.Next
  }
  // not found new head, return nil
  if newHead == nil {
    return nil
  }
  for pre != nil {
    if pre.Val != val {
      newHead.Next = pre
      newHead = newHead.Next
    }
    pre = pre.Next
  }
  newHead.Next = pre
  return head
}
```

