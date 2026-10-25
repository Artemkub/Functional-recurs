// Массивы: присваивание массива и вывод значений и адресов элементов
package main

import "fmt"

func main() {

	var first [5]string
	second := [5]string{"Анна", "Борис", "Вера", "Глеб", "Дарья"}

	first = second

	for i, v := range first {
		fmt.Printf("Значение: %s  Адрес: %p\n", v, &first[i])
	}
}
