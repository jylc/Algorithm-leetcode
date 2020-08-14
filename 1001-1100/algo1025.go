package _001_1100

import "math"

func divisorGame(N int) bool {
	if N < 2 {
		return true
	}
	var tmp int = N
	var flag bool = true
	var i int

	for {
		//tmp为1时结束游戏
		if tmp == 1 {
			break
		}
		//A
		if flag {
			//最大因数
			i = maxYS(tmp)
			if tmp == i {
				return false
			}
			tmp -= i
			flag = false
		} else {
			i = maxYS(tmp)
			if tmp == i {
				return true
			}
			tmp -= i
			flag = true
		}
	}
	if flag {
		return false
	} else {
		return true
	}
}

func maxYS(num int) int {
	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return i
		}
	}
	return -1
}
