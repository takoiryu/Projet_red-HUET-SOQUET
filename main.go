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
	fmt.Println("\n┌──────────────────────────────────────────┐")
	fmt.Println("│               MENU PRINCIPAL             │")
	fmt.Println("├──────────────────────────────────────────┤")
	fmt.Println("│  1. Afficher la fiche du personnage      │")
	fmt.Println("│  2. Ouvrir le sac / inventaire           │")
	fmt.Println("│  3. Retour au jeu                        │")
	fmt.Println("│  4. Quitter l'aventure                   │")
	fmt.Println("└──────────────────────────────────────────┘")
	fmt.Print("➔ Votre choix : ")
	fmt.Scan(&choixmenu)
	switch choixmenu {
	case "1":
		joueur.DisplayInfo()
	case "2":
		inventaire.AfficherInvent(joueur.Inventaire)
	case "3":
		fmt.Println("[i] Reprise de la partie...")
	case "4":
		fmt.Println("\n===========================================")
		fmt.Println("    Merci d'avoir joué ! À bientôt...      ")
		fmt.Println("===========================================")
		return true
	default:
		fmt.Println("[!] Choix invalide. Veuillez saisir 1, 2, 3 ou 4.")
	}
	return false
}
