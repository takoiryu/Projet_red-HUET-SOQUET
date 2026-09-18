package main

import (
	"Projet_red/character"
	"fmt"
)

func main() {
	c1 := []character.Character{
		{
			Nom:        "FJ",
			Classe:     "elfe",
			Niveau:     1,
			Pvmax:      100,
			Pv:         40,
			Inventaire: []string{"potion de soin", "potion de soin", "potion de soin"},
		},
	}
	fmt.Println(c1)
}
