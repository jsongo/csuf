package main

import "fmt"

type Contact struct {
	greeting string
	name     string
	// name     string
	// greeting string
}

func Greet(person Contact) {
	fmt.Println(CreateMessage(person.greeting, person.name))
}

func CreateMessage(greeting, name string) string {
	return greeting + " " + name
}

func main() {

	var t = Contact{"Good to see you,", "Medhi"}
	Greet(t)

	u := Contact{"Glad you're in class,", "Sushant"}
	Greet(u)

	v := Contact{}
	v.greeting = "We're learning great things,"
	v.name = "Marcus"
	Greet(v)
}

// you don't need to return anything (void)
