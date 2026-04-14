package main

import (
	"fmt"
)

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

func BasicAtoi(s string) string {
	
	res:=""
	for i := 0; i < len(s); i++ {
		if s[i] >= '1' && s[i] <= '9' {
			res += s[i]
		}
	}
	return string(res)
}

// 12345
// 12345
// 0
