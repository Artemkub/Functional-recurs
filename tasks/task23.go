// Каналы 4: пул воркеров, 100 нечётных чисел, остановка через shutdown
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	values := make(chan int)

	shutdown := make(chan struct{})

	poolSize := runtime.GOMAXPROCS(0)

	var wg sync.WaitGroup
	wg.Add(poolSize)

	for i := 0; i < poolSize; i++ {
		go func() {
			defer wg.Done()

			for {
				n := rand.Intn(1000)

				select {
				case values <- n:
				case <-shutdown:
					return
				}
			}
		}()
	}

	var numbers []int

	for n := range values {
		if n%2 == 0 {
			continue
		}

		numbers = append(numbers, n)

		if len(numbers) == 100 {
			break
		}
	}

	close(shutdown)

	wg.Wait()

	fmt.Println(len(numbers), "нечётных чисел:")
	fmt.Println(numbers)
}
