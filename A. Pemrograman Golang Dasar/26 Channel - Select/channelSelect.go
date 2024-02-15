package main

import (
	"fmt"
	"runtime"
	"time"
)

func getAverage(numbers []int, ch chan float64) {
	sum := 0
	for _, e := range numbers {
		sum += e
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	max := numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	ch <- max
}
func main() {
	runtime.GOMAXPROCS(2)
	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("Numbers :", numbers)

	chan1 := make(chan float64)
	go getAverage(numbers, chan1)

	chan2 := make(chan int)
	go getMax(numbers, chan2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-chan1:
			fmt.Printf("Nilai Avg : %.2f\n", avg)
		case max := <-chan2:
			fmt.Printf("Nilai Max : %d \n", max)
		}
	}

	time.Sleep(2 * time.Second)
}
