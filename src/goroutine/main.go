package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// goroutine은 함수를 동시에 실행시키는 기능.
// 스레드보다 운영체제의 리소스를 적게 사용하므로 많은 수의 고루틴을 쉽게 생성 할 수있음.

func hello(n int) {
	r := rand.Intn(100) // 100이하의 랜덤한 숫자 생성
	time.Sleep(time.Duration(r))
	fmt.Println("Hello, world!", n)
}

func main() {
	// 멀티코어 활용하기
	// Go언어는 CPU 코어를 하나만 사용하게 되어있음. 시스템의 모든 CPU를 사용하라면?
	runtime.GOMAXPROCS(runtime.NumCPU()) // CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정.
	fmt.Println(runtime.GOMAXPROCS(0))

	s := "Hello, world!"

	// 익명 함수를 고루틴으로 실행
	// 클로저를 고루틴으로 사용할 때 반복문에 의해 바뀌는 변수는 반드시 매개변수로 넘겨줘야 함.
	for i := 0; i < 100; i++ {
		go func(n int) {
			fmt.Println(s, n)
		}(i)
	}

	// 이렇게 되면 변수 i는 0부터 99까지 증가하고, 다시 100이 되면서 반복문이 끝남.
	// 반복문이 완전히 끝난 다음에 고루틴이 생성되므로 고루틴이 생성된 시점의 변수 i의 값은 100이 되어있음.
	// 따라서 모두 Hello,world 100 이 출력됨.
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(s, i)
		}()
	}

	// 클로저를 고루틴으로 실행할 때 반복문에 의해 바뀌는 변수는 반드시 매개변수로 넘겨줘야함.
	// 즉 매개변수로 넘겨주는 시점에 해당 변수의 값이 복사되므로 고루틴이 생성될 때 그대로 사용할 수 있음.
	// 또한 CPU 코어를 하나만 사용하든 여러 개 사용하든 상관없이 반복문에의해 바뀌는 변수는 매개변수로 넘겨줘야 함.

	// go 루틴으로 실행시킬 함수 앞에 go 키워드를 붙이면 됨. 간단?
	for i := 0; i < 100; i++ {
		go hello(i)
	}
	fmt.Scanln() // 메인함수가 종료되지 않도 대기
}
