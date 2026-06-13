package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	intch := make(chan int)

	strch := make(chan string)

	go func() {
		i := 1
		for {
			intch <- i
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		i := 1
		for {
			strch <- "Hello " + strconv.Itoa(i)
			i++
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	for {
		select {
		case number := <-intch:
			fmt.Println(number)
		case str := <-strch:
			fmt.Println(str)
		}
	}
}
