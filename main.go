package main

import (
	"fmt"
	"time"
)

func mine(checkpoint chan int, n int) {

	fmt.Println("работник номер", n, "приступил к работе ")
	time.Sleep(1 * time.Second)
	fmt.Println("работник номер", n, "закончил  ")
	checkpoint <- 10

	//......

	fmt.Println("ура я свободен", n)

}

func main() {

	checkpoint := make(chan int)

	InitTime := time.Now()
	coal := 0
	go mine(checkpoint, 1)
	go mine(checkpoint, 2)
	go mine(checkpoint, 3)

	coal += <-checkpoint

	coal += <-checkpoint

	coal += <-checkpoint

	fmt.Println(coal)

	fmt.Println(time.Since(InitTime))

}
