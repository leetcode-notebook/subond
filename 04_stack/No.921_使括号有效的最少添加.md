## 921 使括号有效的最少添加-简单

题目：

只有满足下面几点之一，括号字符串才是有效的：

- 它是一个空字符串，或者
- 它可以被写成 `AB` （`A` 与 `B` 连接）, 其中 `A` 和 `B` 都是有效字符串，或者
- 它可以被写作 `(A)`，其中 `A` 是有效字符串。

给定一个括号字符串 `s` ，在每一次操作中，你都可以在字符串的任何位置插入一个括号

- 例如，如果 `s = "()))"` ，你可以插入一个开始括号为 `"(()))"` 或结束括号为 `"())))"` 。

返回 *为使结果字符串 `s` 有效而必须添加的最少括号数*。



> **示例 1：**
>
> ```
> 输入：s = "())"
> 输出：1
> ```
>
> **示例 2：**
>
> ```
> 输入：s = "((("
> 输出：3
> ```



分析：

这个问题跟第 20 题（有效括号）类似，第 20 题是判断一个字符串是否是有效的括号，通过对左括号入栈，右括号出栈的操作，判断最后栈是否为空来判断括号是否有效。

这个题目也是类似，能出栈的时候就出栈，不能出栈了，表示需要添加一个，直接计数返回即可。

```go
// date 2023/12/18
func minAddToMakeValid(s string) int {
    res := 0
    stack := make([]rune, 0, 16)

    for _, v := range s {
        if v == '(' {
            stack = append(stack, v)
        } else if v == ')' {
            if len(stack) > 0 {
                stack = stack[:len(stack)-1]
            } else {
                res++
            }
        }
    }
    if len(stack) > 0 {
        res += len(stack)
    }
    return res
}
```

