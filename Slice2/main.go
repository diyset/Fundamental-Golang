package main

import (
	"fmt"
)

func main() {

	var players []Player
	players = []Player{Player{1,"Dian Setiyadi"},Player{2,"Farid"}}

	players = append(players, Player{3,"Dadang"})
	for _, v := range players  {
		fmt.Println(v)
	}
}

type Player struct {
	ID int
	Name string
}

