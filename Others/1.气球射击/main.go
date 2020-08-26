package main

import (
	"fmt"
)

/*作者：Makki
链接：https://www.nowcoder.com/discuss/488619
来源：牛客网

小Q在进行射击气球的游戏，如果小Q在连续T枪中打爆了所有颜色的气球，将得到一只QQ公仔作为奖励。（每种颜色的气球至少被打爆一只）。
这个游戏中有m种不同颜色的气球，编号1到m。小Q一共有n发子弹，然后连续开了n枪。小Q想知道在这n枪中，打爆所有颜色的气球最少用了连续几枪？
输入描述：
第一行两个空格间隔的整数数n，m。n<=1000000 m<=2000
第二行一共n个空格间隔的整数，分别表示每一枪打中的气球的颜色,0表示没打中任何颜色的气球。
输出描述：
一个整数表示小Q打爆所有颜色气球用的最少枪数。如果小Q无法在这n枪打爆所有颜色的气球，则输出-1
示例
输入：
12 5
2 5 3 1 3 2 4 1 0 5 4 3
输出：
6

test
12 5
2 5 3 1 3 2 4 1 5 0 4 3
5*/

//滑动窗口求解
func minContinueShoot(n, m int, balls []int) int {
	visited := make(map[int]int)
	cnt, res := 0, n
	flag := false
	for j, i := 0, 0; j < n; j++ {
		_, ok := visited[balls[j]]
		if !ok && balls[j] != 0 {
			cnt++
		}
		visited[balls[j]]++
		if cnt == m {
			flag = true
			for balls[i] == 0 || visited[balls[i]] > 1 {
				visited[balls[i]]--
				i++
			}
			res = min(res, j-i+1)
		}
	}
	if flag {
		return res
	} else {
		return -1
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	result := minContinueShoot(12, 5, []int{2, 5, 3, 1, 3, 2, 4, 1, 0, 5, 4, 3})
	fmt.Println("result = ", result)
}
