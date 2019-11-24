### String

#### 相关题目

- 14 Longest Common Prefix

#### 14 Longest Common Prefix【最长公共前缀】

思路分析：

算法1：最长公共前缀一定存在最短的那个字符串中，因此先求出最短的字符串的长度minLen，然后依次比较每个字符串中的第i个字符(0 <= i < minLen)，如果相同则追加到结果中，否则退出，返回之前的结果。【垂直扫描】

算法2：水平扫描法，n个字符串的最长公共前缀，一定是前n-1个字符串的最长公共前缀与第n个字符串的公共前缀，即公式LCP(S1...Sn) = LCP(LCP(LCP(S1, S2) S3)...Sn)

```go
// 算法1：直接比较，也称为垂直扫描
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {return ""}
    if len(strs) == 1 {return strs[0]}
    // find the minLen
    minLen := 1 << 31 -1
    for _, str := range strs {
        if len(str) < minLen {
            minLen = len(str)
        }
    }
    var res string
    for i:= 0; i < minLen; i++ {
        r := strs[0][i]
        ok := true
        for j := 1; j < len(strs); j++ {
            if r != strs[j][i] {
                ok = false
                break
            }
        }
        if !ok {
            break
        }
        res = res + string(r)
    }
    return res
}

// 算法2: 水平扫描，利用公式LCP(S1...Sn) = LCP(LCP(LCP(S1,S2), S3)...Sn)
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 { return "" }
    prefix := strs[0]
    for i := 1; i < len(strs); i++ {
        prefix = findPrefix(prefix, strs[i])
        if len(prefix) == 0 {return ""}
    }
    return prefix
}

func findPrefix(s1, s2 string) string {
    minLen := len(s1)
    if len(s2) < minLen {
        minLen = len(s2)
    }
    var res string
    for i := 0; i < minLen; i++ {
        if s1[i] != s2[i] {break}
        res += string(s1[i])
    }
    return res
}
```

