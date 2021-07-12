package main

import "fmt"

//剑指 Offer 05. 替换空格
//https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/

func replaceSpace(s string) string {
	var tmp []rune
	for _,v:=range s{
		if v!=32{
			tmp=append(tmp,v)
		}else{
			tmp=append(tmp,37,50,48)
		}
	}

	return string(tmp)
}

func main() {
	s := "We are happy."
	result := replaceSpace(s)
	fmt.Println("result = ",result)
}
