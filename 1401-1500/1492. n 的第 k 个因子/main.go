package main

import (
	"fmt"
)

//1492. n 的第 k 个因子
//https://leetcode-cn.com/problems/the-kth-factor-of-n/

func kthFactor(n int, k int) int {
	var arr []int
	for i:=1;i<n;i++{
		if n%i==0{
			arr=append(arr,i)
		}
	}
	arr=append(arr,n)
	if k>len(arr){
		return -1
	}
	return arr[k-1]
}

func main() {
	result := kthFactor(1000, 3)
	fmt.Println("result = ",result)
}
