// Дженерики: copyfy с ограничением по типам string и int
package main

import "fmt"

type copyer interface {
	string | int
}

func copyfy[T copyer](s []T) []T {
	result := make([]T, len(s))
	copy(result, s)
	return result
}

func main() {

	names := []string{"Анна", "Борис", "Вера"}
	namesCopy := copyfy(names)
	fmt.Println(names, namesCopy)

	numbers := []int{1, 2, 3}
	numbersCopy := copyfy(numbers)
	fmt.Println(numbers, numbersCopy)
}
