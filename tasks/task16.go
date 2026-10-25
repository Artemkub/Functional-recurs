// Обработка ошибок: пользовательский тип appError и интерфейс temporary
package main

import (
	"errors"
	"fmt"
	"log"
)

type appError struct {
	err     error
	message string
	code    int
}

func (a *appError) Error() string {
	return fmt.Sprintf("App Error[%s] Message[%s] Code[%d]", a.err, a.message, a.code)
}

func (a *appError) Temporary() bool {
	return a.code != 9
}

type temporary interface {
	Temporary() bool
}

func checkFlag(flag bool) error {
	if !flag {
		return &appError{err: errors.New("flag false"), message: "флаг выключен", code: 9}
	}

	return errors.New("default error")
}

func main() {

	for _, flag := range []bool{false, true} {
		err := checkFlag(flag)

		switch e := err.(type) {
		case temporary:
			log.Println("Ошибка:", err)
			if !e.Temporary() {
				log.Println("Ошибка не временная, программа должна завершиться")
			}
		default:
			log.Println("Ошибка:", err)
		}
	}
}
