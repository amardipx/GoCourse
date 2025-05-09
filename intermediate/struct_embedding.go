package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

type employee struct {
	person
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi I am %s and I am %d years old\n", p.name, p.age)
}

func (e employee) introduce() { // Method overriding
	fmt.Printf("Hi I am %s, I am %d years old, my id is %s and I earn %.2f per month\n", e.name, e.age, e.empId, e.salary)
}

func main() {
	emp := employee{
		person: person{name: "John", age: 30},
		empId:  "E001", salary: 50000,
	}

	fmt.Println("Name:", emp.person.name)
	fmt.Println("Age:", emp.age) // Accesssing the embedded struct field directly, unlike above
	fmt.Println("Employee Id:", emp.empId)
	fmt.Println("Salary:", emp.salary)

	emp.introduce()
}
