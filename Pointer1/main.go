package main

import (
	"fmt"
	)
func main() {

	var hello string = "Hello"
	var strPointer *string
	//Mengecek alamat memori
	strPointer = &hello

	fmt.Println(strPointer)

	change(hello)
	fmt.Println(hello)

	changePtr(&hello)
	fmt.Println(hello)
}
//pass by value
func change(v string){
	v = v + " Golang"
}
//pass by reference
func changePtr(v *string){
	*v = *v + " Golang"
}