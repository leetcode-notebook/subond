## Merge Sort \[归并排序\]

  * 1.[算法介绍](#算法介绍)
  * 2.[代码实现](#代码实现)
  * 3.[复杂度分析](#复杂度分析)
  * 4.[归并算法的应用](#归并算法的应用)
    *

### 算法介绍

与快速排序一样，归并排序也是采用分治思想。将待排序列分成两个子序列(直到不可再分)，然后对两个子序列进行有序合并。merge()函数是归并排序的关键。

![归并排序过程演示](http://www.geeksforgeeks.org/wp-content/uploads/gq/2013/03/Merge-Sort.png)

### 代码实现

根据归并算法的原理，很容易想到利用递归的方式进行实现，如下所示：

* C++

```cpp
// 将两个有序序列合并成一个序列
// 需要使用辅助空间
void merge(vector<int> &data, int first, int mid, int last) {
  int i = first, j = mid + 1;
  vector<int> temp;
  while(i <= mid && j <= last) {
    if(data[i] <= data[j])
      temp.push_back(data[i++]);
    else
      temp.push_back(data[j++]);
  }
  while(i <= mid)
    temp.push_back(data[i++]);
  while(j <= last)
    temp.push_back(data[j++]);
  for(int k = 0; k < temp.size(); k++)
    data[first + k] = temp[k];
  return;
}
void mergesort(vector<int> &data, int first, int last) {
  if(first < last) {
    int mid = first + (last - first) / 2;
    mergesort(data, first, mid);
    mergesort(data, mid + 1, last);
    mergedata(data, first, mid, last);
  }
  return;
}
```

* python

```python
def merge(left,right):
    result=[]
    i,j=0,0
    while i<len(left) and j<len(right):
        if left[i]<=right[j]:
            result.append(left[i])
            i+=1
        else:
            result.append(right[j])
            j+=1
    result+=left[i:]
    result+=right[j:]
    return result

def mergesort(arr):
    if len(arr)<=1:
        return arr
    mid=int(len(arr)/2)
    left=mergesort(arr[:mid])
    right=mergesort(arr[mid:])
    return merge(left,right)
```

### 复杂度分析

归并排序属于递归算法，其时间复杂度可用如下公式表示：

$T(n) = 2T(n/2) + \theta(n)$

因为归并排序总是将序列分成两个子序列，直到不可分解，然后在 **线性时间** 内进行合并。所以三种情况(最好，最坏，一般)的时间复杂度均是 $\theta(nLogn)$ 。

从稳定性而言，归并排序比快速排序稳定。

辅助空间需要O(n)。

### 归并算法的应用

* 线性表的归并排序

  对于线性表而言，其与数组的主要区别在于并不需要申请内存空间，可以在O(1)的时间内和O(1)的空间内实现元素的插入。因此，对线性表进行归并排序，合并数据的过程不需要额外的辅助空间。
