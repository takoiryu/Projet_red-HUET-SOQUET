package main

import (
	"fmt"
)

type Equipement struct {
	Tete  string
	Torse string
	Pieds string
}
type Character struct {
	Nom                string
	Classe             string
	Niveau             int
	Pvmax              int
	Pv                 int
	Atk                int
	Inventaire         []Invent
	Equipement         Equipement
	SkillList          []string
	TailleMax          int
	NbrAmeliorationSac int
}

func (c Character) DisplayInfo() {
	affpvmax := c.Pvmax
	affpv := c.Pv
	affbar := 0
	barvid := 20
	var divpv float64
	var porcpv int
	divpv = float64(affpv) / float64(affpvmax)
	porcpv = int(divpv * 100)
	affbar = porcpv / 5
	barvid = 20 - affbar
	fmt.Println("\n╔══════════════════════════════════════════════════╗")
	fmt.Println("║               FICHE DE PERSONNAGE                ║")
	fmt.Println("╠══════════════════════════════════════════════════╣")
	fmt.Printf("║  Nom        : %-34s ║\n", c.Nom)
	fmt.Printf("║  Classe     : %-34s ║\n", c.Classe)
	fmt.Printf("║  Niveau     : %-34d ║\n", c.Niveau)
	fmt.Println("╠══════════════════════════════════════════════════╣")
	fmt.Printf("║  Santé      : ")
	for affbar > 0 {
		fmt.Printf("█")
		affbar--
	}
	for barvid > 0 {
		fmt.Printf("░")
		barvid--
	}
	fmt.Printf("PV             ║\n")
	fmt.Println("╠══════════════════════════════════════════════════╣")
	fmt.Println("║                   ÉQUIPEMENT                     ║")
	fmt.Println("╠══════════════════════════════════════════════════╣")
	fmt.Printf("║  Tête       : %-34s ║\n", c.Equipement.Tete)
	fmt.Printf("║  Torse      : %-34s ║\n", c.Equipement.Torse)
	fmt.Printf("║  Pieds      : %-34s ║\n", c.Equipement.Pieds)
	fmt.Println("╚══════════════════════════════════════════════════╝")
}
func (c *Character) CharacterCreation() {
	fmt.Println("Sélectionnez le nom de votre personnage :")
	fmt.Scan(&c.Nom)
	var choix string
	for {
		fmt.Println("Sélectionnez la classe de votre personnage : \n 1 : Humain \n 2 : Hain \n 3 : Hobbit")
		fmt.Scan(&choix)
		switch choix {
		case "1":
			c.Classe = "Humain"
			c.Pvmax = 100
		case "2":
			c.Classe = "Nain"
			c.Pvmax = 120
		case "3":
			c.Classe = "Hobbit"
			c.Pvmax = 80
		default:
			fmt.Println("\n Choix invalide ! Veuillez saisir 1, 2 ou 3.")
			continue
		}
		break
	}
	c.Pv = 2000
	c.TailleMax = 10
}
