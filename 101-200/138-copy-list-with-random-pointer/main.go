package main

// Node https://leetcode-cn.com/problems/copy-list-with-random-pointer/
//138. 复制带随机指针的链表
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
type RandomNode *Node

//用map
func copyRandomList(head *Node) *Node {
	m := make(map[*Node]RandomNode)
	cur := head
	for cur != nil {
		m[cur] = &Node{Val: cur.Val, Next: cur.Next}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]
}

func main() {

}
