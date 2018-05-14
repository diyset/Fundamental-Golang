package main

import (
	"fmt"
	)
func main() {
 person := &Person{ID:1,Name:"Dian Setiyadi",Age:22,Job:"Programmer"}

 fmt.Println(person.GetID())
 fmt.Println(person.GetName())
 fmt.Println(person.GetAge())
 fmt.Println(person.GetJob())

 fmt.Println("Perubahan Dari Fungsi ChangeName")

 person.ChangeName("Michael Jordan")
 fmt.Println(person.GetName())

}

type Person struct {
	ID int
	Name string
	Age int
	Job string
}
//Function pointer untuk mengganti nama
func (p *Person) ChangeName(newName string){
	p.Name = newName
}

func (p *Person) GetID() int{
	return p.ID
}

func (p *Person) GetName() string{
	return p.Name
}

func (p *Person) GetAge() int{
	return p.Age
}
func (p *Person) GetJob() string{
	return p.Job
}
