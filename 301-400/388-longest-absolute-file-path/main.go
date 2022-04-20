package main

import (
	"fmt"
)

//https://leetcode-cn.com/problems/longest-absolute-file-path/
//388. 文件的最长绝对路径
func lengthLongestPath(input string) int {
	length := len(input)
	hash := make([]int, length)
	ans := 0
	for i := 0; i < length; {
		level := 0
		// 判断层级
		for i < length && input[i] == '\t' {
			level++
			i++
		}
		j := i
		isDir := true
		// 计算长度与是否为目录
		for j < length && input[j] != '\n' {
			if input[j] == '.' {
				isDir = false
			}
			j++
		}
		cur := j - i
		prev := 0
		if level-1 >= 0 {
			prev = hash[level-1]
		} else {
			prev = -1
		}

		path := prev + 1 + cur
		// 自动从层级为0处开始计算；覆盖之前的
		if isDir {
			hash[level] = path
		} else if path > ans {
			ans = path
		}
		i = j + 1
	}
	return ans
}
func main() {
	//input1 := "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
	//input1 := "file1.txt\nfile2.txt\nlongfile.txt"
	//input1 := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
	//input1 := "a"
	//input1 := "a\n\tb\n\t\tc"
	input1 := "a\n\tb.txt\na2\n\tb2.txt"
	ans := lengthLongestPath(input1)
	fmt.Println(ans)
}
