## 栈


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
