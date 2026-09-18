package inventaire

import "fmt"

type Invent struct {
	NomObj   string
	Quantite int
}

func InitInventaire() []Invent {
	return []Invent{
		{NomObj: "pièce d'or", Quantite: 100},
		{NomObj: "potion de soin", Quantite: 0},
		{NomObj: "potion de poison", Quantite: 0},
		{NomObj: "armure en cuir", Quantite: 0},
		{NomObj: "armure en fer forgé", Quantite: 0},
		{NomObj: "armure en mithril", Quantite: 0},
	}
}
func UtiliserObjet(lst []Invent, nomObjet string) bool {
	for i := range lst {
		if lst[i].NomObj == nomObjet {
			if lst[i].Quantite > 0 {
				lst[i].Quantite--
				fmt.Printf("Vous avez utilisé : %s (Reste : %d)\n", nomObjet, lst[i].Quantite)
				return true
			} else {
				fmt.Printf("Vous n'avez plus de %s !\n", nomObjet)
				return false
			}
		}
	}
	fmt.Println("Objet non trouvé dans l'inventaire.")
	return false
}
func AfficherInvent(lst []Invent) {
	fmt.Println("===INVENTAIRE===")
	vide := true
	for _, item := range lst {
		if item.Quantite > 0 {
			fmt.Println(item.NomObj, ":", item.Quantite)
			vide = false
		}
	}
	if vide {
		fmt.Println("Votre inventaire est vide.")
	} else {
		var use string

		fmt.Println("utiliser un object ? y/n")
		fmt.Scan(&use)
		if use == "y" {
			var usechoix string
			fmt.Print("Quel objet ? potions (pot), armure (arm) : ")
			fmt.Scan(&usechoix)

			switch usechoix {
			case "pot":
				var typePot string
				fmt.Print("Soin ou poison ? (s/p) : ")
				fmt.Scan(&typePot)

				switch typePot {
				case "s":
					UtiliserObjet(lst, "potion de soin")
				case "p":
					UtiliserObjet(lst, "potion de poison")
				default:
					fmt.Println("Choix de potion invalide.")
				}
			case "arm":
				var typearm string
				fmt.Print("cuir fer ou mithril ? (c/f/m) : ")
				fmt.Scan(&typearm)

				switch typearm {
				case "c":
					UtiliserObjet(lst, "armure en cuir")
				case "f":
					UtiliserObjet(lst, "armure en fer forgé")
				case "m":
					UtiliserObjet(lst, "armure en mithril")
				default:
					fmt.Println("Choix d'armure invalide.")
				}
			}
		}
	}
}
