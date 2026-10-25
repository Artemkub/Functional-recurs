// Словари: map со строковым ключом и целыми значениями
package main

import "fmt"

func main() {

	ages := make(map[string]int)

	ages["Анна"] = 20
	ages["Борис"] = 21
	ages["Вера"] = 22
	ages["Глеб"] = 23
	ages["Дарья"] = 24

	for k, v := range ages {
		fmt.Println(k, v)
	}
}
