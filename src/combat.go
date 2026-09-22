package main

import (
	"fmt"
)

func goblinPattern(joueur *Character, gobelin *Goblins, tour int) {
	degats := gobelin.Att
	if tour%3 == 0 {
		degats = gobelin.Att * 2
	}
	joueur.Pv -= degats
	if joueur.Pv < 0 {
		joueur.Pv = 0
	}
	fmt.Printf("%s inflige à %s %d de dégâts !\n", gobelin.Nom, joueur.Nom, degats)
	fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
}

func characterTurn(joueur *Character, monstre *Goblins) {
	var choix string

	for {
		fmt.Println("\n===== TOUR DE JOUEUR =====")
		fmt.Println("1 : Attaquer")
		fmt.Println("2 : Inventaire")
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			degats := joueur.Att
			monstre.Pv -= degats
			if monstre.Pv < 0 {
				monstre.Pv = 0
			}

			fmt.Printf("\n%s utilise Attaque basique et inflige %d dégâts à %s !\n", joueur.Nom, degats, monstre.Nom)
			fmt.Printf("PV restants de %s : %d/%d\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
			return

		case "2":
			fmt.Println("\n[Inventaire]")

		default:
			fmt.Println("\nChoix invalide ! Veuillez saisir 1 ou 2.")
		}
	}
}
