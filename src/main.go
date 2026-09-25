package main

import (
	"fmt"
)

func main() {
	fmt.Println(`▓▄▄▄▄▄▄▄▄▄▓ ▓▄▄▓       ▓▄▄▄▄▄▄▄▄▄  ▓▄▄▄▄▄▄▄▄    ▓▄▄▓ ▓▄▄▄▄▄▄▄▄▄ 
▒▀▀▀▀▀▀▀▀▀▒ ▒▓▓▒       ▒▀▀▀▀▀▀▓▓▓▒ ▒▀▀▀▀▀▓▓▓▄▄  ▒▀▀▒ ▒▀▀▀▀▀▀▓▓▓▒
▄▄▄▄        ░▒▒░       ▄▄▄▄   ░▒▒░ ▄▄▄▄  ▀▀▒▒▒░ ▄▄▄▄ ▄▄▄▄   ░▒▒░
█▓▓█ █▄▄█   █░░█       ░░░░   █░░█ ░░░░   ░░░▀▀ █░░█ ░░░░   █░░█
▒  ▒ ▒▄▄▒   ░  ░       ░  ░   ░  ░ ░  ░ ▄▀ ▄▀   ▒  ▒ ░  ░   ░  ░
▓░░▓        ▓  ▓       ▓░░▓   ▓░░▓ ▓░░▓ ▀▄░░▀▄  ▓░░▓ ▓░░▓   ▓░░▓
█▒▒█        ▀▀▀▀       █▒▒█   █▒▒█ █▒▒█   ▀▄▒▒█ █▒▒█ █▒▒█   █▒▒█
█▓▓▀▀▀▀▀▀▀█ █▀▀▀▀▀▀▀▀█ █▓▓▓▀▀▀▓▓▓█ █▓▓█    █▓▓█ █▓▓█ █▓▓▓▀▀▀▓▓▓█
█▄▄▄▄▄▄▄▄▄█ █▄▄▄▄▄▄▄▄█ █▄▄▄▄▄▄▄▄▀  █▄▄█    █▄▄█ █▄▄█ █▄▄▄▄▄▄▄▄▀ `)
	fmt.Println(` 🌋  L'OMBRE DE SAURON SUR LA TERRE DU MILIEU - Eldrid

Le Seigneur Ténébreux Sauron s'éveille à nouveau, menaçant de plonger la Terre du Milieu dans une nuit éternelle. Vous incarnez un jeune aventurier parti de rien, dernier rempart face au chaos grandissant.

⚔️ 	VOTRE MISSION :

PREPAREZ-VOUS : Visitez Gandalf pour acheter potions et livres de sorts, et passez voir Gimli à la forge pour concevoir des armures en fer et en mithril.

COMBATTEZ : Affrontez des hordes de monstres, accumulez de l'or et maîtrisez la magie.

L'ULTIME DEFI : Survivez aux épreuves jusqu'à faire face à Sauron lui-même dans un affrontement final pour le destin du monde.

Serez-vous le HEROS capable de terrasser le Seigneur Ténébreux ?
|`)
	var joueur Character
	joueur.CharacterCreation()
	joueur.Inventaire = InitInventaire()
	joueur.DisplayInfo()
	//magazin := InitGandalf()
	//forge := InitGimly()
	Fight1(&joueur)
	Fight2(&joueur)
	Fight3(&joueur)
	Fight4(&joueur)
	Fight5(&joueur)
	//for !menu(&joueur, magazin, forge) {
	//}

}
func menu(joueur *Character, magazin []Gandalf, forge []Gimly) bool {
	var choixmenu string
	fmt.Println("\n======= MENU =======")
	fmt.Println("1:Informations personages \n 2:Inventaire \n 3:Marchand \n 4:Forgeron \n 5:Retour \n 6:Quitter")
	fmt.Scan(&choixmenu)
	switch choixmenu {
	case "1":
		joueur.DisplayInfo()
	case "2":
		AfficherInvent(joueur)
	case "3":
		AffGandalf(magazin, joueur)
	case "4":
		AffGimly(forge, joueur)
	case "5":
		fmt.Println("Retour au jeu...")
	case "6":
		return true
	default:
		fmt.Println("Choix invalide.")
	}
	return false
}
