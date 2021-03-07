## 单向链表的基本操作

  * 1.[链表的定义](#链表的定义)
  * 2.[链表的基本操作](#链表的基本操作)
    * 2.1[建立链表](#链表建立)
    * 2.2[插入元素](#插入元素)
    * 2.3[删除元素](#删除元素)
    * 2.4[反转链表](#反转链表)
    * 2.5[链表长度](#链表长度)

### 链表的定义

链表节点的定义如下：

```
struct Node {
  int val;
  Node *next;
};
```

### 链表的基本操作

链表的基本操作，包括建立，删除，插入，反转，求取长度。

### 链表建立

单向链表的建立，包括两种建立方式，一种是每次都将新元素放在整个链表的最前面；另一种是每次都将新元素放在整个链表的最后面。

```
/*方式一：每次均将新元素放在整个链表的最前面*/
Node *InsertFront(Node *phead, int value) {
  Node *s;
  s = (Node*)new(Node);
  s->val = value;
  s->next = phead;
  phead = s;
  return phead;
}

/*方式二：每次均将新元素放在整个链表的最后面*/
Node* InsertBack(Node *phead, int value) {
  Node *s;
  s = (Node*)new(Node);
  s->val = value;
  s->next = NULL;
  if(phead == NULL) {
    phead = s;
    return phead;
  } else {
    Node *p = phead;
    while(p->next != NULL) {
      p = p->next;
    }
    p->next = s;
    return phead;
  }
}
```
### 插入元素

链表的插入，包括两种：在某个 **位置前插入** 和 **位置后插入** 两种方式。

```
/*
功能：链表的插入
参数：链表头指针，匹配元素，新元素
说明：当匹配元素存在时，将元素插入到位置的后面；否则插入到整个链表的后面
第一步，建立新节点；第二步，查找匹配元素并插入新元素，无需保存前驱节点
 */
Node *InsertAfter(Node *phead, int value1, int value2) {
  Node *s, *p;
  int flag = 0;
  s = (Node*)new(Node);
  s->val = value2;
  p = phead;
  while(p->next != NULL) {
    if(p->val == value1) {
      Node *ptemp = p->next;
      p->next = s;
      s->next = ptemp;
      flag = 1;
      break;
    } else {
      p = p->next;
    }
  }
  if(!flag) {
    s->next = NULL;
    p->next = s;
  }
  return phead;
}
/*
功能：链表的插入
参数：表链表头指针，匹配元素，新元素
说明：当位置存在时，将元素插入到位置的前面；否则，插入到整个链表的后面
第一步，建立新节点；第二步，查找匹配元素并插入新元素，无需保存前驱节点
 */
Node *InsertBefore(Node *phead, int value1, int value2) {
  Node *s;
  s = (Node*)new(Node);
  s->val = value2;
  int flag = 0;
  if(phead->val == value1) {
    s->next = phead;
    phead = s;
    return phead;
  }
  Node *p = phead;
  while(p->next != NULL) {
    if(p->next->val == value1) {
      Node *ptemp = p->next;
      p->next = s;
      s->next = ptemp;
      flag = 1;
      break;
    } else {
      p = p->next;
    }
  }
  if(!flag) {
    s->next = NULL;
    p->next = s;
  }
  return phead;
}
```
### 删除元素

```
/*
功能：链表的删除
参数：表头，元素
说明：删除某个元素，若无此元素则不进行任何操作
 */
Node *Delete(Node *phead, int target) {
  if(phead == NULL) return phead;
  Node *pNode = phead;
  if(pNode->val == target) {
    Node *ptemp = pNode->next;
    pNode->next = NULL;
    delete pNode;
    phead = ptemp;
    return phead;
  }
  Node *p = phead->next;
  Node *pPre = phead;
  while(p != NULL) {
    if(p->val == target) {
      break;
    } else {
      pPre = p;
      p = p->next;
    }
  }
  pPre->next = p->next;
  p->next = NULL;
  delete p;
  return phead;
}
```
### 反转链表

```
/*
功能：反转链表
参数：链表的头指针，返回链表的头指针
 */
Node *Reverse(Node *phead) {
  Node *pReverseHead = NULL;
  Node *pNode = phead;
  Node *pPre = NULL;
  while(pNode != NULL) {
    Node *pNext = pNode->next;
    pNode->next = pPre;
    pPre = pNode;
    pNode = pNext;
  }
  pReverseHead = pPre;
  return pReverseHead;
}
```
### 链表长度

```
/*
功能：计算链表长度
参数：输入表头，输出无符号整型
 */
unsigned Length(Node *phead) {
  Node *p = phead;
  unsigned len = 0;
  while(p != NULL) {
    len++;
    p = p->next;
  }
  return len;
}
```
