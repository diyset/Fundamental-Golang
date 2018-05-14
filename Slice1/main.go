package main

import (
	"fmt"
)

func main() {
	mySlice := []int{10,20,30,40,50}

	for i, v := range mySlice {
		fmt.Println(i, v)
	}

	mySliceStr := []string{"Dian","Farid","Derosi","Fajar"}
	mySliceStr = append(mySliceStr, "Dadang")
	for _,v := range mySliceStr {
		fmt.Println(v)
	}

}

