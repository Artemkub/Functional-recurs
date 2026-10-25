// Функции: возврат указателя и ошибки
package main

import "fmt"

type user struct {
	name  string
	email string
}

func newUser() (*user, error) {
	u := user{name: "Иван", email: "ivan@example.com"}
	return &u, nil
}

func main() {

	u, err := newUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(*u)

	_, err = newUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Второй вызов без ошибки")
}
