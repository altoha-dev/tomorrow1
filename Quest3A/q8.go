package main

import (
	"fmt"
)

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) string {
	res := ""
	// for i:=12; 12>0; i--{
	// 11>0
	// 10
	// 0>0

	// }

	for i := len(s) - 1; i >= 0; i-- {
		// 12>0. 
		// 11>0
		res += string(s[i])
	}
	return res
}

// !dlroW olleH
// Hello World!
