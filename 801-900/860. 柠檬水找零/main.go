package main

//860. 柠檬水找零
//https://leetcode-cn.com/problems/lemonade-change/
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			} else {
				five--
				ten++
			}
		} else if bill == 20 {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
func main() {

}
