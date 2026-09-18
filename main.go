package main

import (
	"Projet_red/character"
	"Projet_red/inventaire"
	"fmt"
)

func main() {
	var joueur character.Character
	joueur.CharacterCreation()
	for !menu(&joueur) {

	}
}
func menu(joueur *character.Character) bool {
	var choixmenu string
	fmt.Println("\n--- MENU ---")
	fmt.Println("1:Afficher les informations du personnage \n 2:Accéder au contenu de l inventaire \n 3:retour \n 4:Quitter")
	fmt.Scan(&choixmenu)
	switch choixmenu {
	case "1":
		joueur.DisplayInfo()
	case "2":
		inventaire.AfficherInvent()
	case "3":
		fmt.Println("Retour au jeu...")
	case "4":
		return true
	default:
		fmt.Println("Choix invalide.")
	}
	return false
}
