package main

import "sort"

//981. 基于时间的键值存储
//https://leetcode-cn.com/problems/time-based-key-value-store/

type pair struct {
	value     string
	timestamp int
}

type TimeMap struct {
	m map[string][]pair
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{map[string][]pair{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], pair{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	p := this.m[key]
	idx := sort.Search(len(p), func(i int) bool {
		return p[i].timestamp > timestamp
	})
	if idx > 0 {
		return p[idx].value
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {

}
