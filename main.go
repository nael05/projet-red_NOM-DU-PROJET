package main

import (

	"jeu/outils"
)

func main() {
	outils.ClearScreen()

	outils.Menu()

	var joueur outils.Perso
	joueur.InitCharacter()
	outils.Printcharactere(joueur)
	outils.Boutique(&joueur)

	// Création d’un adversaire
	ordi := outils.GenererInventaireOrdi()
	outils.Combat(&joueur, &ordi)
}
