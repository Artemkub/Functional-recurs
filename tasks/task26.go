// Дженерики: обобщённый стек stack
package main

import (
	"errors"
	"fmt"
)

type stack[T any] struct {
	data []T
}

func (s *stack[T]) push(value T) {
	s.data = append(s.data, value)
}

func (s *stack[T]) pop() (T, error) {
	var zero T

	if len(s.data) == 0 {
		return zero, errors.New("stack is empty")
	}

	last := len(s.data) - 1
	value := s.data[last]
	s.data = s.data[:last]

	return value, nil
}

func main() {

	var s stack[int]

	s.push(10)
	s.push(20)

	value, err := s.pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("pop:", value)

	value, err = s.pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("pop:", value)

	_, err = s.pop()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
}
