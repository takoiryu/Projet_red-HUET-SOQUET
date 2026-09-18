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
	fmt.Println("Nom du personnage : ", c.Nom)
	fmt.Println("Classe du personnage : ", c.Classe)
	fmt.Println("Niveau du personnage : ", c.Niveau)
	fmt.Println("Nombre max de points de vie du personnage : ", c.Pvmax)
	fmt.Println("Nombre actuel de points de vie du personnage : ", c.Pv)
	fmt.Println("contenu de l'inventaire : ", c.Inventaire)
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
