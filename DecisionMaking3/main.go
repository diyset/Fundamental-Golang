package main

import ("fmt")

func main() {

	x := 1001
	switch x {
	case 60:
		fmt.Println("X adalah 60")
	case 70:
		fmt.Println("X adalah 70")
	case 80:
		fmt.Println("X adalah 80")
	case 90:
		fmt.Println("X adalah 90")
	case 100:
		fmt.Println("X adalah 100")
	default:
		fmt.Println("Tidak ada di case!")
	}
}
