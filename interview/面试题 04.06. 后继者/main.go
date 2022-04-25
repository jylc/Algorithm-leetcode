package main

//interview 04.06. 后继者
//https://leetcode-cn.com/problems/successor-lcci/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var (
		ans  *TreeNode
		flag bool
	)
	var search func(pre *TreeNode, p int)
	search = func(cur *TreeNode, p int) {
		if cur == nil {
			return
		}
		search(cur.Left, p)
		if cur.Val == p {
			flag = true
		} else if flag {
			ans = cur
			flag = false
		}
		search(cur.Right, p)
	}
	search(root, p.Val)
	return ans
}

func main() {

}
