package main

//https://leetcode-cn.com/problems/goat-latin/
//824. 山羊拉丁文
func toGoatLatin(sentence string) string {
	ans := make([]uint8, 0)
	i := 0
	n := len(sentence)
	index := 1
	for i < n {
		head := sentence[i]
		flag := false
		switch head {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			ans = append(ans, head)
		default:
			flag = true
		}
		j := i + 1
		for j < n && sentence[j] != ' ' {
			ans = append(ans, sentence[j])
			j++
		}
		j++
		if flag {
			ans = append(ans, head)
		}
		ans = append(ans, 'm', 'a')
		for k := 0; k < index; k++ {
			ans = append(ans, 'a')
		}
		if j < n {
			ans = append(ans, ' ')
		}
		index++
		i = j
	}
	return string(ans)
}
func main() {
	sentence := "The quick brown fox jumped over the lazy dog"
	ans := toGoatLatin(sentence)
	print(ans)
}
