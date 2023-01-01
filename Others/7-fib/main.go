package main

import (
	"fmt"
)

func fib(n int64) int64 {
	if n < 2 {
		return n
	}
	p, q, r := int64(0), int64(0), int64(1)
	for i := int64(2); i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}

func main() {
	for {
		var num int64
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		ans := fib(num)
		fmt.Printf("%v\n", ans)
	}
}
