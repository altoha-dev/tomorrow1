package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

func BasicAtoi(s string) int {
	res := ""
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '1' && s[i] <= '9' {
			res += string((s[i]))
			num, _ = strconv.Atoi(res)
		}
	}
	return num
}

// 12345
// 12345
// 0
