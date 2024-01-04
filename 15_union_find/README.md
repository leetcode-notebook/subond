## 并查集Union-Find

### 基本概念

并查集被认为是最简洁而优雅的数据结构之一，主要用于解决一些元素分组问题。

它管理一系列不相交的集合，并支持两种操作：

- 合并（Union）：把两个不相交的集合合并成一个集合
- 查询（Find）：查询两个元素是否在同一个集合中



并查集的应用场景很多，其中最直接的一个场景是：亲戚问题。

> 题目：
>
> 如果某个家族人员过于庞大，要判断两个人是否有亲戚关系，不是很容易，现在给出某个亲戚关系图，求任意给出的两个人是否具有亲戚关系。
>
> 规定，如果x和y是亲戚，y和z是亲戚，那么x和z也是亲戚。
>
> 规定，如果x和y是亲戚，那么x的亲戚都是y的亲戚，y的亲戚也都是x的亲戚。



## 代码设计

通常并查集的数据结构中有两个变量，parent和count。

Parent 表示节点的关系；count 表示连通分量。

初始化时，每个节点指向自己，连通分量就等于节点总数。

合并操作，先判断两个节点的父结点是否相同，如果相同表示已经在同一个集合了，直接返回；否则将其中一个节点的父结点更新为另一个的父结点。

查询操作，如果节点的父结点不是自己，那么一定要继续查找直到找到根节点。

```go
type (
	MyUnionFind struct {
		parent []int // 存储每个节点的父结点
		count  int   // 存储连通分量
	}
)

// n 表示图中一共有多少个节点
func newMyUnionFind(n int) *MyUnionFind {
	u := &MyUnionFind{
		count:  n,
		parent: make([]int, n),
	}
	// 初始化时, 每个节点的父结点都是自己
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

// 将两个节点x y 合并
func (u *MyUnionFind) union(x, y int) {
	xp, yp := u.find(x), u.find(y)
	if xp == yp {
		return
	}
	u.parent[yp] = xp
	u.count--
}

// 查找 x 的父结点
func (u *MyUnionFind) find(x int) int {
  // 如果节点不是自身，表示有关系，继续查找，直到根节点
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

func (u *MyUnionFind) getCount() int {
	return u.count
}

```



最基本的并查集模板

```go
type (
	MyUnionFind struct {
		parent []int // 存储每个节点的父结点
	}
)

// n 表示图中一共有多少个节点
func newMyUnionFind(n int) *MyUnionFind {
	u := &MyUnionFind{
		parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

func (u *MyUnionFind) union(x, y int) {
	xp, yp := u.find(x), u.find(y)
	if xp == yp {
		return
	}
	u.parent[yp] = xp
}

func (u *MyUnionFind) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}
```

