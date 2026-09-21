package combat

import (
	"Projet_red/character"
	"Projet_red/monstres"
	"fmt"
)

func goblinPattern(joueur *character.Character, gobelin *monstres.Goblins, tour int) {
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
