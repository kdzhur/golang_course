package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// #1 Task
	randomNumbersChan := make(chan int)
	averageChan := make(chan float64)
	done := make(chan bool)

	go numbersGenerator(randomNumbersChan)

	go avgValue(randomNumbersChan, averageChan)

	go printAverage(averageChan, done)

	<-done
}

func numbersGenerator(ch chan<- int) {

	for i := 0; i < 10; i++ {
		num := rand.Intn(100) + 1
		ch <- num
	}
	close(ch)
}

func avgValue(inputChan <-chan int, outputChan chan<- float64) {
	sum := 0
	count := 0

	for num := range inputChan {
		sum += num
		count++
	}

	average := float64(sum) / float64(count)
	outputChan <- average
	close(outputChan)
}

func printAverage(ch <-chan float64, doneCh chan<- bool) {
	average := <-ch
	fmt.Printf("Avg value: %.2f\n", average)
	doneCh <- true
}
