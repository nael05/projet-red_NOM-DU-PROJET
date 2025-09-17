package main

import (

	"jeu/outils"
)
//La fonction main appel toutes les fonctions annexe pour le bon deroulement du jeu, dans le bonne ordre
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
