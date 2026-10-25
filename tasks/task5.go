// Константы: типизированные и нетипизированные
package main

import "fmt"

const (
	server = "localhost"

	port int = 8080
)

func main() {

	fmt.Println("server:", server)
	fmt.Println("port:", port)

	var result float64 = 3 / 2.0
	fmt.Println("3 / 2.0 =", result)

	var product int64 = 1000 * 1000
	fmt.Println("1000 * 1000 =", product)
}
