package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 45
	return &p
}
func main() {
	fmt.Print(person{"Bob", 20})
	fmt.Print(newPerson("sean"))
	s := newPerson("sean")
	s.age = 30
	fmt.Print(&s)
}
