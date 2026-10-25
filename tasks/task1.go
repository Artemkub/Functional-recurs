// Переменные: нулевые значения, литералы и преобразование типа
package main

import "fmt"

func main() {

	var s string
	var i int
	var b bool

	fmt.Println("Нулевые значения:")
	fmt.Printf("string: %q\n", s)
	fmt.Printf("int: %d\n", i)
	fmt.Printf("bool: %t\n", b)

	s2 := "hello"
	i2 := 42
	b2 := true

	fmt.Println("С начальными значениями:")
	fmt.Printf("string: %q\n", s2)
	fmt.Printf("int: %d\n", i2)
	fmt.Printf("bool: %t\n", b2)

	pi := float32(3.14)
	fmt.Printf("float32: %v (тип %T)\n", pi, pi)
}
