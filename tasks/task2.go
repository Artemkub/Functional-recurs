// Структуры: именованная и анонимная структура user
package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
}

func main() {

	u := user{
		name:  "Иван",
		email: "ivan@example.com",
		age:   20,
	}

	fmt.Println("Name:", u.name)
	fmt.Println("Email:", u.email)
	fmt.Println("Age:", u.age)

	a := struct {
		name  string
		email string
		age   int
	}{
		name:  "Мария",
		email: "maria@example.com",
		age:   22,
	}

	fmt.Println("Name:", a.name)
	fmt.Println("Email:", a.email)
	fmt.Println("Age:", a.age)
}
