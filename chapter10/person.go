package chapter10

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func PersonCreation() {
	var p1 Person
	p1.firstName = "Edward"
	p1.lastName = "Chen"
	upPerson(&p1)
	fmt.Printf("The name of person is %v\n", p1)

	p2 := new(Person)
	p2.firstName = "Edward"
	p2.lastName = "Chen"
	upPerson(p2)
	fmt.Printf("The name of person is %+v\n", p2)

	p3 := Person{"Edward", "Chen"}
	upPerson(&p3)
	fmt.Printf("The name of person is %#v\n", p3)
}
