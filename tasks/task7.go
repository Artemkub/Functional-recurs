// Массивы: присваивание массива и вывод значений и адресов элементов
package main

import "fmt"

func main() {

	var first [5]string
	second := [5]string{"Анна", "Борис", "Вера", "Глеб", "Дарья"}

	first = second

	for i, v := range first {
		fmt.Println("Значение:", v, "Адрес:", &first[i])
	}
}
