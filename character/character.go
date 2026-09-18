package character

import "fmt"

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	Pvmax      int
	Pv         int
	Inventaire []string
}

func InitCharacter(nom string, classe string, niveau int, pvmax int, pv int, inventaire []string) Character {
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
	fmt.Println("classe de votre personnage :")
	fmt.Scan(&c.Classe)
	fmt.Println("niveau de votre personnage :")
	fmt.Scan(&c.Niveau)
	fmt.Println("nombre de points de vie de votre personnage :")
	fmt.Scan(&c.Pvmax)
	c.Pv = c.Pvmax
}
