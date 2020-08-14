package main

//100. 相同的树
//https://leetcode-cn.com/problems/same-tree/
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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var dfs func(p *TreeNode, q *TreeNode) bool
	dfs = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		} else if p != nil && q != nil {
			l := dfs(p.Left, q.Left)
			r := dfs(p.Right, q.Right)
			if p.Val == q.Val && l && r {
				return true
			} else {
				return false
			}
		}
		return false
	}
	return dfs(p, q)
}
func main() {

}
