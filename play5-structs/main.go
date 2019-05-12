package main

import "fmt"

type contactInfo struct {
	address string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// this is not a good idea because if we change the
	// order inside the struct then the calls will break
	// also this forces you to populate all the fields in the struct
	p1 := person{"Alexis", "Rodrigez", contactInfo{"Chile", 59772}}

	p2 := person{
		firstName: "Radamel",
		lastName:  "Falcao",
		contact: contactInfo{
			"Colombia",
			467788, // if you declaring props in new lines then each line should end with , (this is weird!!!)
		},
	}

	p1.print()
	p2.print()

	var p3 person
	fmt.Println(p3) // will print empty value

	p3.firstName = "Alessandro"
	p3.lastName = "Del Piero"
	p3.contact = contactInfo{address: "Italy", zipCode: 795564}
	p3.print()

	// passing the address of the variable to the function
	(&p3).modifyFirstName("Rui")
	(&p3).modifyLastName("Costa")
	p3.print()

	// even if you directly use the varible, go will send the address if the receiver accepts pointer
	p3.modifyFirstName("Alessandro")
	p3.modifyLastName("Nesta")
	p3.print()

	// everything in go is "pass by value", in this case it acts as refrence is because slices are reference types
	// reference types themselves store addresses to the underlying primitive value data structure
	// so when these reference types are copied inside the function they continue to point to the same underlying data structure
	names := []string{"Maldini", "Nesta", "Cafu", "Stam"}
	editNames(names)
	fmt.Printf("%+v", names)
}

func editNames(names []string) {
	names[0] = "Dida"
}

// function receives the pointer which is directly modifies
func (p *person) modifyFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person) modifyLastName(newLastName string) {
	(*p).lastName = newLastName
}
func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
