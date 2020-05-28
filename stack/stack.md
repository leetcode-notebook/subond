## 栈

目录：

  * [栈的基本操作](#栈的基本操作)
  * [栈的实现](#栈的实现)
    * [C++](#C++)
    * [Python](#Python)
    * [栈的应用](#栈的应用)
  * [栈的反转](#栈的反转)
  * [栈的排序](#栈的排序)
  * [SpecialStack](#SpecialStack)
  * [相关题目](#相关题目)

### 栈的基本操作

  + push(x)：将元素x加入到栈中
  + pop()：弹出栈顶元素，但是，**不会返回弹出元素的值**
  + top()：取出栈顶元素
  + empty()：判断栈是否为空，当栈为空时，返回true
  + size()：返回栈中的元素个数

### 栈的实现

从时间复杂度角度而言，push(), pop(), top()，empty()操作只需要O(1)的时间。从实现角度而言，栈的的实现方式有两种，**数组和线性表**。

### C++

```cpp
#define MAX 1000
class Stack
{
private:
  int top;
public:
  int a[MAX];
  Stack() {top = -1;}
  bool push(int x);
  int pop();
  bool empty();
};
// 入栈
bool Stack::push(int x) {
  if(top >= MAX) {
    cout << "Stack Overflow";
    return false;
  } else {
    a[++top] = x;
    return true;
  }
}
// 出栈，一般情况下不带返回，直接删除元素；而自己写的函数可以允许删除的时候返回元素的值
int Stack::pop() {
  if(top < 0) {
    cout << "Stack Underflow";
    return 0;
  } else {
    int x = a[top--];
    return x;
  }
}
// 空判断
bool Stack::empty() {
  return (top < 0)
}
```

### Python

```python
class Stack:
    """模拟栈的基本操作"""
    # 初始化
    def __init__(self):
        self.item = []
    # 判空
    def isEmpty(self):
        return len(self.item) == 0
    # 压栈
    def push(self, data):
        self.item.append(data)
    # 出栈
    def pop(self):
        if not self.isEmpty():
            return self.item.pop()
    # 去栈顶元素
    def top(self):
        if not self.isEmpty():
            return self.item[len(self.item)-1]
    # 栈的大小
    def size(self):
        return len(self.item)
```

### 栈的应用

栈的主要应用包括：运算符的优先级，字符串反转。

### 栈的反转

利用迭代思想，实现栈的反转，只能使用push(),pop(),empty(),size()操作。

```
栈：                  栈：
4  ->  top            1  ->  top
3            反转后    2
2                     3
1                     4
```

分析：栈的标准push()是向栈顶压入新元素，如果将push()改为向栈底压入新元素，那么就可以利用empty(),pop()来递归实现栈的反转。

```cpp
stack<int> st;
void push_bottom(int x) {
  if(st.size() == 0)
    st.push(x);
  else {
    int temp = st.top();
    st.pop();
    push_bottom(x);
    st.push(temp);
  }
}
void Reverse() {
  if(st.size() > 0) {
    int x = st.top();
    st.pop();
    Reverse();
    push_bottom(x);
  }
}
```

* Python

```Python
def insertAtBottom(stack, item):
    if isEmpty(stack):
        push(stack, item)
    else:
        temp = pop(stack)
        insertAtBottom(stack, item)
        push(stack, temp)

# Below is the function that reverses the given stack
# using insertAtBottom()
def reverse(stack):
    if not isEmpty(stack):
        temp = pop(stack)
        reverse(stack)
        insertAtBottom(stack, temp)
```

### 栈的排序

注意：利用递归思想，只能使用栈的基本操作。

### 算法分析

```
// 排序
sortStack(stack S)
	if stack is not empty:
		temp = pop(S);  
		sortStack(S);
		sortedInsert(S, temp);
// 向已排序的栈中插入元素
sortedInsert(Stack S, element)
	if stack is empty OR element > top element
		push(S, elem)
	else
		temp = pop(S)
		sortedInsert(S, element)
		push(S, temp)
```

### 算法实现

* C++

```cpp
// 递归实现向已排序栈中插入元素
void SortedInsert(stack<int> &s, int x)
{
  if(s.empty() || x > s.top()) {
    s.push(x);
    return;
  }
  int temp = s.top();
  s.pop();
  SortedInsert(s, x);
  s.push(temp);
}
// 排序
void StackSort(stack<int> &s) {
  if(!s.empty()) {
    int temp = s.top();
    s.pop();
    StackSort(s);
    SortedInsert(s, temp);
  }
}
```

* Python

```Python
def SortedInsert(s, x):
    if s.isEmpty() or x > s.top():
        s.push(x)
        return;
    else:
        temp = s.top()
        s.pop()
        SortedInsert(s, x)
        s.push(temp)

def StackSort(s):
    if not s.isEmpty():
        temp = s.top()
        s.pop()
        StackSort(s)
        SortedInsert(s, temp)
```

### SpecialStack

实现一种特殊的栈结构，让其push(),pop(),empty(),getmin()的四种操作的时间复杂度为O(1)。

```
// 继承上面栈的实现部分，Stack类
class SpecialStack: public Stack
{
private:
  Stack min;
public:
  int pop();
  void push(int x);
  int getmin();
};
// 入栈
void SpecialStack::push(int x) {
  if(empty()) {
    Stack::push(x);
    min.push(x);
  } else {
    Stack::push(x);
    int y = min.pop();
    min.push(y);
    if(x < y)
      min.push(x);
    else
      min.push(y);
  }
}
// 出栈
int SpecialStack::pop() {
  int x = Stack::pop();
  min.pop();
  return x;
}
// getmin
int SpecialStack::getmin() {
  int x = min.pop();
  min.push(x);
  return x;
}
```

### 相关题目
- 155 最小栈
    
    
#### 155 最小栈

题目要求：实现栈的push, pop, top以及常数时间内的getMin()操作。
push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。

思路分析：考察辅助栈，有两种方案依次实现。

```go
// date 2020/03/23
// 第一种方案：辅助栈stackMin只保留当前栈中的最小值，即辅助栈和数据栈stackData长度不相等
type MinStack struct {
    stackData []int
    stackMin []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stackData: make([]int, 0),
        stackMin: make([]int, 0),
    }
}


func (this *MinStack) Push(x int)  {
    if len(this.stackMin) == 0 || this.stackMin[len(this.stackMin)-1] >= x {
        this.stackMin = append(this.stackMin, x)
    }
    this.stackData = append(this.stackData, x)
}


func (this *MinStack) Pop()  {
    if len(this.stackData) <= 0 {return}
    if len(this.stackMin) > 0 && this.stackMin[len(this.stackMin)-1] == this.stackData[len(this.stackData)-1] {
        this.stackMin = this.stackMin[:len(this.stackMin)-1]
    }
    this.stackData = this.stackData[:len(this.stackData)-1]
}


func (this *MinStack) Top() int {
    return this.stackData[len(this.stackData)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.stackMin) <= 0 { return 0 }
    return this.stackMin[len(this.stackMin)-1]
}
// 第二种方案：辅助栈stackMin随时保留当前栈的最小值，保持和数据栈stackData长度相等
type MinStack struct {
    stackData []int
    stackMin []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        stackData: make([]int, 0),
        stackMin: make([]int, 0),
    }
}


func (this *MinStack) Push(x int)  {
    if len(this.stackMin) == 0 || this.stackMin[len(this.stackMin)-1] >= x {
        this.stackMin = append(this.stackMin, x)
    } else {
        this.stackMin = append(this.stackMin, this.stackMin[len(this.stackMin)-1])
    }
    this.stackData = append(this.stackData, x)
}


func (this *MinStack) Pop()  {
    if len(this.stackData) <= 0 {return}
    this.stackData = this.stackData[:len(this.stackData)-1]
    this.stackMin = this.stackMin[:len(this.stackMin)-1]
}


func (this *MinStack) Top() int {
    return this.stackData[len(this.stackData)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.stackMin) <= 0 { return 0 }
    return this.stackMin[len(this.stackMin)-1]
}
```

#### 二叉搜索树的前序序列判断

问题描述

给定一个数组，判断其是否为一个二叉搜索树的前序遍历，并且要求时间复杂度为O(n)。

提示：左节点的值 小于 根节点的值 小于或等于 右节点的值，满足二叉搜索树的特征。

算法分析

这个问题与NGE问题类似，可以采用栈进行求解。算法如下：

```
1. 创建一个空栈
2. 初始化root = INT_MIN
3. 遍历每一个节点pre[i]
  a) If pre[i] is smaller than current root, return false.
  b) Keep removing elements from stack while pre[i] is greater than stack top.
     Make the last removed item as new root.
     At this point, pre[i] is greater than the removed root.
  c) push pre[i] to stack (All elements in stack are in decreasing order.
```

算法实现

```cpp
bool BSTCheck(int pre[], int n) {
  stack<int> s;
  int root = INT_MIN;
  for(int i = 0; i < n; i++) {
    if(pre[i] < root)
      return false;
    while(!s.empty() && s.top() < pre[i]) {
      // 进入到while(),则说明存在右子树
      root = s.top();
      s.pop();
    }
    // 压入左子树的根节点
    s.push(pre[i]);
  }
  return true;
}
```