package main

import "fmt"

type Storager interface {
	Get() string
	Set(string)
}

type Person struct {
	name string
}

func (p *Person) Get() string {
	return p.name
}

func (p *Person) Set(name string) {
	p.name = name
}

func NewPerson(name string) *Person {
	return &Person{name}
}

func Exec(s Storager, name string) {
	s.Set(name)
	fmt.Println(s.Get())
}

func main() {
	p := NewPerson("John")
	Exec(p, "jorge")
}
