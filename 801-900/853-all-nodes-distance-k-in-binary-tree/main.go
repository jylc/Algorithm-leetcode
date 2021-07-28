package main

//https://leetcode-cn.com/problems/all-nodes-distance-k-in-binary-tree/
//863. 二叉树中所有距离为 K 的结点
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

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parents := map[int]*TreeNode{}
	var findParents func(node *TreeNode)
	findParents = func(node *TreeNode) {
		if node.Left != nil {
			parents[node.Left.Val] = node
			findParents(node.Left)
		}
		if node.Right != nil {
			parents[node.Right.Val] = node
			findParents(node.Right)
		}
	}
	findParents(root)
	ans := make([]int, 0)
	var dfs func(node *TreeNode, from *TreeNode, depth int)
	dfs = func(node *TreeNode, from *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == k {
			ans = append(ans, node.Val)
		}
		if node.Left != from {
			dfs(node.Left, node, depth+1)
		}
		if node.Right != from {
			dfs(node.Right, node, depth+1)
		}

		if parents[node.Val] != from {
			dfs(parents[node.Val], node, depth+1)
		}
	}
	dfs(target, nil, 0)
	return ans
}
func main() {

}
