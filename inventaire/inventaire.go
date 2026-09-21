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
		{NomObj: "chapeau en cuir", Quantite: 0},
		{NomObj: "tunique en cuir", Quantite: 0},
		{NomObj: "bottes en cuir", Quantite: 0},
		{NomObj: "heaume en fer forgé", Quantite: 0},
		{NomObj: "plastron en fer forgé", Quantite: 0},
		{NomObj: "bottes en fer forgé", Quantite: 0},
		{NomObj: "cote de maille en mithril", Quantite: 0},
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
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║               INVENTAIRE                 ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	vide := true
	for _, item := range lst {
		if item.Quantite > 0 {
			fmt.Printf("║  • %s : %d ║\n", item.NomObj, item.Quantite)
			vide = false
		}
	}
	if vide {
		fmt.Println("║        (Votre sac est totalement vide)   ║")
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
				var piece string
				switch typearm {
				case "c":
					fmt.Println("vous choisiser l'armure en cuir, quelle piece ? (tete/corp/pieds)t/c/p")
					fmt.Scan(&piece)
					switch piece {
					case "t":
						UtiliserObjet(lst, "chapeau en cuir")
					case "c":
						UtiliserObjet(lst, "plastron en fer forgé")
					case "p":
						UtiliserObjet(lst, "bottes en cuir")
					default:
						fmt.Println("Choix de piece invalide")
					}

				case "f":
					fmt.Println("vous choisiser l'armure en fer forgé, quelle piece ? (tete/corp/pieds) t/c/p")
					fmt.Scan(&piece)
					switch piece {
					case "t":
						UtiliserObjet(lst, "heaume en fer forgé")
					case "c":
						UtiliserObjet(lst, "tunique en cuir")
					case "p":
						UtiliserObjet(lst, "bottes en fer forgé")
					default:
						fmt.Println("Choix de piece invalide")
					}
				case "m":
					UtiliserObjet(lst, "cote de maille en mithril")
				default:
					fmt.Println("Choix d'armure invalide.")
				}
			}
		}
	}
	fmt.Println("╚══════════════════════════════════════════╝")
}
