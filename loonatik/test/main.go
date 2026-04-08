package main

import (
	"fmt"
	"strconv"
)

func main() {
	hexStr := "1E"
	decimal, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Println("Hex:", hexStr, "decimal", decimal)

	for i:=0, i<len()
}
