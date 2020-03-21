## Queue

#### 基本定义

队列实现的是一种先进先出（FIFO）策略的线性表。

队列有队头(head)和队尾(tail)，当有一个元素入队时，放入队尾；出队时，即删除队头元素。

#### 相关题目

- 146 LRU缓冲机制
- 621 Task Scheduler【任务调度器】
- 622 Design Circular Queue 
- 641 Design Circular Deque【M】
- 933 Number of Recent Calls

#### 621 Task Scheduler【任务调度器】

题目描述：

给定一个用字符数组表示的 CPU 需要执行的任务列表。其中包含使用大写的 A - Z 字母表示的26 种不同种类的任务。任务可以以任意顺序执行，并且每个任务都可以在 1 个单位时间内执行完。CPU 在任何一个单位时间内都可以执行一个任务，或者在待命状态。

然而，两个相同种类的任务之间必须有长度为 n 的冷却时间，因此至少有连续 n 个单位时间内 CPU 在执行不同的任务，或者在待命状态。

你需要计算完成所有任务所需要的最短时间。

示例1：

```
输入: tasks = ["A","A","A","B","B","B"], n = 2
输出: 8
执行顺序: A -> B -> (待命) -> A -> B -> (待命) -> A -> B.
```

```go
// 算法1：优先安排出现次数最多的任务
func leastInterval(tasks []byte, n int) int {
    m := make([]int, 26)
    for _, t := range tasks {
        m[t - 'A']++
    }
    sort.Slice(m, func(i, j int) bool {
        return m[i] < m[j]
    })
    var res int
    for m[25] > 0 {
        i := 0
        for i <= n {
            if m[25] == 0 {
                break
            }
            if i < 16 && m[25-i] > 0 {
                m[25-i]--
            }
            res++
            i++
        }
        sort.Slice(m, func(i, j int) bool {
            return m[i] < m[j]
        })
    }
    return res
}
// 算法二：设计，利用空闲时间[leetcode-cn]
func leastInterval(tasks []byte, n int) int {
    m := make([]int, 26)
    for _, t := range tasks {
        m[t - 'A']++
    }
    sort.Slice(m, func(i, j int) bool {
        return m[i] < m[j]
    })
    max_val := m[25] - 1
    idle := n * max_val
    for i := 24; i >= 0 && m[i] > 0; i-- {
        if m[i] <= max_val {
            idle -= m[i]
        } else {
            idle -= max_val
        }
    }
    if idle > 0 {
        return idle + len(tasks)
    }
    return len(tasks)
}
```

#### 933 Number of Recent Calls

滑动窗口

```go
type RecentCounter struct {
    count []int
}


func Constructor() RecentCounter {
    return RecentCounter{
        count: make([]int, 0),
    }
}


func (this *RecentCounter) Ping(t int) int {
    this.count = append(this.count, t)
    for index, v := range this.count {
        if v >= t - 3000 {
            this.count = this.count[index:]
            return len(this.count)
        }
    }
    return len(this.count)
}
```



#### 622 Design Circular Queue【设计循环双端队列】【M】

```go
type RecentCounter struct {
    count []int
}


func Constructor() RecentCounter {
    return RecentCounter{
        count: make([]int, 0),
    }
}


func (this *RecentCounter) Ping(t int) int {
    this.count = append(this.count, t)
    var res int
    for i, v := range this.count {
        if v >= t - 3000 {
            res = i
            break
        }
    }
    this.count = this.count[res:]
    return len(this.count)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
```

#### 641 Design Circular Deque

```go
type MyCircularDeque struct {
    data []int
    front, rear, size int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
    return MyCircularDeque{
        data: make([]int, k),
        front: -1,
        rear: -1,
        size: k,
    }
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
    if this.IsFull() {
        return false
    }
    // insert empty deque
    if this.front == -1 {
        this.front = 0
        this.rear = 0
        this.data[this.front] = value
        return true
    }
    if this.front == 0 {
        this.front = this.size - 1
        this.data[this.front] = value
        return true
    }
    this.front -= 1
    this.data[this.front] = value
    return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
    if this.IsFull() {
        return false
    }
    // insert empty deque
    if this.front == -1 {
        this.front = 0
        this.rear = 0
        this.data[this.rear] = value
        return true
    }
    this.rear = (this.rear + 1) % this.size
    this.data[this.rear] = value
    return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
    if this.IsEmpty() {
        return false
    }
    // delete the last element
    if this.front == this.rear {
        this.front = -1
        this.rear = -1
        return true
    }
    this.front = (this.front + 1) % this.size
    return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
    if this.IsEmpty() {
        return false
    }
    // delete the last element
    if this.front == this.rear {
        this.front = -1
        this.rear = -1
        return true
    }
    if this.rear == 0 {
        this.rear = this.size - 1
        return true
    }
    this.rear -= 1
    return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[this.front]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[this.rear]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
    return this.front == -1
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
    return this.front == (this.rear + 1) % this.size
}
```

