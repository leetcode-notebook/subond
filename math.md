### Math

#### 相关题目

- 8 

- 9 Palindrome Number

#### 8 String to Integer

https://leetcode.com/problems/string-to-integer-atoi/

```go
/* 细节题: 需要各种异常场景
1. 忽略开头的空白字符
2. 数字的符号
3. 溢出
4. 无效输入
*/
func myAtoi(str string) int {
    myMAX := 1 << 31 - 1
    myMIN := -1 << 31
    sign := 1
    var res int
    i := 0
    n := len(str)
    for i < n && str[i] == ' ' {i++}
    if i == n {return res}
    if str[i] == '+' {
        i++
    } else if str[i] == '-' {
        i++
        sign = -1
    }
    for ; i < n; i++ {
        if str[i] - '0' < 0 || str[i] - '0' > 9 {break}
        t := int(str[i] - '0')
        if res > myMAX / 10 || res == myMAX / 10 && t > myMAX % 10 {
            if sign == 1 {
                return myMAX
            } else if sign == -1 {
                return myMIN
            }
        }
        res = res * 10 + t
    }
    return res * sign
}
```



#### 9 Palindrome Number

```go
// 算法1：直接比较数字的每一位，更通用的算法
func isPalindrome(x int) bool {
    if x < 0 || (x != 0 && x % 10 == 0) {return false}
    d := 1
    for x / d >= 10 {
        d *= 10
    }
    for x > 0 {
        // 最高位
        q := x / d
        // 最低位
        p := x % 10
        if q != p { return false }
        // 移动至次高位，同时去掉最后一位
        x = x % d / 10
        d /= 100
    }
    return true
}
// 算法2：将数字从中间分开，判断左边和右边的数字是否相等
func isPalindrome(x int) bool {
    if x < 0 || (x != 0 && x % 10 == 0) {return false}
    left := 0
    for x > left {
        left = left * 10  + x % 10
        x /= 10
    }
    return x == left || x == left/10
}
```

思路学习：算法2对于判断链表同样适用。参考题目234https://leetcode.com/problems/palindrome-linked-list/