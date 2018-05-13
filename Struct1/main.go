package main

import (
	"fmt"
)

type Vector struct {
	X int
	Y int
}

type Player struct {
	ID int
	Name string
}
func main() {
 var v Vector
 v.X = 10
 v.Y = 20

 fmt.Println(v)
 fmt.Println("Y = ",v.Y)
 fmt.Println("X = ",v.X)

 var player Player
 player.ID = 1
 player.Name = "Dian Setiyadi"

 player1 := Player{2,"Michael Jordan"}

 fmt.Println(player.ID)
 fmt.Println(player.Name)
 fmt.Println(player1.ID)
 fmt.Println(player1.Name)

}

