## 349 两个数组的交集-中等

题目：

给定两个数组 `nums1` 和 `nums2` ，返回 *它们的交集* 。输出结果中的每个元素一定是 **唯一** 的。我们可以 **不考虑输出结果的顺序** 。



分析：

哈希表存储，如果存在第一个数组中，值为1，如果存在第二个数组，值为2。

最后遍历哈希表即可。

```go
// date 2023/12/05
func intersection(nums1 []int, nums2 []int) []int {
    set := make(map[int]int, 32)
    for _, v := range nums1 {
        _, ok := set[v]
        if !ok {
            set[v] = 1
        }
    }
    for _, v := range nums2 {
        _, ok := set[v]
        if ok {
            set[v] = 2
        }
    }
    res := make([]int, 0, 4)
    for k, v := range set {
        if v == 2 {
            res = append(res, k)
        }
    }
    return res
}
```

思考：如果是已经排序的数组，如何求交集。