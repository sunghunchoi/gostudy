package main

import (
	"fmt"
	"strconv"
)

// 각 값이나 인스턴스의 실제 타입은 상관하지 않고 구현된 메서드로만 타입을 판다하는것을 덕 타이핑이라고함(Duck Typing)
// 만약 어떤 새가 오리처럼 걷고, 헤엄치고 , 꽥꽥거리는 소리를 낸다면 나는 그 새를 오리라고 부르겠다
type Duck struct {
}

func (d Duck) quack() {
	fmt.Println("꽥!")
}

func (d Duck) feathers() {
	fmt.Println("오리는 흰색과 회색의 털을 가지고 있습니다.")
}

type Quacker interface {
	quack()
	feathers()
}

func inTheForest(q Quacker) {
	q.quack()
	q.feathers()
}

func main() {
	var donald Duck // 오리 인스턴스 생성
	var john Person // 사람 인스턴스 생성

	inTheForest(donald)
	inTheForest(john)

	// 특정 인터페이스를 구현하는지를 검사
	if v, ok := interface{}(donald).(Quacker); ok {
		fmt.Println(v, ok)
	}

	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("Hello,world!"))

	andrew := Person{
		name: "Andrew",
		age:  35,
	}

	fmt.Println(formatString(andrew))
	fmt.Println(formatString(Person{"Maria", 20}))
	fmt.Println(formatString(Person{"Jhon", 12}))

	// 인터페이스에 저정된 타입이 특정 타입인지를 검사.
	var t interface{}
	t = Person{name: "show", age: 30}
	if v, ok := t.(Person); ok {
		fmt.Println(v, ok)
	}
}

// 	빈 인터페이스 사용하기
//  인터페이스에 아무 메서드도 정의되어 있지 않다면 모든 타입을 지정할 수 있음.
func f1(arg interface{}) {
	// 모든 타입을 저장할수 있음.
}

type Any interface{}

func f2(arg Any) {}

// 구조체 인스턴스 및 포인터도 빈 인터페이스로 넘길수 있음
type Person struct {
	name string
	age  int
}

func (p Person) quack() {
	fmt.Println("사람은 오리를 흉내냅니다. 꽥~")
}

func (p Person) feathers() {
	fmt.Println("사람은 땅에서 깃털을 주워서 보여줍니다.")
}

// 빈 인터페이스 타입은 함수의 매개변수, 리턴값, 구조체의 필드로 사용할 수 있음.
func formatString(arg interface{}) string {
	// 빈 인터페이스를 사용하여 모든 타입을 받음
	// typeAssertion
	switch arg.(type) {
	case int:
		i := arg.(int)
		return strconv.Itoa(i)
	case float32:
		f := arg.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
	case float64:
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64)
	case string:
		s := arg.(string)
		return s
	case Person:
		p := arg.(Person)
		return p.name + " " + strconv.Itoa(p.age)
	case *Person:
		p := arg.(*Person)
		return p.name + " " + strconv.Itoa(p.age)
	default:
		return "Error"
	}
}
