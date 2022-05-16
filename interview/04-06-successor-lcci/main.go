package main

//https://leetcode.cn/problems/successor-lcci/
//面试题 04.06. 后继者

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) (ans *TreeNode) {
	flag := false
	var search func(node *TreeNode)
	search = func(node *TreeNode) {
		if node == nil {
			return
		}
		search(node.Left)
		if node.Val == p.Val {
			flag = true
		} else if flag {
			p = node
		}
		search(node.Right)
	}
	search(root)
	return
}
func main() {

}
