package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "eW8="
	result, _ := base64.StdEncoding.DecodeString(s)
	fmt.Println(result)
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
	fmt.Println(string(ans))
}
