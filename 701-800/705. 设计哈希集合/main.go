package main

//705. 设计哈希集合
//https://leetcode-cn.com/problems/design-hashset/

const base = 1000009

type MyHashSet struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{data: make([]int, base)}
}

func (this *MyHashSet) Add(key int) {
	this.data[key] = 1
}

func (this *MyHashSet) Remove(key int) {
	this.data[key] = 0
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if v := this.data[key]; v == 1 {
		return true
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

func main() {

}
