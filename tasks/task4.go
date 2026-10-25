// Указатели: изменение поля структуры через функцию
package main

import "fmt"

type user struct {
	name string
	age  int
}

func birthday(u *user) {
	u.age++
}

func main() {

	u := user{name: "Иван", age: 20}

	fmt.Println("До:", u)

	birthday(&u)

	fmt.Println("После:", u)
}
