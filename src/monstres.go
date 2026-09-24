package main

type Goblins struct {
	Nom    string
	Niveau int
	Pvmax  int
	Pv     int
	Att    int
}

func initGobelins() Goblins {
	return Goblins{
		Nom:    "Maxence le goblin",
		Niveau: 1,
		Pvmax:  40,
		Pv:     40,
		Att:    5,
	}
}

type Orcs struct {
	Nom    string
	Niveau int
	Pvmax  int
	Pv     int
	Att    int
}

func initOrcs() Orcs {
	return Orcs{
		Nom:    "Thomas l'orc",
		Niveau: 5,
		Pvmax:  80,
		Pv:     80,
		Att:    10,
	}
}

type Wargs struct {
	Nom    string
	Niveau int
	Pvmax  int
	Pv     int
	Att    int
}

func initWargs() Wargs {
	return Wargs{
		Nom:    "Dylan le warg",
		Niveau: 10,
		Pvmax:  150,
		Pv:     150,
		Att:    25,
	}
}

type Uruk_hai struct {
	Nom    string
	Niveau int
	Pvmax  int
	Pv     int
	Att    int
}

func initUruk_hai() Uruk_hai {
	return Uruk_hai{
		Nom:    "Bob l'uruk",
		Niveau: 25,
		Pvmax:  250,
		Pv:     250,
		Att:    30,
	}
}

type Sauron struct {
	Nom    string
	Niveau int
	Pvmax  int
	Pv     int
	Att    int
}

func initSauron() Sauron {
	return Sauron{
		Nom:    "Sauron",
		Niveau: 80,
		Pvmax:  500,
		Pv:     500,
		Att:    50,
	}
}
