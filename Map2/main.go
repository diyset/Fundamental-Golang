package main

import (
	"fmt"
)

func main() {

	mapPlayer := make(map[int]Player)

	mapPlayer[1] = Player{1,"Dian Setiyadi"}
	mapPlayer[2] = Player{2,"Farid"}
	mapPlayer[3] = Player{3,"Dadang"}

	for _, v := range mapPlayer {
		fmt.Println(v.Name)
	}


}

type Player struct {
	ID int
	Name string
}

