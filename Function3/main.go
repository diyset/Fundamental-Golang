package main

import ("fmt")

func main() {


	nextValue := genNumber()

	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())

	lv := love("Dian")
	fmt.Println(lv("Golang"))
	fmt.Println(lv("Java"))
	fmt.Println(lv("Node JS"))
}

func genNumber() func() int {
	x := 0
	return func() int {
		x = x + 1
		return x
	}
}

func love(name string) func(string) string {
	return func(things string) string {
		return fmt.Sprintf("%s love %s",name,things)
	}
}
