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
	ch := make(chan int)
	var wg sync.WaitGroup

	result := make([]int, 0)

	wg.Add(1)
	go func() {
		for i := 0; i < n; i++ {
			wg.Add(1)
			go func() {
				value := rand.IntN(100) + 1
				ch <- square(value)
				wg.Done()
			}()
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		result = append(result, res)
	}

	fmt.Println(result)
}
