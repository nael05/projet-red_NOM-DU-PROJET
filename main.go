package main

import ("jeu/outils")

func main() {
	outils.ClearScreen()

	outils.Menu()

	var joueur outils.Perso
	joueur.InitCharacter()
	outils.Printcharactere(joueur)
	outils.Boutique(&joueur)
	outils.Combat(j *Perso, ordi *Perso)
}