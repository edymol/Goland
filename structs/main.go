package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contactInfo
	password string
	price    float32
}

type contactInfo struct {
	email   string
	zipcode int
}

func main() {
	/* one way of declaring a type person that will give default values
	var alex person
	alex.firstName = "Alex"
	alex.lastname = "Anderson"
	fmt.Println(alex) this will print and empty array
	fmt.Printf("%+v", alex) // this will print the fields of alex
	*/
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		age:       32,
		contactInfo: contactInfo{
			email:   "email@email.com",
			zipcode: 07201,
		},
		password: "password123",
		price:    20.5,
	}
	fmt.Println(alex)
	//fmt.Printf("%+v", alex)
	/* The code below is a pointer to memory object.
	the & is an operator that gives access to the memory where the object is place. example &alex.
	* is an operator that access directly to the memory address so when we write *pointToPerson we are asking
	for the value that exist in the address &alex*/
	//alexPointer := &alex
	//alexPointer.updateName("Alexa") // this code will not update the name to Alexa because no pointer

	// another more compact way to write the above code.
	alex.updateName("Alexa")
	alex.print()
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	(*pointerToPerson).age = 50
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
