package main

import ("fmt")

func main() {
	fmt.Println("Fungsi Main")
	x := 5
	y := 15

	hasil := add(x,y)
	fmt.Println(hasil)

	fmt.Println(hello("Dian Setiyadi"))
}

func add(x int, y int) int {
	return x + y
}
func hello(name string) string {
	return fmt.Sprintf("Hello %s\n",name)
}