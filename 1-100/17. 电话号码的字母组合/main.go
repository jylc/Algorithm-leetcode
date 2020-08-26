package main

//17. 电话号码的字母组合
//https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var phoneMap map[string]string = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var combinations []string
	var dfs func(nums string, index int, combination string)
	dfs = func(nums string, index int, combination string) {
		if index == len(nums) {
			combinations = append(combinations, combination)
		} else {
			num := string(nums[index])
			letters := phoneMap[num]
			cnt := len(letters)
			for i := 0; i < cnt; i++ {
				dfs(nums, index+1, combination+string(letters[i]))
			}
		}
	}
	dfs(digits, 0, "")
	return combinations
}
func main() {

}
