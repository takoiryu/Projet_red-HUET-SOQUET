package main

import "fmt"

type Gandalf struct {
	Obj  string
	Prix int
}
type Gimly struct {
	Objarm  string
	Prixarm int
}
type Recette struct {
	PrixOr  int
	Fer     int
	Cuir    int
	Mithril int
}

var recettes = map[string]Recette{
	"heaume en fer forgé":       {PrixOr: 2, Fer: 1, Cuir: 0, Mithril: 0},
	"plastron en fer forgé":     {PrixOr: 3, Fer: 2, Cuir: 1, Mithril: 0},
	"bottes en fer forgé":       {PrixOr: 1, Fer: 1, Cuir: 1, Mithril: 0},
	"cote de maille en mithril": {PrixOr: 10, Fer: 0, Cuir: 0, Mithril: 10},
}

func InitGandalf() []Gandalf {
	return []Gandalf{
		{Obj: "potion de soin", Prix: 3},
		{Obj: "potion de poison", Prix: 6},
		{Obj: "livre de sort : boule de feu", Prix: 3},
		{Obj: "chapeau en cuir", Prix: 5},
		{Obj: "plastron en cuir", Prix: 7},
		{Obj: "bottes en cuir", Prix: 4},
		{Obj: "cuir de goblin", Prix: 2},
		{Obj: "lingots de fer", Prix: 10},
		{Obj: "minerais de mithril", Prix: 30},
	}
}
func InitGimly() []Gimly {
	return []Gimly{
		{Objarm: "heaume en fer forgé", Prixarm: 2},
		{Objarm: "plastron en fer forgé", Prixarm: 3},
		{Objarm: "bottes en fer forgé", Prixarm: 1},
		{Objarm: "cote de maille en mithril", Prixarm: 10},
	}
}
func AffGandalf(shop []Gandalf, joueur *Character) {
	fmt.Println("|n╔══════════════════════════════════════════╗")
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
			return

		case choix > 0 && choix <= len(shop):
			AcheterObjet(joueur, shop[choix-1])
		case choix == 484:
			Easteregg()
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
func AffGimly(forge []Gimly, joueur *Character) {
	fmt.Println("|n╔══════════════════════════════════════════╗")
	fmt.Println("║            FORGE DE GIMMLY           ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	for i, item := range forge {
		fmt.Printf("║ %d. %s - %d pièces d'or\n", i+1, item.Objarm, item.Prixarm)
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
			return

		case choix > 0 && choix <= len(forge):
			AcheterForge(joueur, forge[choix-1])
		default:
			fmt.Println("Choix invalide, veuillez réessayer.")
		}
	}
}
func AcheterForge(joueur *Character, armure Gimly) {
	recette := recettes[armure.Objarm]
	// ca stoke les quantité de l'inventaire dans une map
	stock := make(map[string]int)
	for _, item := range joueur.Inventaire {
		stock[item.NomObj] = item.Quantite
	}
	//la ca verifie si on a ce qu'il faut pour le craft
	if stock["pièce d'or"] < recette.PrixOr ||
		stock["lingots de fer"] < recette.Fer ||
		stock["cuir de goblin"] < recette.Cuir ||
		stock["minerais de mithril"] < recette.Mithril {

		fmt.Printf("Ressources insuffisantes pour %s !\n", armure.Objarm)
		fmt.Printf("Requis : %d Or, %d Fer, %d Cuir, %d Mithril\n",
			recette.PrixOr, recette.Fer, recette.Cuir, recette.Mithril)
		return
	}
	// et la si on a le necessaire POUF! y sont plus dans l'inventaire
	trop := LimitInv(joueur.Inventaire, joueur)
	if trop == false {
		retirerIngred := map[string]int{
			"pièce d'or":          recette.PrixOr,
			"lingots de fer":      recette.Fer,
			"cuir de goblin":      recette.Cuir,
			"minerais de mithril": recette.Mithril,
		}
		for i := range joueur.Inventaire {
			nom := joueur.Inventaire[i].NomObj
			if qte, besoin := retirerIngred[nom]; besoin {
				joueur.Inventaire[i].Quantite -= qte
			}
		}
		//la ca ajoutte le reultat a l'inv
		trouve := false
		for i := range joueur.Inventaire {
			if joueur.Inventaire[i].NomObj == armure.Objarm {
				joueur.Inventaire[i].Quantite++
				trouve = true
				break
			}
		}
		if !trouve {
			joueur.Inventaire = append(joueur.Inventaire, Invent{NomObj: armure.Objarm, Quantite: 1})
		}

		fmt.Printf("Gimly a forgé votre %s !\n", armure.Objarm)
	} else {
		fmt.Println("plus de place dans votre inventaire")
	}
}
func AcheterObjet(joueur *Character, article Gandalf) {
	trop := LimitInv(joueur.Inventaire, joueur)
	if trop == false {
		indexOr := -1
		for i := range joueur.Inventaire {
			if joueur.Inventaire[i].NomObj == "pièce d'or" {
				indexOr = i
				break
			}
		}
		if indexOr == -1 || joueur.Inventaire[indexOr].Quantite < article.Prix {
			fmt.Printf("Vous n'avez pas assez d'or pour acheter %s ! (Prix : %d)|n", article.Obj, article.Prix)
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
		fmt.Printf("Achat réussi ! Vous avez acheté : %s (-%d pièces d'or)|n", article.Obj, article.Prix)
	} else {
		fmt.Println("plus de place dans votre inventaire")
	}
}
