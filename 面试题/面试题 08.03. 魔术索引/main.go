package main


//面试题 08.03. 魔术索引
//https://leetcode-cn.com/problems/magic-index-lcci/

func findMagicIndex(nums []int) int {
	for i:=0;i<len(nums);i++{
		if i==nums[i]{
			return i
		}
	}
	return -1
}

func main() {
	
}
