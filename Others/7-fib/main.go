package main

import "fmt"

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	a, b := 1, 1
	for i := 0; i < n-2; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	n := 4
	ans := fib(n)
	fmt.Println(ans)
}
