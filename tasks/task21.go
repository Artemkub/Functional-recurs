// Каналы 2: fan out, 100 горутин генерируют случайные числа
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const goroutines = 100

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	ch := make(chan int, goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			ch <- rand.Intn(1000)
		}()
	}

	wait := goroutines

	var numbers []int

	for wait > 0 {
		n := <-ch
		numbers = append(numbers, n)
		wait--
	}

	fmt.Println(len(numbers), "чисел:")
	fmt.Println(numbers)
}
