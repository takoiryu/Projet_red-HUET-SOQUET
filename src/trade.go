package main

import "fmt"

type Gandalf struct {
	Obj  string
	Prix int
}

func InitGandalf() []Gandalf {
	return []Gandalf{
		{Obj: "potion de soin", Prix: 3},
		{Obj: "potion de poison", Prix: 6},
		{Obj: "livre de sort : boule de feu", Prix: 3},
		{Obj: "chapeau en cuir", Prix: 5},
		{Obj: "plastron en cuir", Prix: 10},
		{Obj: "bottes en cuir", Prix: 4},
		{Obj: "cuir de goblin", Prix: 2},
		{Obj: "lingots de fer", Prix: 15},
		{Obj: "minerais de mithril", Prix: 30},
	}
}
func AffGandalf(shop []Gandalf, joueur *Character) {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║            BOUTIQUE DE GANDALF           ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	for i, item := range shop {
		fmt.Printf("║ %d. %s - %d pièces d'or\n", i+1, item.Obj, item.Prix)
	}
	fmt.Println("║ 0. Quitter la boutique")
	fmt.Println("╚══════════════════════════════════════════╝")
	for {
		var choix int
		fmt.Print("\nQuel objet souhaitez-vous acheter (numéro) ? ")
		fmt.Scan(&choix)

		switch {
		case choix == 0:
			fmt.Println("À bientôt dans ma boutique !")
			return // Quitte la fonction et sort du magasin

		case choix > 0 && choix <= len(shop):
			AcheterObjet(joueur, shop[choix-1])

		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
func AcheterObjet(joueur *Character, article Gandalf) {
	indexOr := -1
	for i := range joueur.Inventaire {
		if joueur.Inventaire[i].NomObj == "pièce d'or" {
			indexOr = i
			break
		}
	}
	if indexOr == -1 || joueur.Inventaire[indexOr].Quantite < article.Prix {
		fmt.Printf("Vous n'avez pas assez d'or pour acheter %s ! (Prix : %d)\n", article.Obj, article.Prix)
		return
	}
	joueur.Inventaire[indexOr].Quantite -= article.Prix
	trouve := false
	for i := range joueur.Inventaire {
		if joueur.Inventaire[i].NomObj == article.Obj {
			joueur.Inventaire[i].Quantite++
			trouve = true
			break
		}
	}
	if !trouve {
		joueur.Inventaire = append(joueur.Inventaire, Invent{NomObj: article.Obj, Quantite: 1})
	}
	fmt.Printf("Achat réussi ! Vous avez acheté : %s (-%d pièces d'or)\n", article.Obj, article.Prix)
}
