## 关于单链表的其他问题 | set 1

[TOC]

### 判断一个链表是否存在环

解法：利用 **快慢指针** 遍历链表，当快慢两个指针在某一时刻指向同一个节点时，则存在环，否则不存在环。

```
bool detectloop(struct Node *phead) {
  struct Node *slow = phead, *fast = phead;
  while(slow && fast && fast->next) {
    slow = slow->next;
    fast = fast->next->next;
    if(slow == fast)
      return true;
  }
  return false;
}
```

### 判断一个元素是否存在链表中

```
// 迭代实现
bool search(Node *phead, int x) {
  Node *p = phead;
  while(p != NULL) {
    if(p->val == x)
      return true;
    p = p->next;
  }
  return false;
}
// 递归实现
bool search(Node *phead, int x) {
  if(phead == NULL)
    return false;
  if(phead->val == x)
    return true;
  return search(phead->next; x);
}
```

### 合并两个有序链表

注意保存新链表的头节点。

```
// 递归实现方式
Node *SortMerge(Node *phead1, Node *phead2) {
  Node *result = NULL;
  if(phead1 == NULL)
    return phead2;
  else if(phead2 == NULL)
    return phead1;

  if(phead1>val <= phead2->val) {
    result = phead1;
    result->next = SortMerge(phead1->next, phead2);
  } else {
    result = phead2;
    result->next = SortMerge(phead1, phead2->next);
  }
  return result;
}
```

### 在不给出链表的头指针条件下，删除链表中的某个节点

解法：将准备删除的节点的下一个节点的值拷贝至将要删除的节点，后删除下一个节点。需要考虑，这个节点是否为尾节点。

```
void delete(Node *node) {
  //尾节点
  if(node->next == NULL) {
    node = NULL;
  } else {
  //非尾节点
  Node *temp = node->next;
  node->val = temp->val;
  node->next = temp->next;
  free(temp);
  }
}
```

### 删除一个已排序链表中的重复元素

解法：比较当前节点与下一个节点的值，如果相同，则将当前节点指向下一个节点的下一个节点，并删除下一个节点。

```
void removeduplicates(Node *phead) {
  if(phead == NULL || phead->next == NULL)
    return phead;
  Node *current = phead, *next_next = NULL;
  while(current->next != NULL) {
    if(current->val == current->next->val) {
      next_next=current->next->next;
      free(current->next);
      current->next= next_next;
    } else {
      current = current->next;
    }
  }
}
```

### 删除一个未排序链表中重复元素

解法：借助辅助空间set，检查元素的是否出现过。

```
void removeDuplicates(Node *phead) {
  unorder_set<int> seen;
  Node *current = phead;
  Node *pre = NULL;
  while(current != NULL) {
    if(seen.find(current->val) != seen.end()) {
      pre->next = current->next;
      free(current);
    } else {
      seen.insert(current->val);
      pre = current;
    }
    current = pre->next;
  }
}
```

### 成对交换给定链表的元素

例如：1->2->3->4->5->6 变为 2->1->4->3->6->5。若最后一个元素落单，则无需改动最后一个元素。

解法：1）不改变原来的数据结构，只是交换节点值；2）改变原来的数据结构，直接交换链表节点。

```
// Pairwise swap elements of a given linked list
// 解法1：交换节点值
Node* Pairwise(Node *phead) {
  if(phead == NULL && phead->next == NULL)
    return phead;
  Node *pre = phead, *after = phead->next;
  while(pre && after) {
    int temp = pre->val;
    pre->val = after->val;
    after->val = temp;
    pre = after->next;
    if(pre != NULL)
      after = pre->next;
  }
  return phead;
}
// 解法2：交换节点
Node* Pairwise2(Node **phead) {
  if((*phead) == NULL && (*phead)->next == NULL)
    return (*phead);
  Node *pre = (*phead), *cur = (*phead)->next;
  *phead = cur;
  while(true) {
    Node *next = cur->next;
    cur->next = pre;
    if(next == NULL || next->next == NULL) {
      pre->next = next;
      break;
    }
    pre->next = next->next;
    pre = next;
    cur = pre->next;
  }
}
```

### 逆序打印链表中的元素

```
// 递归实现
void Print(Node *phead) {
  if(phead == NULL)
    return NULL;
  Print(phead->next);
  printf("%d ", phead->val);
}
```

### 将链表中的最后一个元素移动至链表的最前面

```
Node *Last2Head(Node *phead) {
  if(phead == NULL || phead->next == NULL)
    return phead;
  Node *last = phead->next, *pre = phead;
  while(pre && last && last->next != NULL) {
    last = last->next;
    pre = pre->next;
  }
  pre->next = NULL;
  last->next = phead;
  phead = last;
  return phead;
}
```

### 输出两个有序链表的公共部分

例如：1-2->3->4->6和2->4->6->8，输出2->4->6。

```
// 迭代实现
Node* Intersection(Node *phead1, Node *phead2) {
  if(phead1 == NULL || phead2 == NULL)
    return NULL;
  Node *p1 = phead1, *p2 = phead2;
  Node *pNew = NULL, *p =NULL;
  int flag = 1;
  while(p1 && p2) {
    if(p1->val == p2->val) {
      if(flag) {
        pNew = p1;
        p = pNew;
        flag = 0;
      } else {
        p->next = p1;
        p = p->next;
      }
      p1 = p1->next;
      p2 = p2->next;
    } else if(p1->val < p2->val) {
      p1 = p1->next;
    } else {
      p2 = p2->next;
    }
  }
  p->next = NULL;
  return pNew;
}
// 递归实现
Node *Intersection(Node *phead1, Node *phead2) {
  if(phead1 == NULL || phead2 == NULL)
    return NULL;
  if(phead1->val < phead2->val)
    return Intersection(phead1->next, phead2);
  if(phead1->val > phead2->val)
    return Intersection(phead1, phead2->next);
  Node *result = (struct Node*)malloc(sizeof(struct Node));
  result->val = phead1->val;
  result->next = Intersection(phead1->next, phead2->next);
  return result;
}
```

### 删除偶数位置上的节点

```
Node* DelteAlternate(Node *phead) {
  if(phead == NULL || phead->next == NULL)
    return phead;
  Node *pnew = phead;
  while(pnew && pnew->next) {
    Node *pnext = pnew->next->next;
    free(pnew->next);
    pnew->next = pnext;
    pnew = pnew->next;
  }
  return phead;
}
```

### 反转固定个数的节点

例如：1->2->3->4->5->6->7->8->NULL;当k=3时，反转后为3->2->1->6->5->4->8->7->NULL。

```
Node* Reverse(Node *phead, int k) {
  Node *p = phead;
  Node *pnext = NULL;
  Node *pre = NULL;
  int count = 0;
  while(p != NULL && count < k) {
    pnext = p->next;
    p->next = pre;
    pre = p;
    p = pnext;
    count++;
  }
  if(pnext != NULL)
    phead->next = Reverse(next, k);
  return pre;
}
```

### 交替反转链表中固定个数的节点

输入：1->2->3->4->5->6->7->8 k=3
输出：3->2->1->4->5->6->8->7

```
Node *KReverse(Node *phead, int k) {
  Node *p = phead;
  Node *pnext = NULL;
  Node *pre = NULL;
  count = 0;
  while(p != NULL && count < k) {
    pnext = p->next;
    p->next = pre;
    pre = p;
    p = pnext;
    count++;
  }
  if(phead != NULL)
    phead->next = p;
  count = 0;
  while(count < (k - 1) && p != NULL) {
    p = p->next;
    count++;
  }
  if(p != NULL)
    p->next = KReverse(p->next, k);
  return pre;
}
```

### 单调递减链表

解释：Delete nodes which have a greater value on right side

保证原有链表节点相对位置不变的情况下，使之变为单调递减链表。

输入：12->15->10->11->5->6->2->3
输出：15->11->6->3

如果当前节点的值比其下一个节点的值小，则删除当前节点，保留其下一个节点。

解法：先将链表反转，这样程序的逻辑就变为如果后一个节点小于前一个节点，则删除后一个节点。这样做的目的是 **确定新链表的头部**。原来的程序逻辑不太容易确定链表的头部，但是尾部相对容易确定，

```
// 这个函数对反转后的链表进行处理
Node DeleteSmallRight(Node *phead) {
  Node *pcur = phead;
  Node *maxnode = phead;
  Node *temp = NULL;
  while(pcur != NULL && pcur->next != NULL) {
    if(pcur->next->val < maxnode->val) {
      temp = pcur->next;
      pcur->next = temp->next;
      free(temp);
    } else {
      pcur = pcur->next;
      maxnode = pcur;
    }
  }
}
```
