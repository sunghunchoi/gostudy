package main

import (
	"fmt"
	"time"
)

// select {case <- 채널:코드 }
func main() {
	single()
	c1 := make(chan int)
	c2 := make(chan string)
	c3 := make(chan int)

	go func() {
		for {
			c1 <- 10
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c2 <- "Hello World!"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			i := <-c3
			fmt.Println("c3 : ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case c3 <- 3:
			case i := <-c1:
				fmt.Println("c1 : ", i)
			case s := <-c2:
				fmt.Println("c2 : ", s)
			case <-time.After(50 * time.Millisecond):
				fmt.Println("timeout")
			}
		}
	}()
	time.Sleep(10 * time.Second)
}

// 채널 한개로 select 사용하기
func single() {
	c1 := make(chan int)

	go func() {
		for {
			i := <-c1
			fmt.Println("c1 : ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			c1 <- 20
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			select {
			case c1 <- 10:
			case i := <-c1:
				fmt.Println("c1 : ", i)

			}
		}
	}()
	time.Sleep(10 * time.Second)
}
