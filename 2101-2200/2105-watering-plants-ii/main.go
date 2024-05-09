package main

import "fmt"

// https://leetcode.cn/problems/watering-plants-ii/?envType=daily-question&envId=2024-05-09
// 2105. 给植物浇水 II

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	count := 0
	xA, xB := 0, len(plants)-1
	waterA, waterB := capacityA, capacityB
	for xA < xB {
		if waterA >= plants[xA] {
			waterA -= plants[xA]
		} else {
			waterA = capacityA - plants[xA]
			count++
		}

		if waterB >= plants[xB] {
			waterB -= plants[xB]
		} else {
			waterB = capacityB - plants[xB]
			count++
		}
		xA++
		xB--
	}

	if xA == xB {
		if waterA >= waterB && waterA < plants[xA] {
			count++
		}
		if waterA < waterB && waterB < plants[xB] {
			count++
		}
	}
	return count
}

func main() {
	plants1 := []int{2, 2, 3, 3}
	capacityA1 := 5
	capacityB1 := 5
	fmt.Println(minimumRefill(plants1, capacityA1, capacityB1))
	plants2 := []int{2, 2, 3, 3}
	capacityA2 := 3
	capacityB2 := 4
	fmt.Println(minimumRefill(plants2, capacityA2, capacityB2))
	plants3 := []int{5}
	capacityA3 := 10
	capacityB3 := 8
	fmt.Println(minimumRefill(plants3, capacityA3, capacityB3))
	plants4 := []int{1, 2, 4, 4, 5}
	capacityA4 := 6
	capacityB4 := 5
	fmt.Println(minimumRefill(plants4, capacityA4, capacityB4))
}
