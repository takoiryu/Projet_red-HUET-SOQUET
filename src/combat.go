package main

import (
	"fmt"
)

func goblinTurn(joueur *Character, gobelin *Goblins, tour int) {
	degats := gobelin.Att
	if tour%3 == 0 {
		degats = gobelin.Att * 2
	}
	joueur.Pv -= degats
	if joueur.Pv < 0 {
		joueur.Pv = 0
	}
	fmt.Printf("%s inflige à %s %d de dégâts !\n", gobelin.Nom, joueur.Nom, degats)
	fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
}

func characterTurn1(joueur *Character, monstre *Goblins) {
	var choix string
	magazin := InitGandalf()
	forge := InitGimly()
	for {
		if joueur.Pv <= 0 {
			fmt.Printf("\nDéfaite..vous êtes KO ..t'est pas très fort hein\n")
			Fight1(joueur)
			return
		}
		if monstre.Pv <= 0 {
			for i := range joueur.Inventaire {
				if joueur.Inventaire[i].NomObj == "pièce d'or" {
					joueur.Inventaire[i].Quantite += 5
					fmt.Println("\nVous gagnez 5 pièces d'or !")
					break
				}
			}
			fmt.Printf("\nVictoire ! Prochain combat.\n")
			Fight2(joueur)
			return
		}
		fmt.Println("\n===== TOUR DE JOUEUR =====")
		fmt.Println("1 : Attaquer")
		fmt.Println("2 : Menu")
		fmt.Print("Choix : ")
		fmt.Scan(&choix)
		switch choix {
		case "1":
			var typesort string
			fmt.Println("\nChoisissez votre sort :")
			fmt.Println("1 : Coup de poing")
			fmt.Println("2 : Boule de feu")
			fmt.Print("Votre choix : ")
			fmt.Scan(&typesort)

			degat := 0

			switch typesort {
			case "1":
				fmt.Println("\nCoup de poing")
				degat = 10
				monstre.Pv = monstre.Pv - degat
				joueur.Pv -= 5
				fmt.Printf("%s inflige à %s %d de dégâts !\n", joueur.Nom, monstre.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
				fmt.Printf("%s inflige à %s %d de dégâts !\n", monstre.Nom, joueur.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
				degat = 0
			case "2":
				fmt.Println("\nBoule de feu")
				degat = 20
				monstre.Pv = monstre.Pv - degat
				joueur.Pv -= 5
				fmt.Printf("%s inflige à %s %d de dégâts !\n", joueur.Nom, monstre.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
				fmt.Printf("%s inflige à %s %d de dégâts !\n", monstre.Nom, joueur.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
				degat = 0
			default:
				fmt.Println("Choix invalide.")
				continue
			}
			monstre.Pv -= degat
		case "2":
			menu(joueur, magazin, forge)
			return
		default:
			fmt.Println("\nChoix invalide ! Veuillez saisir 1 ou 2.")
		}
	}
}

func Fight1(joueur *Character) {
	gobelin := initGobelins()
	tour := 1
	fmt.Println("\n==========================================")
	fmt.Printf("      DÉBUT DU COMBAT contre %s\n", gobelin.Nom)
	fmt.Println("==========================================")
	for joueur.Pv > 0 && gobelin.Pv > 0 {
		fmt.Printf("\n======TOUR %d ======\n", tour)

		characterTurn1(joueur, &gobelin)

		if gobelin.Pv <= 0 {
			fmt.Printf("\nVictoire ! Vous avez terrassé %s !\n", gobelin.Nom)
			fmt.Println("Retour au menu principal...")
			return
		}

		goblinTurn(joueur, &gobelin, tour)

		tour++
	}
}

func orcsTurn(joueur *Character, orcs *Orcs, tour int) {
	degats := orcs.Att
	if tour%3 == 0 {
		degats = orcs.Att * 2
	}
	joueur.Pv -= degats
	if joueur.Pv < 0 {
		joueur.Pv = 0
	}
	fmt.Printf("%s inflige à %s %d de dégâts !\n", orcs.Nom, joueur.Nom, degats)
	fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
}

func characterTurn2(joueur *Character, monstre *Orcs) {
	var choix string
	magazin := InitGandalf()
	forge := InitGimly()
	for {
		if joueur.Pv <= 0 {
			fmt.Printf("\nDéfaite..vous êtes KO ..t'est pas très fort hein\n")
			Fight2(joueur)
			return
		}
		if monstre.Pv <= 0 {
			for i := range joueur.Inventaire {
				if joueur.Inventaire[i].NomObj == "pièce d'or" {
					joueur.Inventaire[i].Quantite += 10
					fmt.Println("\nVous gagnez 10 pièces d'or !")
					break
				}
			}
			fmt.Printf("\nVictoire ! Prochain combat.\n")
			Fight3(joueur)
			return
		}
		fmt.Println("\n===== TOUR DE JOUEUR =====")
		fmt.Println("1 : Attaquer")
		fmt.Println("2 : Menu")
		fmt.Print("Choix : ")
		fmt.Scan(&choix)
		switch choix {
		case "1":
			var typesort string
			fmt.Println("\nChoisissez votre sort :")
			fmt.Println("1 : Coup de poing")
			fmt.Println("2 : Boule de feu")
			fmt.Print("Votre choix : ")
			fmt.Scan(&typesort)

			degat := 0

			switch typesort {
			case "1":
				fmt.Println("\nCoup de poing")
				degat = 10
				monstre.Pv = monstre.Pv - degat
				joueur.Pv -= 10
				fmt.Printf("%s inflige à %s %d de dégâts !\n", joueur.Nom, monstre.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
				fmt.Printf("%s inflige à %s %d de dégâts !\n", monstre.Nom, joueur.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
				degat = 0
			case "2":
				fmt.Println("\nBoule de feu")
				degat = 20
				monstre.Pv = monstre.Pv - degat
				joueur.Pv -= 10
				fmt.Printf("%s inflige à %s %d de dégâts !\n", joueur.Nom, monstre.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
				fmt.Printf("%s inflige à %s %d de dégâts !\n", monstre.Nom, joueur.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
				degat = 0
			default:
				fmt.Println("Choix invalide.")
				continue
			}
			monstre.Pv -= degat
		case "2":
			menu(joueur, magazin, forge)
			return
		default:
			fmt.Println("\nChoix invalide ! Veuillez saisir 1 ou 2.")
		}
	}
}

func Fight2(joueur *Character) {
	orcs := initOrcs()
	tour := 1
	fmt.Println("\n==========================================")
	fmt.Printf("      DÉBUT DU COMBAT contre %s\n", orcs.Nom)
	fmt.Println("==========================================")
	for joueur.Pv > 0 && orcs.Pv > 0 {
		fmt.Printf("\n======TOUR %d ======\n", tour)

		characterTurn2(joueur, &orcs)

		if orcs.Pv <= 0 {
			fmt.Printf("\nVictoire ! Vous avez terrassé %s !\n", orcs.Nom)
			fmt.Println("Retour au menu principal...")
			return
		}

		orcsTurn(joueur, &orcs, tour)

		tour++
	}
}

func wargsTurn(joueur *Character, wargs *Wargs, tour int) {
	degats := wargs.Att
	if tour%3 == 0 {
		degats = wargs.Att * 2
	}
	joueur.Pv -= degats
	if joueur.Pv < 0 {
		joueur.Pv = 0
	}
	fmt.Printf("%s inflige à %s %d de dégâts !\n", wargs.Nom, joueur.Nom, degats)
	fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
}

func characterTurn3(joueur *Character, monstre *Wargs) {
	var choix string
	magazin := InitGandalf()
	forge := InitGimly()
	for {
		if joueur.Pv <= 0 {
			fmt.Printf("\nDéfaite..vous êtes KO ..t'est pas très fort hein\n")
			Fight3(joueur)
			return
		}
		if monstre.Pv <= 0 {
			for i := range joueur.Inventaire {
				if joueur.Inventaire[i].NomObj == "pièce d'or" {
					joueur.Inventaire[i].Quantite += 20
					fmt.Println("\nVous gagnez 20 pièces d'or !")
					break
				}
			}
			fmt.Printf("\nVictoire ! Prochain combat.\n")
			Fight4(joueur)
			return
		}
		fmt.Println("\n===== TOUR DE JOUEUR =====")
		fmt.Println("1 : Attaquer")
		fmt.Println("2 : Menu")
		fmt.Print("Choix : ")
		fmt.Scan(&choix)
		switch choix {
		case "1":
			var typesort string
			fmt.Println("\nChoisissez votre sort :")
			fmt.Println("1 : Coup de poing")
			fmt.Println("2 : Boule de feu")
			fmt.Print("Votre choix : ")
			fmt.Scan(&typesort)

			degat := 0

			switch typesort {
			case "1":
				fmt.Println("\nCoup de poing")
				degat = 10
				monstre.Pv = monstre.Pv - degat
				joueur.Pv -= 25
				fmt.Printf("%s inflige à %s %d de dégâts !\n", joueur.Nom, monstre.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
				fmt.Printf("%s inflige à %s %d de dégâts !\n", monstre.Nom, joueur.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
				degat = 0
			case "2":
				fmt.Println("\nBoule de feu")
				degat = 30
				monstre.Pv = monstre.Pv - degat
				joueur.Pv -= 25
				fmt.Printf("%s inflige à %s %d de dégâts !\n", joueur.Nom, monstre.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
				fmt.Printf("%s inflige à %s %d de dégâts !\n", monstre.Nom, joueur.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
				degat = 0
			default:
				fmt.Println("Choix invalide.")
				continue
			}
			monstre.Pv -= degat
		case "2":
			menu(joueur, magazin, forge)
			return
		default:
			fmt.Println("\nChoix invalide ! Veuillez saisir 1 ou 2.")
		}
	}
}

func Fight3(joueur *Character) {
	wargs := initWargs()
	tour := 1
	fmt.Println("\n==========================================")
	fmt.Printf("      DÉBUT DU COMBAT contre %s\n", wargs.Nom)
	fmt.Println("==========================================")
	for joueur.Pv > 0 && wargs.Pv > 0 {
		fmt.Printf("\n======TOUR %d ======\n", tour)

		characterTurn3(joueur, &wargs)

		if wargs.Pv <= 0 {
			fmt.Printf("\nVictoire ! Vous avez terrassé %s !\n", wargs.Nom)
			fmt.Println("Retour au menu principal...")
			return
		}

		wargsTurn(joueur, &wargs, tour)

		tour++
	}
}

func uruk_haiTurn(joueur *Character, uruk_hai *Uruk_hai, tour int) {
	degats := uruk_hai.Att
	if tour%3 == 0 {
		degats = uruk_hai.Att * 2
	}
	joueur.Pv -= degats
	if joueur.Pv < 0 {
		joueur.Pv = 0
	}
	fmt.Printf("%s inflige à %s %d de dégâts !\n", uruk_hai.Nom, joueur.Nom, degats)
	fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
}

func characterTurn4(joueur *Character, monstre *Uruk_hai) {
	var choix string
	magazin := InitGandalf()
	forge := InitGimly()
	for {
		if joueur.Pv <= 0 {
			fmt.Printf("\nDéfaite..vous êtes KO ..t'est pas très fort hein\n")
			Fight4(joueur)
			return
		}
		if monstre.Pv <= 0 {
			for i := range joueur.Inventaire {
				if joueur.Inventaire[i].NomObj == "pièce d'or" {
					joueur.Inventaire[i].Quantite += 40
					fmt.Println("\nVous gagnez 40 pièces d'or !")
					break
				}
			}
			fmt.Printf("\nVictoire ! Prochain combat.\n")
			Fight5(joueur)
			return
		}
		fmt.Println("\n===== TOUR DE JOUEUR =====")
		fmt.Println("1 : Attaquer")
		fmt.Println("2 : Menu")
		fmt.Print("Choix : ")
		fmt.Scan(&choix)
		switch choix {
		case "1":
			var typesort string
			fmt.Println("\nChoisissez votre sort :")
			fmt.Println("1 : Coup de poing")
			fmt.Println("2 : Boule de feu")
			fmt.Print("Votre choix : ")
			fmt.Scan(&typesort)

			degat := 0

			switch typesort {
			case "1":
				fmt.Println("\nCoup de poing")
				degat = 10
				monstre.Pv = monstre.Pv - degat
				joueur.Pv -= 10
				fmt.Printf("%s inflige à %s %d de dégâts !\n", joueur.Nom, monstre.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
				fmt.Printf("%s inflige à %s %d de dégâts !\n", monstre.Nom, joueur.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
				degat = 0
			case "2":
				fmt.Println("\nBoule de feu")
				degat = 30
				monstre.Pv = monstre.Pv - degat
				joueur.Pv -= 30
				fmt.Printf("%s inflige à %s %d de dégâts !\n", joueur.Nom, monstre.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
				fmt.Printf("%s inflige à %s %d de dégâts !\n", monstre.Nom, joueur.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
				degat = 0
			default:
				fmt.Println("Choix invalide.")
				continue
			}
			monstre.Pv -= degat
		case "2":
			menu(joueur, magazin, forge)
			return
		default:
			fmt.Println("\nChoix invalide ! Veuillez saisir 1 ou 2.")
		}
	}
}

func Fight4(joueur *Character) {
	uruk_hai := initUruk_hai()
	tour := 1
	fmt.Println("\n==========================================")
	fmt.Printf("      DÉBUT DU COMBAT contre %s\n", uruk_hai.Nom)
	fmt.Println("==========================================")
	for joueur.Pv > 0 && uruk_hai.Pv > 0 {
		fmt.Printf("\n======TOUR %d ======\n", tour)

		characterTurn4(joueur, &uruk_hai)

		if uruk_hai.Pv <= 0 {
			fmt.Printf("\nVictoire ! Vous avez terrassé %s !\n", uruk_hai.Nom)
			fmt.Println("Retour au menu principal...")
			return
		}

		uruk_haiTurn(joueur, &uruk_hai, tour)

		tour++
	}
}

func sauronTurn(joueur *Character, sauron *Sauron, tour int) {
	degats := sauron.Att
	if tour%3 == 0 {
		degats = sauron.Att * 2
	}
	joueur.Pv -= degats
	if joueur.Pv < 0 {
		joueur.Pv = 0
	}
	fmt.Printf("%s inflige à %s %d de dégâts !\n", sauron.Nom, joueur.Nom, degats)
	fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
}

func characterTurn5(joueur *Character, monstre *Sauron) {
	var choix string
	magazin := InitGandalf()
	forge := InitGimly()
	for {
		if joueur.Pv <= 0 {
			fmt.Printf("\nDéfaite..vous êtes KO ..t'est pas très fort hein.\n")
			Fight4(joueur)
			return
		}
		if monstre.Pv <= 0 {
			for i := range joueur.Inventaire {
				if joueur.Inventaire[i].NomObj == "pièce d'or" {
					joueur.Inventaire[i].Quantite += 100
					fmt.Println("\nVous gagnez 100 pièces d'or !")
					break
				}
			}
			fmt.Printf("\nVictoire ! La terre du milieu est sauvée !\n")
			return
		}
		fmt.Println("\n===== TOUR DE JOUEUR =====")
		fmt.Println("1 : Attaquer")
		fmt.Println("2 : Menu")
		fmt.Print("Choix : ")
		fmt.Scan(&choix)
		switch choix {
		case "1":
			var typesort string
			fmt.Println("\nChoisissez votre sort :")
			fmt.Println("1 : Coup de poing")
			fmt.Println("2 : Boule de feu")
			fmt.Print("Votre choix : ")
			fmt.Scan(&typesort)

			degat := 0

			switch typesort {
			case "1":
				fmt.Println("\nCoup de poing")
				degat = 10
				monstre.Pv = monstre.Pv - degat
				joueur.Pv -= 10
				fmt.Printf("%s inflige à %s %d de dégâts !\n", joueur.Nom, monstre.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
				fmt.Printf("%s inflige à %s %d de dégâts !\n", monstre.Nom, joueur.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
				degat = 0
			case "2":
				fmt.Println("\nBoule de feu")
				degat = 30
				monstre.Pv = monstre.Pv - degat
				joueur.Pv -= 50
				fmt.Printf("%s inflige à %s %d de dégâts !\n", joueur.Nom, monstre.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", monstre.Nom, monstre.Pv, monstre.Pvmax)
				fmt.Printf("%s inflige à %s %d de dégâts !\n", monstre.Nom, joueur.Nom, degat)
				fmt.Printf("Santé de %s : %d/%d PV\n", joueur.Nom, joueur.Pv, joueur.Pvmax)
				degat = 0
			default:
				fmt.Println("Choix invalide.")
				continue
			}
			monstre.Pv -= degat
		case "2":
			menu(joueur, magazin, forge)
			return
		default:
			fmt.Println("\nChoix invalide ! Veuillez saisir 1 ou 2.")
		}
	}
}

func Fight5(joueur *Character) {
	sauron := initSauron()
	tour := 1
	fmt.Println("\n==========================================")
	fmt.Printf("      DÉBUT DU COMBAT contre %s\n", sauron.Nom)
	fmt.Println("==========================================")
	for joueur.Pv > 0 && sauron.Pv > 0 {
		fmt.Printf("\n======TOUR %d ======\n", tour)

		characterTurn5(joueur, &sauron)

		if sauron.Pv <= 0 {
			fmt.Printf("\nVictoire ! Vous avez terrassé %s !\n", sauron.Nom)
			fmt.Println("Retour au menu principal...")
			return
		}

		sauronTurn(joueur, &sauron, tour)

		tour++
	}
}
