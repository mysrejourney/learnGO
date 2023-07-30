package main

import "fmt"

var name string = "Sats"

const age = 40

func main() {
	var place string = "Bangalore"

	fmt.Printf("My name is %v and the type is %T\n", name, name)
	fmt.Printf("My age is %v and the type is %T\n", age, age)
	fmt.Printf("My place is %v and the type is %T\n", place, place)
}
