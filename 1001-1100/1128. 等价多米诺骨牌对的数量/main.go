package main

//1128. 等价多米诺骨牌对的数量
//https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs/

func numEquivDominoPairs(dominoes [][]int) int {
	cnt := [100]int{}
	res := 0
	for _, dominoe := range dominoes {
		if dominoe[0] > dominoe[1] {
			dominoe[1], dominoe[0] = dominoe[0], dominoe[1]
		}
		tmp := dominoe[0]*10 + dominoe[1]
		res += cnt[tmp]
		cnt[tmp]++
	}
	return res
}

func main() {
}
