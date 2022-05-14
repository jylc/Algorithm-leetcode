package main

//https://leetcode.cn/problems/count-nodes-equal-to-average-of-subtree/
//2265. 统计值等于子树平均值的节点数

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) (ans int) {
	ans = 0
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		left, l := dfs(node.Left)
		right, r := dfs(node.Right)
		sum := left + right + node.Val
		cnt := l + r + 1
		if sum/cnt == node.Val {
			ans++
		}
		return sum, cnt
	}

	dfs(root)
	return
}
func main() {

}
