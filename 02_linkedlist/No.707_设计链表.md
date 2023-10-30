## 707 设计链表-中等

题目：

你可以选择使用单链表或者双链表，设计并实现自己的链表。

单链表中的节点应该具备两个属性：`val` 和 `next` 。`val` 是当前节点的值，`next` 是指向下一个节点的指针/引用。

如果是双向链表，则还需要属性 `prev` 以指示链表中的上一个节点。假设链表中的所有节点下标从 **0** 开始。

实现 `MyLinkedList` 类：

- `MyLinkedList()` 初始化 `MyLinkedList` 对象。
- `int get(int index)` 获取链表中下标为 `index` 的节点的值。如果下标无效，则返回 `-1` 。
- `void addAtHead(int val)` 将一个值为 `val` 的节点插入到链表中第一个元素之前。在插入完成后，新节点会成为链表的第一个节点。
- `void addAtTail(int val)` 将一个值为 `val` 的节点追加到链表中作为链表的最后一个元素。
- `void addAtIndex(int index, int val)` 将一个值为 `val` 的节点插入到链表中下标为 `index` 的节点之前。如果 `index` 等于链表的长度，那么该节点会被追加到链表的末尾。如果 `index` 比长度更大，该节点将 **不会插入** 到链表中。
- `void deleteAtIndex(int index)` 如果下标有效，则删除链表中下标为 `index` 的节点。



分析：

使用 双链表 方便在尾部增加，同时把双向链表的 head 和 tail 用哑结点 表示，这样可以减少不必要的判断。

size 表示链表中的元素个数，对 index 进行校验。

```go
// date 2023/10/17
type MyLinkedList struct {
    size int
    head, tail *MyNode
}

type MyNode struct {
    val int
    prev, next *MyNode
}


func Constructor() MyLinkedList {
    ll := MyLinkedList{
        size: -1,
        head: &MyNode{},
        tail: &MyNode{},
    }
    ll.head.next = ll.tail
    ll.tail.prev = ll.head
    return ll
}


func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index > this.size {
        return -1
    }
    pre := this.head
    for index >= 0 {
        pre = pre.next
        index--
    }
    return pre.val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    node := &MyNode{val, nil, nil}
    node.next = this.head.next
    this.head.next.prev = node

    this.head.next = node
    node.prev = this.head
    this.size++
}


func (this *MyLinkedList) AddAtTail(val int)  {
    node := &MyNode{val, nil, nil}
    this.tail.prev.next = node
    node.prev = this.tail.prev
    
    node.next = this.tail
    this.tail.prev = node
    this.size++
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 || index > this.size + 1 {
        return
    }
    if index == 0 {
        this.AddAtHead(val)
        return
    }
    if index == this.size+1 {
        this.AddAtTail(val)
        return
    }
    node := &MyNode{val, nil, nil}
    pre := this.head
    for index > 0 {
        pre = pre.next
        index--
    }
    after := pre.next

    pre.next = node
    after.prev = node

    node.prev = pre
    node.next = after
    this.size++
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index > this.size {
        return
    }
    pre := this.head
    for index > 0 {
        index--
        pre = pre.next
    }
    after := pre.next.next
    pre.next = after
    after.prev = pre
    this.size--
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

