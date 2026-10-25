// Каналы 3: до 100 горутин, нечётные числа попадают в срез
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const goroutines = 100

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()

			n := rand.Intn(1000)

			if n%2 == 0 {
				return
			}

			ch <- n
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var numbers []int

	for n := range ch {
		numbers = append(numbers, n)
	}

	fmt.Println(len(numbers), "нечётных чисел:")
	fmt.Println(numbers)
}
