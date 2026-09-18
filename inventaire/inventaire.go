package inventaire

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
