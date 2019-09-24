package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)

	go func() {
		for {
			c1 <- 20
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case i := <-c1:
				fmt.Println(i)
				time.Sleep(100 * time.Millisecond)

			default:
				fmt.Println("waiting.....")
				time.Sleep(100 * time.Millisecond)

			}
		}
	}()
	fmt.Scanln()
}
