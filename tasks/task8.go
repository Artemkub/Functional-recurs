// Срезы: nil-срез с append и срез строк
package main

import "fmt"

func main() {

	var numbers []int

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i*10)
	}

	for i, v := range numbers {
		fmt.Println(i, v)
	}

	names := []string{"Анна", "Борис", "Вера", "Глеб", "Дарья"}

	for i, v := range names {
		fmt.Println(i, v)
	}

	part := names[1:3]

	for i, v := range part {
		fmt.Println(i, v)
	}
}
