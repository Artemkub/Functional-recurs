// Экспорт: программа, использующая пакет toy
package main

import (
	"fmt"

	"play.ground/toy"
)

func main() {

	t := toy.New("Машинка", 300)

	t.UpdateOnHand(100)
	t.UpdateSold(25)

	fmt.Println("Name:", t.Name)
	fmt.Println("Weight:", t.Weight)
	fmt.Println("OnHand:", t.OnHand())
	fmt.Println("Sold:", t.Sold())
}
