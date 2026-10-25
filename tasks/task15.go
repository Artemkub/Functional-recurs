// Обработка ошибок: ErrInvalidValue и ErrAmountTooLarge
package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidValue   = errors.New("invalid value")
	ErrAmountTooLarge = errors.New("amount too large")
)

func checkAmount(amount float64) error {
	if amount == 0 {
		return ErrInvalidValue
	}

	if amount > 1000 {
		return ErrAmountTooLarge
	}

	return nil
}

func main() {

	amounts := []float64{0, 1500, 250}

	for _, amount := range amounts {
		err := checkAmount(amount)

		switch err {
		case ErrInvalidValue:
			fmt.Println(amount, "- неверное значение:", err)
		case ErrAmountTooLarge:
			fmt.Println(amount, "- сумма слишком большая:", err)
		case nil:
			fmt.Println(amount, "- всё хорошо")
		default:
			fmt.Println(amount, "- неизвестная ошибка:", err)
		}
	}
}
