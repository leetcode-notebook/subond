package template

type (
	// UnionFind define
	// 路径压缩 + 秩优化
	UnionFind struct {
		parent []int // 存储每个节点的父结点
		rank   []int //
		count  int   // 存储连通分量
	}
)

// NewUnionFind define
func NewUnionFind(n int) *UnionFind {
	u := &UnionFind{
		count:  n,
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	// 初始化时, 每个节点的根节点都是自身
	for i := 0; i < n; i++ {
		u.parent[i] = i
	}
	return u
}

// Union define
// merge x, y to the single set
func (u *UnionFind) Union(x, y int) {
	xp, yp := u.Find(x), u.Find(y)
	if xp == yp {
		return
	}
	// depth of tree_x < depth of tree_y
	if u.rank[xp] < u.rank[yp] {
		u.parent[xp] = yp
	} else {
		u.parent[yp] = xp
		// 如果深度相同，那么新的根节点秩加1
		if u.rank[xp] == u.rank[yp] {
			u.rank[xp]++
		}
	}
	u.count--
}

// Find define
// search the root of x
func (u *UnionFind) Find(x int) int {
	root := x
	for root != u.parent[root] {
		root = u.parent[root]
	}
	//compress path
	for x != u.parent[x] {
		temp := u.parent[x]
		u.parent[x] = root
		x = temp
	}
	return root
}

// GetCount define
// the total counts
func (u *UnionFind) GetCount() int {
	return u.count
}
