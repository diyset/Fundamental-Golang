package main

import ("fmt")

func main() {
	fmt.Println("Fungsi Main")
	x := 5
	y := 15

	hasil := add(x, y)
	fmt.Println(hasil)

	fmt.Println(hello("Dian Setiyadi"))

	fmt.Println("Fungsi Swap")
	swapX := "Dian"
	swapY := "Setiyadi"
	fmt.Println("Sebelum Swap")
	fmt.Println(swapX,swapY)
	resultX, resultY := swap(swapX,swapY)
	fmt.Println(resultX, resultY)

}

func add(x int, y int) int {
	return x + y
}
func hello(name string) string {
	return fmt.Sprintf("Hello %s\n",name)
}

func swap(x string, y string) (string, string) {
	return y,x
}
