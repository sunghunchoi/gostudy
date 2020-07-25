package main

import "fmt"

// 전역변수 a 선언
var a int = 1

func localVarTest() int {
	// localVarTest()의 지역변수 a 선언
	var a int = 10
	a += 3
	return a
}

func globalVar() int {
	a += 3
	return a
}

func main() {
	fmt.Println("지역변수 a의 연산: ", localVarTest())
	fmt.Println("전역변수 a의 연산: ", globalVar())
}
