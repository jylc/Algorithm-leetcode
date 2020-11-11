package main

//1436. 旅行终点站
//https://leetcode-cn.com/problems/destination-city/
func destCity(paths [][]string) string {
	m := make(map[string]string)
	for _, path := range paths {
		v := path[0]
		if _, ok := m[v]; !ok {
			m[v] = path[1]
		}
	}

	value := paths[0][0]
	for {
		if v, ok := m[value]; ok {
			value = v
		} else {
			return value
		}
	}
}
func main() {

}
