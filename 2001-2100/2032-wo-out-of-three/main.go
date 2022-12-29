package main

// https://leetcode.cn/problems/two-out-of-three/
// 2032. 至少在两个数组中出现的值

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m1, m2, m3 := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	ans := make([]int, 0)
	for _, i := range nums1 {
		if _, ok := m1[i]; !ok {
			m1[i] = true
		}
	}

	for _, i := range nums2 {
		if _, ok := m2[i]; !ok {
			m2[i] = true
		}
	}

	for _, i := range nums3 {
		if _, ok := m3[i]; !ok {
			m3[i] = true
		}
	}

	for i := 1; i <= 100; i++ {
		count := 0
		if m1[i] {
			count++
		}

		if m2[i] {
			count++
		}

		if m3[i] {
			count++
		}

		if count >= 2 {
			ans = append(ans, i)
		}
	}
	return ans
}
func main() {

}
