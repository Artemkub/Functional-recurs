// Переменные: нулевые значения, литералы и преобразование типа
package main

import "fmt"

func main() {

	var s string
	var i int
	var b bool

	fmt.Println("Нулевые значения:")
	fmt.Println("string:", s)
	fmt.Println("int:", i)
	fmt.Println("bool:", b)

	s2 := "hello"
	i2 := 42
	b2 := true

	fmt.Println("С начальными значениями:")
	fmt.Println("string:", s2)
	fmt.Println("int:", i2)
	fmt.Println("bool:", b2)

	pi := float32(3.14)
	fmt.Println("float32:", pi)
}
