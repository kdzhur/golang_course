package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbChan := make(chan int)
	resultChan := make(chan [2]int)
	doneChan := make(chan bool)

	go numbersGenerator(numbChan, resultChan, doneChan)
	go findMinMax(numbChan, resultChan)

	<-doneChan
}

func numbersGenerator(ch chan int, result chan [2]int, done chan bool) {

	for i := 0; i < 7; i++ {
		num := rand.Intn(100) + 1
		ch <- num
	}

	close(ch)
	fmt.Println(<-result)
	done <- true
}

func findMinMax(ch <-chan int, result chan<- [2]int) {
	min := 999999
	max := -999999

	for num := range ch {
		fmt.Println("Number: ", num)
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	result <- [2]int{min, max}
}
