package main

import (
	"fmt"
)

func main() {
	var joueur Character
	joueur.CharacterCreation()
	joueur.Inventaire = InitInventaire()
	for !menu(&joueur) {
	}
}
func menu(joueur *Character) bool {
	var choixmenu string
	fmt.Println("\n--- MENU ---")
	fmt.Println("1:Afficher les informations du personnage \n 2:Accéder au contenu de l inventaire \n 3:retour \n 4:Quitter")
	fmt.Scan(&choixmenu)
	switch choixmenu {
	case "1":
		joueur.DisplayInfo()
	case "2":
		AfficherInvent(joueur)
	case "3":
		fmt.Println("Retour au jeu...")
	case "4":
		return true
	default:
		fmt.Println("Choix invalide.")
	}
	return false
}
