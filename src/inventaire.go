package main

import (
	"fmt"
	"time"
)

type Invent struct {
	NomObj   string
	Quantite int
}

func InitInventaire() []Invent {
	return []Invent{
		{NomObj: "pièce d'or", Quantite: 100},
		{NomObj: "potion de soin", Quantite: 2},
		{NomObj: "potion de poison", Quantite: 0},
		{NomObj: "chapeau en cuir", Quantite: 0},
		{NomObj: "tunique en cuir", Quantite: 0},
		{NomObj: "bottes en cuir", Quantite: 0},
		{NomObj: "heaume en fer forgé", Quantite: 0},
		{NomObj: "plastron en fer forgé", Quantite: 0},
		{NomObj: "bottes en fer forgé", Quantite: 0},
		{NomObj: "cote de maille en mithril", Quantite: 0},
		{NomObj: "cuir de goblin", Quantite: 0},
		{NomObj: "lingots de fer", Quantite: 0},
		{NomObj: "minerais de mithril", Quantite: 0},
		{NomObj: "Sort: Boule de feu", Quantite: 0},
	}
}
func UtiliserObjet(joueur *Character, nomObjet string) bool {
	for i := range joueur.Inventaire {
		if joueur.Inventaire[i].NomObj == nomObjet {
			if joueur.Inventaire[i].Quantite > 0 {
				joueur.Inventaire[i].Quantite--
				fmt.Printf("Vous avez utilisé : %s (Reste : %d)\n", nomObjet, joueur.Inventaire[i].Quantite)
				if nomObjet == "potion de soin" {
					UtilPot(joueur)
				}
				if nomObjet == "Sort: Boule de feu" {
					joueur.Inventaire[i].Quantite++
				}
				var ancObj string
				switch nomObjet {
				case "chapeau en cuir":
					ancObj = joueur.Equipement.Tete
					joueur.Equipement.Tete = "chapeau en cuir"
					joueur.Pvmax += 10
				case "heaume en fer forgé":
					ancObj = joueur.Equipement.Tete
					joueur.Equipement.Tete = "heaume en fer forgé"
					joueur.Pvmax += 20
				case "tunique en cuir":
					ancObj = joueur.Equipement.Torse
					joueur.Equipement.Torse = "tunique en cuir"
					joueur.Pvmax += 25
				case "plastron en fer forgé":
					ancObj = joueur.Equipement.Torse
					joueur.Equipement.Torse = "plastron en fer forgé"
					joueur.Pvmax += 40
				case "cote de maille en mithril":
					ancObj = joueur.Equipement.Torse
					joueur.Equipement.Torse = "cote de maille en mithril"
					joueur.Pvmax += 200
				case "bottes en fer forgé":
					ancObj = joueur.Equipement.Pieds
					joueur.Equipement.Pieds = "bottes en fer forgé"
					joueur.Pvmax += 30
				case "bottes en cuir":
					ancObj = joueur.Equipement.Pieds
					joueur.Equipement.Pieds = "bottes en cuir"
					joueur.Pvmax += 15
				case "potion de poison":
					PoisonPot(joueur)
				}
				switch ancObj {
				case "chapeau en cuir":
					joueur.Pvmax -= 10
				case "heaume en fer forgé":
					joueur.Pvmax -= 20
				case "tunique en cuir":
					joueur.Pvmax -= 25
				case "plastron en fer forgé":
					joueur.Pvmax -= 40
				case "cote de maille en mithril":
					joueur.Pvmax -= 200
				case "bottes en cuir":
					joueur.Pvmax -= 15
				case "bottes en fer forgé":
					joueur.Pvmax -= 30
				}
				trouve := false
				for i := range joueur.Inventaire {
					if joueur.Inventaire[i].NomObj == ancObj {
						joueur.Inventaire[i].Quantite++
						trouve = true
						break
					}
				}

				if !trouve {
					joueur.Inventaire = append(joueur.Inventaire, Invent{NomObj: ancObj, Quantite: 1})
				}
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
func AfficherInvent(joueur *Character) {
	fmt.Println("\n╔══════════════════════════════════════════╗")
	fmt.Println("║               INVENTAIRE                 ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	vide := true
	for _, item := range joueur.Inventaire {
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
			fmt.Print("Quel objet ? potions (pot), armure (arm) :")
			fmt.Scan(&usechoix)

			switch usechoix {
			case "pot":
				var typePot string
				fmt.Print("Soin ou poison ? (s/p) : ")
				fmt.Scan(&typePot)

				switch typePot {
				case "s":
					UtiliserObjet(joueur, "potion de soin")
				case "p":
					UtiliserObjet(joueur, "potion de poison")
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
						UtiliserObjet(joueur, "chapeau en cuir")
					case "c":
						UtiliserObjet(joueur, "tunique en cuir")
					case "p":
						UtiliserObjet(joueur, "bottes en cuir")
					default:
						fmt.Println("Choix de piece invalide")
					}

				case "f":
					fmt.Println("vous choisiser l'armure en fer forgé, quelle piece ? (tete/corp/pieds) t/c/p")
					fmt.Scan(&piece)
					switch piece {
					case "t":
						UtiliserObjet(joueur, "heaume en fer forgé")
					case "c":
						UtiliserObjet(joueur, "tunique en cuir")
					case "p":
						UtiliserObjet(joueur, "bottes en fer forgé")
					default:
						fmt.Println("Choix de piece invalide")
					}
				case "m":
					UtiliserObjet(joueur, "cote de maille en mithril")
				default:
					fmt.Println("Choix d'armure invalide.")
				}
			}
		}
	}
	fmt.Println("╚══════════════════════════════════════════╝")
}
func UtilPot(c *Character) {
	if c.Pv >= c.Pvmax {
		fmt.Println("Vos PV sont déjà au max !")

		return
	}
	c.Pv += 50
	if c.Pv > c.Pvmax {
		c.Pv = c.Pvmax
	}
}
func LimitInv(inv []Invent, joueur *Character) bool {
	somme := 0
	for i := range inv {
		if inv[i].NomObj == "pièce d'or" {
			continue
		}
		somme += inv[i].Quantite
	}
	return somme >= joueur.TailleMax
}
func PoisonPot(joueur *Character) {
	fmt.Println("\n🧪 Vous buvez une potion de poison ! Le poison s'infiltre dans votre sang...")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		joueur.Pv -= 10
		if joueur.Pv < 0 {
			joueur.Pv = 0
		}
		fmt.Printf("🤢 Degâts du poison (%ds/3s) : -10 PV | PV actuels : %d/%d\n", i, joueur.Pv, joueur.Pvmax)
		if joueur.Pv == 0 {
			fmt.Println("💀 Vous avez succombé au poison...")
			joueur.Niveau = 1
			joueur.Pv = joueur.Pvmax / 2
			return
		}
	}
}
