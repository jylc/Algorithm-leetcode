package main

import "fmt"

// TreeNode https://leetcode.cn/problems/serialize-and-deserialize-bst/
//449. 序列化和反序列化二叉搜索树
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
type Codec struct {
	arrs []rune
}

func Constructor() Codec {
	return Codec{
		arrs: make([]rune, 0),
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var dfs func(root *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		this.arrs = append(this.arrs, rune(node.Val))
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return string(this.arrs)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	this.arrs = []rune(data)

	var insert func(l, r int) *TreeNode
	insert = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		j := l + 1
		for j <= r && this.arrs[j] <= this.arrs[l] {
			j++
		}
		node := &TreeNode{
			Val:   int(this.arrs[l]),
			Left:  insert(l+1, j-1),
			Right: insert(j, r),
		}
		return node
	}
	return insert(0, len(this.arrs)-1)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */

func create(tree []int) (root *TreeNode) {
	n := len(tree)
	if n == 0 {
		return nil
	}
	root = &TreeNode{Val: tree[0], Left: nil, Right: nil}
	var insert func(prev, node *TreeNode, val int)
	insert = func(prev, node *TreeNode, val int) {
		if node == nil {
			node = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
			if val <= prev.Val {
				prev.Left = node
			} else {
				prev.Right = node
			}
			return
		}

		if val <= node.Val {
			insert(node, node.Left, val)
		} else if val > node.Val {
			insert(node, node.Right, val)
		}
	}
	for i := 1; i < n; i++ {
		tmp := root
		insert(tmp, tmp, tree[i])
	}
	return
}

func main() {
	var root *TreeNode
	arrs := []int{12, 5, 2, 9, 18, 15, 17, 16, 19}
	root = create(arrs)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		fmt.Println(node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	//dfs(root)
	ser := Constructor()
	deser := Constructor()
	tree := ser.serialize(root)
	ans := deser.deserialize(tree)
	dfs(ans)
}
