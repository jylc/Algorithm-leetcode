package main

//117. 填充每个节点的下一个右侧节点指针 II
//https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/

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
	var nodeArr []*Node
	nodeArr = append(nodeArr, root)

	for len(nodeArr) > 0 {
		tmpArr := nodeArr
		nodeArr = nil
		for i, node := range tmpArr {
			if i+1 < len(tmpArr) {
				node.Next = tmpArr[i+1]
			}
			if node.Left != nil {
				nodeArr = append(nodeArr, node.Left)
			}
			if node.Right != nil {
				nodeArr = append(nodeArr, node.Right)
			}
		}
	}
	return root
}
func main() {

}
