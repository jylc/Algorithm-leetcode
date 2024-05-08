package main

import "fmt"

// https://leetcode.cn/problems/watering-plants/?envType=daily-question&envId=2024-05-08
// 2079. 给植物浇水

func wateringPlants(plants []int, capacity int) int {
	step := 0
	tmp := capacity
	for i := 0; i < len(plants); i++ {
		if capacity < plants[i] {
			step += 2*i + 1
			capacity = tmp - plants[i]
		} else {
			capacity -= plants[i]
			step++
		}
	}
	return step
}

func main() {
	plants1 := []int{2, 2, 3, 3}
	capacity1 := 5
	fmt.Println(wateringPlants(plants1, capacity1))

	plants2 := []int{1, 1, 1, 4, 2, 3}
	capacity2 := 4
	fmt.Println(wateringPlants(plants2, capacity2))

	plants3 := []int{7, 7, 7, 7, 7, 7, 7}
	capacity3 := 8
	fmt.Println(wateringPlants(plants3, capacity3))
}
