// Дженерики: обобщённый тип keymap на основе map
package main

import "fmt"

type keymap[T any] map[string]T

func (k keymap[T]) set(key string, value T) {
	k[key] = value
}

func (k keymap[T]) get(key string) (T, bool) {
	value, found := k[key]
	return value, found
}

func main() {

	km := make(keymap[int])

	km.set("jack", 10)
	km.set("jill", 20)

	value, found := km.get("jack")
	if !found {
		fmt.Println("jack не найден")
		return
	}
	fmt.Println("jack:", value)

	value, found = km.get("jill")
	if !found {
		fmt.Println("jill не найден")
		return
	}
	fmt.Println("jill:", value)
}
