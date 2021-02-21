## 基本排序算法

目录：

  * [冒泡排序](#冒泡排序)
  * [选择排序](#选择排序)
  * [插入排序](#插入排序)


### 冒泡排序

冒泡排序(Bubble Sort)的原理是这样的：每次比较两个相邻的两个元素，按照大小顺序进行交换，这样在一次的遍历中，可以将最大或最小的值交换值序列的最后一个。

```
// 冒泡排序算法
void swap(int *a, int *b) {
  int temp = *a;
  *a = *b;
  *b = temp;
}
void BubbleSort(int array[], int n) {
  for(int i = 0; i < n - 1; i++) {
    for(int j = 0; j < n - i - 1; j++) {
      if(array[j] > array[j + 1])
        swap(&array[j], &array[j+1]);
    }
  }
}
```

### 选择排序

选择排序的思想是：每次选择未排序序列中最值，然后放到已排序序列的尾部。

```
// 选择排序算法
void swap(int *a, int *b) {
  int temp = *a;
  *a = *b;
  *b = temp;
}
void SelectionSort(int array[], int n) {
  if(n == 1)
    return;
  int i, j, min_index;
  for(i = 0; i < n - 1; i++) {
    min_index = i;
    for(j = i + 1; j < n; j++) {
      if(array[j] < array[min_index])
        min_index = j;
    }
    swap(&array[min_index], &array[i]);
  }
}
```

### 插入排序

插入排序的思想是：依次选择未排序序列中的元素，并将其插入到已排序序列中的合适位置。

```
Loop from i = 1 to n-1.
  a) Pick element arr[i] and insert it into sorted sequence arr[0…i-1]
```

```
// 插入排序
void InsertSort(int array[], int n) {
  int i, j, temp;
  for(i = 1; i < n; i++) {
    temp = array[i];
    j = i - 1;
    while(j >= 0 && array[j] > temp) {
      array[j+1] = array[j];
      j--;
    }
    array[j+1] = temp;
  }
}
```
