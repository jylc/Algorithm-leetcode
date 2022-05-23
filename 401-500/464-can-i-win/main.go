package main

import "fmt"

//https://leetcode.cn/problems/can-i-win/
//464. 我能赢吗
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	//等差数列求和公式，所有数字相加都不能达到目标值
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	//记忆化搜索
	memo := make(map[int]bool, 1<<maxChoosableInteger)
	//state表示数字的使用状态，curTotal为当前和
	var dfs func(int, int) bool
	dfs = func(state, curTotal int) (res bool) {
		if _, ok := memo[state]; !ok {
			for i := 0; i < maxChoosableInteger; i++ {
				//数字 i + 1 已经被使用，跳过
				if (state>>i)&1 == 1 {
					continue
				}
				//先手能赢
				if curTotal+i+1 >= desiredTotal {
					res = true
					break
				}
				//i + 1被选择后，先手对手会输
				if !dfs(state|(1<<i), curTotal+i+1) {
					res = true
					break
				}
			}
			//记忆化搜索
			memo[state] = res
		}
		return memo[state]
	}
	return dfs(0, 0)
}
func main() {
	maxChoosableInteger := 10
	desiredTotal := 300
	/*cnt := 0
	for i := 1; i <= 10; i++ {
		cnt += i
	}
	fmt.Println(cnt)*/
	ans := canIWin(maxChoosableInteger, desiredTotal)
	fmt.Println(ans)
}
