// Указатели: адрес и значение переменной
package main

import "fmt"

func main() {

	count := 20

	fmt.Println("Адрес count:", &count)
	fmt.Println("Значение count:", count)

	var p *int = &count

	fmt.Println("Адрес p:", &p)
	fmt.Println("Значение p (адрес count):", p)
	fmt.Println("Значение, на которое указывает p:", *p)
}
