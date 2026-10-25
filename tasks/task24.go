// Дженерики: общая функция marshal для любого типа
package main

import (
	"encoding/json"
	"fmt"
)

func marshal[T any](v T) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type User struct {
	Name string
	Age  int
}

func main() {

	u := User{Name: "Bill", Age: 30}

	data, err := marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
