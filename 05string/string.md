## String

[TOC]

### Golang中的字符串

golang语言中字符串中所包含的类型

```go
	str := "su"
	for i, s := range str {
		fmt.Printf("i %T, v %T, s[i] %T\n", i, s, str[i]) // int, int32, uint8
	}
	for _, v := range str {
		fmt.Printf("%T \n", v) // int32
	}
	for i := range str {
		fmt.Printf("%T \n", i) // int
	}
byte是uint8的别名, rune是int32的别名。
```

### 相关题目

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

#### 20 有效的字符串/有效的括号

思路分析：stack数据结构，压栈及出栈

算法如下：时间复杂度O(n) 空间复杂度O(1)

```go
// date 2020/01/06
/* 算法：
 0. 如果字符串长度为单数，直接返回false.
 1. 遇到左括号，入栈
 2. 遇到右括号，判断栈是否为空，如果为空，返回false
 3. 如果不为空，出栈
 4. 遍历结束后，判断栈是否为空。
*/
func isValid(s string) bool {
  if len(s) & 0x1 == 1 {return false}
  stack := make([]rune, 0)
  var c rune
  for _, v := range s {
    if v == '(' || v == '{' || v == '[' {
      stack = append(stack, v)
    } else {
      if len(stack) == 0 {return false}
      c = stack[len(stack)-1]
      if v == ')' && c == '(' || v == '}' && c == '{' || v == ']' && c == '[' {
        stack = stack[:len(stack)-1]
      } else {
        return false
      }
    }
  }
  return len(stack) == 0
}
```

#### 43 字符串相乘

题目要求：

思路分析：任务拆解，把大任务拆解成一个一个的小任务。

```go
// date 2020/03/28
/*
1. 先实现两个字符串相加；2.实现单个数字的相乘，3.实现零的追加。
*/
func multiply(num1 string, num2 string) string {
    if len(num1) == 0 { return num2 }
    if len(num2) == 0 { return num1 }
    if len(num1) == 1 && num1[0] == '0' { return "0" }
    if len(num2) == 1 && num2[0] == '0' { return "0" }
    var res string
    for i := len(num2)-1; i >= 0; i-- {
        t := product(num1, num2[i])
        t = appendZeros(t, len(num2)-1-i)
        res = addString(res, t)
    }
    return res
}

func appendZeros(s string, c int) string {
    if c <= 0 { return s }
    for c > 0 {
        s += fmt.Sprintf("%d", 0)
        c--
    }
    return s
}

func product(s1 string, s2 byte) string {
    var res string
    var carry, n1, t int
    i, n2 := len(s1)-1, int(s2 - '0')
    for i >= 0 || carry > 0 {
        n1 = 0
        if i >= 0 { n1 = int(s1[i]-'0') }
        t = n1 * n2 + carry
        carry = t / 10
        res = fmt.Sprintf("%d", t%10) + res
        i--
    }
    return res
}

func addString(s1, s2 string) string {
    var res string
    var carry, n1, n2, t int
    i, j := len(s1)-1, len(s2)-1
    for i >= 0 || j >= 0 || carry > 0 {
        n1, n2 = 0, 0
        if i >= 0 { n1 = int(s1[i] - '0') }
        if j >= 0 { n2 = int(s2[j] - '0') } 
        t = n1 + n2 + carry
        carry = t / 10
        res = fmt.Sprintf("%d", t%10) + res
        i--
        j--
    }
    return res
}
```



#### 有效的字母异位词

思路分析

算法1：题目中说明了只包含小写字母，因此可以用数组判断。

```go
func isAnagram(s string, t string) bool {
  if len(s) != len(t) {return false}
  ms, mt := [26]int{}, [26]int{}
  for i := 0; i < len(s); i++ {
    ms[s[i] - 'a']++
    mt[t[i] - 'a']++
  }
  return ms == mt
}
```

#### 实现strStr()

思路分析

算法1：找到起始点，并判断是否相同。

```go
func strStr(haystack, needle string) int {
  if len(needle) <= 0 {return 0}
  n1, n2 := len(haystack), len(needle)
  for i := 0; i < n1; i++ {
    if haystack[i] != needle[0] {
      continue
    }
    if i + n2 > n1 {
      return -1
    }
    if string(haystack[i:i+n]) == needle {return i}
  }
  return -1
}
```

#### 反转字符串

思路分析

算法1：双指针

```go
func reverseStrings(s []byte) {
  i, j := 0, len(s)-1
  for i < j {
    t := s[i]
    s[i] = s[j]
    s[j] = t
    i++
    j--
  }
}
```

算法2：递归

```go
func reverseStrings(s []byte) {
  // basic case
  if len(s) <= 1 {
      return
  } else {
      t := s[0]
      s[0] = s[len(s)-1]
      s[len(s)-1] = t
      reverseString(s[1:len(s)-1])
  }  
}
```

#### 151 翻转字符串里的单词

题目要求：

思路分析：

```go
// date 2020/03/29
/*
1. 从字符串尾部开始遍历，依次翻转每个单词
2. 遍历字符串的时候主要去掉多余空格，最后的结果也需要去掉多余的空格
*/
func reverseWords(s string) string {
  res, word_res := make([]byte, 0), make([]byte, 0)
  end := len(s)-1
  // 去掉尾部多余的空格
  for end >= 0 && s[end] == ' ' { end-- }
  for i := end; i >= 0; i-- {
    if s[i] != ' ' {
      word_res = append(word_res, s[i])
    } else {
      // 找到一个单词
      res = append(res, reverseWord(word_res)...)
      res = append(res, ' ')
      // 去掉中间多余的空格
      for i >= 0 && s[i] == ' ' { i-- }
      i++
      word_res = word_res[:0]
    }
  }
  if len(word_res) > 0 {
    res = append(res, reverseWord(word_res)...)
  }
  if len(res) > 0 && res[len(res)-1] == ' ' {
    res = res[:len(res)-1]
  }
  return string(res)
}

func reverseWord(word []byte) []byte {
  i, j := 0, len(word)-1
  var c byte
  for i < j {
    c = word[i]
    word[i] = word[j]
    word[j] = c
    i++
    j--
  }
  return word
}
```



#### 246 中心对称数

题目要求：https://leetcode-cn.com/problems/strobogrammatic-number/

```go
// date 2020/03/21
func isStrobogrammatic(num string) bool {
    m := map[byte]byte{
        '6': '9',
        '9': '6',
        '8': '8',
        '0': '0',
        '1': '1',
    }
    l, r, n := 0, len(num)-1, len(num)
    for l < r {
        if m[num[l]] == num[r] {
            l++
            r--
        } else {
            return false
        }
    }
    // 如果是奇数个元素
    if n & 0x1 == 1 && m[num[n/2]] != num[n/2] { return false }
    return true
}
```



#### 58 最后一个单词的长度

题目要求： 给定一个字符串，返回其最后一个单词的长度。

算法: 注意字符串末尾的空字符串。

```go
// date 2020/01/19
func lengthOfLastWord(s string) int {
  var res int
  for i := len(s) - 1; i >= 0; i-- {
    if res == 0 && s[i] == ' ' {continue}
    if s[i] == ' ' {return res}
    res++
  }
  return res
}
```

#### 67 二进制求和

题目要求：给定两个二进制字符串，返回其和，结果用字符串存储。

算法1：直接操作。

算法2：用0补齐较短的字符串。

```go
// date 2020/01/19
func addBinary(a, b string) string {
  var n int
  if len(a) < len(b) {
    n = len(b) - len(a)
  } else {
    n = len(a) - len(b)
  }
  var res string
  for n > 0 {
    res += "0"
    n--
  }
  
}
```



#### 验证回文串

思路分析

```go
// date 2019/12/28
// 
func isPalindrome(s string) bool {
  s = strings.ToLower(s)
  var res string
  for i := 0; i < len(s); i++ {
    if 'a' <= s[i] && s[i] <= 'z' || '0' <= s[i] && s[i] <= '9' {
      res += string(s[i])
    }
  }
  i, j := 0, len(res) - 1
  for i < j {
    if res[i] != res[j] {
      return false
    } else {
      i++
      j--
    }
  }
  return true
}
```

#### 345 反转字符串中的元音字母

```go
// date 2019/12/28
// 双指针
func reverseVowels(s string) {
  vs := make([byte]int)
  vs['a'] = 1
  vs['e'] = 1
  vs['i'] = 1
  vs['o'] = 1
  vs['u'] = 1
  vs['A'] = 1
  vs['E'] = 1
  vs['I'] = 1
  vs['O'] = 1
  vs['U'] = 1
  i, j := 0, len(s) - 1
  var temp byte
  for i < j {
    _, ok := vs[s[i]]
    if ok {
      _, ok2 := vs[s[j]]
      if ok2 {
        temp = s[j]
        s[j] = s[i]
        s[i] = temp
        i++
        j--
      } else {
        j--
      }
    } else {
      i++
    }
  }
}
```

#### 415 字符串相加

题目要求：https://leetcode-cn.com/problems/add-strings/

思路分析：思路并不难，关键是如何让自己的代码更优雅

```go
// date 2020/03/28
func addStrings(num1 string, num2 string) string {
    var res string
    var carry, n1, n2, t int
    i, j := len(num1)-1, len(num2)-1
    for i >= 0 || j >= 0 || carry > 0 {
        n1, n2 = 0, 0
        if i >= 0 { n1 = int(num1[i] - '0') }
        if j >= 0 { n2 = int(num2[j] - '0') } 
        t = n1 + n2 + carry
        carry = t / 10
        res = fmt.Sprintf("%d", t%10) + res
        i--
        j--
    }
    return res
}
```

#### 521 最长特殊序列I

题目要求：https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/

思路分析：题意难以理解

```go
// date 2020/03/08
func findLUSlength(a, b string) int {
  if a == b { return len(a) }
  if len(a) == len(b) || len(a) > len(b) { return len(a) }
  return len(b)
}
```

#### 576 字符串的排列

题目要求：https://leetcode-cn.com/problems/permutation-in-string/

思路分析：

```go
// date 2020/03/28
/*
1. 注意题目要求，字符串问题很多都可以使用数组，好处就是golang中数组可以直接比较的
2. 第一个字符串的排列之一是第二个字符串的子串，所以字符出现的顺序很重要
*/
func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) { return false }
    // 题目中只包含小写字母，所以使用数组更方便[26]int
    var s1map [26]int
    for _, c := range s1 { s1map[c - 'a']++ }
    for i := 0; i <= len(s2)-len(s1); i++ {
        var s2map [26]int
        for j := 0; j < len(s1); j++ {
            s2map[s2[i+j] - 'a']++
        }
        if s1map == s2map { return true }
    }
    return false
}
```



#### 594 最长和谐子序列

题目要求：https://leetcode-cn.com/problems/longest-harmonious-subsequence/

算法分析

```go
// date 2020/03/08
// 算法一：两层循环，超时
// 算法二：利用map,时间复杂度O(N),空间复杂度O(N)
func findLHS(nums []int) int {
  m := make(map[int]int, 0)
  for _, v := range nums {
    if _, ok := m[v]; ok {
      m[v]++
    } else {
      m[v] = 1
    }
  }
  res := 0
  for k, v1 := range m {
    if v2, ok := m[k+1]; ok && v1+v2 > res { res = v1+v2 }
  }
  return res
}
// 算法三，算法二的优化版
func findLHS(nums []int) int{
  res, m := 0, make(map[int]int, 0)
  for _, v := range nums {
    if _, ok := m[v]; ok {
      m[v]++
    } else {
      m[v] = 1
    }
    // cal
    if v1, ok := m[v-1]; ok && m[v]+v1 > res { res = v1+m[v] }
    if v1, ok := m[v+1]; ok && m[v]+v1 > res { res = v1+m[v] }
  }
  return res
}
```



#### 674 最长连续递增序列

题目要求：https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/

算法分析：

```go
// date 2020/03/08
func findLengthOfLCIS(nums []int) int {
  if len(nums) == 0 { return 0 }
  res, temp := 1, 1
  for i := 0; i < len(nums)-1; i++ {
    if nums[i+1] > nums[i] {
      temp++
    } else {
      temp = 1
    }
    if temp > res { res = temp }
  }
  return res
}
```

#### 字符串中第一个唯一的字符

```go
func firstUniqChar(s string) int {
  m := make(map[uint8]int, len(s))
  for i := 0; i < len(s); i++ {
    if _, ok := m[s[i]]; ok {
      m[s[i]] = -1
    } else {
      m[s[i]] = i
    }
  }
  for i := 0; i < len(s); i++ {
    if v, ok := m[s[i]]; ok && v != -1 {
      return v
    }
  }
  return -1
}
```

#### 翻转字符串里的单词

题目要求：给定一个字符串，返回其按单词翻转，并去除多余空格的结果。

算法：从字符串的尾部遍历，找到单个单词，并翻转放入其结果，注意空格。

```go
// date 2020/02/02
func reverseWords(s string) string {
  res, temp := make([]byte, 0), make([]byte, 0) // 存放结果和单个单词
  end := len(s)-1
  for end >= 0 && s[end] == ' ' { end-- }
  for i := end; i >= 0; i-- {
    if s[i] != ' ' {
      temp = append(temp, s[i])
    } else {
      // find the word
      res = append(res, reverseWord(temp)...)
      res = append(res, ' ')
      for i >= 0 && s[i] == ' ' { i-- }
      i++
      temp = make([]byte, 0)
    }
  }
  if 0 != len(temp) {
    res = append(res, reverseWord(temp)...)
  }
  if len(res) > 0 && res[len(res)-1] == ' ' {
    res = res[:len(res)-1]
  }
  return string(res[:len(res)])
}
func reverseWord(word []byte) []byte {
  s, e := 0, len(word)-1
  var t byte
  for s < e {
    t = word[s]
    word[s] = word[e]
    word[e] = t
    s++
    e--
  }
  return word
}
```

#### 820 单词的压缩编码

题目要求：https://leetcode-cn.com/problems/short-encoding-of-words/

思路分析：考察知识点，字典树

如果一个单词是另一个单词的后缀，则该单词不需要重新编码。

```go
// date 2020/03/28
func minimumLengthEncoding(words []string) int {
  // 将单词存入哈希表，方便删除
  set := make(map[string]string, len(words))
  for _, w := range words {
    set[w] = w
  }
  // 查找可能存在的单词后缀，并删除它，注意单词的索引从1开始
  for _, word := range words {
    for i := 1; i < len(word); i++ {
      key := string(word[i:])
      if _, ok := set[key]; ok {
        delete(set, key)
      }
    }
  }
  // 统计并返回结果
  var res int
  for word := range set {
    res += len(word)+1
  }
  return res
}
```

