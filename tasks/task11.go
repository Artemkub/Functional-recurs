// Интерфейсы: speaker, english и chinese
package main

import "fmt"

type speaker interface {
	speak()
}

type english struct{}

func (english) speak() {
	fmt.Println("Hello World")
}

type chinese struct{}

func (*chinese) speak() {
	fmt.Println("你好世界")
}

func sayHello(s speaker) {
	s.speak()
}

func main() {

	var s speaker

	e := english{}
	s = e
	s.speak()

	c := chinese{}
	s = &c
	s.speak()

	sayHello(english{})
	sayHello(&chinese{})
}
