// Горутины: счёт вниз от 100 до 0 и вверх от 0 до 100
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(2)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 100; i >= 0; i-- {
			fmt.Println("[вниз]", i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Println("[вверх]", i)
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Terminating Program")
}
