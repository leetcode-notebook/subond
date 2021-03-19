## Quick Sort  \[快速排序\]

  * 1.[算法介绍](#算法介绍)
  * 2.[复杂度分析](#复杂度分析)
  * 3.[算法流程](#算法流程)

### 算法介绍

快速排序与归并排序均属于分治思想算法(Divide and Conquer algorithm)，即选取一个基准值(pivot)，然后将其他给定值置于基准值的左右，从而实现排序过程。基准值的选取可以有多种方式：

1. 选取第一个元素作为基准值
2. 选取最后一个元素作为基准值
3. 随机选取一个元素作为基准值
4. 选取中间值作为基准值

快速排序算法的关键是partition()函数，其实现的目标是在给定的序列和基准值x的情况下，将基准值x放在正确的位置上，所有小于x的元素放在x的左边，所有大于等于x的元素放在x的右边。

### 复杂度分析

通常情况下，快速排序的时间复杂度可用如下公式表示：

$T(n) = T(k) + T(n-k-1) + \theta(n)$

其中，前两项表示递归调用，第三项表示分区过程(即，partition)。k代表序列中小于等于基准的元素个数。因此，快速排序的时间复杂度与 **待排序列** 和 **基准值的选取** 有关。下面分三种情况分别讨论：

* 最差的情况

  当分区过程总是选取最大或最小的元素作为基准值时，则发生最差的情况。此时，待排序列实际上已经是有序序列(递增，或递减)。因此，上述时间复杂度公式变为：

  $T(n) = T(0) + T(n-1) + \theta(n)$，即时间复杂度为$\theta(n^2)$

* 最好的情况

  当分区过程每次都能选出中间值作为基准值时，则发生最好的情况，此时，时间复杂度公式为：

  $T(n) = 2T(n/2) + \theta(n)$，即时间复杂度为 $\theta(nLogn)$

* 一般情况

  其时间复杂度为 $\theta(nLogn)$

### 算法流程

递归形式的快速排序的伪代码如下：

```
// low-起始索引；high-结束索引
quicksort(arry[], low, high) {
  if(low < high) {
    // pi是partition()产生的基准值的索引
    pi = partition(arry, low, high);
    quicksort(arry, low, pi - 1);
    quicksort(arry, pi + 1, high);
  }
}
```

partition()函数实现序列的划分，其逻辑是总是选取最左边的元素作为基准值x(也可以选取其他元素最为基准值)，然后追踪小于基准值x的索引i；遍历整个序列，若当前元素小于基准值x，则交换当前值值与i所在的值，否则忽略当前值。其伪代码如下：

```golang
// 函数选取最后一个值作为基准值
// 将所有小于等于基准值的元素放在基准值的左边，将所有大于基准值的元素放在基准值的右边
// index表示元素应该被放在的位置,for循环之后区间[left, index]均小于基准值,区间[index+1, right]均大于基准值,再交换一次swap(index+1, right)
// 则index+1为基准值,区间[index+2, right]为大于等于基准值
// 最后返回基准值的位置index+1
partition(nums []int, left, right int) {
  p := nums[right]
  index := left - 1
  for left <= right-1 {
    if nums[left] <= p {
      index++
      swap(nums, index, left)
    }
    left++
  }
  swap(nums, index+1, rigth)
  return index+1
}
```

```golang
// 另一种实现方式
func quickSort(nums []int, left, right int) {
  p := division(nums, left, right)
  quickSort(nums, left, p-1)
  quickSort(nums, p+1, right)
}

func division(nums []int, left, right int) int {
  base := nums[left]
  for left < right {
    for left < right && nums[right] >= base { right-- }
    nums[left] = nums[right]
    for left < right && nums[left] <= base { left++ }
    nums[right] = nums[left]
    nums[left] = base
  }
  return left
}
```
