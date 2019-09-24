package main

import (
	"fmt"
	"runtime"
)

// 채널은 고루틴끼리 데이터를 주고 받고, 실행 흐름을 제어하는 기능.
// 모든 타입은 채널을 이용할수 있음.
// 채널 자체는 값이아닌 레퍼런스
func sum(a int, b int, c chan int) {
	c <- a + b
}

func main() {
	doMajorChanger()
	//checkClose()
	//useClose()
	//doChannel2()
	// 동기채널은 값을 보내는쪽에서는 값을 받는쪽에서 값을 받을때 까지 대기.
	// 값을 받는쪽은 값을 보내는쪽에서 값을 보낼때까지 대기.
	//done := make(chan bool)
	//count := 3
	//
	//go func() {
	//	for i:=0; i < count; i++{
	//		done <- true
	//		fmt.Println("고루틴 : ", i)
	//		time.Sleep(1 * time.Second)
	//	}
	//}()
	//
	//for i := 0; i< count; i++{
	//	<- done
	//	fmt.Println("main func : " , i)
	//}

	//c :=  make(chan int)
	//
	//// 채널을 매개변수로 받는 함수는 반드시 고루틴으로 실행되어야함.
	//go sum(1,2,c)
	//
	//n := <- c
	//
	//fmt.Println(n)
}

func doChannel2() {
	// 채널 버퍼링( 채널에서 버퍼가 가득 차면 값을 꺼내서 출력
	runtime.GOMAXPROCS(1)

	// 버퍼가 2개인 비동기 채널 생성
	done := make(chan bool, 1)
	count := 4

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("go routine : ", i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("main func : ", i)
	}
}

func useClose() {
	c := make(chan int)

	// 채널을 가져온 뒤 두 번째 리턴값으로 채널이 닫혔는지 확인할 수 있음.
	go func() {
		for i := 0; i < 5; i++ {
			c <- i // 채널에 값을 보냄
		}
		close(c) // 채널을 닫음.
	}()

	// range 는 채널에 값이 몇 개나 들어올지 모르기 때문에 값이 들어올 때마다 계속 꺼내기 위해 사용.
	for i := range c { // range 를 사용하여 채널에서 값을 꺼냄.
		fmt.Println(i) // 꺼낸 값을 출력
	}
}

func checkClose() {

	c := make(chan int, 1)

	go func() {
		c <- 1
	}()

	if a, ok := <-c; ok {
		fmt.Println(a, ok)
	}

	close(c)

	if a, ok := <-c; ok {
		fmt.Println(a, ok)
	} else {
		fmt.Println("Channel c is closed")
	}
}

// 보내기 전용 채널
func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	c <- 100
	// fmt.Println(<-c) 보내기 전용채널에서 값을 꺼내면 컴파일에러!
}

// 받기 전용 채널
func consumer(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(<-c)
	// c<- 1  받기 전용채널에서 값을 보내면 컴파일 에러
}

func doMajorChanger() {
	c := make(chan int)
	go producer(c)
	go consumer(c)

	fmt.Scanln()
}
