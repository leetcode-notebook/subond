## 5 最长回文字符串-中等

题目：

给你一个字符串 `s`，找到 `s` 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。



分析：

从中心到两边的双指针。

```go
// date 2023/12/09
func longestPalindrome(s string) string {
    n := len(s)
    res := ""
    for i:= 0; i < n; i++ {
        s1 := palindrome(s, i, i)
        s2 := palindrome(s, i, i+1)
        if len(s1) > len(res) {
            res = s1
        }
        if len(s2) > len(res) {
            res = s2
        }
    }
    return res
}

func palindrome(s string, l, r int) string {
    for l >= 0 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }
    return string(s[l+1:r])
}
```

