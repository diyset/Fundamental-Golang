package main

import (
	"fmt"
)

func main() {

	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	for i := 0;i < len(arr) ;i++  {
		fmt.Println(arr[i])
	}

	fmt.Println(arr[3])

	arrStr := [5] string{"Dian","Dadang","Babay","Rudy","Dugong"}
	for i := 0;i < len(arrStr) ;i++  {
		fmt.Println(arrStr[i])
	}

	fmt.Println(arrStr[4])
	//						Dimensi 0 | Dimensi 1
	arrMulti := [2][3] int{{1, 2, 3},{4, 5, 6}}
	fmt.Println(arrMulti[0][2])
	for i := 0;i < len(arrMulti) ;i++  {
		fmt.Println(arrMulti[0][i])
	}

}