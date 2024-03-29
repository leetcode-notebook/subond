## 165 比较版本好-中等

题目：

给你两个版本号 `version1` 和 `version2` ，请你比较它们。

版本号由一个或多个修订号组成，各修订号由一个 `'.'` 连接。每个修订号由 **多位数字** 组成，可能包含 **前导零** 。每个版本号至少包含一个字符。修订号从左到右编号，下标从 0 开始，最左边的修订号下标为 0 ，下一个修订号下标为 1 ，以此类推。例如，`2.5.33` 和 `0.1` 都是有效的版本号。

比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较 **忽略任何前导零后的整数值** 。也就是说，修订号 `1` 和修订号 `001` **相等** 。如果版本号没有指定某个下标处的修订号，则该修订号视为 `0` 。例如，版本 `1.0` 小于版本 `1.1` ，因为它们下标为 `0` 的修订号相同，而下标为 `1` 的修订号分别为 `0` 和 `1` ，`0 < 1` 。

返回规则如下：

- 如果 `*version1* > *version2*` 返回 `1`，
- 如果 `*version1* < *version2*` 返回 `-1`，
- 除此之外返回 `0`。



分析：

双指针，通过将版本号转成数字的方式进行比较。

这样做的好处就是变相地把前导零去掉了。

```go
// date 2023/12/10
func compareVersion(version1 string, version2 string) int {
    n1, n2 := len(version1), len(version2)
    i, j := 0, 0
    for i < n1 || j < n2 {
        v1 := 0
        for i < n1 && version1[i] != '.' {
            v1 = v1 * 10 + int(version1[i] - '0')
            i++
        }
        i++
        v2 := 0
        for j < n2 && version2[j] != '.' {
            v2 = v2 * 10 + int(version2[j] - '0')
            j++
        }
        j++
        if v1 > v2 {
            return 1
        } else if v1 < v2 {
            return -1
        }
    }
    return 0
}
```

