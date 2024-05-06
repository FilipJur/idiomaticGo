package main

import "fmt"

func main() {
	exercises()
}

func exercises() {
	exercise1()
	exercise2()
	exercise3()
}

func exercise1() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	fmt.Println(greetings, len(greetings), cap(greetings))
	fmt.Println(greetings[:2])
	fmt.Println(greetings[1:4])
	fmt.Println(greetings[3:])
}

func exercise2() {
	message := "Hi  and "

	fmt.Printf("%c\n", message[3])
}

func exercise3() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	var arthur = Employee{
		"arthur",
		"morgan",
		1,
	}
	var john = Employee{
		firstName: "John",
		lastName:  "Marston",
		id:        2,
	}
	var dutch Employee
	dutch.firstName = "Dutch"
	dutch.lastName = "Van Der Linde"
	dutch.id = 3

	fmt.Println(arthur, john, dutch)
}
