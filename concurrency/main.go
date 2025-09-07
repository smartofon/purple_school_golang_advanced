package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func square(n int) int {
	return n * n
}

func main() {

	n := 10
	ch1 := make(chan int)
	ch2 := make(chan int)
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	result := make([]int, 0)

	wg1.Add(1)
	go func() {
		numbers := make([]int, 0)
		for i := 0; i < n; i++ {
			numbers = append(numbers, rand.IntN(100)+1)
		}
		for _, value := range numbers {
			ch1 <- value
		}
		wg1.Done()
	}()

	wg2.Add(1)
	go func() {
		for value := range ch1 {
			ch2 <- square(value)
		}
		wg2.Done()
	}()

	go func() {
		wg1.Wait()
		close(ch1)
	}()

	go func() {
		wg2.Wait()
		close(ch2)
	}()

	for res := range ch2 {
		result = append(result, res)
	}

	fmt.Println(result)

}
