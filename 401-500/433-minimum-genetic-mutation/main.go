package main

import "fmt"

//https://leetcode-cn.com/problems/minimum-genetic-mutation/
//433. 最小基因变化
//单向由start到end寻找,直到变化后的字符换等于end结束。
func minMutation(start string, end string, bank []string) int {
	if start == end {
		return 0
	}

	bankSet := map[string]struct{}{}
	for _, s := range bank {
		bankSet[s] = struct{}{}
	}

	if _, ok := bankSet[end]; !ok {
		return -1
	}
	q := []string{start}
	for step := 0; q != nil; step++ {
		tmp := q
		q = nil
		for _, cur := range tmp {
			for i, x := range cur {
				//当前字符变为其他字符
				for _, y := range "ACGT" {
					if y != x {
						nxt := cur[:i] + string(y) + cur[i+1:]
						if _, ok := bankSet[nxt]; ok { //如果变换字符后的字符串在bank中，以此为基础继续搜索
							if nxt == end {
								return step + 1
							}
							delete(bankSet, nxt)
							q = append(q, nxt)
						}
					}
				}
			}
		}
	}
	return -1
}
func main() {
	start := "AAAAACCC"
	end := "AACCCCCC"
	bank := []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}
	ans := minMutation(start, end, bank)
	fmt.Println(ans)
}
