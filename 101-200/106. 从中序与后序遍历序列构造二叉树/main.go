package main

//106. 从中序与后序遍历序列构造二叉树
//https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
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

func buildTree(inorder []int, postorder []int) *TreeNode {
	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}

	var buildNode func(is, ie, ps, pe int) *TreeNode
	buildNode = func(is, ie, ps, pe int) *TreeNode {
		if ie < is || pe < ps {
			return nil
		}
		root := postorder[pe]
		ri := m[root]

		node := &TreeNode{
			Val:   root,
			Left:  buildNode(is, ri-1, ps, ps+ri-is-1),
			Right: buildNode(ri+1, ie, ps+ri-is, pe-1),
		}
		return node
	}
	return buildNode(0, len(inorder)-1, 0, len(postorder)-1)
}
func main() {
	inorder := []int{
		9, 3, 15, 20, 7,
	}
	postorder := []int{
		9, 15, 7, 20, 3,
	}

	buildTree(inorder, postorder)

}
