package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var myString = "Hello Golang"
	var myStringTwo = "10"

	fmt.Println(myString)
	fmt.Println(len(myString))
	fmt.Println("Upper String")
	//Upper String
	myStringUpper := strings.ToUpper(myString)
	fmt.Println(myStringUpper)
	fmt.Println("Lower String")
	//Lower String
	myStringLower := strings.ToLower(myString)
	fmt.Println(myStringLower)

	fmt.Println("Contains Dari String")
	//Mengecek isi dari string jika ada return nya adalah "true"
	resultContains := strings.Contains(myString,"Go")
	fmt.Println(resultContains)

	fmt.Println("Split String")
	//Split String
	splitString := strings.Split(myString," ")
	fmt.Println(splitString)
	for _, v := range splitString  {
		fmt.Println(v)
	}

	fmt.Println("Parsing dari String ke Integer")
	//Parse dari String ke Integer
	resultParsingInt, err := strconv.Atoi(myStringTwo)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultParsingInt * 10)
	}

	//Cara mengkonversi dari integer menjadi string
	resultParsingStr := strconv.Itoa(10)
	fmt.Println(resultParsingStr)
}
