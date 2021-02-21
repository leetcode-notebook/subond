## Quick Sort  \[快速排序\]

  * 1.[算法介绍](#算法介绍)
  * 2.[复杂度分析](#复杂度分析)
  * 3.[算法流程](#算法流程)
  * 4.[代码实现](#代码实现)

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

```
// 函数选取最后一个值作为基准值，然后将所有小于等于基准值的元素放在基准值的左边，将所有大于基准值的元素放在基准值的右边，
// 即在排序的序列中给基准值找到正确的位置

partition(arry[], low, high) {
  pi = arry[high];
  i = low - 1;
  for(j = low; j <= high - 1; j++) {
    if(arry[j] <= pi) {
      i++;
      swap arry[i] and arry[j];
    }
  }
  swap arry[i + 1] and arry[high]
  return i + 1;
}
```

### 代码实现

* C++

```cpp
// 第一种
void swap(int* a, int* b) {
  int t = *a;
  *a = *b;
  *b = t;
}

int partition(int arry[], int low, int high) {
  int p = arry[high];
  int i = low - 1;
  for(int j = low; j <= high - 1; j++) {
    if(arry[j] <= p) {
      i++;
      swap(arry[i], arry[j]);
    }
  }
  swap(arry[i + 1], arry[high]);
  return (i + 1);
}

void QuickSort(int arry[], int low, int high) {
  if(low < high) {
    int pi = partition(arry, low, high);
    QuickSort(arry, low, pi - 1);
    QuickSort(arry, pi + 1, high);
  }
}

// 第二种
int division(vector<int> &data, int low, int high) {
  int base = data[low];
  while(low < high) {
    while(low < high && data[high] >= base) high--;
    data[low] = data[high];
    while(low < high && data[low] <= base) low++;
    data[high] = data[low];
    data[low] = base;
  }
  return low;
}
void quicksort(vector<int> &data, int low, int high) {
  if(low < high) {
    int mid = division(data, low, high);
    quicksort(data, low, mid - 1);
    quicksort(data, mid + 1, high);
  }
}
```
* python

```python
def partition(arry, low, high):
    i = low - 1
    pivot = arry[high]
    for j in range(low, high):
        if arry[j] <= pivot:
            i = i + 1
            arry[i], arry[j] = arry[j], arry[i]

    arry[i+1], arry[high] = arry[high], arry[i+1]
    return (i+1)

def QuickSort(arry, low, high):
    if low < high:
        pi = partition(arry, low, high)
        QuickSort(arry, low, pi - 1)
        QuickSort(arry, pi + 1, high)
```
