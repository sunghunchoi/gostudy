package main

import "fmt"

// return 값은 int형 받기전용 채널
//func sum(a,b int) <-chan int{
//	out := make(chan int)
//	go func(){
//		out <- a+b
//	}()
//	return out
//}

func main() {
	c := num(1, 2)
	out := sum(c)
	fmt.Println(<-out)
}

// int형 받기 전용채널을 리턴
func num(a, b int) <-chan int {
	out := make(chan int) // int 형 채널 생성
	go func() {
		out <- a   // 채널에 a의 값을 보냄.
		out <- b   // 채널에 b의 값을 보냄.
		close(out) // 채널을 닫음.
	}()
	return out
}

// 매개변수는 int형 받기전용채널, 리턴값도 int형 받기전용 채널
func sum(c <-chan int) <-chan int {
	out := make(chan int) // int형 채널 생성
	go func() {
		r := 0
		for i := range c { // range를 사용하여 채널이 닫힐 때까지 값을 꺼냄.
			r = r + i // 꺼낸 값을 모두 더함.
		}
		out <- r // 더한 결과를 채널에 보냄.
	}()
	return out
}
