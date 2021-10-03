package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo // we don't need to specify the field name
}

func main() {
	// alex := person{"Alex", "Anderson"} // same as bellow but depending on order
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// var alex person // same as above but with zero values, in this case ""
	// fmt.Println(alex)
	// fmt.Printf("%+v\n", alex)

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.print()

	// access memory address of jim with &
	// convert value to pointer, unnecessary due to shortcuts
	// jimPointer := &jim
	jim.updateName("Jimmy")
	jim.print()
}

// Go is a pass by value language
// we must use * to specify a pointer, otherwise the assignment is ineffective
// if we don't specify a pointer p will be a copy of the receiver
// when using * on a type it means you are using a pointer to that type
func (p *person) updateName(newFirstName string) {
	// access the value of pointer with *
	// convert pointer to value, unnecessary due to shortcut
	// (*p).firstName = newFirstName
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
