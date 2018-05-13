package main

import (
	"fmt"
)

func main() {
	//function dalam bentuk variable
	add := func(x,y int)int {
		return x + y
	}

	hello := func(name string) string {
		return fmt.Sprintf("Hello %s",name)
	}

	fmt.Println(add(10,20))
	fmt.Println(hello("Dian Setiyadi"))
}


