package mock

import "fmt"

type Talker interface {
	SayHello(word string) (res string)
}

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{
		name: name,
	}
}

func (p *Person) SayHello(word string) (rep string) {
	return fmt.Sprintf("%s said %s", p.name, word)
}
