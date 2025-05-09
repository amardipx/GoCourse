package intermediate

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
}

type Address struct {
	city    string
	country string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAge() {
	p.age++
}

func main() {

	p1 := Person{
		firstName: "James",
		lastName:  "Bond",
		age:       30,
		address: Address{
			city:    "London",
			country: "U.K",
		},
	}

	// p2 := Person{
	// 	firstName: "Sherlock",
	// 	age:       25,
	// }

	fmt.Println(p1.firstName)
	// fmt.Println(p1.lastName)
	fmt.Println(p1.fullName())

	fmt.Println("Before increment:", p1.age)
	p1.incrementAge()
	fmt.Println("After increment:", p1.age)

	// Anonymous structs

	// user := struct {
	// 	username string
	// 	email    string
	// }{
	// 	username: "user123",
	// 	email:    "email@email.com",
	// }
}
