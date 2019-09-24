package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

type Plane struct {
	val string
}

func MakeTire(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Trie_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Trie_P, "
			outPlaneChan <- plane
		}
	}
}

func MakeEngine(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlanChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Engine_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Engine_P, "
			outPlanChan <- plane
		}
	}
}

func StartCarWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car" + strconv.Itoa(i)}
		i++
	}
}

func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane" + strconv.Itoa(i)}
		i++
	}
}

func main() {
	carchan1 := make(chan Car)
	carchan2 := make(chan Car)
	carchan3 := make(chan Car)

	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)

	go StartCarWork(carchan1)
	go StartPlaneWork(planeChan1)
	go MakeTire(carchan1, planeChan1, carchan2, planeChan2)
	go MakeEngine(carchan2, planeChan2, carchan3, planeChan3)

	for {
		select {
		case resultCar := <-carchan3:
			fmt.Println(resultCar.val)
		case resultPlane := <-planeChan3:
			fmt.Println(resultPlane.val)
		}
	}
}
