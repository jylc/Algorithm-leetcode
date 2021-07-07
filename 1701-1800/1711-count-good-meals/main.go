package main

//1711. 大餐计数
//https://leetcode-cn.com/problems/count-good-meals/
func countPairs(deliciousness []int) int {
	maxValue := deliciousness[0]
	for _, v := range deliciousness[1:] {
		if v > maxValue {
			maxValue = v
		}
	}

	maxSum := maxValue * 2
	cnt := make(map[int]int)
	ans := 0
	for _, v := range deliciousness {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			//哈希表中已有元素下标一定小于当前元素的下标
			ans += cnt[sum-v]
		}
		cnt[v]++
	}
	return ans % (1e9 + 7)
}

func main() {

}
