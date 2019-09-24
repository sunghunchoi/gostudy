package main

import "fmt"

func pop(c chan int) {
	fmt.Println("Pop func")
	v := <-c
	fmt.Println(v)
}

func main() {
	var c chan int
	c = make(chan int)

	go pop(c)

	c <- 10

	fmt.Println("end of program")

}
