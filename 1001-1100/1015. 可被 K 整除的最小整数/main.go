package main

import "fmt"

//1015. 可被 K 整除的最小整数
//https://leetcode-cn.com/problems/smallest-integer-divisible-by-k/

func smallestRepunitDivByK(K int) int {
	if K%2 == 0 ||K%5==0{
		return -1
	}
	N := 1
	count := 1
	for {
		if N%K == 0 {
			return count
		}
		N=N%K
		N=N*10+1
		count++
	}
}

func main() {
	result := smallestRepunitDivByK(2)
	fmt.Println("result = ",result)
}
