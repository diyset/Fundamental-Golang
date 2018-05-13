package main

import (
	"fmt"
	)
func main() {
	person1 := Person{ID:1,Name:"Dian Setiyadi",Age:22}
	printPerson(person1)


}

type Person struct {
	ID int
	Name string
	Age int
}

func printPerson(p Person){
	fmt.Println("ID = ",p.ID)
	fmt.Println("Name = ",p.Name)
	fmt.Println("Age = ",p.Age)
}

