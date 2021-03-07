## Linked List[链表]

[TOC]

### [基本操作](linkedlist_basic.md)

### 常用技巧

#### 快慢指针

在数组章节我们学习的双指针技巧，通常是两个指针从前后一起遍历，向中间逼近；或者读写指针维护不同的数据元素；而在链表的运用中，由于单向链表只能单向遍历，通过调整两个指针的遍历速度完成某种特定任务，因此也称为”快慢指针“技巧。

参见题目19, 141, 142, 160。

回顾：运用双指针技巧的几点注意事项，1）在运用Next节点时一定要判断节点是否为空，举例`fast.Next.Next`一定要判断`fast != nil && fast.Next != nil`；2）注意循环结束的条件。

#### 哑结点

哑结点，也称为哨兵结点，主要应用与某些头节点不明确的场景，例如题目21合并两个有序链表。

#### 经典题目

关于的链表的经典问题，主要包括题目206反转链表，203移除链表元素，234回文链表和328奇偶链表。

### 相关题目

#### 2 两数相加

题目要求：https://leetcode-cn.com/problems/add-two-numbers/

思路分析：

算法1：直接计算。将结果存入其中一个链表。

```go
// 算法1
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil { return l2 }
    if l2 == nil { return l1 }
    h1, h2 := l1, l2
    var temp, carry int
    for h1 != nil && h2 != nil {
        temp = h1.Val + h2.Val + carry
        carry = temp / 10
        h1.Val = temp % 10
        // 两个链表同时走到尾部，查看进位
        if h1.Next == nil && h2.Next == nil {
            if carry == 1 {
                h1.Next = &ListNode{Val: 1}
            }
            return l1
        }
        // 默认l1走到尾部，查看l2的剩余结点
        if h1.Next == nil && h2.Next != nil {
            h1.Next = h2.Next
            h1 = h1.Next
            break
        }
        h1 = h1.Next
        h2 = h2.Next
    }
    for h1 != nil {
        temp = h1.Val + carry
        carry = temp / 10
        h1.Val = temp % 10
        if h1.Next == nil {
            if carry == 1 {
                h1.Next = &ListNode{Val: 1}
            }
            return l1
        }
        h1 = h1.Next
    }
    return l1
}
// date 2020/03/29
// 算法二：直接计算，并先开辟空间存储结果
// 这个解法更好
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{Val:-1}
    pre := dummy
    temp_res, n1, n2, carry := 0, 0, 0, 0
    for l1 != nil || l2 != nil || carry > 0 {
        n1, n2 = 0, 0
        if l1 != nil {
            n1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            n2 = l2.Val
            l2 = l2.Next
        }
        temp_res = n1 + n2 + carry
        carry = temp_res / 10
        pre.Next = &ListNode{Val: temp_res%10}
        pre = pre.Next
    }
    return dummy.Next
}
```

#### 19 Remove Nth Node From End of List【M】

题目链接：https://leetcode.com/problems/remove-nth-node-from-end-of-list/

题目要求：

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
    // 保护逻辑，如果n大于零，说明欲要删除的节点超过链表长度
    if n > 0 { return head }
    // fast走到链表的尾部，则n刚好为链表的长度，即删除头节点
    if fast == nil {
        return head.Next
    }
    // fast指针继续走到链表的尾部，则slow指针刚好为欲要删除的节点的前继节点
    for fast.Next != nil {
        slow, fast = slow.Next, fast.Next
    }
    slow.Next = slow.Next.Next
    return head
}
```

#### 21 合并两个有序链表

题目要求：https://leetcode-cn.com/problems/merge-two-sorted-lists/

思路分析：利用哑结点

```go
// date 2020/03/29 2020/05/08
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    dummy := &ListNode{Val: -1}
    pre := dummy
    for l1 != nil || l2 != nil {
        if l1 == nil {
            pre.Next = l2
            break
        }
        if l2 == nil {
            pre.Next = l1
            break
        }
        if l1.Val < l2.Val {
            pre.Next = l1
            l1 = l1.Next
        } else {
            pre.Next = l2
            l2 = l2.Next
        }
        pre = pre.Next
    }
    return dummy.Next
}
```

算法二：递归版

每个节点只会访问一次，因此时间复杂度为O(n+m)；只有到达l1或l2的底部，递归才会返回，所以空间复杂度为O(n+m)

```go
// 递归版
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
  if l1 == nil {return l2}
  if l2 == nil {return l1}
  var newHead *ListNode
  if l1.Val < l2.Val {
    newHead = l1
    newHead.Next = mergeTwoLists(l1.Next, l2)
  } else {
    newHead = l2
    newHead.Next = mergeTwoLists(l1, l2.Next)
  }
  return newHead
}
```



#### 23 合并K个排序链表

题目链接：https://leetcode.com/problems/merge-k-sorted-lists/submissions/

题目要求：给定K个有序链表，返回合并后的排序链表。

思路分析：复习递归算法，掌握自底向上和自上向下的递归方法。

算法1：递归算法，自底向上的两两合并。每个链表遍历一遍，所示时间复杂度O(n*k)，空间复杂度O(1)。

```go
// date 2020/01/06; 2020/02/21
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 { return nil }
    if len(lists) == 1 { return lists[0] }
    l1, l2 := mergeKLists(lists[:len(lists)-1]), lists[len(lists)-1]
    dummy := &ListNode{Val: -1}
    pre := dummy
    for l1 != nil || l2 != nil {
        if l1 == nil {
            pre.Next = l2
            break
        }
        if l2 == nil {
            pre.Next = l1
            break
        }
        if l1.Val < l2.Val {
            pre.Next = l1
            l1 = l1.Next
        } else {
            pre.Next = l2
            l2 = l2.Next
        }
        
        pre = pre.Next
    }
    return dummy.Next
}
```



#### 24 两两交换链表中的节点

思路分析：

算法1：先找出新的头节点，并保存两两交换后的第二个节点，用于更新其Next指针。

```go
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    var newHead *ListNode
    pre, p1, p2 := head, head, head.Next
    after := p2.Next
    p1.Next = after
    p2.Next = p1
    pre, p1 = p1, after
    newHead = p2
    for p1 != nil && p1.Next != nil {
        p2 = p1.Next
        after = p2.Next
        p1.Next = after
        p2.Next = p1
        pre.Next = p2
        pre, p1 = p1, after
    }
    return newHead
}
```

算法2：递归实现两两交换

```go
func swapPairs(head *ListNode) *ListNode {
  var nHead *ListNode
  // basic case
  if head == nil || head.Next == nil {
    return head
  } else {
    n := head.Next.Next
    nHead = head.Next
    nHead.Next = head
    head.Next = swapPairs(n)
  }
  return nHead
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
  // -1是因为当前节点也算一个
  pre.Next = head
  pre = head
  step := l - k - 1
  for step > 0 {
    pre = pre.Next
    step--
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

#### 92 反转链表II

题目要求：1<= m <= n <= len(LinkedList),一趟扫描完成反转。

思路分析：

1. 找到m节点及其前驱节点
2. 反转m->n的所有节点
3. 更新m节点的下一跳，更新头节点

```go
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil || head.Next == nil || m >= n {return head}
    var mpre, mnode, mtail *ListNode
    pre := head
    for m > 1 && pre != nil {
        mpre, pre = pre, pre.Next
        m--
        n--
    }
    if pre == nil {return head}
    mnode, mtail = pre, pre
    pre = pre.Next
    n--
    for n > 1 && pre != nil {
        after := pre.Next
        pre.Next = mtail
        mtail = pre
        pre = after
        n--
    }
    // update m node
    mnode.Next = pre.Next
    // update n node
    pre.Next = mtail
    // update head
    if mpre == nil {
        head = pre
    } else {
        mpre.Next = pre
    }
    return head
}
```

#### 141 Linked List Cycle【E】

题目链接：https://leetcode.com/problems/linked-list-cycle/

题目要求：给定一个单向链表，判断链表中是否存在环。

思路分析：通过快慢指针是否相遇来判断链表中是否存在环。

```go
// 算法1：快慢指针
// 时间复杂度分析
// 如果链表中不存在环，则fast指针最多移动N/2次；如果链表中存在环，则fast指针需要移动M次才能与slow指针相遇，M为环的元素个数。
// 总体来讲，时间复杂度为O(N)
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow, fast = slow.Next, fast.Next.Next
        if slow == fast {return true}
    }
    return false
}
```

#### 142 Linked List Cycle II【M】

题目链接：https://leetcode.com/problems/linked-list-cycle-ii/

题目要求：给定一个单向链表，判断链表是否存在环，如果存在，则返回环开始的节点，否则返回空。

思路分析：

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

#### 206 反转链表【E】

题目要求：https://leetcode.com/problems/reverse-linked-list/

思路分析：

直接迭代，时间复杂度O(n)，空间复杂度O(1)。

```go
// 算法1：迭代
func reverseList(head *ListNode) *ListNode {
    var pre, after *ListNode
    for head != nil {
        after = head.Next
        head.Next = pre
        pre, head = head, after
    }
    return pre
}
```

算法2：假设链表为n1->n2->...->nk-1->nk->nk+1->...nm；如果从节点nk+1到结点nm均已经反转，即

n1->n2->...->nk-1->nk->nk+1<-...<-nm那么nk的操作则需要nk.Next.Next = nk

时间复杂度O(n)，空间复杂度O(n)，因为递归会达到n层。

```go
// 算法2：递归
func reverseList(head *ListNode) *ListNode {
  if head == nil || head.Next == nil {return head}
  pre := reverseList(head.Next)
  head.Next.Next = head
  head.Next = nil
  return pre
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

