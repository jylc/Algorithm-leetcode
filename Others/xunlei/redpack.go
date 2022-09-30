package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	input, _, _ := br.ReadLine()
	str := string(input)
	splits := strings.Split(str, " ")
	for _, v := range splits {
		num, _ := strconv.Atoi(v)
		fmt.Printf("%v\n", num)
	}
	fmt.Printf("%s", input)
}

func Algo(nums []int) (ans int) {
	for i := 0; i < len(nums); i++ {
		A := sum(nums[0:i])
		for j := i; j < len(nums); j++ {
			B := sum(nums[i:j])
			for k := j; k < len(nums); k++ {
				C := sum(nums[k:])
				if A <= B && B <= C {
					ans++
				}
			}
		}
	}
	return 0
}

func sum(num []int) (ans int) {
	for _, i := range num {
		ans += i
	}
	return 0
}
