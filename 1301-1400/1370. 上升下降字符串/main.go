package main

import "sort"

//1370. 上升下降字符串
//https://leetcode-cn.com/problems/increasing-decreasing-string/
func sortString(s string) string {
	//直接创建‘a’到‘z’的数组，记录每个字母出现的次数
	cnt := ['z' + 1]int{}
	for _, ch := range s {
		cnt[ch]++
	}

	n := len(s)
	ans := make([]byte, 0, n)
	for len(ans) < n {
		for i := byte('a'); i < 'z'+1; i++ {
			if cnt[i] > 0 {
				ans = append(ans, i)
				cnt[i]--
			}
		}

		for i := byte('z'); i >= 'a'; i-- {
			if cnt[i] > 0 {
				ans = append(ans, i)
				cnt[i]--
			}
		}
	}
	return string(ans)
}

func sortString2(s string) string {
	//利用map记录每个字母出现的次数将问题复杂化了
	m := make(map[byte]int)
	var tmp []byte
	for _, v := range s {
		k := byte(v)
		if _, ok := m[k]; !ok {
			tmp = append(tmp, k)
		}
		m[k]++
	}

	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i] < tmp[j] {
			return true
		}
		return false
	})

	n := len(s)
	ans := make([]byte, 0, n)
	for len(ans) < n {
		for _, k := range tmp {
			if v, _ := m[k]; v != 0 { //注意：当m[k]==0时，ok不为nil!!!
				ans = append(ans, k)
				m[k]--
			}
		}

		for i := len(tmp) - 1; i >= 0; i-- {
			if v, _ := m[tmp[i]]; v != 0 {
				ans = append(ans, tmp[i])
				m[tmp[i]]--
			}
		}
	}

	return string(ans)
}
func main() {

}
