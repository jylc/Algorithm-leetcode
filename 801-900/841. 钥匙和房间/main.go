package main

import "fmt"

//841. 钥匙和房间
//https://leetcode-cn.com/problems/keys-and-rooms/
func canVisitAllRooms(rooms [][]int) bool {
	arr := make(map[int][]int)
	visited := make([]bool, len(rooms))
	for i, room := range rooms {
		for _, v := range room {
			arr[i] = append(arr[i], v)
		}
		visited[i] = false
	}
	fmt.Println("arr = ", arr)
	var dfs func(curr int)
	dfs = func(curr int) {
		for {
			if v, ok := arr[curr]; !ok || len(v) == 0 {
				break
			}
			tmp := arr[curr][0]
			arr[curr] = arr[curr][1:]
			dfs(tmp)
		}
		visited[curr] = true
	}
	dfs(0)
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
func main() {
	arr := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	result := canVisitAllRooms(arr)
	fmt.Println("result = ", result)
}
