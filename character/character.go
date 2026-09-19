package character

import (
	"Projet_red/inventaire"
	"fmt"
)

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	Pvmax      int
	Pv         int
	Inventaire []inventaire.Invent
}

func InitCharacter(nom string, classe string, niveau int, pvmax int, pv int, inventaire []inventaire.Invent) Character {
	return Character{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		Pvmax:      pvmax,
		Pv:         pv,
		Inventaire: inventaire,
	}
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
	fmt.Println("╚══════════════════════════════════════════════════╝")
}
func (c *Character) CharacterCreation() {
	fmt.Println("nom de votre personnage :")
	fmt.Scan(&c.Nom)
	var choix string
	for {
		fmt.Println("classe de votre personnage : \n 1:humain \n 2:nain \n 3:hobbit")
		fmt.Scan(&choix)
		switch choix {
		case "1":
			c.Classe = "humain"
		case "2":
			c.Classe = "nain"
		case "3":
			c.Classe = "hobbit"
		default:
			fmt.Println("\n Choix invalide ! Veuillez saisir 1, 2 ou 3.")
			continue
		}
		break
	}
	fmt.Println("niveau de votre personnage :")
	fmt.Scan(&c.Niveau)
	fmt.Println("nombre de points de vie de votre personnage :")
	fmt.Scan(&c.Pvmax)
	c.Pv = c.Pvmax
}
