package main

//190. 颠倒二进制位
//https://leetcode-cn.com/problems/reverse-bits/

func reverseBits(num uint32) uint32 {
	var result uint32
	result = 0
	for i := 0; i < 32; i++ {
		tmp := num & 1
		num = num >> 1
		result = result << 1
		if tmp == 1 {
			result = result | 1
		}
	}
	return result
}
func main() {

}
