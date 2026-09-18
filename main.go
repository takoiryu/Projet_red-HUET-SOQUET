package main

import (
	"Projet_red/character"
	"fmt"
)

func main() {
	c1 := character.InitCharacter(
		"FJ", "elfe", 1, 100, 40, []string{"potion de soin", "potion de soin", "potion de soin"},
	)
	fmt.Println(c1)
}
