package main

import "fmt"

// StockSpanner https://leetcode-cn.com/problems/online-stock-span/
//901. 股票价格跨度
type stack struct {
	nums []int
}

func (s *stack) pop() int {
	tmp := s.nums[len(s.nums)-1]
	s.nums = s.nums[:len(s.nums)-1]
	return tmp
}

func (s *stack) push(num int) {
	s.nums = append(s.nums, num)
}

func (s stack) peek() int {
	return s.nums[len(s.nums)-1]
}

func (s stack) isEmpty() bool {
	if len(s.nums) == 0 {
		return true
	}
	return false
}

type StockSpanner struct {
	prices  stack
	weights stack
}

func Constructor() StockSpanner {
	return StockSpanner{
		prices:  stack{nums: make([]int, 0)},
		weights: stack{nums: make([]int, 0)},
	}

}

func (this *StockSpanner) Next(price int) int {
	w := 1
	for !this.prices.isEmpty() && this.prices.peek() <= price {
		this.prices.pop()
		w += this.weights.pop()
	}
	this.prices.push(price)
	this.weights.push(w)
	return w
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
func main() {
	call := []string{"StockSpanner", "next", "next", "next", "next", "next", "next", "next", "next", "next", "next"}
	prices := []int{0, 28, 14, 28, 35, 46, 53, 66, 80, 87, 88}
	var stockSpanner StockSpanner
	for i, c := range call {
		if c == "StockSpanner" {
			stockSpanner = Constructor()
		}
		if c == "next" {
			ans := stockSpanner.Next(prices[i])
			fmt.Println(ans)
		}
	}
}
