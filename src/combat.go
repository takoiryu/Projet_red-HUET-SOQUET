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
		if joueur.Pv <= 0 {
			fmt.Printf("\nDéfaite...  vous êtes KO. Retour au menu principal...vous êtes pas très fort\n")
			menu(joueur, magazin, forge)
			return
		}
		if monstre.Pv <= 0 {
			fmt.Printf("\nVictoire ! Retour au menu principal.")
			menu(joueur, magazin, forge)

		}
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

			degat := 0

			switch typesort {
			case "1":
				fmt.Println("\nCoup de poing")
				degat = 10
				monstre.Pv = monstre.Pv - degat
				joueur.Pv -= 10
				fmt.Printf("%s inflige à %s %d de dégâts !\n", joueur.Nom, monstre.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
				fmt.Printf("%s inflige à %s %d de dégâts !\n", monstre.Nom, joueur.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
				degat = 0
			case "2":
				fmt.Println("\nBoule de feu")
				degat = 20
				monstre.Pv = monstre.Pv - degat
				joueur.Pv -= 10
				fmt.Printf("%s inflige à %s %d de dégâts !\n", joueur.Nom, monstre.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
				fmt.Printf("%s inflige à %s %d de dégâts !\n", monstre.Nom, joueur.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
				degat = 0
			default:
				fmt.Println("Choix invalide.")
				continue
			}
			monstre.Pv -= degat
		case "2":
			menu(joueur, magazin, forge)
			return
		default:
			fmt.Println("\nChoix invalide ! Veuillez saisir 1 ou 2.")
		}
	}
}

func isDead(joueur *Character, monstre *Goblins) {
	magazin := InitGandalf()
	forge := InitGimly()
	if joueur.Pv <= 0 {
		fmt.Printf("\nDéfaite...  vous êtes KO.\n")
		fmt.Println("Retour au menu principal...vous êtes pas très fort")
		menu(joueur, magazin, forge)
		return
	}
	if monstre.Pv <= 0 {
		fmt.Printf("\nVictoire !")
		fmt.Println("Retour au menu principal.")
		menu(joueur, magazin, forge)
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
