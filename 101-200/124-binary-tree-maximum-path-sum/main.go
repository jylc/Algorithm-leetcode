package main

import "math"

//https://leetcode.cn/problems/binary-tree-maximum-path-sum/
//124. 二叉树中的最大路径和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(0, dfs(node.Left))
		right := max(0, dfs(node.Right))
		ans = max(ans, left+right+node.Val)
		/*
			问：为什么返回max(left,right)+node.Val
			答：返回的是当前节点与到某个叶节点的路径的最大值，为当前节点的父节点提供某一边的最大值；
				而当前节点到两个叶节点的最大路径由max(ans, left+right+node.Val)计算
		*/
		return max(left, right) + node.Val
	}
	return dfs(root)
}
func main() {

}
