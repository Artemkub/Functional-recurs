// Каналы 1: две горутины передают друг другу число десять раз
package main

import (
	"fmt"
	"sync"
)

func main() {

	share := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		goroutine("Горутина 1", share)
		wg.Done()
	}()

	go func() {
		goroutine("Горутина 2", share)
		wg.Done()
	}()

	share <- 1

	wg.Wait()
}

func goroutine(name string, share chan int) {
	for {
		value, ok := <-share
		if !ok {
			fmt.Println(name, "канал закрыт")
			return
		}

		fmt.Println(name, "получила", value)

		if value == 10 {
			close(share)
			return
		}

		share <- value + 1
	}
}
