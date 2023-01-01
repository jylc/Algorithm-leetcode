package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	br := bufio.NewReader(os.Stdin)
	input, _, _ := br.ReadLine()
	result, _ := base64.StdEncoding.DecodeString(string(input))
	ans := make([]byte, 0)
	for {
		tmp := make([]byte, 0)
		reminder := 0
		for i := 0; i < len(result); i++ {
			t := int(result[i]) + reminder*256
			tmp = append(tmp, byte(t/36))
			reminder = t % 36
		}
		tmp = append(tmp, byte(reminder))
		if tmp[len(tmp)-1] < 10 {
			ans = append(ans, tmp[len(tmp)-1]+'0')
		} else {
			ans = append(ans, tmp[len(tmp)-1]-10+'A')
		}
		result = tmp[:len(tmp)-1]
		for len(result) != 0 {
			if result[0] == 0 {
				result = result[1:]
			} else {
				break
			}
		}

		if len(result) == 0 {
			break
		}
	}
	out := make([]byte, len(ans))
	for i := len(ans) - 1; i >= 0; i-- {
		out = append(out, ans[i])
	}
	fmt.Println(string(out))
}
