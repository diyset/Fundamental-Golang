package main

import (
	"fmt"
)

func main() {
	var mapPerson map[int] string
	mapPerson = make(map[int]string)
	mapPerson[1] = "Dian Setiyadi"
	mapPerson[2] = "Farid"
	mapPerson[3] = "Dadang"

	fmt.Println(mapPerson[1])

	for k, v := range mapPerson  {
		fmt.Println("Key Map Is =",k,"Value Map Is =",v)
	}

	farid, ok := mapPerson[4]
	if !ok {
		fmt.Println("Dadang tidak ada di mapPerson")
	} else {
		fmt.Println(farid)
	}

}

