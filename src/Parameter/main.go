package main

import "fmt"

// 가변인자로 string형태의 값 여러개를 받음
func addCaseOne(names ...string) string {
	var result string
	// for문을 이용한 num[i]의 순차적인 접근
	for i := 0; i < len(names); i++ {
		result = result + names[i]
		if i+1 != len(names) {
			result = result + ","
		}
	}
	return result
}

// 가변인자로 string형태의 값 여러개를 받음
func addCaseTwo(names ...string) string {
	var result string
	// for문을 이용한 num[i]의 순차적인 접근
	for i, v := range names {
		result = result + v
		if i+1 != len(names) {
			result = result + ","
		}
	}
	return result
}

func main() {
	s1, s2, s3, s4, s5 := "picachu", "raichu", "pairi", "kame", "hennaHana"
	ss := []string{"a1", "a2", "a3", "a4"}
	fmt.Println(addCaseOne(s1, s2))
	fmt.Println(addCaseOne(s1, s2, s4))
	fmt.Println(addCaseOne(ss...))
	fmt.Println(addCaseTwo(s3, s4, s5))
	fmt.Println(addCaseTwo(s1, s3, s4, s5))
	fmt.Println(addCaseTwo(ss...))
}
