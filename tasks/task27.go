// Дженерики: marshal только для значений с интерфейсом json.Marshaler
package main

import (
	"encoding/json"
	"fmt"
)

func marshal[T json.Marshaler](v T) ([]byte, error) {
	data, err := v.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return data, nil
}

type user struct {
	name  string
	email string
}

func (u user) MarshalJSON() ([]byte, error) {
	v := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{
		Name:  u.name,
		Email: u.email,
	}

	return json.Marshal(v)
}

func main() {

	u := user{name: "Bill", email: "bill@ardanlabs.com"}

	data, err := marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}
