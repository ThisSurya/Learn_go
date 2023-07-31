package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email  string
	Alamat int
}

func main() {
	//how to declare struct and initialize it
	surya := person{
		firstName:   "Surya",
		lastName:    "Michael",
		contactInfo: contactInfo{"adfadsf", 213},
	}

	//another ways:
	// surya1 := person{firstName: "Surya", lastName: "Michael"}

	fmt.Println(surya)

	//another again
	var surya2 person
	//the default value firstName and lasName surya2 is ""

	//to fill it
	surya2.firstName = "Surya"
	surya2.lastName = "Michael"
	surya2.update("jotaro")
	surya2.print()
}

func (p *person) update(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
