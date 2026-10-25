// Методы: средний показатель бейсболиста hits / atBats
package main

import (
	"fmt"
	"math"
)

type player struct {
	name   string
	atBats int
	hits   int
}

func (p player) average() float64 {
	return float64(p.hits) / float64(p.atBats)
}

func main() {

	players := []player{
		{name: "Иванов", atBats: 100, hits: 30},
		{name: "Петров", atBats: 80, hits: 28},
		{name: "Сидоров", atBats: 120, hits: 25},
	}

	for _, p := range players {
		fmt.Println(p.name+":", math.Round(p.average()*1000)/1000)
	}
}
