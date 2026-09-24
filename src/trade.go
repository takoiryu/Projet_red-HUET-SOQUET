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
		{Obj: "Potion de soin", Prix: 3},
		{Obj: "Potion de poison", Prix: 6},
		{Obj: "Livre de sort : boule de feu", Prix: 3},
		{Obj: "Chapeau en cuir", Prix: 5},
		{Obj: "Plastron en cuir", Prix: 7},
		{Obj: "Bottes en cuir", Prix: 4},
		{Obj: "Cuir de goblin", Prix: 2},
		{Obj: "Lingots de fer", Prix: 10},
		{Obj: "Minerais de mithril", Prix: 30},
		{Obj: "Amelioration d'inventaire", Prix: 30},
		{Obj: "Sort: Boule de feu", Prix: 10},
	}
}
func InitGimly() []Gimly {
	return []Gimly{
		{Objarm: "Heaume en fer forgé", Prixarm: 2},
		{Objarm: "Plastron en fer forgé", Prixarm: 3},
		{Objarm: "Bottes en fer forgé", Prixarm: 1},
		{Objarm: "Cote de maille en mithril", Prixarm: 10},
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
		fmt.Print("\n Quel objet souhaitez-vous acheter (numéro) ? ")
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
		fmt.Println("Plus de place dans votre inventaire")
	}
}
func AcheterObjet(joueur *Character, article Gandalf) {
	if article.Obj == "Amelioration d'inventaire" {
		if joueur.NbrAmeliorationSac >= 3 {
			fmt.Println("Vous avez déjà amélioré votre sac 3 fois ! C'est le maximum.")
			return
		}
		indexOr := -1
		for i := range joueur.Inventaire {
			if joueur.Inventaire[i].NomObj == "Pièce d'or" {
				indexOr = i
				break
			}
		}

		if indexOr == -1 || joueur.Inventaire[indexOr].Quantite < article.Prix {
			fmt.Printf("Vous n'avez pas assez d'or pour acheter %s !\n", article.Obj)
			return
		}
		joueur.Inventaire[indexOr].Quantite -= article.Prix
		joueur.TailleMax = joueur.TailleMax + 10
		joueur.NbrAmeliorationSac = joueur.NbrAmeliorationSac + 1

		fmt.Printf("Sac agrandi ! Vous pouvez porter %d objets. (Amélioration %d sur 3)\n", joueur.TailleMax, joueur.NbrAmeliorationSac)
		return
	}
	if LimitInv(joueur.Inventaire, joueur) {
		fmt.Println("Plus de place dans votre inventaire pour un nouvel objet !")
		return
	}
	indexOr := -1
	for i := range joueur.Inventaire {
		if joueur.Inventaire[i].NomObj == "Pièce d'or" {
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
		if article.Obj == "Sort: Boule de feu" {
		}
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
}
