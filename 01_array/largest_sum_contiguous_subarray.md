## 连续最大子序列和

[TOC]


### 问题定义

给定一个数组，输出其最大连续子序列和的大小。

例如：数组为{-2, -3, 4, -1, -2, 1, 5, -3}，输出为7（即4,-1,-2,1,5}。

### 算法分析

```cpp
Initialize:
    max_so_far = 0
    max_ending_here = 0

Loop for each element of the array
  (a) max_ending_here = max_ending_here + a[i]
  (b) if(max_ending_here < 0)
            max_ending_here = 0
  (c) if(max_so_far < max_ending_here)
            max_so_far = max_ending_here
return max_so_far
```

### 算法实现

```cpp
int MaxSubArraySum(int a[], int n) {
  int max_so_far = 0, max_ending_here = 0;
  int first = 0, last = 0;   // 起始点坐标
  for(int i = 0; i < n; i++) {
    max_ending_here = max_ending_here + a[i];
    if(max_ending_here < 0) {
      max_ending_here = 0;
      first = i + 1;
    } else if(max_ending_here > max_so_far) {
      max_so_far = max_ending_here;
      last = i;
    }
  }
  cout << "From " << first << " to " << last << endl;
  return max_so_far;
}
```

### 二维空间问题

Maximum sum rectangle in a 2D matrix，即二维矩阵中的最大和矩形

![二维矩阵中的最大和矩阵](http://on64c9tla.bkt.clouddn.com/Algorithm/2weijuzhensum.png)

如图所示，图中蓝色矩形框内就是最大的矩阵和。

### 算法分析

借助上面的算法原理，我们可以找到一行或者一列中的最大连续子序列和，以及子序列的起点和终点。

```cpp
Initialize:
    start = -1
    finish = -1
    sum = 0

for i = 0; i < n; i++
  1) sum += a[i]
  2) if(sum < 0) {
    sum = 0;start = i + 1;
  } else if(sum > maxsum) {
    maxsum = sum; finish = i;
  }
  3) if(finish != -1)
    return maxsum;
```

同样，对二维矩阵而言，子矩阵的和需要从 **左上角** 开始，计算到 **右下角** 元素，因此，子矩阵的和的累计过程即向下延伸，也向右延伸。所以，用temp[i]计算每列中自顶行至尾行的和，同时将列作为基本的移动单元。

时间复杂度为O(n^3)

### 算法实现

```
// 找到一列中的最大连续子序列的和，并返回起点和终点
int FindOne(int a[], int &start, int &finish, int n) {
  int sum = 0, res = 0;
  int stemp = -1;
  start = -1;
  finish = -1;
  for(int i = 0; i < n; i++) {
    sum += a[i];
    if(sum < 0) {
      sum = 0;
      stemp = i + 1;
    } else if(sum > res) {
      res = sum;
      start = stemp;
      finish = i;
    }
  }
  if(finish != -1)
    return res;
  // 特殊情况，即所有元素均小于零
  res = a[0];
  start = 0;
  finish = 0;
  for(int i = 0; i < n; i++) {
    if(a[i] > res) {
      res = a[i];
      start = finish = i;
    }
  }
  return res;
}
//
int FindTow(int a[R][C]) {
  int res, finalleft, finalright, finaltop, finaldown;
  int left, right, i;
  int temp[R], sum, start, finish;
  for(left = 0; left < C; left++) {
    for(right = left; right < C; right++) {
      for(i = 0; i < R; i++)
        temp[i] += a[i][right];
      sum = FindOne(temp, start, finish, R);
      if(sum > res) {
        res = sum;
        finalleft = left;
        finalright = right;
        finaltop = start;
        finaldown = finish;
      }
    }
  }
  printf("(Top, Left) (%d, %d)n", finalTop, finalLeft);
  printf("(Bottom, Right) (%d, %d)n", finalBottom, finalRight);
  printf("Max sum is: %dn", res);
}
```
