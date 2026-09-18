package character

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
