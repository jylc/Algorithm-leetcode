package main

import (
	"math"
)

//https://leetcode-cn.com/contest/weekly-contest-252/problems/three-divisors/
//5830. 三除数
func isThree(n int) bool {
	cnt := 0
	for i := 1; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			if n/i != i {
				cnt += 2
			} else {
				cnt += 1
			}
		}
	}
	if cnt==3{
		return true
	}else{
		return false
	}
}
func main() {

}
