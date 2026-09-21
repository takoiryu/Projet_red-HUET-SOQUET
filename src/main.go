package main

import (
	"fmt"
)

func main() {
	var joueur Character
	joueur.CharacterCreation()
	joueur.Inventaire = InitInventaire()
	magazin := InitGandalf()
	for !menu(&joueur, magazin) {
	}
}
func menu(joueur *Character, magazin []Gandalf) bool {
	var choixmenu string
	fmt.Println("\n--- MENU ---")
	fmt.Println("1:Afficher les informations du personnage \n 2:Accéder au contenu de l inventaire \n 3:marchand\n 4:retour \n 5:Quitter")
	fmt.Scan(&choixmenu)
	switch choixmenu {
	case "1":
		joueur.DisplayInfo()
	case "2":
		AfficherInvent(joueur)
	case "3":
		AffGandalf(magazin, joueur)
	case "4":
		fmt.Println("Retour au jeu...")
	case "5":
		return true
	default:
		fmt.Println("Choix invalide.")
	}
	return false
}
