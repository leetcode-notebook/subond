## 116 填充每个节点的下一个右侧节点-中等




```go
// date 2023/10/23
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
	queue := make([]*Node, 0, 16)
    queue = append(queue, root)

    for len(queue) != 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            cur := queue[i]
            if i + 1 < n {
                cur.Next = queue[i+1]
            }
            if cur.Left != nil {
                queue = append(queue, cur.Left)
            }
            if cur.Right != nil {
                queue = append(queue, cur.Right)
            }
        }
        queue = queue[n:]
    }
    return root
}
```
