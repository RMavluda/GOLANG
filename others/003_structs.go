package main

import "fmt"

type employee struct {
	name   string
	sex    string
	age    int
	salary int
}

func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}
func main() {
	employee1 := newEmployee("Hanna", "Female", 25, 2000)

	fmt.Printf("Employee: %s\nAge: %d\nSalary: %d\n", employee1.name, employee1.age, employee1.salary)
}
