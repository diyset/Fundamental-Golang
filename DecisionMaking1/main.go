package main

import (
	"fmt"
)


func main(){
	x := 10
	y := 5
	fmt.Println(x>=y)
	if x>= y {
		fmt.Println("Ini akan di eksekusi jika true")
	} else {
		fmt.Println("Tidak Dieksekusi!")
	}

}
