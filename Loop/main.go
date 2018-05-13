package main

import ("fmt")

func main (){

	x := 0
	var y int
	yy := 20
	var xx int
	for  i := 0;i < 10 ;i++  {
		x++
		fmt.Println(x)
		fmt.Printf("Perulangan Ke %d\n", x)
	}

	for y < yy {
		y++
		fmt.Printf("Perulangan Ke %d\n",x)
	}

	for {

		xx++

		if xx == 5 {
			fmt.Println("Lanjut")
			continue
		}

		fmt.Printf("Hello go %d\n",xx)

		if xx==10 {
			break
		}
	}

}
