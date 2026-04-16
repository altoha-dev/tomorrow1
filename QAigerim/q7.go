package main

import (
	"fmt"
)

func main() {
	var temikGay bool = false

	istina(&temikGay)

	fmt.Println(temikGay)
}

func istina(n *bool) {
	*n = true
}
