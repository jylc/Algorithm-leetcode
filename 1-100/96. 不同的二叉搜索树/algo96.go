package _6__不同的二叉搜索树

//96. 不同的二叉搜索树
//https://leetcode-cn.com/problems/unique-binary-search-trees/
func numTrees(n int) int {
	g:=make([]int,n+1)
	g[0]=1
	g[1]=1
	for i:=2;i<=n;i++{
		for j:=1;j<=i;j++{
			g[i]+=g[j-1]*g[i-j]
		}
	}
	return g[n]
}
