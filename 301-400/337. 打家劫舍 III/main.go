package main

//337. 打家劫舍 III
//https://leetcode-cn.com/problems/house-robber-iii/
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

func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) []int

	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}
		l,r:=dfs(node.Left),dfs(node.Right)
		selected:= node.Val + l[1] + r[1]
		noselected := max(l[0], l[1]) + max(r[0], r[1])
		return []int{selected,noselected}
	}

	values := dfs(root)
	return max(values[0], values[1])
}

func max(f int, g int) int {
	if f > g {
		return f
	} else {
		return g
	}
}

func main() {

}
