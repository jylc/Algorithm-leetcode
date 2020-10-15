package main

//116. 填充每个节点的下一个右侧节点指针
//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		tmp := queue
		queue = nil

		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
func main() {

}
