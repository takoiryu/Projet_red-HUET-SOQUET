package main

import (
	"fmt"
)

func goblinTurn(joueur *Character, gobelin *Goblins, tour int) {
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

func characterTurn1(joueur *Character, monstre *Goblins) {
	var choix string
	magazin := InitGandalf()
	forge := InitGimly()
	for {
		fmt.Println("\n===== TOUR DE JOUEUR =====")
		fmt.Println("1 : Attaquer")
		fmt.Println("2 : Menu")
		fmt.Print("Choix : ")
		fmt.Scan(&choix)

		switch choix {
		case "1":
			var typesort string
			fmt.Println("\nChoisissez votre sort :")
			fmt.Println("1 : Coup de poing")
			fmt.Println("2 : Boule de feu")
			fmt.Print("Votre choix : ")
			fmt.Scan(&typesort)

			var sort string
			degat := 0

			switch typesort {
			case "1":
				sort = "Coup de poing"
				degat = 10
			case "2":
				sort = "Boule de feu"
				degat = 20
			default:
				fmt.Println("Choix invalide.")
				continue
			}
		case "2":
			menu(joueur, magazin, forge)
			return
		default:
			fmt.Println("\nChoix invalide ! Veuillez saisir 1 ou 2.")
		}
	}
}

func isDead(joueur *Character) {
	magazin := InitGandalf()
	forge := InitGimly()
	if joueur.Pv <= 0 {
		fmt.Printf("\nDéfaite...  vous êtes KO.\n")
		fmt.Println("Retour au menu principal...")
		menu(joueur, magazin, forge)
		return
	}
}

func trainingFight(joueur *Character) {
	gobelin := initGobelins()
	tour := 1
	fmt.Println("\n==========================================")
	fmt.Printf("      DÉBUT DU COMBAT contre %s\n", gobelin.Nom)
	fmt.Println("==========================================")
	for joueur.Pv > 0 && gobelin.Pv > 0 {
		fmt.Printf("\n======TOUR %d ======\n", tour)

		characterTurn1(joueur, &gobelin)

		if gobelin.Pv <= 0 {
			fmt.Printf("\nVictoire ! Vous avez terrassé %s !\n", gobelin.Nom)
			fmt.Println("Retour au menu principal...")
			return
		}

		goblinTurn(joueur, &gobelin, tour)

		tour++
	}
}
